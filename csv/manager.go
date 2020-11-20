package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"gorm.io/gorm"
)

type Manager struct {
	fileConfigs []FileConfiguration
	fileWatcher *fsnotify.Watcher
	mux         sync.Mutex
}

func (m *Manager) getConfigForFile(filename string) (FileConfiguration, error) {
	for _, conf := range m.fileConfigs {
		if conf.FilePattern == "" {
			continue
		}

		match, err := regexp.MatchString(conf.FilePattern, filename)

		if err != nil {
			return conf, err
		}

		if match {
			return conf, nil
		}
	}

	return FileConfiguration{}, errors.New("could not find config for file")
}

func (m *Manager) getVal(col ColumnMapping, row RowData) (interface{}, error) {
	v := row[col.Ordinal-1]

	switch col.Type {
	case Float64Column:
		return m.parseString(v)
	case StrColumn:
		return m.parseString(v)
	case TimestampColumn:
		return m.parseTimestamp(v, col.Args)
	default:
		return nil, errors.New("invalid column mapping type")
	}
}

func (m *Manager) parseTimestamp(v interface{}, args JSONB) (t time.Time, err error) {
	s, err := m.parseString(v)

	if err != nil {
		return t, err
	}

	// Parse timestamp format from args
	format, ok := args["timestamp_format"].(string)

	if !ok {
		format = "2006-01-02"
	}

	t, err = time.Parse(format, s)

	return t, err
}

func (m *Manager) parseString(v interface{}) (string, error) {
	return fmt.Sprintf("%v", v), nil
}

func (m *Manager) parseFloat(v interface{}) (float64, error) {
	s, err := m.parseString(v)

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(s, 64)
}

func (m *Manager) parseRow(conf FileConfiguration, row RowData) (Record, error) {
	record := make(map[string]interface{})

	for _, col := range conf.ColumnMappings {
		if col.Ordinal-1 > len(record) {
			return nil, fmt.Errorf("Column mapping ordinal outside record index")
		}
		key := col.Name
		val, err := m.getVal(col, row)

		if err != nil {
			return nil, err
		}

		record[key] = val
	}

	return record, nil
}

func (m *Manager) ParseFile(file string) (chan Record, chan error) {
	// make channels for sending data / errors to consumer
	outCh := make(chan Record)
	errCh := make(chan error)
	var f *os.File

	go func() {
		ext := filepath.Ext(file)

		defer func() {
			if f != nil {
				f.Close()
			}
			close(outCh)
			close(errCh)
		}()

		if ext != ".csv" {
			errCh <- fmt.Errorf("file of type \"%s\" is not a valid CSV", ext)
			return
		}

		fileConfig, err := m.getConfigForFile(file)

		if err != nil {
			errCh <- fmt.Errorf("Could not get file configration for file %s: %v", file, err)
			return
		}

		// Read file
		f, err := os.Open(file)

		if err != nil {
			errCh <- fmt.Errorf("Unable to read input file %v", err)
			return
		}

		r := csv.NewReader(f)

		for {
			record, err := r.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				errCh <- err
				return
			}
			datum, err := m.parseRow(fileConfig, record)
			if err != nil {
				errCh <- err
			} else {
				outCh <- datum
			}
		}
	}()

	return outCh, errCh
}

func NewManager(db *gorm.DB) (m *Manager, err error) {
	m = &Manager{}

	// Load all csv file configurations
	db.Preload("ColumnMappings").Find(&m.fileConfigs)

	return m, nil
}

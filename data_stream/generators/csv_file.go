package generators

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type CSVFileGenerator struct {
	fileDefinitions []CSVFileDefinition
	inputDir        string
	fileWatcher     *fsnotify.Watcher
}

type CSVFileGeneratorConfig struct {
	fileDefinitions []CSVFileDefinition
	inputDir        string
}

type JSONB map[string]interface{}

type CSVColumnType string

const (
	CSVIntColumn       CSVColumnType = "int"
	CSVFloat64Column                 = "float64"
	CSVStrColumn                     = "string"
	CSVTimestampColumn               = "timestamp"
)

type CSVFileDefinition struct {
	Name        string
	FilePattern string
	Columns     []CSVColumn
}

type CSVColumn struct {
	Name    string        `json:"name"`
	Key     string        `json:"key"`
	Ordinal int           `json:"ordinal"`
	Type    CSVColumnType `json:"type"`
	Args    JSONB         `json:"args"`
}

func (gen *CSVFileGenerator) Next() (Chunk, error) {
	return nil, nil
}

func NewCSVFileGenerator(conf CSVFileGeneratorConfig) (Generator, error) {
	fw, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, err
	}

	gen := &CSVFileGenerator{
		fileDefinitions: conf.fileDefinitions,
		inputDir:        conf.inputDir,
		fileWatcher:     fw,
	}
	// Call start on generator
	gen.startGenerator()

	return gen, nil
}

// start watching input directory for new csv files
// right now this only applies to new created files
func (gen *CSVFileGenerator) startGenerator() error {
	// Start go-routine to watch for input files
	go func() {
		for {
			select {
			case event, ok := <-gen.fileWatcher.Events:
				if !ok {
					return
				}
				if event.Op != fsnotify.Create {
					continue
				}
				if filepath.Ext(event.Name) == ".csv" {
					fmt.Println(event.Name)
				}
			case err, ok := <-gen.fileWatcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// add csv input directory to file watcher
	return gen.fileWatcher.Add(gen.inputDir)
}

type parseableJobDefinition struct {
	Name        string `json:"name"`
	FilePattern string `json:"file_pattern"`
	Columns     []parseableCSVColumn
}

type parseableCSVColumn struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	Ordinal int    `json:"ordinal"`
	Type    string `json:"type"`
	Args    JSONB  `json:"args"`
}

func (p *parseableCSVColumn) normalize() (c CSVColumn, err error) {
	ct, err := parseColumnType(p.Type)

	if err != nil {
		return
	}

	return CSVColumn{
		Name:    p.Name,
		Key:     p.Key,
		Ordinal: p.Ordinal,
		Type:    ct,
		Args:    p.Args,
	}, err
}

func (p *parseableJobDefinition) normalize() (def JobDefinition, err error) {
	jt, err := parseJobtype(p.Type)

	if err != nil {
		return
	}

	cols := make([]CSVColumn, len(p.Columns))

	for idx, c := range p.Columns {
		col, err := c.normalize()
		if err != nil {
			return def, err
		}
		cols[idx] = col
	}

	return JobDefinition{
		Name:        p.Name,
		Type:        jt,
		FilePattern: p.FilePattern,
		Columns:     cols,
	}, err
}

func parseColumnType(s string) (ColumnType, error) {
	switch s {
	case "int":
		return IntColumn, nil
	case "float64":
		return Float64Column, nil
	case "string":
		return StrColumn, nil
	case "timestamp":
		return TimestampColumn, nil
	default:
		return "", fmt.Errorf("invalid column type \"%s\"", s)
	}
}

func parseJobtype(s string) (JobType, error) {
	switch s {
	case "csv_file_load":
		return CSVFileLoad, nil
	default:
		return "", fmt.Errorf("invalid job type \"%s\"", s)
	}
}

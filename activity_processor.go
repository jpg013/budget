package budget

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jpg013/budget/models"

	"gorm.io/gorm"
)

type rowMapper func([]string) (*models.Activity, error)

func makeDiscoverMapper() rowMapper {
	ordinalMap := make(map[string]int)

	return func(row []string) (data *models.Activity, err error) {
		s := strings.Split(row[0], "/")
		s1 := fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
		td, err := time.Parse("2006-01-02", s1)

		if err != nil {
			return data, err
		}

		s = strings.Split(row[1], "/")
		s1 = fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
		pd, err := time.Parse("2006-01-02", s1)

		if err != nil {
			return data, err
		}

		amount, err := strconv.ParseFloat(row[3], 32)

		if err != nil {
			return data, err
		}

		key := strings.Join(row, "-")
		fmt.Println(key)

		_, ok := ordinalMap[key]

		if !ok {
			ordinalMap[key] = 0
		}

		ordinal := ordinalMap[key]
		ordinalMap[key]++

		code := fmt.Sprintf("discover:%s-%d", key, ordinal)

		data = &models.Activity{
			TransactionDate: td,
			PostedDate:      pd,
			Description:     row[2],
			Amount:          float32(amount),
			Category:        row[4],
			Code:            code,
			SourceID:        1,
		}

		return data, err
	}
}

var fileTypes = map[string]bool{
	"discover": true,
	"capfed":   true,
}

type ActivityProcessor struct {
	db *gorm.DB
	fw *FileWatcher
	in chan *models.ActivityFile
}

func NewActivityProcessor(db *gorm.DB) *ActivityProcessor {
	return &ActivityProcessor{
		db: db,
		fw: nil,
		in: make(chan *models.ActivityFile),
	}
}

func (a *ActivityProcessor) processFile(af *models.ActivityFile) error {
	f, err := os.Open(af.Name)

	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV", err)
	}

	// Remove the first record (these are the column names)
	records = records[1:]
	mapper := makeDiscoverMapper()

	for _, r := range records {
		model, err := mapper(r)

		if err != nil {
			return err
		}

		if a.db.Model(model).Where("code = ?", model.Code).Updates(&model).RowsAffected == 0 {
			result := a.db.Create(&model)

			if result.Error != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

func (a *ActivityProcessor) handleFile(path string) error {
	fileName := filepath.Base(path)

	// parse file name
	s := strings.Split(fileName, "-")

	if len(s) < 1 {
		return errors.New(fmt.Sprintf("invalid file %s", path))
	}

	t := strings.ToLower(s[0])

	if _, ok := fileTypes[t]; !ok {
		return errors.New(fmt.Sprintf("invalid file type %s", t))
	}

	activityFile := &models.ActivityFile{
		Name:      path,
		Type:      t,
		Extension: filepath.Ext(fileName),
	}

	result := a.db.Create(activityFile)

	if result.Error != nil {
		return result.Error
	}

	go func() {
		a.in <- activityFile
	}()

	return nil
}

func (a *ActivityProcessor) Start() error {
	// Create a new file watcher
	a.fw = NewFileWatcher()
	err := a.fw.Run("./tmp")

	if err != nil {
		return err
	}

	go func() {
		for {
			evt := <-a.fw.Data

			if evt.op == createFile {
				err := a.handleFile(evt.file)

				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	go func() {
		for {
			obj := <-a.in
			a.processFile(obj)
		}
	}()

	return nil
}

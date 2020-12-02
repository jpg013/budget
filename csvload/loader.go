package csvload

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fsnotify/fsnotify"
)

var InputDir = filepath.Join("./csvload", "csv_input_files")

type Loader struct {
	db             *sql.DB
	jobDefinitions []JobDefinition
	fileWatcher    *fsnotify.Watcher
}

func NewLoader(db *sql.DB) (m *Loader, err error) {
	fw, err := fsnotify.NewWatcher()

	if err != nil {
		return m, err
	}

	m = &Loader{
		db:          db,
		fileWatcher: fw,
	}

	// load job definitions
	err = m.loadJobDefinitions()

	if err != nil {
		return m, err
	}

	return m, err
}

func (l *Loader) loadJobDefinitions() (err error) {
	dir := filepath.Join("./csvload", "job_definitions")
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return
	}

	for _, f := range files {
		path := filepath.Join(dir, f.Name())
		ext := filepath.Ext(path)

		// ignore files that are not json
		if ext != ".json" {
			continue
		}

		// open json file
		jsonFile, err := os.Open(path)

		if err != nil {
			return err
		}

		// define holder for parseable job definition
		var jobDef parseableJobDefinition

		// read json bytes
		jsonBytes, err := ioutil.ReadAll(jsonFile)

		if err != nil {
			return err
		}

		// load json bytes into parseable job definition
		err = json.Unmarshal(jsonBytes, &jobDef)

		if err != nil {
			return err
		}

		// normalize job definition
		res, err := jobDef.normalize()

		if err != nil {
			return err
		}

		// append the result to the job definitions
		l.jobDefinitions = append(l.jobDefinitions, res)
	}

	return err
}

// func (l *Loader) watchForInputFiles() {
// 	inDir := filepath.Join("./csvload", "csv_input_files")

// 	go func() {
// 		for {
// 			select {
// 			case event, ok := <-l.fileWatcher.Events:
// 				if !ok {
// 					return
// 				}
// 				if event.Op != fsnotify.Create {
// 					continue
// 				}
// 				if filepath.Ext(event.Name) == ".csv" {
// 					go l.handleCSVFile(event.Name)
// 				}
// 			case err, ok := <-l.fileWatcher.Errors:
// 				if !ok {
// 					return
// 				}
// 				log.Println("error:", err)
// 			}
// 		}
// 	}()

// 	// add csv input directory to file watcher
// 	l.fileWatcher.Add(inDir)
// }

func (l *Loader) getMatchingJobDefinition(file string) (j JobDefinition, err error) {
	// search for matching job definition
	matchingDefs := make([]JobDefinition, 0)

	for _, j := range l.jobDefinitions {
		match, _ := regexp.MatchString(j.FilePattern, file)

		if match {
			matchingDefs = append(matchingDefs, j)
		}
	}
	if len(matchingDefs) == 0 {
		return j, fmt.Errorf("Could not find matching job definition for \"%s\"", file)
	}
	if len(matchingDefs) > 1 {
		return j, fmt.Errorf("Multiple job definitions found for \"%s\"", file)
	}
	return matchingDefs[0], nil
}

// main handler for csv file
// func (l *Loader) handleCSVFile(file string) {
// 	jobDefinition, err := l.getMatchingJobDefinition(file)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Read file
// 	f, err := os.Open(file)
// 	defer f.Close()

// 	r := csv.NewReader(f)

// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			return
// 		}
// 		if err != nil {
// 			errCh <- err
// 			return
// 		}
// 		datum, err := m.parseRow(config, record)
// 		if err != nil {
// 			errCh <- err
// 		} else {
// 			code := fmt.Sprintf("%s", datum["code"])
// 			val, ok := codes[code]
// 			val++
// 			if !ok {
// 				codes[code] = val
// 			}
// 			code = fmt.Sprintf("%s:%d", code, val)
// 			datum["code"] = code
// 			outCh <- datum
// 		}
// 	}
// }

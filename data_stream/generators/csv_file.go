package generators

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type CSVFileGenerator struct {
	inputDir    string
	fileWatcher *fsnotify.Watcher
}

type CSVFileGeneratorConfig struct {
	inputDir string
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
		inputDir:    conf.inputDir,
		fileWatcher: fw,
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
				go gen.processFile(event.Name)
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

func (get *CSVFileGenerator) processFile(fileName string) {
		ext := filepath.Ext(fileName)

		if ext != ".csv" {
			return
		}

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

		// Read file
		f, err := os.Open(fileName)

		if err != nil {
			errCh <- fmt.Errorf("Unable to read input file %v", err)
			return
		}

		r := csv.NewReader(f)
		codes := make(map[string]int)
		// for {
		// 	record, err := r.Read()
		// 	if err == io.EOF {
		// 		return
		// 	}
		// 	if err != nil {
		// 		errCh <- err
		// 		return
		// 	}
		// 	datum, err := m.parseRow(config, record)
		// 	if err != nil {
		// 		errCh <- err
		// 	} else {
		// 		code := fmt.Sprintf("%s", datum["code"])
		// 		val, ok := codes[code]
		// 		val++
		// 		if !ok {
		// 			codes[code] = val
		// 		}
		// 		code = fmt.Sprintf("%s:%d", code, val)
		// 		datum["code"] = code
		// 		outCh <- datum
		// 	}
		// }
	}()

	return outCh, errCh
}

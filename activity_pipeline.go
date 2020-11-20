package budget

import (
	"fmt"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"gorm.io/gorm"
)

const TmpFileDir = "tmp_files"

// ActivityPipeline watches for new input files
// and parses them into activity.
type ActivityPipeline struct {
	fileWatcher *fsnotify.Watcher
	Done        chan bool
	db          *gorm.DB
	mux         sync.Mutex
	isRunning   bool
}

func NewActivityPipeline(db *gorm.DB) (*ActivityPipeline, error) {
	return &ActivityPipeline{db: db}, nil
	// err := a.startWatcher()

	// return a, err
}

func (ap *ActivityPipeline) Start() error {
	ap.mux.Lock()
	defer ap.mux.Unlock()

	if ap.isRunning == true {
		return nil
	}

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		return err
	}

	ap.fileWatcher = watcher

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println(event)
				// f.data <- event
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	return ap.fileWatcher.Add(TmpFileDir)
}

func (ap *ActivityPipeline) Stop() {

}

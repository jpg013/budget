package budget

import (
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type fileOP string

const (
	createFile fileOP = "create"
)

type fileEvent struct {
	op   fileOP
	file string
}

type FileWatcher struct {
	watcher *fsnotify.Watcher
	mu      sync.Mutex
	Done    chan bool
	Data    chan *fileEvent
}

func (f *FileWatcher) close() error {
	f.watcher.Close()
	close(f.Done)

	return nil
}

func (f *FileWatcher) Run(dirs ...string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.watcher != nil {
		log.Printf("file watcher already in progress")
		return nil
	}

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		return err
	}

	f.watcher = watcher

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				fileEvent := &fileEvent{
					file: event.Name,
				}

				if event.Op == fsnotify.Create {
					fileEvent.op = createFile
				}

				f.Data <- fileEvent

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	for _, dir := range dirs {
		err = f.watcher.Add(dir)

		if err != nil {
			defer f.close()
			return err
		}
	}

	return nil
}

func NewFileWatcher() *FileWatcher {
	return &FileWatcher{
		watcher: nil,
		Done:    make(chan bool),
		Data:    make(chan *fileEvent),
	}
}

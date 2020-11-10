package budget

import (
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type FileWatcher struct {
	watcher *fsnotify.Watcher
	mu      sync.Mutex
	Done    chan bool
	Data    chan fsnotify.Event
}

func (f *FileWatcher) close() error {
	f.watcher.Close()
	close(f.Done)

	return nil
}

func (f *FileWatcher) WatchDirs(dirs ...string) error {
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
				f.Data <- event
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
		Data:    make(chan fsnotify.Event),
	}
}

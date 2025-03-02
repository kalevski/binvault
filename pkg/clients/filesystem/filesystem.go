package filesystem

import (
	"binvault/pkg/cfg"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var FileQueue = make(chan string, 10)

func initQueue(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			path := filepath.Join(dirPath, file.Name())
			if err != nil {
				log.Println("error getting absolute path:", err)
				continue
			}
			FileQueue <- path
		}
	}
}

func WatchFolder(name string) {
	path := filepath.Join(cfg.GetPath("DATA_PATH"), name)
	initQueue(path)
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println("error creating folder watcher", err)
	}

	defer watcher.Close()
	log.Default().Println("Watching folder:", path)
	watcher.Add(path)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				filePath := event.Name
				FileQueue <- filePath
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Watcher error:", err)
		}
	}
}

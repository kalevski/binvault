package filesystem

import (
	"binvault/pkg/cfg"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type FileEventHandler func(path string, action string)

func WatchFolder(name string, handler FileEventHandler) {
	path := filepath.Join(cfg.GetPath("DATA_PATH"), name)
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
				handler(filePath, "create")
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Watcher error:", err)
		}
	}
}

func GetFiles(path string) []string {
	dirPath := filepath.Join(cfg.GetPath("DATA_PATH"), path)
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	var fileList []string
	for _, file := range files {
		if !file.IsDir() {
			path := filepath.Join(dirPath, file.Name())
			if err != nil {
				log.Println("error getting absolute path:", err)
				continue
			}
			fileList = append(fileList, path)
		}
	}
	return fileList
}

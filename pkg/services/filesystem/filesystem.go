package filesystem

import (
	"binvault/pkg/env"
	"binvault/pkg/models"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

const (
	FOLDER_TEMP    = "temp"
	FOLDER_BUCKETS = "buckets"
	FOLDER_PUBLIC  = "public"
)

type FileEventHandler func(path string, action string)

func WatchFolder(path string, handler FileEventHandler) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Println("error creating folder watcher", err)
	}

	defer watcher.Close()
	log.Println("Watching folder:", path)
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
			log.Println("Watcher error:", err)
		}
	}
}

func GetFiles(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var fileList []string
	for _, file := range files {
		if !file.IsDir() {
			path := filepath.Join(path, file.Name())
			if err != nil {
				log.Println("error getting absolute path:", err)
				continue
			}
			fileList = append(fileList, path)
		}
	}
	return fileList
}

func SaveFile(path string, content []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func GetFolderPath(dir string) string {
	dataPath, err := filepath.Abs(env.GetPath("DATA_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dataPath, dir)
}

func GetBucketPath(bucket string) string {
	path := filepath.Join(GetFolderPath(FOLDER_BUCKETS), bucket)
	return path
}

func GetPublicBucketPath(bucket string) string {
	path := filepath.Join(GetFolderPath(FOLDER_PUBLIC), bucket)
	return path
}

func SetBucketVisibility(bucket string, visibility models.Visibility) {
	path := GetBucketPath(bucket)
	publicPath := GetPublicBucketPath(bucket)
	if visibility == models.Visibility_Public {
		err := os.Symlink(path, publicPath)
		if err != nil {
			log.Println("error creating symlink:", err)
		}
	} else {
		if _, err := os.Lstat(publicPath); err == nil {
			err := os.Remove(publicPath)
			if err != nil {
				log.Println("error removing symlink:", err)
			}
		}
	}
}

func CreateFolder(path string, name string) {
	path = filepath.Join(path, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

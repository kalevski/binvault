package processors

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
	"binvault/pkg/services/filesystem"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CompressionFunc func(path string, target string) error

func handleFile(path string) {

	pathParts := strings.Split(path, "/")
	filename := pathParts[len(pathParts)-1]
	fileType := filesystem.GetFileType(filename)
	if fileType != models.FileType_Image {
		log.Printf("File %s is not an image, skipping", filename)
		return
	}

	file := &database.File{}
	result := database.ObtainConnection().Preload("Bucket").First(file, "name = ?", filename)

	if result.Error != nil {
		log.Printf("File %s not found in database, skipping", filename)
	}

	bucketPath := filesystem.GetBucketPath(file.Bucket.Name)

	extension := filesystem.GetFileExtension(filename)
	target := filepath.Join(bucketPath, filename)

	err := execute(extension, path, target)
	if err != nil {
		log.Printf("Error compressing file %s: %s", filename, err.Error())
	} else {
		if _, err := os.Stat(path); err == nil {
			os.Remove(path)
		}
	}
}

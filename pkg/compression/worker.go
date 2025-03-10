package compression

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
	"binvault/pkg/services/filesystem"
	"binvault/pkg/services/guetzli"
	"binvault/pkg/services/pngquant"
	"log"
	"path/filepath"
	"strings"
)

// example of path /data/temp/room-123.jpg
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

	if extension == ".jpg" {
		guetzli.Compress(path, target)
	}

	if extension == ".png" {
		pngquant.Compress(path, target)
	}

}

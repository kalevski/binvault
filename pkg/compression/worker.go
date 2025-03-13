package compression

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
	"binvault/pkg/services/filesystem"
	"binvault/pkg/services/guetzli"
	"binvault/pkg/services/pngquant"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CompressionFunc func(path string, target string) error

var compressionFuncs = map[string]CompressionFunc{
	".jpg": guetzli.Compress,
	".png": pngquant.Compress,
}

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

	if compressionFunc, ok := compressionFuncs[extension]; ok {
		err := compressionFunc(path, target)
		if err != nil {
			log.Printf("Error compressing file %s: %s", filename, err.Error())
		} else {
			log.Printf("File %s compressed successfully", filename)
		}
		err = os.Remove(path)
		if err != nil {
			log.Printf("Error deleting file %s: %s", path, err.Error())
		}
	}

}

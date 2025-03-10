package files

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
)

func FileGetMany(bucketName string, limit int, offset int) []models.File {
	db := database.ObtainConnection()
	var bucket database.Bucket
	db.First(&bucket, "name = ?", bucketName)
	var entries []database.File
	db.Limit(limit).Offset(offset).Where("bucket_id = ?", bucket.ID).Find(&entries)
	var files []models.File
	for _, entry := range entries {
		files = append(files, models.File{
			Bucket:     bucket.Name,
			Name:       entry.Name,
			Size:       entry.Size,
			Extension:  entry.Extension,
			Type:       entry.Type,
			Visibility: bucket.Visibility,
		})
	}
	return files
}

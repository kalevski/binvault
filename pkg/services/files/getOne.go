package files

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
)

func FileGetOne(bucketName string, fileName string) (*models.File, error) {
	db := database.ObtainConnection()

	var bucket database.Bucket
	result := db.First(&bucket, "name = ?", bucketName)
	if result.Error != nil {
		return nil, result.Error
	}

	var entry database.File
	result = db.First(&entry, "name = ? bucketID = ?", fileName, bucket.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.File{
		Bucket:     bucket.Name,
		Name:       entry.Name,
		Size:       entry.Size,
		Extension:  entry.Extension,
		Type:       entry.Type,
		Visibility: bucket.Visibility,
	}, nil
}

package buckets

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
)

func BucketGetOne(bucketName string) (*models.Bucket, error) {
	db := database.ObtainConnection()
	var entry database.Bucket
	tx := db.First(&entry, "name = ?", bucketName)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &models.Bucket{
		Name:       entry.Name,
		CreatedBy:  entry.CreatedBy,
		CreatedAt:  entry.CreatedAt,
		Visibility: entry.Visibility,
	}, nil
}

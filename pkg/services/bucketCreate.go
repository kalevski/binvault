package services

import (
	"binvault/pkg/clients/filesystem"
	"binvault/pkg/database"
	"binvault/pkg/models"
)

func BucketCreate(name string, visibility models.Visibility, createdBy string) (*models.Bucket, error) {
	bucket := &database.Bucket{
		Name:       name,
		Visibility: visibility,
		CreatedBy:  createdBy,
	}
	tx := database.ObtainConnection().Create(bucket)
	if tx.Error != nil {
		return nil, tx.Error
	}
	entry := &models.Bucket{
		Name:       bucket.Name,
		Visibility: bucket.Visibility,
		CreatedBy:  bucket.CreatedBy,
		CreatedAt:  bucket.CreatedAt,
	}
	filesystem.CreateFolder(filesystem.GetFolderPath(filesystem.FOLDER_BUCKETS), entry.Name)

	return entry, nil
}

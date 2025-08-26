package buckets

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
	"binvault/pkg/services/filesystem"
	"fmt"
)

func BucketCreate(name string, visibility models.Visibility, createdBy string) (*models.Bucket, error) {
	db := database.ObtainConnection()

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	var existingBucket database.Bucket
	if err := tx.Where("name = ?", name).First(&existingBucket).Error; err == nil {
		tx.Rollback()
		return nil, fmt.Errorf("bucket with name '%s' already exists", name)
	}

	bucket := &database.Bucket{
		Name:       name,
		Visibility: visibility,
		CreatedBy:  createdBy,
	}
	if err := tx.Create(bucket).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create bucket: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	entry := &models.Bucket{
		Name:       bucket.Name,
		Visibility: bucket.Visibility,
		CreatedBy:  bucket.CreatedBy,
		CreatedAt:  bucket.CreatedAt,
	}

	bucketsPath := filesystem.GetFolderPath(filesystem.FOLDER_BUCKETS)
	filesystem.CreateFolder(bucketsPath, entry.Name)
	filesystem.SetBucketVisibility(entry.Name, entry.Visibility)

	return entry, nil
}

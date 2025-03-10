package buckets

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
	"binvault/pkg/services/filesystem"
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
	bucketsPath := filesystem.GetFolderPath(filesystem.FOLDER_BUCKETS)
	filesystem.CreateFolder(bucketsPath, entry.Name)
	filesystem.SetBucketVisibility(entry.Name, entry.Visibility)

	return entry, nil
}

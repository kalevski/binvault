package buckets

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
)

func BucketGetMany(limit int, offset int) []models.Bucket {
	db := database.ObtainConnection()
	var buckets []models.Bucket = make([]models.Bucket, 0)
	var entries []database.Bucket
	db.Limit(limit).Offset(offset).Find(&entries)
	for _, entry := range entries {
		buckets = append(buckets, models.Bucket{
			Name:       entry.Name,
			CreatedBy:  entry.CreatedBy,
			CreatedAt:  entry.CreatedAt,
			Visibility: entry.Visibility,
		})
	}
	return buckets

}

package buckets

import "binvault/pkg/database"

func BucketDelete(bucketName string) error {
	db := database.ObtainConnection()
	var bucket database.Bucket
	tx := db.First(&bucket, "name = ?", bucketName)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Delete(&bucket)
	return tx.Error
}

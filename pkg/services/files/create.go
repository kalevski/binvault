package files

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
	"binvault/pkg/services/filesystem"
	"binvault/pkg/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path/filepath"
)

func FileCreate(bucketName string, filename string, content []byte, strict bool) (*models.File, error) {
	db := database.ObtainConnection()
	var bucket database.Bucket
	tx := db.First(&bucket, "name = ?", bucketName)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var count int64
	db.Model(&database.File{}).Where("bucket_id = ? AND name = ?", bucket.ID, filename).Count(&count)

	if count > 0 && strict {
		return nil, fmt.Errorf("file %s already exists in bucket %s", filename, bucketName)
	} else if count > 0 && !strict {
		hash := utils.RandomHash(5)
		basename := filename[:len(filename)-len(filepath.Ext(filename))]
		extension := filesystem.GetFileExtension(filename)
		filename = fmt.Sprintf("%s-%s.%s", basename, hash, extension)
	}

	path := filepath.Join(filesystem.GetFolderPath(filesystem.FOLDER_TEMP), filename)

	fileType := filesystem.GetFileType(filename)
	extension := filesystem.GetFileExtension(filename)

	hash := md5.Sum(content)
	hashString := hex.EncodeToString(hash[:])

	file := &database.File{
		Name:      filename,
		BucketID:  bucket.ID,
		Size:      int64(len(content)),
		Extension: extension,
		Type:      fileType,
		Hash:      hashString,
	}

	db.Create(file)

	entry := &models.File{
		Name:       file.Name,
		Bucket:     bucketName,
		Size:       file.Size,
		Type:       file.Type,
		Extension:  file.Extension,
		CreatedAt:  file.CreatedAt,
		Visibility: bucket.Visibility,
	}

	err := filesystem.SaveFile(path, content)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

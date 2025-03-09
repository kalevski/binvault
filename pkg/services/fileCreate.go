package services

import (
	"binvault/pkg/cfg"
	"binvault/pkg/clients/filesystem"
	"binvault/pkg/database"
	"binvault/pkg/models"
	"io"
	"mime/multipart"
	"path/filepath"
)

func FileCreate(bucketName string, header multipart.FileHeader, file multipart.File) (*models.File, error) {
	db := database.ObtainConnection()
	var entry database.Bucket
	tx := db.First(&entry, "name = ?", bucketName)
	if tx.Error != nil {
		return nil, tx.Error
	}

	content := make([]byte, header.Size)
	_, err := file.Read(content)
	if err != nil && err != io.EOF {
		return nil, err
	}

	path := filepath.Join(cfg.GetVars().TEMP_DIR_NAME, header.Filename)
	err = filesystem.SaveFile(path, content)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

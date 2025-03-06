package services

import (
	"binvault/pkg/models"
	"mime/multipart"
)

func FileCreate(bucketName string, header multipart.FileHeader, file multipart.File) (*models.File, error) {
	return nil, nil
}

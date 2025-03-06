package database

import (
	"binvault/pkg/models"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bucket struct {
	Base
	DeletedAt  gorm.DeletedAt `gorm:"uniqueIndex:bucket_unique"`
	Name       string         `gorm:"uniqueIndex:bucket_unique"`
	CreatedBy  string
	Visibility models.Visibility
	Files      []File `gorm:"foreignKey:BucketID"`
}
type File struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:file_unique"`
	BucketID  uint
	Bucket    Bucket `gorm:"foreignKey:BucketID"`
	Name      string `gorm:"uniqueIndex:file_unique"`
	Size      int64
	Extension string
	Type      models.FileType
	Hash      string
}

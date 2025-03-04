package database

import (
	"binvault/pkg/cfg"
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() {
	db := ObtainConnection()
	db.AutoMigrate(&File{})
	db.AutoMigrate(&Bucket{})
}

func OpenConnection() *gorm.DB {
	log.Default().Println("Opening connection to database")
	path := filepath.Join(cfg.GetPath("DATA_PATH"), cfg.GetVars().DB_NAME)
	log.Println("Database path: ", path)
	sqlite := sqlite.Open(path)
	db, err := gorm.Open(sqlite, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func ObtainConnection() *gorm.DB {
	if database == nil {
		database = OpenConnection()
	}
	return database
}

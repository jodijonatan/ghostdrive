package database

import (
	"log"

	"drive-mini/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("drive.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err) // biar kelihatan error aslinya
	}

	DB.AutoMigrate(&models.User{}, &models.FileMeta{})
}

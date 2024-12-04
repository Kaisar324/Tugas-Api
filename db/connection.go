package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
}

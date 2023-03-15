package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&Product{})
	Db = db
}

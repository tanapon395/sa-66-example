package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectDB() (*gorm.DB, error) {
	var err error
	var database *gorm.DB
	database, err = gorm.Open(sqlite.Open("sa-66.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&User{},
		&Gender{},
	)

	// SetUp Gender
	database.Where(Gender{Name: "ชาย"}).FirstOrCreate(&Gender{Name: "ชาย"})
	database.Where(Gender{Name: "หญิง"}).FirstOrCreate(&Gender{Name: "หญิง"})

	return database, nil
}

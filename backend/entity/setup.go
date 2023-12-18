package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-66.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	database.AutoMigrate(
		&User{},
		&Gender{},
	)

	db = database

	// SetUp Gender
	// Gender Data
	male := Gender{
		Name: "ชาย",
	}
	db.Model(&Gender{}).Create(&male)

	female := Gender{
		Name: "หญิง",
	}
	db.Model(&Gender{}).Create(&female)

}

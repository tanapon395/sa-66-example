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
	db.Where(Gender{Name: "ชาย"}).FirstOrCreate(&Gender{Name: "ชาย"})
	db.Where(Gender{Name: "หญิง"}).FirstOrCreate(&Gender{Name: "หญิง"})

}

package entity

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// func DB() *gorm.DB {
// 	return db
// }

func ConnectDB() (*gorm.DB, error) {
	var err error
	var database *gorm.DB
	database, err = gorm.Open(sqlite.Open("sa-66.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return database, nil
}

func SetupDatabase() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// database, err := gorm.Open(sqlite.Open("sa-66.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// Migrate the schema
	db.AutoMigrate(
		&User{},
		&Gender{},
	)

	// db = database

	// SetUp Gender
	db.Where(Gender{Name: "ชาย"}).FirstOrCreate(&Gender{Name: "ชาย"})
	db.Where(Gender{Name: "หญิง"}).FirstOrCreate(&Gender{Name: "หญิง"})

}

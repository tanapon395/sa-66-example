package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Profile   string `gorm:"type:longtext"`

	// GenderID ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id"`
}

type Gender struct {
	gorm.Model
	Name string
	User []User `gorm:"foreignKey:GenderID"`
}

package data

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId     string `gorm:"primary_key"`
	Password   string
	RegistDate string
	LastLogin  string
}

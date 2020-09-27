package data

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId     string `gorm:"primary_key"`
	UserName   string
	Password   string
	RegistDate string
	LastLogin  string
}

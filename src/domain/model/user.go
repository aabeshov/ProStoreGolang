package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	login    string `gorm:"unique"`
	password string
}

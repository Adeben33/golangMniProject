package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type Todo struct {
	gorm.Model
	Todo   string `json:"todo"`
	UserID uint
}

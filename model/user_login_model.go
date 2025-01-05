package model

import "gorm.io/gorm"

type UserLogin struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

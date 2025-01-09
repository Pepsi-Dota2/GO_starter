package entities_user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string
	Password        string
	Email           string
	Role            string
	Token           string
	PhoneNumber     string
	IsPhoneVerified bool `gorm:"default:false"`
	Address         string
}

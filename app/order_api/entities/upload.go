package entities

import "gorm.io/gorm"

type UploadFile struct {
	gorm.Model
	File string
}

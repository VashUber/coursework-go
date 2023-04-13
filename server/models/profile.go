package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model

	UserID uint
	Avatar string
	Age    string
}

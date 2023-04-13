package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string
	Email    string
	Password string
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE;"`
}

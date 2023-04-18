package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string
	Email    string `gorm:"unique"`
	Password string
	Profile  *Profile `gorm:"constraint:OnDelete:CASCADE;"`
	// Ticket   Ticket  `gorm:"constraint:OnDelete:CASCADE;"`
}

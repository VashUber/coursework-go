package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string
	Email    string `gorm:"unique"`
	Password string
	Profile  *Profile `gorm:"constraint:OnDelete:CASCADE;unique"`
	Ticket   *Ticket  `gorm:"constraint:OnDelete:CASCADE;unique"`
}

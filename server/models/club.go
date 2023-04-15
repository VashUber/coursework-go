package models

import "gorm.io/gorm"

type Club struct {
	gorm.Model

	ClubAddress ClubAddress `gorm:"constraint:OnDelete:CASCADE;"`
	Thumb       string
	Image       string
}

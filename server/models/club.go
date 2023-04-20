package models

import "gorm.io/gorm"

type Club struct {
	gorm.Model

	Thumb        string
	Image        string
	Info         string
	Name         string
	ClubAddress  *ClubAddress  `gorm:"constraint:OnDelete:CASCADE;"`
	ClubSchedule *ClubSchedule `gorm:"constraint:OnDelete:CASCADE;"`
}

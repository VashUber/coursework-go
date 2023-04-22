package models

import "gorm.io/gorm"

type Club struct {
	gorm.Model

	Thumb        string        `json:"thumb"`
	Image        string        `json:"image"`
	Info         string        `json:"info"`
	Name         string        `json:"name"`
	ClubAddress  *ClubAddress  `gorm:"constraint:OnDelete:CASCADE;" json:"address"`
	ClubSchedule *ClubSchedule `gorm:"constraint:OnDelete:CASCADE;" json:"schedule"`
}

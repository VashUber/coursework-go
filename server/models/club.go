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
	Equipment    []*Equipment  `gorm:"many2many:club_equipment;" json:"equipment"`
	Ticket       []Ticket      `gorm:"constraint:OnDelete:CASCADE;" json:"tickets"`
}

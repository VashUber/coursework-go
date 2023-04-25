package models

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model

	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Image string  `json:"image"`
	Clubs []*Club `gorm:"many2many:club_equipment;" json:"clubs"`
}

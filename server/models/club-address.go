package models

import "gorm.io/gorm"

type ClubAddress struct {
	gorm.Model

	ClubID uint   `gorm:"unique"`
	Street string `json:"street"`
	Home   string `json:"home"`
	Subway string `json:"subway"`
}

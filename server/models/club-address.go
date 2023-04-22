package models

import "gorm.io/gorm"

type ClubAddress struct {
	gorm.Model

	ClubID uint
	Street string `json:"street"`
	Home   string `json:"home"`
}

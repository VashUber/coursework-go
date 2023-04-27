package models

import (
	"github.com/jackc/pgx/pgtype"
	"gorm.io/gorm"
)

type Trainer struct {
	gorm.Model

	Name       string      `json:"name"`
	Services   pgtype.JSON `gorm:"type:json;default:'[]'" json:"services"`
	Club       []*Club     `gorm:"many2many:club_trainer;" json:"club"`
	Image      string      `json:"image"`
	Thumb      string      `json:"thumb"`
	Experience uint        `json:"experience"`
}

package models

import (
	"github.com/jackc/pgx/pgtype"
	"gorm.io/gorm"
)

type Schedule struct {
	Start string
	End   string
}

type ClubSchedule struct {
	gorm.Model

	ClubId   uint        `gorm:"unique"`
	Schedule pgtype.JSON `gorm:"type:json;default:'[]'" json:"schedule"`
}

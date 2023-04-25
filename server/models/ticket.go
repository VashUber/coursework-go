package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model

	UserID      uint `gorm:"unique"`
	ClubID      uint
	Price       uint      `json:"price"`
	StartDate   time.Time `json:"start_date"`
	ExpiredDate time.Time `json:"expired_date"`
	Time        uint      `json:"time"`
	Club        *Club     `json:"club"`
}

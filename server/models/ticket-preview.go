package models

import "gorm.io/gorm"

type TicketPreview struct {
	gorm.Model

	Price uint
	Time  uint
	Title string
	Info  string
}

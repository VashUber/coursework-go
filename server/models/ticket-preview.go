package models

import "gorm.io/gorm"

type TicketPreview struct {
	gorm.Model

	Price uint   `json:"price"`
	Time  uint   `json:"time"`
	Title string `json:"title"`
	Info  string `json:"info"`
}

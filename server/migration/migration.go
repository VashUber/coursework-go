package main

import (
	"log"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db.CreateConnection()
	RunAutoMigration()
}

func RunAutoMigration() {
	db.Database.AutoMigrate(&models.User{}, &models.Profile{}, &models.Club{}, &models.ClubAddress{}, &models.ClubSchedule{}, &models.Equipment{}, &models.TicketPreview{}, &models.Ticket{}, &models.Trainer{})
}

package main

import (
	"log"

	"github.com/VashUber/coursework-go/server/cron/tasks"
	"github.com/VashUber/coursework-go/server/db"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	db.CreateConnection()

	c := cron.New()
	tasks.InitDeleteExpiredTicketCron(c)

	c.Run()
}

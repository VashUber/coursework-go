package main

import (
	"log"
	"os"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	db.CreateConnection()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = ":3000"
	}

	app.Listen(port)
}

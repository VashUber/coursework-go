package main

import (
	"log"
	"os"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Static("/static", "./static")
	middleware.InitSessionStorage()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4020",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	router.InitAllRoutes(app)
	db.CreateConnection()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = ":3000"
	}

	app.Listen(port)
}

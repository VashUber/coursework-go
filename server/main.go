package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	api.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "unknown")

		return c.SendString(fmt.Sprintf("Hello, %s", name))
	})

	app.Listen(":3000")
}

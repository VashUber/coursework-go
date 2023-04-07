package router

import "github.com/gofiber/fiber/v2"

func DefineAuthRoutes(r fiber.Router) {
	authGroup := r.Group("/auth")

	authGroup.Post("/signin", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "signin",
		})
	})
	authGroup.Post("/signup", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "signup",
		})
	})
	authGroup.Get("/get-user-info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "get-user-info",
		})
	})
}

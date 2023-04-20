package router

import (
	"github.com/gofiber/fiber/v2"
)

func InitAllRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")

	DefineAuthRoutes(apiGroup)
	DefineProfileRoutes(apiGroup)
	DefineTicketsPreviewRoutes(apiGroup)
}

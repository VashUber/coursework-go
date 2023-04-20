package router

import (
	"github.com/VashUber/coursework-go/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func DefineTicketsPreviewRoutes(r fiber.Router) {
	ticketsPreviewGroup := r.Group("/tickets-preview")

	ticketsPreviewGroup.Get("/get-all", controllers.GetTicketsPreview)
}

package router

import (
	"github.com/VashUber/coursework-go/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func DefineTicketRoutes(r fiber.Router) {
	ticketGroup := r.Group("/ticket")

	ticketGroup.Post("/buy-ticket", controllers.BuyTicket)
}

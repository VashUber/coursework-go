package controllers

import (
	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

func GetTicketsPreview(c *fiber.Ctx) error {
	var tickets []models.TicketPreview
	db.Database.Find(&tickets)

	return c.JSON(fiber.Map{
		"tickets": tickets,
	})
}

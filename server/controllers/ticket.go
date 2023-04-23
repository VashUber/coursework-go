package controllers

import (
	"time"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

func BuyTicket(c *fiber.Ctx) error {
	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	id, ok := sess.Get("id").(uint)
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	type BuyTicketBody struct {
		Ticket uint `json:"ticket"`
		Club   uint `json:"club"`
	}
	var body BuyTicketBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	start := time.Now()
	var expired time.Time

	if body.Ticket == 12 {
		expired = start.AddDate(1, 0, 0)
	} else {
		expired = start.AddDate(0, int(body.Ticket), 0)
	}

	ticket := &models.Ticket{
		ClubID:      body.Club,
		UserID:      id,
		StartDate:   start,
		ExpiredDate: expired,
	}

	db.Database.Create(ticket)

	return c.SendStatus(fiber.StatusOK)
}

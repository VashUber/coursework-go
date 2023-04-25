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

	var user models.User
	db.Database.Preload("Ticket").Where("id = ?", id).First(&user)

	if user.Ticket != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Your already have a ticket")
	}

	type BuyTicketBody struct {
		Ticket uint `json:"ticket"`
		Club   uint `json:"club"`
	}
	var body BuyTicketBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var ticketPreview models.TicketPreview
	err = db.Database.Where("time = ?", body.Ticket).First(&ticketPreview).Error
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
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
		Price:       ticketPreview.Price,
		Time:        ticketPreview.Time,
	}

	db.Database.Create(ticket)

	return c.SendStatus(fiber.StatusOK)
}

func GetUserTicket(c *fiber.Ctx) error {
	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	id := sess.Get("id")

	var user models.User
	db.Database.Preload("Ticket").Where("id = ?", id).First(&user)

	if user.Ticket == nil {
		return c.JSON(fiber.Map{
			"ticket": nil,
		})
	}

	var ticket models.Ticket
	db.Database.Preload("Club").Preload("Club.ClubAddress", db.Database.Select("ID")).Where("user_id = ?", id).First(&ticket)

	type Address struct {
		Street string `json:"street"`
		Home   string `json:"home"`
	}
	type Result struct {
		ID          uint      `json:"id"`
		Price       uint      `json:"price"`
		StartDate   time.Time `json:"start_date"`
		ExpiredDate time.Time `json:"expired_date"`
		Time        uint      `json:"time"`
		Address     Address   `json:"address"`
	}

	return c.JSON(fiber.Map{
		"ticket": Result{
			ID:          ticket.ID,
			Price:       ticket.Price,
			StartDate:   ticket.StartDate,
			ExpiredDate: ticket.ExpiredDate,
			Time:        ticket.Time,
			Address: Address{
				Home:   ticket.Club.ClubAddress.Home,
				Street: ticket.Club.ClubAddress.Street,
			},
		},
	})
}

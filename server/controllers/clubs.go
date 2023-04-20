package controllers

import (
	"strconv"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

const perPage = 5

func GetClubsPerPage(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if page <= 0 {
		page = 1
	}

	var clubs []models.Club
	offset := (page - 1) * perPage
	db.Database.Preload("ClubAddress").Preload("ClubSchedule").Offset(offset).Limit(perPage).Find(&clubs)

	return c.JSON(fiber.Map{
		"clubs": clubs,
	})
}

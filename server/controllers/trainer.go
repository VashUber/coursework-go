package controllers

import (
	"strconv"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

func GetTrainersPerPage(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if page <= 0 {
		page = 1
	}

	var trainers []models.Trainer
	var count int64
	offset := (page - 1) * perPage
	db.Database.Offset(offset).Limit(perPage).Find(&trainers)
	db.Database.Table("trainers").Count(&count)

	pages := (count + perPage - 1) / perPage

	return c.JSON(fiber.Map{
		"trainers": trainers,
		"pages":    pages,
	})
}

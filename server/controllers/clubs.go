package controllers

import (
	"strconv"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

const perPage = 6

func GetClubsPerPage(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if page <= 0 {
		page = 1
	}

	var clubs []models.Club
	var count int64
	offset := (page - 1) * perPage
	db.Database.Preload("ClubAddress").Preload("ClubSchedule").Offset(offset).Limit(perPage).Find(&clubs)
	db.Database.Table("clubs").Count(&count)

	pages := (count + perPage - 1) / perPage

	return c.JSON(fiber.Map{
		"clubs": clubs,
		"pages": pages,
	})
}

func GetClub(c *fiber.Ctx) error {
	clubID, err := strconv.Atoi(c.Query("id", "1"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var club models.Club
	err = db.Database.Preload("ClubAddress").Preload("ClubSchedule").Where("id = ?", clubID).Find(&club).Error

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"club": club,
	})
}

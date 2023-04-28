package controllers

import (
	"strconv"

	"github.com/VashUber/coursework-go/server/db"
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

	search := c.Query("search", "")
	experience, err := strconv.Atoi(c.Query("experience", "0"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	type Result struct {
		Id         uint   `json:"id"`
		Name       string `json:"name"`
		Thumb      string `json:"thumb"`
		Experience uint   `json:"experience"`
	}

	var trainers []Result
	var count int64
	offset := (page - 1) * perPage

	db.Database.Table("trainers").Where("name LIKE ?", search+"%").Where("experience > ?", experience).Offset(offset).Limit(perPage).Find(&trainers)
	db.Database.Table("trainers").Where("name LIKE ?", search+"%").Where("experience > ?", experience).Count(&count)

	pages := (count + perPage - 1) / perPage

	return c.JSON(fiber.Map{
		"trainers": trainers,
		"pages":    pages,
	})
}

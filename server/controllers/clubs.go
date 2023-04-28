package controllers

import (
	"strconv"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

func GetClubsPerPage(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if page <= 0 {
		page = 1
	}

	search := c.Query("search", "")
	subway := c.Query("subway", "")

	var clubs []models.Club
	var count int64
	offset := (page - 1) * perPage

	db.Database.Table("clubs").Joins("ClubAddress").Where("name LIKE ?", search+"%").Where("subway LIKE ?", subway+"%").Offset(offset).Limit(perPage).Find(&clubs)
	db.Database.Table("clubs").Joins("ClubAddress").Where("name LIKE ?", search+"%").Where("subway LIKE ?", subway+"%").Find(&[]models.Club{}).Count(&count)

	pages := (count + perPage - 1) / perPage

	return c.JSON(fiber.Map{
		"clubs": clubs,
		"pages": pages,
	})
}

func GetAllClubsLightWeight(c *fiber.Ctx) error {
	type Result struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	}
	var res []Result
	err := db.Database.Model(&models.Club{}).Select("name", "id").Find(&res).Error
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"clubs": res,
	})
}

func GetClub(c *fiber.Ctx) error {
	clubID, err := strconv.Atoi(c.Query("id", "1"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var club models.Club
	err = db.Database.Preload("ClubAddress").Preload("ClubSchedule").Preload("Equipment").Where("id = ?", clubID).Find(&club).Error

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"club": club,
	})
}

func GetAllSubwayStations(c *fiber.Ctx) error {
	type Result struct {
		ID     uint   `json:"id"`
		Subway string `json:"subway"`
	}

	var subwayStations []Result
	err := db.Database.Model(&models.ClubAddress{}).Distinct("subway").Find(&subwayStations).Error

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"subways": subwayStations,
	})
}

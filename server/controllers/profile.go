package controllers

import (
	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

func GetProfileInfo(c *fiber.Ctx) error {
	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		panic(err)
	}

	id := sess.Get("id")

	var profile models.Profile
	err = db.Database.Where("id = ?", id).Find(&profile).Error
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"profile": fiber.Map{
			"avatar":   profile.Avatar.String,
			"birthday": profile.Birthday,
		},
	})
}

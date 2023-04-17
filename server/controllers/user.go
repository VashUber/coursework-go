package controllers

import (
	"fmt"
	"time"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/VashUber/coursework-go/server/utils"
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

func UpdateUserInfo(c *fiber.Ctx) error {
	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		panic(err)
	}

	id := sess.Get("id")

	fmt.Println(id)

	type UpdateUserInfoBody struct {
		Email    string    `json:"email" validate:"required,email"`
		Password string    `json:"password" validate:"required,min=6"`
		Name     string    `json:"name" validate:"required"`
		Birthday time.Time `json:"birthday" validate:"required"`
	}
	var body UpdateUserInfoBody
	c.BodyParser(&body)

	err = utils.ValidateStructInRequest(body, c)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

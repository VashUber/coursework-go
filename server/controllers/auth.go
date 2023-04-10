package controllers

import (
	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	type SignupBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"string"`
		Age      uint8  `json:"age"`
	}

	body := SignupBody{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db.Database.Create(&models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Age:      body.Age,
	})

	return c.SendStatus(fiber.StatusAccepted)
}

func Signin(c *fiber.Ctx) error {
	type SigninBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		panic(err)
	}

	body := SigninBody{}
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	if body.Password != "123456" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	sess.Set("id", body.Email)
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUserInfo(c *fiber.Ctx) error {
	// TODO
	return c.JSON(fiber.Map{
		"message": "get-user-info",
	})
}

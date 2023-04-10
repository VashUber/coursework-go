package controllers

import (
	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser() {
	db.Database.Create(&models.User{
		Name:     "string",
		Email:    "string",
		Password: "string",
	})
}

func Signup(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "signup",
	})
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

	sess.Set("email", body.Email)
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUserInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "get-user-info",
	})
}

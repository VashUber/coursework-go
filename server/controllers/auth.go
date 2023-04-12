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
		Name     string `json:"name"`
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
	})

	return c.SendStatus(fiber.StatusAccepted)
}

func Signin(c *fiber.Ctx) error {
	type SigninBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	body := SigninBody{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	var user models.User

	db.Database.Where("email = ?", body.Email).Find(&user)

	if user.Password != body.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		panic(err)
	}

	sess.Set("id", user.ID)
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUserInfo(c *fiber.Ctx) error {
	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		panic(err)
	}

	id := sess.Get("id")
	if id == nil {
		return c.JSON(fiber.Map{
			"user": nil,
		})
	}

	var user models.User
	db.Database.Where("id = ?", id).Find(&user)

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"email": user.Email,
			"name":  user.Name,
		},
	})
}

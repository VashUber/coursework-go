package controllers

import (
	"time"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/VashUber/coursework-go/server/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	type SignupBody struct {
		Email    string    `json:"email" validate:"required,email"`
		Password string    `json:"password" validate:"required,min=6"`
		Name     string    `json:"name" validate:"required"`
		Birthday time.Time `json:"birthday" validate:"required"`
	}

	var body SignupBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = utils.ValidateStructInRequest(body, c)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	profile := models.Profile{
		Birthday: body.Birthday,
	}

	db.Database.Create(&models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(password),
		Profile:  profile,
	})

	return c.SendStatus(fiber.StatusAccepted)
}

func Signin(c *fiber.Ctx) error {
	type SigninBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body SigninBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	var user models.User
	db.Database.Where("email = ?", body.Email).Find(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
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

func Signout(c *fiber.Ctx) error {
	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		panic(err)
	}

	sess.Destroy()
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

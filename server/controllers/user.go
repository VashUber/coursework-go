package controllers

import (
	"database/sql"
	"time"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/VashUber/coursework-go/server/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	id := sess.Get("id")

	type UpdateUserInfoBody struct {
		Email    string    `json:"email"`
		Password string    `json:"password"`
		Name     string    `json:"name"`
		Birthday time.Time `json:"birthday"`
	}
	var body UpdateUserInfoBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := utils.ValidateStructInRequest(body, c); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var user models.User
	if err := db.Database.Preload("Profile").Where("id = ?", id).First(&user).Error; err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if len(body.Password) != 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		user.Password = string(hashedPassword)
	}

	user.Email = body.Email
	user.Name = body.Name

	user.Profile.Birthday = body.Birthday

	tx := db.Database.Begin()

	if err := tx.Model(&user).Updates(user).Error; err != nil {
		tx.Rollback()
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err := tx.Model(&user.Profile).Updates(user.Profile).Error; err != nil {
		tx.Rollback()
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tx.Commit()

	return c.SendStatus(fiber.StatusOK)
}

func UploadAvatar(c *fiber.Ctx) error {
	sess, err := middleware.SessionStorage.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	id := sess.Get("id")

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	c.SaveFile(file, "./static/avatars/"+file.Filename)

	var user models.User
	if err := db.Database.Preload("Profile").Where("id = ?", id).First(&user).Error; err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Profile.Avatar = sql.NullString{
		String: "http://localhost:4000/static/avatars/" + file.Filename,
		Valid:  true,
	}

	err = db.Database.Model(&user.Profile).Updates(user.Profile).Error
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

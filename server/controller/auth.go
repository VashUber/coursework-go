package controller

import (
	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
)

func CreateUser() {
	db.Database.Create(&models.User{
		Name:     "string",
		Email:    "string",
		Password: "string",
	})
}

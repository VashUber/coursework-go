package router

import (
	"github.com/VashUber/coursework-go/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func DefineTrainersRoutes(r fiber.Router) {
	trainersGroup := r.Group("/trainers")

	trainersGroup.Get("/get-trainers", controllers.GetTrainersPerPage)
}

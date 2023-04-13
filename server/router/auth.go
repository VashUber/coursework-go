package router

import (
	"github.com/VashUber/coursework-go/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func DefineAuthRoutes(r fiber.Router) {
	authGroup := r.Group("/auth")

	authGroup.Post("/signup", controllers.Signup)
	authGroup.Post("/signin", controllers.Signin)
	authGroup.Get("/signout", controllers.Signout)
	authGroup.Get("/get-user-info", controllers.GetUserInfo)
}

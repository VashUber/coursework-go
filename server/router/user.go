package router

import (
	"github.com/VashUber/coursework-go/server/controllers"
	"github.com/VashUber/coursework-go/server/middleware"
	"github.com/gofiber/fiber/v2"
)

func DefineProfileRoutes(r fiber.Router) {
	userGroup := r.Group("/user")

	userGroup.Get("/get-profile-info", middleware.SessionChecker, controllers.GetProfileInfo)
	userGroup.Post("/update-profile-info", middleware.SessionChecker, controllers.UpdateUserInfo)
	userGroup.Post("/upload-avatar", middleware.SessionChecker, controllers.UploadAvatar)
}

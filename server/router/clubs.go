package router

import (
	"github.com/VashUber/coursework-go/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func DefineClubsRoutes(r fiber.Router) {
	clubsGroup := r.Group("/clubs")

	clubsGroup.Get("/get-clubs", controllers.GetClubsPerPage)
	clubsGroup.Get("/get-all-clubs", controllers.GetAllClubsLightWeight)
	clubsGroup.Get("/get-club", controllers.GetClub)
	clubsGroup.Get("/get-subway-stations", controllers.GetAllSubwayStations)
}

package routes

import (
	"rghdrizzle/gameLog/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App){
	app.Get("/Recommended",controllers.GetRecommendedGames)
	app.Get("/Search",controllers.SearchGameDatabaseBasedOnText)
}
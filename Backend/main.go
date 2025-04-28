package main

import(
	"github.com/gofiber/fiber/v2"
	"rghdrizzle/gameLog/routes"
)

func main(){
	app := fiber.New()
	routes.UserRoutes(app)
	app.Listen(":3333")
}
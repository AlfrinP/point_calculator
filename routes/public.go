package routes

import (
	"github.com/AlfrinP/point_calculator/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/signup", controllers.SignUp)
	auth.Post("/signin", controllers.SignIn)
}

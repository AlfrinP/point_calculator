package routes

import (
	"github.com/AlfrinP/point_calculator/controllers"
	"github.com/AlfrinP/point_calculator/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupProtectedRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/securitycheck", middleware.DeserializeUser, controllers.SecurityCheck)
}

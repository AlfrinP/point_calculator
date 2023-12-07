package routes

import (
	"github.com/Levantate-Labs/sainterview-backend/auth-service/controllers"
	"github.com/Levantate-Labs/sainterview-backend/auth-service/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupProtectedRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/securitycheck", middleware.DeserializeUser, controllers.SecurityCheck)
}

package routes

import (
	"github.com/Levantate-Labs/sainterview-backend/auth-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/healthcheck", controllers.HealthCheck)
}

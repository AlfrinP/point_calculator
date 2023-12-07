package routes

import (
<<<<<<< HEAD
	"github.com/Levantate-Labs/sainterview-backend/auth-service/controllers"
=======
	"github.com/AlfrinP/point_calculator/controllers"
>>>>>>> 52a2cfba8417f30f47f3a85feb3c92850e82f352
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/healthcheck", controllers.HealthCheck)
}

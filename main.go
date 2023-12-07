package main

import (
	"log"

<<<<<<< HEAD
	"github.com/Levantate-Labs/sainterview-backend/auth-service/config"
	"github.com/Levantate-Labs/sainterview-backend/auth-service/routes"
	"github.com/Levantate-Labs/sainterview-backend/auth-service/storage"
=======
	"github.com/AlfrinP/point_calculator/config"
	"github.com/AlfrinP/point_calculator/routes"
	"github.com/AlfrinP/point_calculator/storage"
>>>>>>> 52a2cfba8417f30f47f3a85feb3c92850e82f352
	"github.com/gofiber/fiber/v2"
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	storage.ConnectDB(&config)
}
func main() {

	app := fiber.New()
	routes.SetupRoutes(app)
	routes.SetupPublicRoutes(app)
	routes.SetupProtectedRoutes(app)

	app.Listen(":3000")
}

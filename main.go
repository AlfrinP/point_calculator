package main

import (
	"log"

	"github.com/Levantate-Labs/sainterview-backend/auth-service/config"
	"github.com/Levantate-Labs/sainterview-backend/auth-service/routes"
	"github.com/Levantate-Labs/sainterview-backend/auth-service/storage"
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

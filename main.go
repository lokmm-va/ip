package main

import (
	"plant-api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	api := app.Group("/api/v1")
	{
		api.Get("/plants", handlers.GetPlants)
		api.Get("/plants/:id", handlers.GetPlant)
		api.Post("/plants", handlers.CreatePlant)
		api.Put("/plants/:id", handlers.UpdatePlant)
		api.Delete("/plants/:id", handlers.DeletePlant)
	}

	app.Listen(":3000")
}

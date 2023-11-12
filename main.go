package main

import (
	"log"

	"github.com/artisademi/gorm-fiber-api/database"
	"github.com/artisademi/gorm-fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to this awesome api")
}

func setUpRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)

	// User Endpoint
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

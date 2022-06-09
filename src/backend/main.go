package main

import (
	"log"

	"github.com/NQuirmbach/todo-app/database"
	"github.com/NQuirmbach/todo-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/joho/godotenv/autoload" // load .env file
)

func main() {
	app := fiber.New()

	// Init database
	db := database.Init()

	app.Use(cors.New())
	app.Use(recover.New())

	// Register routes
	routes.RegisterApi(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	// Startup app
	log.Fatal(app.Listen(":3000"))
}

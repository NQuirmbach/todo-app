package main

import (
	"log"

	"github.com/NQuirmbach/todo-app/database"
	"github.com/NQuirmbach/todo-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/joho/godotenv/autoload" // load .env file
)

func main() {
	app := fiber.New()

	app.Use(recover.New())

	// Init database
	db := database.Init()

	// Register routes
	routes.RegisterApi(app, db)

	// Startup app
	log.Fatal(app.Listen(":5000"))
}

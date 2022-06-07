package routes

import (
	"log"

	"github.com/NQuirmbach/todo-app/api/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApi(app *fiber.App, db *gorm.DB) {
	router := app.Group("/api")

	log.Println("Register api routes")

	controllers.Todos(router, db)
}

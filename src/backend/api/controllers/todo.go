package controllers

import (
	"log"

	"github.com/NQuirmbach/todo-app/api/dto"
	"github.com/NQuirmbach/todo-app/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Todos(r fiber.Router, db *gorm.DB) {
	r.Get("/todos", func(c *fiber.Ctx) error {
		log.Println("Get todos")

		var todos []entity.Todo
		db.Find(&todos)

		result := []dto.TodoDto{}

		for _, el := range todos {
			dto := dto.NewTodoDto(el)
			result = append(result, dto)
		}

		return c.JSON(result)
	})

	r.Post("/todos", func(c *fiber.Ctx) error {
		log.Println("Creating new todo")

		t := new(dto.TodoDto)
		if err := c.BodyParser(t); err != nil {
			log.Fatal(err)
			return fiber.ErrBadRequest
		}

		entity := &entity.Todo{
			Text:    t.Text,
			DueDate: t.DueDate,
		}

		db.Create(entity)

		dto := dto.NewTodoDto(*entity)

		return c.Status(201).JSON(dto)
	})

	r.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var t entity.Todo
		result := db.First(&t, id)

		if result.Error != nil {
			return result.Error
		}

		db.Delete(&t)

		return c.SendStatus(200)
	})
}

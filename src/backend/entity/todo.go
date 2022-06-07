package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Text    string
	IsDone  bool
	DueDate time.Time
}

func NewTodo(text string, dueDate time.Time) *Todo {
	t := &Todo{
		Text:    text,
		IsDone:  false,
		DueDate: dueDate,
	}

	return t
}

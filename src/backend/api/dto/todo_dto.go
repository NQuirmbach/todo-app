package dto

import (
	"time"

	"github.com/NQuirmbach/todo-app/entity"
)

type TodoDto struct {
	ID      uint      `json:"id"`
	Text    string    `json:"text"`
	IsDone  bool      `json:"isDone"`
	DueDate time.Time `json:"dueDate"`
}

func NewTodoDto(t entity.Todo) TodoDto {
	return TodoDto{
		ID:      t.ID,
		Text:    t.Text,
		IsDone:  t.IsDone,
		DueDate: t.DueDate,
	}
}

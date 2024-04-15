package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	TargetDate   time.Time `json:"target_date"`
	DoneDate     time.Time `json:"done_date"`
	Is_completed bool      `json:"is_completed"`
}

func NewTodo(id, name string, target_date, done_date time.Time, is_completed bool) *Todo {
	return &Todo{
		ID:           NewTodoID(),
		Name:         name,
		TargetDate:   target_date,
		DoneDate:     done_date,
		Is_completed: is_completed,
	}
}

func NewTodoID() string {
	return uuid.NewString()
}

func IsValidTodoID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

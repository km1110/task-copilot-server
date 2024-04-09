package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID           string
	Name         string
	TargetDate   time.Time
	DoneDate     time.Time
	Is_completed bool
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

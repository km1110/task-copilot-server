package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID         string
	Name       string
	TargetDate time.Time
	DoneDate   time.Time
	Status     bool
}

func NewTodo(id, name string, targetdate, donedate time.Time, status bool) *Todo {
	return &Todo{
		ID:         NewTodoID(),
		Name:       name,
		TargetDate: targetdate,
		DoneDate:   donedate,
		Status:     status,
	}
}

func NewTodoID() string {
	return uuid.NewString()
}

func IsValidTodoID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

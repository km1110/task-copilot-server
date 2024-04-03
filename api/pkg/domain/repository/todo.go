package repository

import (
	"context"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
)

type Todo interface {
	GetTodos(ctx context.Context, userID string) ([]*model.Todo, error)
	CreateTodo(ctx context.Context, userID string, todo *model.Todo) (*model.Todo, error)
	UpdateTodo(ctx context.Context, id string, t *model.Todo) (*model.Todo, error)
	DeleteTodo(ctx context.Context, id string) error
}

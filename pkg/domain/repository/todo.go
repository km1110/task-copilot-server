package repository

import (
	"context"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
)

type Todo interface {
	GetTodos(ctx context.Context, userID string) ([]*model.Todo, error)
}

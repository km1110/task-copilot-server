package usecase

import (
	"context"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type getTodo struct {
	repoTodo repository.Todo
}

func NewGetTodo(rt repository.Todo) *getTodo {
	return &getTodo{repoTodo: rt}
}

func (gt *getTodo) Exec(ctx context.Context, userID string) ([]*model.Todo, error) {
	if !model.IsValidUserID(userID) {
		return nil, xerrors.Errorf("!model.IsValidUserID: userID is invalid")
	}

	t, err := gt.repoTodo.GetTodos(ctx, userID)
	if err != nil {
		return nil, xerrors.Errorf("gt.repoTodo.GetTodos: %v", err)
	}

	return t, nil
}

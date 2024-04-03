package usecase

import (
	"context"
	"time"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type createTodo struct {
	repoTodo repository.Todo
}

func NewCreateTodo(rt repository.Todo) *createTodo {
	return &createTodo{repoTodo: rt}
}

func (ct *createTodo) Exec(ctx context.Context, userID, name string, target_date, done_date time.Time, status bool) (*model.Todo, error) {
	t := model.NewTodo(model.NewTodoID(), name, target_date, done_date, status)

	newTodo, err := ct.repoTodo.CreateTodo(ctx, userID, t)
	if err != nil {
		return nil, xerrors.Errorf("ct.repoTodo.CreateTodo: %v", err)
	}

	return newTodo, nil
}

package usecase

import (
	"context"
	"time"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type updateTodo struct {
	repoTodo repository.Todo
}

func NewUpdateTodo(rt repository.Todo) *updateTodo {
	return &updateTodo{repoTodo: rt}
}

func (ut *updateTodo) Exec(ctx context.Context, id, name string, target_date, done_date time.Time, is_completed bool) (*model.Todo, error) {
	if !model.IsValidTodoID(id) {
		return nil, xerrors.Errorf("!model.IsValidTodoID: todoID is invalid")
	}

	t := model.NewTodo(id, name, target_date, done_date, is_completed)

	res, err := ut.repoTodo.UpdateTodo(ctx, id, t)
	if err != nil {
		return nil, xerrors.Errorf("ut.repoTodo.UpdateTodo: %v", err)
	}

	return res, nil
}

package usecase

import (
	"context"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type deleteTodo struct {
	repoTodo repository.Todo
}

func NewDeleteTodo(rt repository.Todo) *deleteTodo {
	return &deleteTodo{repoTodo: rt}
}

func (dt *deleteTodo) Exec(ctx context.Context, id string) error {
	if !model.IsValidTodoID(id) {
		return xerrors.Errorf("!model.IsValidTodoID: todoID in invalid")
	}

	err := dt.repoTodo.DeleteTodo(ctx, id)
	if err != nil {
		return xerrors.Errorf("dt.repoTodo.DeleteTodo: %v", err)
	}

	return nil
}

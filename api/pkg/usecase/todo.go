package usecase

import (
	"context"

	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/validation"
)

type ITodoUsecase interface {
	GetAllTodos(ctx context.Context, userID string) ([]model.Todo, error)
	GetTodobyId(ctx context.Context, userID, todoID string) (model.Todo, error)
	CreateTodo(ctx context.Context, userID string, t model.Todo) (model.Todo, error)
	UpdateTodo(ctx context.Context, todoID string, t model.Todo) (model.Todo, error)
	DeleteTodo(ctx context.Context, todoID string) error
}

type todoUsecase struct {
	tr repository.ITodoRepository
	tv validation.ITodoValidator
}

func NewTodoUsecase(tr repository.ITodoRepository, tv validation.ITodoValidator) ITodoUsecase {
	return &todoUsecase{tr, tv}
}

func (tu todoUsecase) GetAllTodos(ctx context.Context, userID string) ([]model.Todo, error) {
	todos := []model.Todo{}
	if err := tu.tr.GetAllTodos(ctx, userID, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (tu todoUsecase) GetTodobyId(ctx context.Context, userID, todoID string) (model.Todo, error) {
	todo := model.Todo{}
	if err := tu.tr.GetTodobyId(ctx, userID, todoID, &todo); err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func (tu todoUsecase) CreateTodo(ctx context.Context, userID string, todo model.Todo) (model.Todo, error) {
	if err := tu.tv.TodoValidate(todo); err != nil {
		return model.Todo{}, nil
	}

	if err := tu.tr.CreateTodo(ctx, userID, &todo); err != nil {
		return model.Todo{}, nil
	}

	return todo, nil
}

func (tu todoUsecase) UpdateTodo(ctx context.Context, todoID string, todo model.Todo) (model.Todo, error) {
	if err := tu.tr.UpdateTodo(ctx, todoID, &todo); err != nil {
		return model.Todo{}, nil
	}

	return todo, nil
}

func (tu todoUsecase) DeleteTodo(ctx context.Context, todoID string) error {
	if err := tu.tr.DeleteTodo(ctx, todoID); err != nil {
		return nil
	}

	return nil
}

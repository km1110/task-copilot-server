package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
)

type ITodoValidator interface {
	TodoValidate(todo model.Todo) error
}

type todoValidator struct{}

func NewTodoValidator() ITodoValidator {
	return &todoValidator{}
}

func (tv *todoValidator) TodoValidate(todo model.Todo) error {
	return validation.ValidateStruct(&todo,
		validation.Field(
			&todo.Name,
			validation.Required.Error("name is requierd"),
			validation.RuneLength(1, 255).Error("limited max 255 char"),
		),
	)
}

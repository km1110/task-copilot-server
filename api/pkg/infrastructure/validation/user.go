package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Name,
			validation.Required.Error("name is requierd"),
			validation.RuneLength(1, 255).Error("limited max 255 char"),
		),
	)
}

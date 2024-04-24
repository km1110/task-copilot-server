package usecase

import (
	"context"

	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/validation"
)

type IUserUsecase interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, uid string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, uid string, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, uid string) error
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validation.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validation.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu userUsecase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	if err := uu.ur.GetAllUsers(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (uu userUsecase) GetUserById(ctx context.Context, uid string) (model.User, error) {
	user := model.User{}
	if err := uu.ur.GetUserById(ctx, uid, &user); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (uu userUsecase) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.User{}, err
	}

	if err := uu.ur.CreateUser(ctx, &user); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (uu userUsecase) UpdateUser(ctx context.Context, uid string, user model.User) (model.User, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.User{}, err
	}

	if err := uu.ur.UpdateUser(ctx, uid, &user); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (uu userUsecase) DeleteUser(ctx context.Context, uid string) error {
	if err := uu.ur.DeleteUser(ctx, uid); err != nil {
		return err
	}

	return nil
}

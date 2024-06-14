package usecase

import (
	"context"

	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/cache"
)

type IAuthUsecase interface {
	Login(ctx context.Context, uid string) (string, error)
	Register(ctx context.Context, user model.User) (string, error)
}

type authUsecase struct {
	ar repository.IAuthRepository
	uc cache.IUserCache
}

func NewAuthUsecase(ar repository.IAuthRepository, uc cache.IUserCache) IAuthUsecase {
	return &authUsecase{ar, uc}
}

func (au *authUsecase) Login(ctx context.Context, uid string) (string, error) {
	user := model.User{}
	if err := au.ar.Login(ctx, uid, &user); err != nil {
		return "", err
	}

	au.uc.Set("role", user.Role.Name)

	return "Login Success", nil
}

func (au *authUsecase) Register(ctx context.Context, user model.User) (string, error) {
	if err := au.ar.Register(ctx, user.FirebaseUID, &user); err != nil {
		return "", err
	}
	return "Register Success", nil
}

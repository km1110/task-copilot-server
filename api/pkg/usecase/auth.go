package usecase

import (
	"context"
	"log"
	"time"

	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/cache"
)

type IAuthUsecase interface {
	Login(ctx context.Context, uid string) (string, error)
	Register(ctx context.Context, user model.User) (string, error)
	Logout(ctx context.Context, uid string) (string, error)
}

type authUsecase struct {
	ar repository.IAuthRepository
	rr repository.IRedisRepository
	uc cache.IUserCache
}

func NewAuthUsecase(ar repository.IAuthRepository, rr repository.IRedisRepository, uc cache.IUserCache) IAuthUsecase {
	return &authUsecase{ar, rr, uc}
}

func (au *authUsecase) Login(ctx context.Context, uid string) (string, error) {
	user := model.User{}
	if err := au.ar.Login(ctx, uid, &user); err != nil {
		return "", err
	}
	key := "user:" + uid
	au.uc.Set(key, user.Role.Name)

	go func() {
		time.Sleep(30 * time.Second)
		role := au.uc.Get(key)
		if role == "" {
			log.Println("role is not found")
			return
		}
		au.rr.Set(key, role)
	}()

	return "Login Success", nil
}

func (au *authUsecase) Register(ctx context.Context, user model.User) (string, error) {
	if err := au.ar.Register(ctx, user.FirebaseUID, &user); err != nil {
		return "", err
	}
	return "Register Success", nil
}

func (au *authUsecase) Logout(ctx context.Context, uid string) (string, error) {
	key := "user:" + uid
	au.uc.Delete(key)
	au.rr.Delete(key)
	return "Logout Success", nil
}

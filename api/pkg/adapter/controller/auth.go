package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/firebase"
	"github.com/km1110/task-copilot-server/pkg/usecase"
)

type IAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	au          usecase.IAuthUsecase
	firebaseApp firebase.IFirebaseApp
}

func NewAuthController(au usecase.IAuthUsecase, firebaseApp firebase.IFirebaseApp) IAuthController {
	return &authController{au: au, firebaseApp: firebaseApp}
}

func (ac *authController) Login(c *gin.Context) {
	token, err := validateAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := ac.firebaseApp.VerifyIDToken(c, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	msg, err := ac.au.Login(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, msg)
}

func (ac *authController) Register(c *gin.Context) {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := validateAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := ac.firebaseApp.VerifyIDToken(c, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	user.FirebaseUID = uid

	msg, err := ac.au.Register(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, msg)
}

func validateAuthHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header required")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return "", errors.New("Invalid token format")
	}

	return token, nil
}

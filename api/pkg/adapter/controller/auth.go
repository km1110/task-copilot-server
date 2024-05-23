package controller

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/usecase"
	"google.golang.org/api/option"
)

type IAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	au usecase.IAuthUsecase
}

func NewAuthController(au usecase.IAuthUsecase) IAuthController {
	return &authController{au: au}
}

func (ac *authController) Login(c *gin.Context) {
	opt := option.WithCredentialsFile("./service_accout_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header required"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
		return
	}

	verifiedToken, err := firebaseAuth.VerifyIDToken(c, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	firebaseUID := verifiedToken.UID
	uid := firebaseUID

	msg, err := ac.au.Login(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, msg)
}

func (ac *authController) Register(c *gin.Context) {
	firebaseUID, _ := c.Get("firebaseUID")
	uid := firebaseUID.(string)

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/usecase"
)

type IUserController interface {
	GetAllUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	uc usecase.IUserUsecase
}

func NewUserController(uc usecase.IUserUsecase) IUserController {
	return &userController{uc}
}

func (uc *userController) GetAllUsers(c *gin.Context) {
	res, err := uc.uc.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *userController) GetUserById(c *gin.Context) {
	firebaseUID, _ := c.Get("firebaseUID")
	uid := firebaseUID.(string)

	res, err := uc.uc.GetUserById(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *userController) CreateUser(c *gin.Context) {
	firebaseUID, _ := c.Get("firebaseUID")
	uid := firebaseUID.(string)

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.FirebaseUID = uid

	res, err := uc.uc.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	firebaseUID, _ := c.Get("firebaseUID")
	uid := firebaseUID.(string)

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := uc.uc.UpdateUser(c, uid, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteUser(c *gin.Context) {
	firebaseUID, _ := c.Get("firebaseUID")
	uid := firebaseUID.(string)

	err := uc.uc.DeleteUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

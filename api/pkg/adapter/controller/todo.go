package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/usecase"
)

type ITodoController interface {
	GetAllTodos(c *gin.Context)
	GetTodobyId(c *gin.Context)
	CreateTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type todoController struct {
	tu usecase.ITodoUsecase
}

func NewTodoController(tu usecase.ITodoUsecase) ITodoController {
	return &todoController{tu}
}

func (tc *todoController) GetAllTodos(c *gin.Context) {
	//TODO: get user_id
	userID := "b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1"

	res, err := tc.tu.GetAllTodos(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc todoController) GetTodobyId(c *gin.Context) {
	//TODO: get user_id
	userID := "b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1"

	id := c.Param("todo_id")

	res, err := tc.tu.GetTodobyId(c, userID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc todoController) CreateTodo(c *gin.Context) {
	//TODO: get user_id
	userID := "b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1"

	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//TODO: generate new todo_id

	res, err := tc.tu.CreateTodo(c, userID, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (tc *todoController) UpdateTodo(c *gin.Context) {
	id := c.Param("todo_id")

	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := tc.tu.UpdateTodo(c, id, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc *todoController) DeleteTodo(c *gin.Context) {
	id := c.Param("todo_id")

	err := tc.tu.DeleteTodo(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/adapter/controller"
)

func initTodoRouter(router *gin.RouterGroup, tc controller.ITodoController) {
	r := router.Group("/todos")
	r.GET("", tc.GetAllTodos)
	r.GET("", tc.GetTodobyId)
	r.POST("", tc.CreateTodo)
	r.PATCH("/:todo_id", tc.UpdateTodo)
	r.DELETE("/:todo_id", tc.DeleteTodo)
}

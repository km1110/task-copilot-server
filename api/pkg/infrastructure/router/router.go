package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/adapter/controller"
	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/usecase"
)

func InitRouter(db *sql.DB) *gin.Engine {
	g := gin.Default()

	// health chack
	g.GET("/health", controller.Health)

	// todo DPI
	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoController := controller.NewTodoController(todoUsecase)

	initTodoRouter(&g.RouterGroup, todoController)

	return g
}

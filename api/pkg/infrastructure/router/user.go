package router

import (
	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/adapter/controller"
	"github.com/km1110/task-copilot-server/pkg/adapter/middleware"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/cache"
)

func initUserRouter(router *gin.RouterGroup, uc controller.IUserController, c cache.IUserCache) {
	r := router.Group("/users")
	r.GET("", uc.GetAllUsers)
	r.GET("/:user_id", uc.GetUserById)
	r.Use(middleware.AuthorizationMiddleware(c))
	{
		r.POST("", uc.CreateUser)
		r.PATCH("/:user_id", uc.UpdateUser)
		r.DELETE("/:user_id", uc.DeleteUser)
	}
}

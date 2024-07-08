package router

import (
	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/adapter/controller"
)

func initAuthRouter(router *gin.RouterGroup, ac controller.IAuthController) {
	r := router.Group("/auth")
	r.POST("/login", ac.Login)
	r.POST("/register", ac.Register)
	r.DELETE("/logout", ac.Logout)
}

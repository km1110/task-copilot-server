package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/cache"
)

// AuthorizationMiddleware is a middleware to check if the user is authorized to access the endpoint

func AuthorizationMiddleware(uc cache.IUserCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user from the context
		role := uc.Get("role")
		if role == "" {
			// ここでredisにアクセスしてroleを取得する
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Check if the user is authorized
		if role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

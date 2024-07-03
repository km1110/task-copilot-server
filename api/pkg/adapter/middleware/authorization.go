package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/cache"
)

func AuthorizationMiddleware(rr repository.IRedisRepository, uc cache.IUserCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := c.Get("firebaseUID")
		key := "user:" + uid.(string)
		role := uc.Get(key)
		if role == "" {
			r, err := rr.Get(key)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}
			uc.Set(key, r)
			role = r
		}

		if role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func FirebaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}

		auth, err := app.Auth(c.Request.Context())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		if idToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		token, err := auth.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("firebaseUID", token.UID)
		c.Next()
	}
}

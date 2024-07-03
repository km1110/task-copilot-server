package main

import (
	"log"

	"github.com/km1110/task-copilot-server/pkg/infrastructure/cache"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/database/postgresql"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/database/redis"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/firebase"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/router"
)

func main() {
	// init db
	db, err := postgresql.NewPostgresConnector()
	if err != nil {
		panic(err)
	}

	// init redis
	rc, err := redis.NewRedisConnector()
	if err != nil {
		panic(err)
	}

	// init firebase
	fbApp, err := firebase.NewFirebaseApp()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// init cache
	c := cache.NewUserCache()

	r := router.InitRouter(db, rc, fbApp, c)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

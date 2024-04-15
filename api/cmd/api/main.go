package main

import (
	"github.com/km1110/task-copilot-server/pkg/infrastructure/database/postgresql"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/router"
)

func main() {
	db, err := postgresql.NewPostgresConnector()
	if err != nil {
		panic(err)
	}

	r := router.InitRouter(db)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

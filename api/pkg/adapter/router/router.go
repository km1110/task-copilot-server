package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/km1110/task-copilot-server/pkg/adapter/controller"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/database/postgresql"
)

func InitRouter() {
	db, err := postgresql.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: newHandlers(db)}

	// TODO: graceful shutdown

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func newHandlers(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// health check
	r.Get("/health", controller.Health(db))

	initTodoRouter(r, db)

	return r
}

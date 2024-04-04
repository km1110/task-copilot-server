package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/database/postgresql"
	"github.com/km1110/task-copilot-server/pkg/web/http/controller"
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

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", controller.GetTodos(db))
		r.Post("/", controller.CreateTodo(db))
		r.Patch("/{todoID}", controller.UpdateTodo(db))
		r.Delete("/{todoID}", controller.DeleteTodo(db))
	})

	return r
}

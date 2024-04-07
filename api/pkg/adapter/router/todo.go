package router

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/km1110/task-copilot-server/pkg/adapter/controller"
)

func initTodoRouter(r *chi.Mux, db *sql.DB) {
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", controller.GetTodos(db))
		r.Post("/", controller.CreateTodo(db))
		r.Patch("/{todoID}", controller.UpdateTodo(db))
		r.Delete("/{todoID}", controller.DeleteTodo(db))
	})
}

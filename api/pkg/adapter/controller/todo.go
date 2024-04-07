package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/repository"
	"github.com/km1110/task-copilot-server/pkg/usecase"
)

type todoResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	TargetDate time.Time `json:"target_date"`
	DoneDate   time.Time `json:"done_date"`
	Status     bool      `json:"status"`
}

func newTodoResponse(t *model.Todo) *todoResponse {
	return &todoResponse{
		ID:         t.ID,
		Name:       t.Name,
		TargetDate: t.TargetDate,
		DoneDate:   t.DoneDate,
		Status:     t.Status,
	}
}

func newTodosResponse(ts []*model.Todo) []*todoResponse {
	var r []*todoResponse
	for _, t := range ts {
		r = append(r, newTodoResponse(t))
	}
	return r
}

func GetTodos(db *sql.DB) http.HandlerFunc {
	repoTodo := repository.NewTodo(db)
	ucGetTodo := usecase.NewGetTodo(repoTodo)

	handler := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ユーザーIDの取得方法を作成
		userID := "b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1"

		todo, err := ucGetTodo.Exec(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := newTodosResponse(todo)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	return handler
}

func CreateTodo(db *sql.DB) http.HandlerFunc {
	repoTodo := repository.NewTodo(db)
	ucTodo := usecase.NewCreateTodo(repoTodo)

	handler := func(w http.ResponseWriter, r *http.Request) {
		var reqBody createTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID := "b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1"

		todo, err := ucTodo.Exec(r.Context(), userID, reqBody.Name, reqBody.TargetDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := newTodoResponse(todo)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
	}
	return handler
}

func UpdateTodo(db *sql.DB) http.HandlerFunc {
	repoTodo := repository.NewTodo(db)
	ucTodo := usecase.NewUpdateTodo(repoTodo)

	handler := func(w http.ResponseWriter, r *http.Request) {
		var reqBody updateTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todoID := chi.URLParam(r, "todo_id")

		todo, err := ucTodo.Exec(r.Context(), todoID, reqBody.Name, reqBody.TargetDate, reqBody.DoneDate, reqBody.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := newTodoResponse(todo)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	}

	return handler
}

func DeleteTodo(db *sql.DB) http.HandlerFunc {
	repoTodo := repository.NewTodo(db)
	ucTodo := usecase.NewDeleteTodo(repoTodo)

	handler := func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "todoID")

		err := ucTodo.Exec(r.Context(), todoID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	}

	return handler
}

type createTodoRequest struct {
	Name       string    `json:"name"`
	TargetDate time.Time `json:"target_date"`
	DoneDate   time.Time `json:"done_date"`
	Status     bool      `json:"status"`
}

type updateTodoRequest struct {
	Name       string    `json:"name"`
	TargetDate time.Time `json:"target_date"`
	DoneDate   time.Time `json:"done_date"`
	Status     bool      `json:"status"`
}

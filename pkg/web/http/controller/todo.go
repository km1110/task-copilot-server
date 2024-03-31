package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/repository"
	"github.com/km1110/task-copilot-server/pkg/usecase"
)

type todoResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	TargetDate time.Time `json:"targetdate"`
	DoneDate   time.Time `json:"donedate"`
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
		userID := "hoge"

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

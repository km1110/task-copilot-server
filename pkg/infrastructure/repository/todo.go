package repository

import (
	"context"
	"database/sql"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
	"golang.org/x/xerrors"
)

type Todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) *Todo {
	return &Todo{db: db}
}

func (trepo *Todo) GetTodos(ctx context.Context, userID string) ([]*model.Todo, error) {
	query := `SELECT "id", "name", "target_date", "done_date", "status" FROM "todos" WHERE "user_id" = $1;`

	rows, err := trepo.db.QueryContext(ctx, query, userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, xerrors.Errorf("trepo.db.QueryRowContext: %w", ErrNotFound)
		default:
			return nil, xerrors.Errorf("trepo.db.QueryRowContext: %v", err)
		}
	}
	defer rows.Close()

	var todos []*model.Todo

	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Name, &t.TargetDate, &t.DoneDate, &t.Status); err != nil {
			return nil, xerrors.Errorf("rows.Scan: %v", err)
		}
		todos = append(todos, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, xerrors.Errorf("rows.Err: %v", err)
	}
	return todos, nil
}

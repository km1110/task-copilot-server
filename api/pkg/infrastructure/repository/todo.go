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

func (trepo *Todo) CreateTodo(ctx context.Context, userID string, t *model.Todo) (*model.Todo, error) {
	query := `INSERT INTO todos("id", "user_id", "name", "target_date", "done_date", "status") VALUES($1, $2, $3, $4, $5, $6);`

	_, err := trepo.db.ExecContext(ctx, query, t.ID, userID, t.Name, t.TargetDate, t.TargetDate, t.Status)
	if err != nil {
		return nil, xerrors.Errorf("trepo.db.ExecContext: %v", err)
	}

	return t, nil
}

func (trepo *Todo) UpdateTodo(ctx context.Context, id string, t *model.Todo) (*model.Todo, error) {
	query := `UPDATE todos SET "name"=$1, "target_date"=$2, "done_date"=$3, "status"=$4 WHERE "id"=$5`

	_, err := trepo.db.ExecContext(ctx, query, t.Name, t.TargetDate, t.DoneDate, t.Status, id)
	if err != nil {
		return nil, xerrors.Errorf("repo.db.ExecContext: %v", err)
	}

	return t, nil

}

func (trepo *Todo) DeleteTodo(ctx context.Context, id string) error {
	query := `DELETE FROM todos WHERE "id" = $1`

	_, err := trepo.db.ExecContext(ctx, query, id)
	if err != nil {
		return xerrors.Errorf("repo.db.ExecContext: %v", err)
	}

	return nil
}

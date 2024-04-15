package repository

import (
	"context"
	"database/sql"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
)

type ITodoRepository interface {
	GetAllTodos(ctx context.Context, userID string, t *[]model.Todo) error
	GetTodobyId(ctx context.Context, userID, todoID string, t *model.Todo) error
	CreateTodo(ctx context.Context, userID string, t *model.Todo) error
	UpdateTodo(ctx context.Context, todoID string, t *model.Todo) error
	DeleteTodo(ctx context.Context, todoID string) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) ITodoRepository {
	return &todoRepository{db: db}
}

func (tr *todoRepository) GetAllTodos(ctx context.Context, userID string, todos *[]model.Todo) error {
	query := `SELECT "id", "name", "target_date", "done_date", "is_completed" FROM "todos" WHERE "user_id" = $1;`

	rows, err := tr.db.QueryContext(ctx, query, userID)
	// write error function
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.TargetDate, &todo.DoneDate, &todo.Is_completed); err != nil {
			return err
		}
		*todos = append(*todos, todo)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (tr *todoRepository) GetTodobyId(ctx context.Context, userID, todoID string, todo *model.Todo) error {
	query := `SELECT "name", "target_date", "done_date", "is_completed" FROM "todos" WHERE "id" = $1 AND "user_id" = $2;`

	row := tr.db.QueryRowContext(ctx, query, todoID, userID)

	var t model.Todo

	if err := row.Scan(&t.Name, &t.TargetDate, &t.DoneDate, &t.Is_completed); err != nil {
		return err
	}

	*todo = t

	return nil
}

func (tr *todoRepository) CreateTodo(ctx context.Context, userID string, todo *model.Todo) error {
	query := `INSERT INTO todos("id", "user_id", "name", "target_date", "done_date", "is_completed") VALUES($1, $2, $3, $4, $5, $6);`

	_, err := tr.db.ExecContext(ctx, query, todo.ID, userID, todo.Name, todo.TargetDate, todo.TargetDate, todo.Is_completed)
	if err != nil {
		return err
	}

	return nil

}

func (tr *todoRepository) UpdateTodo(ctx context.Context, todoID string, todo *model.Todo) error {
	query := `UPDATE todos SET "name"=$1, "target_date"=$2, "done_date"=$3, "is_completed"=$4 WHERE "id"=$5`

	_, err := tr.db.ExecContext(ctx, query, todo.Name, todo.TargetDate, todo.DoneDate, todo.Is_completed, todoID)
	if err != nil {
		return err
	}

	return nil
}

func (tr todoRepository) DeleteTodo(ctx context.Context, todoID string) error {
	query := `DELETE FROM todos WHERE "id" = $1`

	_, err := tr.db.ExecContext(ctx, query, todoID)
	if err != nil {
		return err
	}

	return nil
}

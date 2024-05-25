package repository

import (
	"context"
	"database/sql"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
)

type IAuthRepository interface {
	Login(ctx context.Context, uid string, user *model.User) error
	Register(ctx context.Context, uid string, user *model.User) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) IAuthRepository {
	return &authRepository{db: db}
}

func (ar *authRepository) Login(ctx context.Context, uid string, user *model.User) error {
	query := `SELECT "id", "name", "role_id", "is_active" FROM "users" WHERE "uid" = $1;`

	row := ar.db.QueryRowContext(ctx, query, uid)

	if err := row.Scan(&user.ID, &user.Name, &user.RoleID, &user.Status); err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) Register(ctx context.Context, uid string, user *model.User) error {
	query := `INSERT INTO "users" ("id", "uid", "name", "role_id", "is_active") VALUES ($1, $2, $3, $4, $5);`

	_, err := ar.db.ExecContext(ctx, query, user.ID, user.FirebaseUID, user.Name, user.RoleID, user.Status)
	if err != nil {
		return err
	}

	return nil
}

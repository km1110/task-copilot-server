package repository

import (
	"context"
	"database/sql"

	"github.com/km1110/task-copilot-server/pkg/domain/model"
)

type IUserRepository interface {
	GetAllUsers(ctx context.Context, u *[]model.User) error
	GetUserById(ctx context.Context, uid string, u *model.User) error
	CreateUser(ctx context.Context, u *model.User) error
	UpdateUser(ctx context.Context, uid string, u *model.User) error
	DeleteUser(ctx context.Context, uid string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetAllUsers(ctx context.Context, users *[]model.User) error {
	query := `SELECT "id", "firebase_uid", "name", "role_id", "status" FROM "users";`

	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FirebaseUID, &user.Name, &user.Role.ID, &user.Status); err != nil {
			return err
		}
		*users = append(*users, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetUserById(ctx context.Context, uid string, user *model.User) error {
	query := `SELECT "id", "name", "role_id", "status" FROM "users" WHERE "firebase_uid" = $1;`

	row := ur.db.QueryRowContext(ctx, query, uid)

	if err := row.Scan(&user.ID, &user.Name, &user.Role.ID, &user.Status); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO "users" ("id", "firebase_uid", "name", "role_id", "status") VALUES ($1, $2, $3, $4, $5);`

	_, err := ur.db.ExecContext(ctx, query, user.ID, user.FirebaseUID, user.Name, user.Role.ID, user.Status)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, uid string, user *model.User) error {
	query := `UPDATE "users" SET "name" = $1, "role_id" = $2, "status" = $3 WHERE "firebase_uid" = $4;`

	_, err := ur.db.ExecContext(ctx, query, user.Name, user.Role.ID, user.Status, uid)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, uid string) error {
	query := `DELETE FROM "users" WHERE "firebase_uid" = $1;`

	_, err := ur.db.ExecContext(ctx, query, uid)
	if err != nil {
		return err
	}

	return nil
}

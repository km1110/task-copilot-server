package model

import "github.com/google/uuid"

type User struct {
	ID     string
	Name   string
	RoleID string
	Status bool
}

func NewUser(id, name, role_id string, status bool) *User {
	return &User{
		ID:     id,
		Name:   name,
		RoleID: role_id,
		Status: status,
	}
}

func NewUserID() string {
	return uuid.NewString()
}

func IsValidUserID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

type UserWithTodos struct {
	ID    string
	Todos []*Todo
}

func NewUserWithTodos(u *User, todos []*Todo) *UserWithTodos {
	return &UserWithTodos{
		ID:    u.ID,
		Todos: todos,
	}
}

package database

import (
	"errors"

	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
)

var (
	errInvalidSQLDatabeseInstance = errors.New("invalid sql db instance")
)

const (
	InstanceMySQL int = iota
)

func NewDatabaseSQLFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstanceMySQL:
		return NewMySQLHandler(newConfigMySQL())
	default:
		return nil, errInvalidSQLDatabeseInstance
	}
}

package router

import (
	"errors"
	"time"

	"github.com/km1110/task-copilot-server/pkg/adapter/logger"
	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/adapter/validator"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGin int = iota
)

func NewWebServerFactory(
	instance int,
	db repository.SQL,
	log logger.Logger,
	validator validator.Validator,
	port Port,
	ctxTimeout time.Duration,

) (Server, error) {
	switch instance {
	case InstanceGin:
		return newGinServer(db, log, validator, port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}

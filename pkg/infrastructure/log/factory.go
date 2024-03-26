package log

import (
	"errors"

	"github.com/km1110/task-copilot-server/pkg/adapter/logger"
)

const InstanceZapLogger int = iota

var errInvalidLoggerInstance = errors.New("invalid log instance")

func NewLoggerFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceZapLogger:
		return NewZapLogger()
	default:
		return nil, errInvalidLoggerInstance
	}
}

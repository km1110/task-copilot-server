package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/km1110/task-copilot-server/pkg/adapter/logger"
	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/adapter/validator"
)

type ginEngin struct {
	router     *gin.Engine
	db         repository.SQL
	log        logger.Logger
	validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func newGinServer(
	db repository.SQL,
	log logger.Logger,
	validator validator.Validator,
	port Port,
	t time.Duration,
) *ginEngin {
	return &ginEngin{
		router:     gin.New(),
		db:         db,
		log:        log,
		validator:  validator,
		port:       port,
		ctxTimeout: t,
	}
}

func (g *ginEngin) Listen() {

}

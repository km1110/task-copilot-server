package infrastructure

import (
	"strconv"
	"time"

	"github.com/km1110/task-copilot-server/pkg/adapter/logger"
	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/adapter/validator"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/database"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/log"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/router"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/validation"
)

type config struct {
	appName       string
	logger        logger.Logger
	validator     validator.Validator
	dbSQL         repository.SQL
	ctxTimeout    time.Duration
	webServerPost router.Port
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Name(name string) *config {
	c.appName = name
	return c
}

func (c *config) Logger(instance int) *config {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Successfully configured log")
	return c
}

func (c *config) Validator(instance int) *config {
	v, err := validation.NewValidatorFactory(instance)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured validator")

	c.validator = v
	return c
}

func (c *config) DbSQL(instance int) *config {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		c.logger.Fatalln(err, "Could not make connection to the database")
	}

	c.logger.Infof("Successfully connected to the SQL database")

	c.dbSQL = db
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.dbSQL,
		c.logger,
		c.validator,
		c.webServerPost,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully conofigured router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPost(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.webServerPost = router.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}

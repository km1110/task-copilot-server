package main

import (
	"os"

	"github.com/km1110/task-copilot-server/pkg/infrastructure"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/database"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/log"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/router"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		Logger(log.InstanceZapLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstanceMySQL)

	app.WebServerPost(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin).
		Start()
}

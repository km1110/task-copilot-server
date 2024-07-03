package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/km1110/task-copilot-server/pkg/adapter/controller"
	"github.com/km1110/task-copilot-server/pkg/adapter/middleware"
	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/cache"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/firebase"
	"github.com/km1110/task-copilot-server/pkg/infrastructure/validation"
	"github.com/km1110/task-copilot-server/pkg/usecase"
)

func InitRouter(db *sql.DB, rc *redis.Client, fbApp firebase.IFirebaseApp, c cache.IUserCache) *gin.Engine {
	g := gin.Default()

	// health chack
	g.GET("/health", controller.Health)

	// redis DPI
	redisRepository := repository.NewRedisRepository(rc)

	// auth DPI
	authRepository := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository, redisRepository, c)
	authController := controller.NewAuthController(authUsecase, fbApp)

	// user DPI
	userRepository := repository.NewUserRepository(db)
	userValidator := validation.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	// todo DPI
	todoRepository := repository.NewTodoRepository(db)
	todoValidator := validation.NewTodoValidator()
	todoUsecase := usecase.NewTodoUsecase(todoRepository, todoValidator)
	todoController := controller.NewTodoController(todoUsecase)

	// public router
	initAuthRouter(g.Group("/"), authController)

	authGroup := g.Group("/")
	authGroup.Use(middleware.FirebaseAuth())
	{
		initUserRouter(authGroup, userController, redisRepository, c)
		initTodoRouter(authGroup, todoController)
	}

	return g
}

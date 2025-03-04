package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/todo-web/internal/configs"
	middlewareConfigs "github.com/wisaitas/todo-web/internal/middlewares/configs"
	"github.com/wisaitas/todo-web/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func init() {
	configs.LoadEnv()
}

type App struct {
	App    *fiber.App
	DB     *gorm.DB
	Redis  *redis.Client
	routes func()
}

func InitializeApp() *App {
	app := fiber.New()
	db := configs.ConnectDB()
	redis := configs.ConnectRedis()

	redisClient := utils.NewRedisClient(redis)

	repositories := initializeRepositories(db)
	services := initializeServices(repositories, redisClient)
	handlers := initializeHandlers(services)
	validates := initializeValidates()
	middlewares := initializeMiddlewares(redisClient)

	apiRoutes := app.Group("/api/v1")
	appRoutes := initializeRoutes(apiRoutes, handlers, validates, middlewares)

	return &App{
		App:   app,
		DB:    db,
		Redis: redis,
		routes: func() {
			appRoutes.SetupRoutes()
		},
	}
}

func (r *App) SetupRoutes() {
	r.routes()
}

func (r *App) Run() {
	go func() {
		if err := r.App.Listen(fmt.Sprintf(":%s", configs.ENV.PORT)); err != nil {
			log.Fatalf("error starting server: %v\n", err)
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdown
	r.close()
}

func (r *App) close() {
	sqlDB, err := r.DB.DB()
	if err != nil {
		log.Fatalf("error getting database: %v\n", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", err)
	}

	if err := r.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", err)
	}

	log.Println("gracefully shutdown")
}

func (r *App) SetupMiddlewares() {
	r.App.Use(
		middlewareConfigs.Limiter(),
		middlewareConfigs.CORS(),
		middlewareConfigs.Healthz(),
		middlewareConfigs.Logger(),
		middlewareConfigs.Recovery(),
	)
}

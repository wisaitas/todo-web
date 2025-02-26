package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/todo-web/internal/configs"
	"github.com/wisaitas/todo-web/internal/utils"

	"github.com/redis/go-redis/v9"

	"github.com/wisaitas/todo-web/internal/handlers"
	"github.com/wisaitas/todo-web/internal/middlewares"
	middlewareConfigs "github.com/wisaitas/todo-web/internal/middlewares/configs"
	"github.com/wisaitas/todo-web/internal/repositories"
	"github.com/wisaitas/todo-web/internal/routes"
	"github.com/wisaitas/todo-web/internal/services"
	"github.com/wisaitas/todo-web/internal/validates"

	"github.com/gofiber/fiber/v2"
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

	// Initialize utils
	redisClient := utils.NewRedisClient(redis)

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository, redisClient)
	authService := services.NewAuthService(userRepository, redisClient)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize validates
	userValidate := validates.NewUserValidate()
	authValidate := validates.NewAuthValidate()

	// Initialize middlewares
	authMiddleware := middlewares.NewAuthMiddleware(redis)
	userMiddleware := middlewares.NewUserMiddleware()

	// Initialize routes
	apiRoutes := app.Group("/api/v1")
	userRoutes := routes.NewUserRoutes(apiRoutes, userHandler, userValidate, authMiddleware, userMiddleware)
	authRoutes := routes.NewAuthRoutes(apiRoutes, authHandler, authValidate, authMiddleware)

	return &App{
		App:   app,
		DB:    db,
		Redis: redis,
		routes: func() {
			userRoutes.UserRoutes()
			authRoutes.AuthRoutes()
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
	)
}

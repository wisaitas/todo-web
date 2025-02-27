package initial

import (
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/todo-web/internal/middlewares"
)

func initializeMiddlewares(redis *redis.Client) *Middlewares {
	return &Middlewares{
		AuthMiddleware: *middlewares.NewAuthMiddleware(redis),
		UserMiddleware: *middlewares.NewUserMiddleware(),
	}
}

type Middlewares struct {
	AuthMiddleware middlewares.AuthMiddleware
	UserMiddleware middlewares.UserMiddleware
}

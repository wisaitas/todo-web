package initial

import (
	"github.com/wisaitas/todo-web/internal/middlewares"
	"github.com/wisaitas/todo-web/internal/utils"
)

func initializeMiddlewares(redisUtil utils.RedisClient) *Middlewares {
	return &Middlewares{
		AuthMiddleware: *middlewares.NewAuthMiddleware(redisUtil),
		UserMiddleware: *middlewares.NewUserMiddleware(redisUtil),
	}
}

type Middlewares struct {
	AuthMiddleware middlewares.AuthMiddleware
	UserMiddleware middlewares.UserMiddleware
}

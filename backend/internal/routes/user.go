package routes

import (
	"github.com/wisaitas/todo-web/internal/handlers"
	"github.com/wisaitas/todo-web/internal/middlewares"
	"github.com/wisaitas/todo-web/internal/validates"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	app            fiber.Router
	userHandler    *handlers.UserHandler
	userValidate   *validates.UserValidate
	authMiddleware *middlewares.AuthMiddleware
	userMiddleware *middlewares.UserMiddleware
}

func NewUserRoutes(
	app fiber.Router,
	userHandler *handlers.UserHandler,
	userValidate *validates.UserValidate,
	authMiddleware *middlewares.AuthMiddleware,
	userMiddleware *middlewares.UserMiddleware,
) *UserRoutes {
	return &UserRoutes{
		app:            app,
		userHandler:    userHandler,
		userValidate:   userValidate,
		authMiddleware: authMiddleware,
		userMiddleware: userMiddleware,
	}
}

func (r *UserRoutes) UserRoutes() {
	users := r.app.Group("/users")
	users.Get("/", r.authMiddleware.AuthToken, r.userMiddleware.GetUsers, r.userValidate.ValidateGetUsersRequest, r.userHandler.GetUsers)
	users.Post("/", r.userValidate.ValidateCreateUserRequest, r.userHandler.CreateUser)
}

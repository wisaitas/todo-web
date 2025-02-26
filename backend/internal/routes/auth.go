package routes

import (
	"github.com/wisaitas/todo-web/internal/handlers"
	"github.com/wisaitas/todo-web/internal/middlewares"
	"github.com/wisaitas/todo-web/internal/validates"

	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	app            fiber.Router
	authHandler    *handlers.AuthHandler
	authValidate   *validates.AuthValidate
	authMiddleware *middlewares.AuthMiddleware
}

func NewAuthRoutes(
	app fiber.Router,
	authHandler *handlers.AuthHandler,
	authValidate *validates.AuthValidate,
	authMiddleware *middlewares.AuthMiddleware,

) *AuthRoutes {
	return &AuthRoutes{
		app:            app,
		authHandler:    authHandler,
		authValidate:   authValidate,
		authMiddleware: authMiddleware,
	}
}

func (r *AuthRoutes) AuthRoutes() {
	r.app.Post("/login", r.authValidate.ValidateLoginRequest, r.authHandler.Login)
	r.app.Post("/logout", r.authMiddleware.AuthToken, r.authHandler.Logout)
	r.app.Post("/register", r.authValidate.ValidateRegisterRequest, r.authHandler.Register)
	r.app.Post("/refresh-token", r.authMiddleware.AuthToken, r.authHandler.RefreshToken)
}

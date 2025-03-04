package routes

import "github.com/gofiber/fiber/v2"

type TodoRoutes struct {
	app fiber.Router
}

func NewTodoRoutes(
	app fiber.Router,
) *TodoRoutes {
	return &TodoRoutes{
		app: app,
	}
}

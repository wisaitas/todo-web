package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wisaitas/todo-web/internal/dtos/response"
)

func Recovery() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Message: "Internal Server Error",
			})
		},
	})
}

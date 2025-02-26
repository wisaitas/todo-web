package middlewares

import (
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/models"

	"github.com/gofiber/fiber/v2"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (r *UserMiddleware) GetUsers(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "user context not found",
		})
	}

	// model user not have role LOL XD
	if userContext.Username != "test" {
		return c.Status(fiber.StatusForbidden).JSON(response.ErrorResponse{
			Message: "you are not authorized to access this resource",
		})
	}

	return c.Next()
}

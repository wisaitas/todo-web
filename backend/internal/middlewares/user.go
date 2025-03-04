package middlewares

import (
	"github.com/wisaitas/todo-web/internal/constants"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/models"
	"github.com/wisaitas/todo-web/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserMiddleware struct {
	redisUtil utils.RedisClient
}

func NewUserMiddleware(redisUtil utils.RedisClient) *UserMiddleware {
	return &UserMiddleware{
		redisUtil: redisUtil,
	}
}

func (r *UserMiddleware) GetUsers(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "user context not found",
		})
	}

	if userContext.Role.Name != constants.ROLE.ADMIN {
		return c.Status(fiber.StatusForbidden).JSON(response.ErrorResponse{
			Message: "you are not authorized to access this resource",
		})
	}

	return c.Next()
}

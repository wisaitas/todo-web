package validates

import (
	"github.com/wisaitas/todo-web/internal/dtos/request"
	"github.com/wisaitas/todo-web/internal/dtos/response"

	"github.com/gofiber/fiber/v2"
)

type AuthValidate struct {
}

func NewAuthValidate() *AuthValidate {
	return &AuthValidate{}
}

func (r *AuthValidate) ValidateLoginRequest(c *fiber.Ctx) error {
	req := request.LoginRequest{}

	if err := validateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *AuthValidate) ValidateRegisterRequest(c *fiber.Ctx) error {
	req := request.RegisterRequest{}

	if err := validateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

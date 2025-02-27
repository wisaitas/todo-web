package handlers

import (
	"github.com/wisaitas/todo-web/internal/dtos/queries"
	"github.com/wisaitas/todo-web/internal/dtos/request"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(
	userService services.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (r *UserHandler) GetUsers(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get query",
		})
	}

	resp, statusCode, err := r.userService.GetUsers(query)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "users fetched successfully",
		Data:    resp,
	})
}

func (r *UserHandler) CreateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.CreateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get request",
		})
	}

	resp, statusCode, err := r.userService.CreateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "user created successfully",
		Data:    resp,
	})
}

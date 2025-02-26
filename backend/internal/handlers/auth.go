package handlers

import (
	"github.com/wisaitas/todo-web/internal/dtos/request"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/models"
	"github.com/wisaitas/todo-web/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (r *AuthHandler) Login(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.LoginRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get request",
		})
	}

	resp, statusCode, err := r.authService.Login(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "login successfully",
		Data:    resp,
	})
}

func (r *AuthHandler) Register(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.RegisterRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get request",
		})
	}

	resp, statusCode, err := r.authService.Register(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "user registered successfully",
		Data:    resp,
	})
}

func (r *AuthHandler) Logout(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "user context not found",
		})
	}

	statusCode, err := r.authService.Logout(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "logout successfully",
	})
}

func (r *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "user context not found",
		})
	}

	resp, statusCode, err := r.authService.RefreshToken(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "refresh token successfully",
		Data:    resp,
	})
}

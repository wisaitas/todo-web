package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/todo-web/internal/dtos/queries"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/services"
)

type ProvinceHandler struct {
	provinceService services.ProvinceService
}

func NewProvinceHandler(provinceService services.ProvinceService) *ProvinceHandler {
	return &ProvinceHandler{provinceService: provinceService}
}

func (r *ProvinceHandler) GetProvinces(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get query",
		})
	}

	resp, statusCode, err := r.provinceService.GetProvinces(query)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "provinces fetched successfully",
		Data:    resp,
	})
}

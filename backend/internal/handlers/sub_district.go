package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/todo-web/internal/dtos/queries"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/services"
)

type SubDistrictHandler struct {
	subDistrictService services.SubDistrictService
}

func NewSubDistrictHandler(subDistrictService services.SubDistrictService) *SubDistrictHandler {
	return &SubDistrictHandler{subDistrictService: subDistrictService}
}

func (r *SubDistrictHandler) GetSubDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.SubDistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get query",
		})
	}

	resp, statusCode, err := r.subDistrictService.GetSubDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "sub districts fetched successfully",
		Data:    resp,
	})
}

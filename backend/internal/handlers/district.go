package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/todo-web/internal/dtos/queries"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/services"
)

type DistrictHandler struct {
	districtService services.DistrictService
}

func NewDistrictHandler(districtService services.DistrictService) *DistrictHandler {
	return &DistrictHandler{districtService: districtService}
}

func (r *DistrictHandler) GetDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.DistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get query",
		})
	}

	resp, statusCode, err := r.districtService.GetDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "districts fetched successfully",
		Data:    resp,
	})
}

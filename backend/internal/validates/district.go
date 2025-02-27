package validates

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/todo-web/internal/dtos/queries"
	"github.com/wisaitas/todo-web/internal/dtos/response"
)

type DistrictValidate struct {
}

func NewDistrictValidate() *DistrictValidate {
	return &DistrictValidate{}
}

func (r *DistrictValidate) ValidateGetDistrictsRequest(c *fiber.Ctx) error {
	query := queries.DistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: fmt.Sprintf("failed to validate request: %s", err.Error()),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

package response

import "github.com/wisaitas/todo-web/internal/models"

type RoleResponse struct {
	BaseResponse
	Name string `json:"name"`
}

func (r *RoleResponse) ModelToResponse(role *models.Role) RoleResponse {
	r.ID = role.ID
	r.CreatedAt = role.CreatedAt
	r.UpdatedAt = role.UpdatedAt
	r.Name = role.Name

	return *r
}

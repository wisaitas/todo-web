package response

import "github.com/wisaitas/todo-web/internal/models"

type GetProvincesResponse struct {
	ID     int    `json:"id"`
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}

func (r *GetProvincesResponse) ModelToResponse(province models.Province) GetProvincesResponse {
	r.ID = province.ID
	r.NameTH = province.NameTH
	r.NameEN = province.NameEN

	return *r
}

package response

import "github.com/wisaitas/todo-web/internal/models"

type GetDistrictsResponse struct {
	ID         int    `json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	ProvinceID int    `json:"province_id"`
}

func (r *GetDistrictsResponse) ModelToResponse(district models.District) GetDistrictsResponse {
	r.ID = district.ID
	r.NameTH = district.NameTH
	r.NameEN = district.NameEN
	r.ProvinceID = district.ProvinceID

	return *r
}

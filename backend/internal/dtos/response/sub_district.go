package response

import "github.com/wisaitas/todo-web/internal/models"

type GetSubDistrictsResponse struct {
	ID         int    `json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	DistrictID int    `json:"district_id"`
	ZipCode    int    `json:"zip_code"`
}

func (r *GetSubDistrictsResponse) ModelToResponse(subDistrict models.SubDistrict) GetSubDistrictsResponse {
	r.ID = subDistrict.ID
	r.NameTH = subDistrict.NameTH
	r.NameEN = subDistrict.NameEN
	r.DistrictID = subDistrict.DistrictID
	r.ZipCode = subDistrict.ZipCode

	return *r
}

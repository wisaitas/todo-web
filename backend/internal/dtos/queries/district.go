package queries

type DistrictQuery struct {
	PaginationQuery
	ProvinceID string `query:"province_id"`
}

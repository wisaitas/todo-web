package queries

type SubDistrictQuery struct {
	PaginationQuery
	DistrictID string `query:"district_id"`
}

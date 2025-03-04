package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/todo-web/internal/dtos/queries"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/models"
	"github.com/wisaitas/todo-web/internal/repositories"
	"github.com/wisaitas/todo-web/internal/utils"
)

type DistrictService interface {
	GetDistricts(query queries.DistrictQuery) (resp []response.GetDistrictsResponse, statusCode int, err error)
}

type districtService struct {
	districtRepository repositories.DistrictRepository
	redis              utils.RedisClient
}

func NewDistrictService(
	districtRepository repositories.DistrictRepository,
	redis utils.RedisClient,
) DistrictService {
	return &districtService{
		districtRepository: districtRepository,
		redis:              redis,
	}
}

func (r *districtService) GetDistricts(query queries.DistrictQuery) (resp []response.GetDistrictsResponse, statusCode int, err error) {
	districts := []models.District{}

	cacheKey := fmt.Sprintf("get_districts:%v:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order, query.ProvinceID)

	cache, err := r.redis.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []response.GetDistrictsResponse{}, http.StatusInternalServerError, err
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []response.GetDistrictsResponse{}, http.StatusInternalServerError, err
		}
	}

	if err := r.districtRepository.GetAll(&districts, &query.PaginationQuery, map[string]interface{}{"province_id": query.ProvinceID}); err != nil {
		return []response.GetDistrictsResponse{}, http.StatusInternalServerError, err
	}

	for _, district := range districts {
		respGetDistrict := response.GetDistrictsResponse{}
		resp = append(resp, respGetDistrict.ModelToResponse(district))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []response.GetDistrictsResponse{}, http.StatusInternalServerError, err
	}

	if err := r.redis.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []response.GetDistrictsResponse{}, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil
}

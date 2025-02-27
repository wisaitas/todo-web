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

type SubDistrictService interface {
	GetSubDistricts(query queries.SubDistrictQuery) (resp []response.GetSubDistrictsResponse, statusCode int, err error)
}

type subDistrictService struct {
	subDistrictRepository repositories.SubDistrictRepository
	redis                 utils.RedisClient
}

func NewSubDistrictService(subDistrictRepository repositories.SubDistrictRepository, redis utils.RedisClient) SubDistrictService {
	return &subDistrictService{
		subDistrictRepository: subDistrictRepository,
		redis:                 redis,
	}
}

func (s *subDistrictService) GetSubDistricts(query queries.SubDistrictQuery) (resp []response.GetSubDistrictsResponse, statusCode int, err error) {
	subDistricts := []models.SubDistrict{}

	cacheKey := fmt.Sprintf("get_sub_districts:%v:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order, query.DistrictID)

	cache, err := s.redis.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []response.GetSubDistrictsResponse{}, http.StatusInternalServerError, err
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []response.GetSubDistrictsResponse{}, http.StatusInternalServerError, err
		}

		return resp, http.StatusOK, nil
	}

	if err := s.subDistrictRepository.GetAllBy("district_id", query.DistrictID, &subDistricts); err != nil {
		return []response.GetSubDistrictsResponse{}, http.StatusInternalServerError, err
	}

	for _, subDistrict := range subDistricts {
		respGetSubDistrict := response.GetSubDistrictsResponse{}
		resp = append(resp, respGetSubDistrict.ModelToResponse(subDistrict))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []response.GetSubDistrictsResponse{}, http.StatusInternalServerError, err
	}

	if err := s.redis.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []response.GetSubDistrictsResponse{}, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil
}

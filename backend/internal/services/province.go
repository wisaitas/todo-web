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

type ProvinceService interface {
	GetProvinces(queries queries.PaginationQuery) (resp []response.GetProvincesResponse, statusCode int, err error)
}

type provinceService struct {
	provinceRepository repositories.ProvinceRepository
	redis              utils.RedisClient
}

func NewProvinceService(
	provinceRepository repositories.ProvinceRepository,
	redis utils.RedisClient,
) ProvinceService {
	return &provinceService{
		provinceRepository: provinceRepository,
		redis:              redis,
	}
}

func (r *provinceService) GetProvinces(query queries.PaginationQuery) (resp []response.GetProvincesResponse, statusCode int, err error) {
	provinces := []models.Province{}

	cacheKey := fmt.Sprintf("get_provinces:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redis.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []response.GetProvincesResponse{}, http.StatusInternalServerError, err
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []response.GetProvincesResponse{}, http.StatusInternalServerError, err
		}

		return resp, http.StatusOK, nil
	}

	if err := r.provinceRepository.GetAll(&provinces, &query); err != nil {
		return []response.GetProvincesResponse{}, http.StatusInternalServerError, err
	}

	for _, province := range provinces {
		respGetProvince := response.GetProvincesResponse{}
		resp = append(resp, respGetProvince.ModelToResponse(province))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []response.GetProvincesResponse{}, http.StatusInternalServerError, err
	}

	if err := r.redis.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []response.GetProvincesResponse{}, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil
}

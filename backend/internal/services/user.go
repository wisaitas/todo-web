package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wisaitas/todo-web/internal/constants"
	"github.com/wisaitas/todo-web/internal/dtos/queries"
	"github.com/wisaitas/todo-web/internal/dtos/request"
	"github.com/wisaitas/todo-web/internal/dtos/response"
	"github.com/wisaitas/todo-web/internal/models"
	"github.com/wisaitas/todo-web/internal/repositories"
	"github.com/wisaitas/todo-web/internal/utils"

	"github.com/redis/go-redis/v9"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers(userContext models.UserContext, query queries.PaginationQuery) (resp []response.GetUsersResponse, statusCode int, err error)
	CreateUser(req request.CreateUserRequest) (resp response.CreateUserResponse, statusCode int, err error)
}

type userService struct {
	userRepository repositories.UserRepository
	redisUtil      utils.RedisClient
}

func NewUserService(
	userRepository repositories.UserRepository,
	redisUtil utils.RedisClient,
) UserService {
	return &userService{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *userService) GetUsers(userContext models.UserContext, query queries.PaginationQuery) (resp []response.GetUsersResponse, statusCode int, err error) {
	users := []models.User{}

	cacheKey := fmt.Sprintf("get_users:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, err
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []response.GetUsersResponse{}, http.StatusInternalServerError, err
		}

		return resp, http.StatusOK, nil
	}

	if userContext.Role.Name == constants.ROLE.ADMIN {
		if err := r.userRepository.GetAll(&users, &query, nil, "Role"); err != nil {
			return []response.GetUsersResponse{}, http.StatusInternalServerError, err
		}
	} else {
		if err := r.userRepository.GetAll(&users, &query, map[string]interface{}{"role_id": userContext.Role.ID}, "Role"); err != nil {
			return []response.GetUsersResponse{}, http.StatusInternalServerError, err
		}
	}

	for _, user := range users {
		respGetUser := response.GetUsersResponse{}
		resp = append(resp, respGetUser.ModelToResponse(user))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, err
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil

}

func (r *userService) CreateUser(req request.CreateUserRequest) (resp response.CreateUserResponse, statusCode int, err error) {
	user := req.ToModel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	user.Password = string(hashedPassword)

	if err = r.userRepository.Create(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return resp, http.StatusBadRequest, errors.New("username already exists")
		}

		return resp, http.StatusInternalServerError, err
	}

	return resp.ToResponse(user), http.StatusCreated, nil
}

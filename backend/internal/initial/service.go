package initial

import (
	"github.com/wisaitas/todo-web/internal/services"
	"github.com/wisaitas/todo-web/internal/utils"
)

func initializeServices(repos *Repositories, redisClient utils.RedisClient) *Services {
	return &Services{
		UserService:        services.NewUserService(repos.UserRepository, redisClient),
		AuthService:        services.NewAuthService(repos.UserRepository, redisClient),
		ProvinceService:    services.NewProvinceService(repos.ProvinceRepository, redisClient),
		DistrictService:    services.NewDistrictService(repos.DistrictRepository, redisClient),
		SubDistrictService: services.NewSubDistrictService(repos.SubDistrictRepository, redisClient),
	}
}

type Services struct {
	UserService        services.UserService
	AuthService        services.AuthService
	ProvinceService    services.ProvinceService
	DistrictService    services.DistrictService
	SubDistrictService services.SubDistrictService
}

package initial

import (
	"github.com/wisaitas/todo-web/internal/repositories"
	"gorm.io/gorm"
)

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:        repositories.NewUserRepository(db),
		ProvinceRepository:    repositories.NewProvinceRepository(db),
		DistrictRepository:    repositories.NewDistrictRepository(db),
		SubDistrictRepository: repositories.NewSubDistrictRepository(db),
	}
}

type Repositories struct {
	UserRepository        repositories.UserRepository
	ProvinceRepository    repositories.ProvinceRepository
	DistrictRepository    repositories.DistrictRepository
	SubDistrictRepository repositories.SubDistrictRepository
}

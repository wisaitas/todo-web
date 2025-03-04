package initial

import (
	"github.com/wisaitas/todo-web/internal/models"
	"github.com/wisaitas/todo-web/internal/repositories"
	"gorm.io/gorm"
)

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:        repositories.NewUserRepository(db, repositories.NewBaseRepository[models.User](db)),
		RoleRepository:        repositories.NewRoleRepository(db, repositories.NewBaseRepository[models.Role](db)),
		ProvinceRepository:    repositories.NewProvinceRepository(db, repositories.NewBaseRepository[models.Province](db)),
		DistrictRepository:    repositories.NewDistrictRepository(db, repositories.NewBaseRepository[models.District](db)),
		SubDistrictRepository: repositories.NewSubDistrictRepository(db, repositories.NewBaseRepository[models.SubDistrict](db)),
		TodoRepository:        repositories.NewTodoRepository(db, repositories.NewBaseRepository[models.Todo](db)),
	}
}

type Repositories struct {
	UserRepository        repositories.UserRepository
	RoleRepository        repositories.RoleRepository
	ProvinceRepository    repositories.ProvinceRepository
	DistrictRepository    repositories.DistrictRepository
	SubDistrictRepository repositories.SubDistrictRepository
	TodoRepository        repositories.TodoRepository
}

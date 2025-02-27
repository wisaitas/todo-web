package repositories

import (
	"github.com/wisaitas/todo-web/internal/models"
	"gorm.io/gorm"
)

type SubDistrictRepository interface {
	BaseRepository[models.SubDistrict]
}

type subDistrictRepository struct {
	BaseRepository[models.SubDistrict]
}

func NewSubDistrictRepository(db *gorm.DB) SubDistrictRepository {
	return &subDistrictRepository{
		BaseRepository: NewBaseRepository[models.SubDistrict](db),
	}
}

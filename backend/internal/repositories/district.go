package repositories

import (
	"github.com/wisaitas/todo-web/internal/models"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	BaseRepository[models.District]
}

type districtRepository struct {
	BaseRepository[models.District]
}

func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &districtRepository{
		BaseRepository: NewBaseRepository[models.District](db),
	}
}

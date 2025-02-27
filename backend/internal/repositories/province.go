package repositories

import (
	"github.com/wisaitas/todo-web/internal/models"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	BaseRepository[models.Province]
}

type provinceRepository struct {
	BaseRepository[models.Province]
}

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceRepository{
		BaseRepository: NewBaseRepository[models.Province](db),
	}
}

package repositories

import (
	"github.com/wisaitas/todo-web/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[models.Role]
}

type roleRepository struct {
	BaseRepository[models.Role]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB, baseRepository BaseRepository[models.Role]) RoleRepository {
	return &roleRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

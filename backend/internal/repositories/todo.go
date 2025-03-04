package repositories

import (
	"github.com/wisaitas/todo-web/internal/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
}

type todoRepository struct {
	BaseRepository[models.Todo]
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB, baseRepository BaseRepository[models.Todo]) TodoRepository {
	return &todoRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

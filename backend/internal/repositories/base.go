package repositories

import (
	"github.com/wisaitas/todo-web/internal/dtos/queries"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	WithTx(tx *gorm.DB) BaseRepository[T]
	GetAll(items *[]T, pagination *queries.PaginationQuery, conditions interface{}, relations ...string) error
	GetBy(conditions interface{}, item *T, relations ...string) error
	Create(item *T) error
	CreateMany(items *[]T) error
	Updates(item *T) error
	UpdateMany(items *[]T) error
	Save(item *T) error
	SaveMany(items *[]T) error
	Delete(item *T) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

func (r *baseRepository[T]) WithTx(tx *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: tx,
	}
}

func (r *baseRepository[T]) GetAll(items *[]T, pagination *queries.PaginationQuery, conditions interface{}, relations ...string) error {
	query := r.db

	if conditions != nil {
		query = query.Where(conditions)
	}

	if pagination.Page != nil && pagination.PageSize != nil {
		offset := *pagination.Page * *pagination.PageSize
		query = query.Offset(offset).Limit(*pagination.PageSize)
	}

	if pagination.Sort != nil && pagination.Order != nil {
		orderClause := *pagination.Sort + " " + *pagination.Order
		query = query.Order(orderClause)
	}

	for _, relation := range relations {
		query = query.Preload(relation)
	}

	return query.Find(items).Error
}

func (r *baseRepository[T]) GetBy(conditions interface{}, item *T, relations ...string) error {
	query := r.db

	for _, relation := range relations {
		query = query.Preload(relation)
	}

	return query.Where(conditions).First(item).Error
}

func (r *baseRepository[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r *baseRepository[T]) CreateMany(items *[]T) error {
	return r.db.Create(items).Error
}

func (r *baseRepository[T]) Updates(item *T) error {
	return r.db.Updates(item).Error
}

func (r *baseRepository[T]) UpdateMany(items *[]T) error {
	return r.db.Updates(items).Error
}

func (r *baseRepository[T]) Save(item *T) error {
	return r.db.Save(item).Error
}

func (r *baseRepository[T]) SaveMany(items *[]T) error {
	return r.db.Save(items).Error
}

func (r *baseRepository[T]) Delete(item *T) error {
	return r.db.Delete(item).Error
}

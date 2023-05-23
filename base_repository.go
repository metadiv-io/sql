package sql

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct{}

func (r *BaseRepository[T]) Save(tx *gorm.DB, entity *T) (*T, error) {
	return Save(tx, entity)
}

func (r *BaseRepository[T]) SaveAll(tx *gorm.DB, entities []T) ([]T, error) {
	return SaveAll(tx, entities)
}

func (r *BaseRepository[T]) Delete(tx *gorm.DB, entity *T) error {
	return Delete(tx, entity)
}

func (r *BaseRepository[T]) DeleteAll(tx *gorm.DB, entities []T) error {
	return DeleteAll(tx, entities)
}

func (r *BaseRepository[T]) DeleteBy(tx *gorm.DB, clause *Clause) error {
	return DeleteAllByClause[T](tx, clause)
}

func (r *BaseRepository[T]) FindOne(tx *gorm.DB, clause *Clause) (*T, error) {
	return FindOne[T](tx, clause)
}

func (r *BaseRepository[T]) FindAll(tx *gorm.DB, clause *Clause) ([]T, error) {
	return FindAll[T](tx, clause)
}

func (r *BaseRepository[T]) FindAllComplex(tx *gorm.DB, clause *Clause, sort *Sort, page *Pagination) ([]T, *Pagination, error) {
	return FindAllComplex[T](tx, clause, sort, page)
}

func (r *BaseRepository[T]) Count(tx *gorm.DB, clause *Clause) (int64, error) {
	return Count[T](tx, clause)
}

func (r *BaseRepository[T]) FindByID(tx *gorm.DB, id uint) (*T, error) {
	return FindOne[T](tx, Eq("id", id))
}

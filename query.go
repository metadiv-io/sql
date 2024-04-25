package sql

import (
	"gorm.io/gorm"
)

/*
Paginate creates a new *Pagination.
*/
func Paginate(page, size int) *Pagination {
	return &Pagination{
		Page: page,
		Size: size,
	}
}

/*
Sort creates a new *Sorting.
*/
func Sort(by string, asc bool) *Sorting {
	return &Sorting{
		By:  by,
		Asc: asc,
	}
}

/*
FindOne finds one record from the database.
*/
func FindOne[T any](tx *gorm.DB, cls *Clause) (*T, error) {
	if cls != nil {
		tx = consumeClause(tx, cls)
	}
	output := new(T)
	err := tx.First(output).Error
	if err != nil {
		return nil, err
	}
	return output, err
}

/*
FindAll finds all records from the database.
*/
func FindAll[T any](tx *gorm.DB, cls *Clause) ([]T, error) {
	if cls != nil {
		tx = consumeClause(tx, cls)
	}
	output := make([]T, 0)
	err := tx.Find(&output).Error
	if err != nil {
		return nil, err
	}
	return output, err
}

/*
FindAllComplex finds all records from the database with pagination and sorting.
*/
func FindAllComplex[T any](tx *gorm.DB, cls *Clause, p *Pagination, s *Sorting) ([]T, *Pagination, error) {

	// part 1: find all records
	tx1 := tx.Begin()
	if cls != nil {
		tx1 = consumeClause(tx1, cls)
	}
	if p != nil && p.Page > 0 && p.Size > 0 {
		tx1 = consumePagination(tx1, p)
	}
	if s != nil && s.By != "" {
		tx1 = consumeSorting(tx1, s)
	}
	output := make([]T, 0)
	err := tx1.Find(&output).Error
	if err != nil {
		tx1.Rollback()
		return nil, nil, err
	}
	tx1.Commit()

	// part 2: count total records
	if p != nil && p.Page > 0 && p.Size > 0 {
		tx2 := tx.Begin()
		if cls != nil {
			tx2 = consumeClause(tx2, cls)
		}
		var total int64
		err = tx2.Model(new(T)).Count(&total).Error
		if err != nil {
			tx2.Rollback()
			return nil, nil, err
		}
		tx2.Commit()
		p.Total = total
	} else {
		p = nil
	}

	return output, p, err
}

/*
Count counts the number of records from the database.
*/
func Count[T any](tx *gorm.DB, cls *Clause) (int64, error) {
	if cls != nil {
		tx = consumeClause(tx, cls)
	}
	var total int64
	err := tx.Model(new(T)).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, err
}

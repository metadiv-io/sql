package sql

import "gorm.io/gorm"

// FindOne finds one record that matches the given clause.
func FindOne[T any](tx *gorm.DB, clause *Clause) (*T, error) {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	output := new(T)
	err := tx.First(output).Error
	if err != nil {
		return nil, err
	}
	return output, err
}

// FindAll finds all records that match the given clause.
func FindAll[T any](tx *gorm.DB, clause *Clause) ([]T, error) {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	output := make([]T, 0)
	err := tx.Find(&output).Error
	if err != nil {
		return nil, err
	}
	return output, err
}

// Count counts the number of records that match the given clause.
func Count[T any](tx *gorm.DB, clause *Clause) (int64, error) {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	var count int64
	err := tx.Model(new(T)).Count(&count).Error
	return count, err
}

// FindAllComplex finds all records that match the given clause and applies the given sort and pagination.
func FindAllComplex[T any](tx *gorm.DB, clause *Clause, sort *Sort, pagination *Pagination) ([]T, *Pagination, error) {

	tx1 := tx.Begin()
	if clause != nil {
		tx1 = clause.Consume(tx1)
	}
	if sort != nil && sort.By != "" {
		sort.By = safeField(sort.By)
		if sort.By != "" {
			tx1 = sort.Consume(tx1)
		}
	}
	if pagination != nil && pagination.Page > 0 && pagination.Size > 0 {
		tx1 = pagination.Consume(tx1)
	}
	output := make([]T, 0)
	err := tx1.Find(&output).Error
	if err != nil {
		tx1.Rollback()
		return nil, nil, err
	}
	tx1.Commit()

	if pagination != nil && pagination.Page > 0 && pagination.Size > 0 {
		tx2 := tx.Begin()
		if clause != nil {
			tx2 = clause.Consume(tx2)
		}
		var count int64
		err := tx2.Model(new(T)).Count(&count).Error
		if err != nil {
			tx2.Rollback()
			return nil, nil, err
		}
		pagination.Total = count
		tx2.Commit()
	} else {
		pagination = nil
	}

	return output, pagination, err
}

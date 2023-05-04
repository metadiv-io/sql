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
	if clause != nil {
		tx = clause.Consume(tx)
	}
	if sort != nil && sort.By != "" {
		sort.By = safeField(sort.By)
		if sort.By != "" {
			tx = sort.Consume(tx)
		}
	}
	if pagination != nil && pagination.Page > 0 && pagination.Size > 0 {
		tx = pagination.Consume(tx)
	}
	output := make([]T, 0)
	err := tx.Find(&output).Error
	if err != nil {
		return nil, nil, err
	}

	tx1 := tx.Session(&gorm.Session{NewDB: true})
	if pagination != nil && pagination.Page > 0 && pagination.Size > 0 {
		pagination.Total, err = Count[T](tx1, clause)
		if err != nil {
			return nil, nil, err
		}
	} else {
		pagination = nil
	}
	return output, pagination, err
}

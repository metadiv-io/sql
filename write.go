package sql

import "gorm.io/gorm"

// Save saves the given model.
func Save[T any](tx *gorm.DB, model *T) (*T, error) {
	err := tx.Save(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// SaveAll saves all the given models.
func SaveAll[T any](tx *gorm.DB, models []T) ([]T, error) {
	err := tx.Save(models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

// Delete deletes the given model.
func Delete[T any](tx *gorm.DB, model *T) error {
	return tx.Delete(model).Error
}

// DeleteAll deletes all the given models.
func DeleteAll[T any](tx *gorm.DB, models []T) error {
	return tx.Delete(models).Error
}

// DeleteAllByClause deletes all the models that match the given clause.
func DeleteAllByClause[T any](tx *gorm.DB, clause *Clause) error {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	return tx.Delete(new(T)).Error
}

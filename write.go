package sql

import (
	"gorm.io/gorm"
)

/*
Save saves a record to the database.
*/
func Save[T any](tx *gorm.DB, record *T) (*T, error) {
	err := tx.Save(record).Error
	if err != nil {
		return nil, err
	}
	return record, err
}

/*
SaveAll saves all records to the database.
*/
func SaveAll[T any](tx *gorm.DB, records []T) ([]T, error) {
	err := tx.Save(records).Error
	if err != nil {
		return nil, err
	}
	return records, err
}

/*
Delete deletes a record from the database.
*/
func Delete[T any](tx *gorm.DB, record *T) error {
	return tx.Delete(record).Error
}

/*
DeleteAll deletes all records from the database.
*/
func DeleteAll[T any](tx *gorm.DB, records []T) error {
	tx1 := tx.Begin()
	for i := range records {
		err := tx1.Delete(&records[i]).Error
		if err != nil {
			tx1.Rollback()
			return err
		}
	}
	return tx1.Commit().Error
}

/*
DeleteBy deletes records from the database by a clause.
*/
func DeleteBy[T any](tx *gorm.DB, cls *Clause) error {
	if cls != nil {
		tx = consumeClause(tx, cls)
	}
	return tx.Delete(new(T)).Error
}

package sql

import "gorm.io/gorm"

type Pagination struct {
	Page  int   `form:"page" json:"page"`
	Size  int   `form:"size" json:"size"`
	Total int64 `json:"total"` // this field is used in response, not in query
}

func Page(page int, size int) *Pagination {
	return &Pagination{
		Page: page,
		Size: size,
	}
}

func (p *Pagination) Consume(tx *gorm.DB) *gorm.DB {
	if p.Page > 0 {
		tx = tx.Offset((p.Page - 1) * p.Size)
	}
	if p.Size > 0 {
		tx = tx.Limit(p.Size)
	}
	return tx
}

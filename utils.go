package sql

import (
	"strings"

	"gorm.io/gorm"
)

/*
consumePagination consumes pagination information from models.IPagination and
returns a *gorm.DB.
*/
func consumePagination(tx *gorm.DB, p *Pagination) *gorm.DB {
	if p == nil {
		return tx
	}
	if p.Page > 0 {
		tx = tx.Offset((p.Page - 1) * p.Size)
	}
	if p.Size > 0 {
		tx = tx.Limit(p.Size)
	}
	return tx
}

/*
ConsumeSorting consumes sorting information from models.ISorting and returns a
*gorm.DB.
*/
func consumeSorting(tx *gorm.DB, s *Sorting) *gorm.DB {
	if s == nil {
		return tx
	}
	if s.By != "" {
		if s.Asc {
			tx = tx.Order(s.By)
		} else {
			tx = tx.Order(s.By + " DESC")
		}
	}
	return tx
}

/*
consumeClause consumes a clause and returns *gorm.DB
*/
func consumeClause(tx *gorm.DB, cls *Clause) *gorm.DB {
	if cls == nil {
		return tx
	}
	stm, values := build(cls, make([]any, 0))
	return tx.Where(stm, values...)
}

/*
build is a utility function that builds a clause.
*/
func build(cls *Clause, values []any) (string, []any) {
	var stm string
	switch strings.ToUpper(cls.Operator) {
	case OP_AND, OP_OR:
		stm, values = buildHasChildren(cls, values)
	case OP_IS_NULL, OP_NOT_NULL:
		stm = cls.Field + " " + cls.Operator
	case OP_IN, OP_NOT_IN:
		if cls.Encrypted {
			stm = "AES_DECRYPT(" + cls.Field + ", '" + db_encrypt_key + "') " + cls.Operator + " (" + strings.TrimRight(strings.Repeat("?,", len(cls.Value.([]interface{}))), ",") + ")"
		} else {
			stm = cls.Field + " " + cls.Operator + " (" + strings.TrimRight(strings.Repeat("?,", len(cls.Value.([]interface{}))), ",") + ")"
		}
		values = append(values, cls.Value.([]interface{})...)
	default:
		if cls.Encrypted {
			stm = "AES_DECRYPT(" + cls.Field + ", '" + db_encrypt_key + "') " + cls.Operator + " ?"
		} else {
			stm = cls.Field + " " + cls.Operator + " ?"
		}
		values = append(values, cls.Value)
	}
	return stm, values
}

/*
buildHasChildren is a utility function that builds a clause with children.
We no need to export this function.
*/
func buildHasChildren(cls *Clause, values []interface{}) (string, []interface{}) {
	var buf []string
	for _, child := range cls.Children {
		s, v := build(child, values)
		buf = append(buf, s)
		values = v
	}
	stm := "(" + strings.Join(buf, " "+cls.Operator+" ") + ")"
	return stm, values
}

/*
NewClause creates a new clause.
*/
func newClause(field, operator string, value interface{}, encrypted bool, children ...*Clause) *Clause {
	return &Clause{
		Field:     safeField(field),
		Operator:  operator,
		Value:     value,
		Children:  children,
		Encrypted: encrypted,
	}
}

/*
safeField is a utility function that makes a field safe.
We append ` to the field name to avoid conflict with SQL keywords.
*/
func safeField(field string) string {
	field = strings.ReplaceAll(field, "`", "")
	field = strings.ReplaceAll(field, ".", "`.`")
	return "`" + field + "`"
}

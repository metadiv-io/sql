package sql

import (
	"strings"

	"gorm.io/gorm"
)

const (
	OPERATOR_EQUAL         = "="
	OPERATOR_NOT_EQUAL     = "<>"
	OPERATOR_GREATER       = ">"
	OPERATOR_LESS          = "<"
	OPERATOR_GREATER_EQUAL = ">="
	OPERATOR_LESS_EQUAL    = "<="
	OPERATOR_LIKE          = "LIKE"
	OPERATOR_NOT_LIKE      = "NOT LIKE"
	OPERATOR_IN            = "IN"
	OPERATOR_NOT_IN        = "NOT IN"
	OPERATOR_IS_NULL       = "IS NULL"
	OPERATOR_IS_NOT_NULL   = "IS NOT NULL"
	OPERATOR_AND           = "AND"
	OPERATOR_OR            = "OR"
)

// Clause is a struct that represents a clause in a SQL statement.
type Clause struct {
	Field    string
	Operator string
	Value    interface{}
	Children []*Clause
}

// Build builds the clause into a SQL statement and returns the statement and the values.
func (c *Clause) Build() (string, []interface{}) {
	return build(c, make([]interface{}, 0))
}

// Consume consumes the clause and applies it to the given transaction.
func (s *Clause) Consume(tx *gorm.DB) *gorm.DB {
	if s != nil {
		stm, values := s.Build()
		tx = tx.Where(stm, values...)
	}
	return tx
}

// NewClause creates a new Clause.
func NewClause(field, operator string, value interface{}, children ...*Clause) *Clause {
	return &Clause{
		Field:    safeField(field),
		Operator: operator,
		Value:    value,
		Children: children,
	}
}

func build(clause *Clause, values []interface{}) (string, []interface{}) {
	var stm string
	switch strings.ToUpper(clause.Operator) {
	case OPERATOR_AND, OPERATOR_OR:
		stm, values = buildHasChildren(clause, values)
	case OPERATOR_IS_NULL, OPERATOR_IS_NOT_NULL:
		stm = clause.Field + " " + clause.Operator
	case OPERATOR_IN, OPERATOR_NOT_IN:
		stm = clause.Field + " " + clause.Operator + " (" + strings.TrimRight(strings.Repeat("?,", len(clause.Value.([]interface{}))), ",") + ")"
		values = append(values, clause.Value.([]interface{})...)
	default:
		stm = clause.Field + " " + clause.Operator + " ?"
		values = append(values, clause.Value)
	}
	return stm, values
}

func buildHasChildren(clause *Clause, values []interface{}) (string, []interface{}) {
	var buf []string
	for _, child := range clause.Children {
		s, v := build(child, values)
		buf = append(buf, s)
		values = v
	}
	stm := "(" + strings.Join(buf, " "+clause.Operator+" ") + ")"
	return stm, values
}

func safeField(field string) string {
	field = strings.ReplaceAll(field, "`", "")
	field = strings.ReplaceAll(field, ".", "`.`")
	return "`" + field + "`"
}

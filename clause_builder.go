package sql

// Eq creates a new Clause with the operator "=".
func Eq(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_EQUAL, value)
}

// Neq creates a new Clause with the operator "<>".
func Neq(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_NOT_EQUAL, value)
}

// Gt creates a new Clause with the operator ">".
func Gt(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_GREATER, value)
}

// Gte creates a new Clause with the operator ">=".
func Gte(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_GREATER_EQUAL, value)
}

// Lt creates a new Clause with the operator "<".
func Lt(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_LESS, value)
}

// Lte creates a new Clause with the operator "<=".
func Lte(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_LESS_EQUAL, value)
}

// Like creates a new Clause with the operator "LIKE".
func Like(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_LIKE, value)
}

// NotLike creates a new Clause with the operator "NOT LIKE".
func NotLike(field string, value interface{}) *Clause {
	return NewClause(field, OPERATOR_NOT_LIKE, value)
}

// Similar creates a new Clause with the operator "LIKE" and the value "%value%".
func Similar(field string, value string) *Clause {
	return NewClause(field, OPERATOR_LIKE, "%"+value+"%")
}

// NotSimilar creates a new Clause with the operator "NOT LIKE" and the value "%value%".
func NotSimilar(field string, value string) *Clause {
	return NewClause(field, OPERATOR_NOT_LIKE, "%"+value+"%")
}

// In creates a new Clause with the operator "IN".
func In(field string, value ...interface{}) *Clause {
	return NewClause(field, OPERATOR_IN, value)
}

// NotIn creates a new Clause with the operator "NOT IN".
func NotIn(field string, value ...interface{}) *Clause {
	return NewClause(field, OPERATOR_NOT_IN, value)
}

// IsNull creates a new Clause with the operator "IS NULL".
func IsNull(field string) *Clause {
	return NewClause(field, OPERATOR_IS_NULL, nil)
}

// IsNotNull creates a new Clause with the operator "IS NOT NULL".
func IsNotNull(field string) *Clause {
	return NewClause(field, OPERATOR_IS_NOT_NULL, nil)
}

// And creates a new Clause with the operator "AND".
func And(children ...*Clause) *Clause {
	return NewClause("", OPERATOR_AND, nil, children...)
}

// Or creates a new Clause with the operator "OR".
func Or(children ...*Clause) *Clause {
	return NewClause("", OPERATOR_OR, nil, children...)
}

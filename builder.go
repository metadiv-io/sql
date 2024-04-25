package sql

/*
Eq creates a new clause with operator `=`.
*/
func Eq(field string, value any) *Clause {
	return newClause(field, OP_EQ, value, false)
}

/*
DecryptEq creates a new clause with operator `=` and the value decrypted.
*/
func DecryptEq(field string, value any) *Clause {
	return newClause(field, OP_EQ, value, true)
}

/*
Neq creates a new clause with operator `<>`.
*/
func Neq(field string, value any) *Clause {
	return newClause(field, OP_NEQ, value, false)
}

/*
DecryptNeq creates a new clause with operator `<>` and the value decrypted.
*/
func DecryptNeq(field string, value any) *Clause {
	return newClause(field, OP_NEQ, value, true)
}

/*
Gt creates a new clause with operator `>`.
*/
func Gt(field string, value any) *Clause {
	return newClause(field, OP_GT, value, false)
}

/*
Gte creates a new clause with operator `>=`.
*/
func Gte(field string, value any) *Clause {
	return newClause(field, OP_GTE, value, false)
}

/*
Lt creates a new clause with operator `<`.
*/
func Lt(field string, value any) *Clause {
	return newClause(field, OP_LT, value, false)
}

/*
Lte creates a new clause with operator `<=`.
*/
func Lte(field string, value any) *Clause {
	return newClause(field, OP_LTE, value, false)
}

/*
Like creates a new clause with operator `LIKE`.
*/
func Like(field string, value any) *Clause {
	return newClause(field, OP_LIKE, value, false)
}

/*
DecryptLike creates a new clause with operator `LIKE` and the value decrypted.
*/
func DecryptLike(field string, value any) *Clause {
	return newClause(field, OP_LIKE, value, true)
}

/*
NotLike creates a new clause with operator `NOT LIKE`.
*/
func NotLike(field string, value any) *Clause {
	return newClause(field, OP_NOT_LIKE, value, false)
}

/*
DecryptNotLike creates a new clause with operator `NOT LIKE` and the value decrypted.
*/
func DecryptNotLike(field string, value any) *Clause {
	return newClause(field, OP_NOT_LIKE, value, true)
}

/*
Similar creates a new clause with operator "LIKE" and the value "%value%".
*/
func Similar(field string, value any) *Clause {
	return newClause(field, OP_LIKE, "%"+value.(string)+"%", false)
}

/*
DecryptSimilar creates a new clause with operator "LIKE" and the value "%value%" decrypted.
*/
func DecryptSimilar(field string, value any) *Clause {
	return newClause(field, OP_LIKE, "%"+value.(string)+"%", true)
}

/*
NotSimilar creates a new clause with operator "NOT LIKE" and the value "%value%".
*/
func NotSimilar(field string, value any) *Clause {
	return newClause(field, OP_NOT_LIKE, "%"+value.(string)+"%", false)
}

/*
DecryptNotSimilar creates a new clause with operator "NOT LIKE" and the value "%value%" decrypted.
*/
func DecryptNotSimilar(field string, value any) *Clause {
	return newClause(field, OP_NOT_LIKE, "%"+value.(string)+"%", true)
}

/*
In creates a new clause with operator `IN`.
*/
func In(field string, value ...any) *Clause {
	return newClause(field, OP_IN, value, false)
}

/*
DecryptIn creates a new clause with operator `IN` and the values decrypted.
Please only use on string values.
*/
func DecryptIn(field string, value ...any) *Clause {
	return newClause(field, OP_IN, value, true)
}

/*
NotIn creates a new clause with operator `NOT IN`.
*/
func NotIn(field string, value ...any) *Clause {
	return newClause(field, OP_NOT_IN, value, false)
}

/*
DecryptNotIn creates a new clause with operator `NOT IN` and the values decrypted.
Please only use on string values.
*/
func DecryptNotIn(field string, value ...any) *Clause {
	return newClause(field, OP_NOT_IN, value, true)
}

/*
IsNull creates a new clause with operator `IS NULL`.
*/
func IsNull(field string) *Clause {
	return newClause(field, OP_IS_NULL, nil, false)
}

/*
IsNotNull creates a new clause with operator `IS NOT NULL`.
*/
func IsNotNull(field string) *Clause {
	return newClause(field, OP_NOT_NULL, nil, false)
}

/*
And creates a new clause with operator `AND`.
*/
func And(children ...*Clause) *Clause {
	return newClause("", OP_AND, nil, false, children...)
}

/*
Or creates a new clause with operator `OR`.
*/
func Or(children ...*Clause) *Clause {
	return newClause("", OP_OR, nil, false, children...)
}

/*
Between creates a new clause that checks if a field is between two values.
*/
func Between(field string, from, to any, includeEdgeTo ...bool) *Clause {
	if len(includeEdgeTo) > 0 && includeEdgeTo[0] {
		return And(Gte(field, from), Lte(field, to))
	}
	return And(Gt(field, from), Lt(field, to))
}

/*
NotBetween creates a new clause that checks if a field is not between two values.
*/
func NotBetween(field string, from, to any, includeEdgeTo ...bool) *Clause {
	if len(includeEdgeTo) > 0 && includeEdgeTo[0] {
		return Or(Lt(field, from), Gt(field, to))
	}
	return Or(Lte(field, from), Gte(field, to))
}

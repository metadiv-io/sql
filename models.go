package sql

type Pagination struct {
	Page  int   `form:"page" json:"page"`
	Size  int   `form:"size" json:"size"`
	Total int64 `form:"-" json:"total"` // this field is used in response, not in query
}

type Sorting struct {
	By  string `form:"by" json:"by"`
	Asc bool   `form:"asc" json:"asc"`
}

type Clause struct {
	Field     string    `json:"field"`
	Operator  string    `json:"operator"`
	Value     any       `json:"value"`
	Encrypted bool      `json:"encrypted"`
	Children  []*Clause `json:"children"`
}

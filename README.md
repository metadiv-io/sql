# sql

## Installation

```bash
go get -u github.com/metadiv-io/sql
```

## Highlights

### Connectors

* sql.Connector.MySQL(host, port, username, password, database string) (*gorm.DB, error)

* sql.Connector.Sqlite(path string) (*gorm.DB, error)

* sql.Connector.SqliteMemory() (*gorm.DB, error)

### Clause Builder

* sql.Eq(field string, value interface{}) *Clause

* sql.Neq(field string, value interface{}) *Clause

* sql.Gt(field string, value interface{}) *Clause

* sql.Gte(field string, value interface{}) *Clause

* sql.Lt(field string, value interface{}) *Clause

* sql.Lte(field string, value interface{}) *Clause

* sql.Like(field string, value interface{}) *Clause

* sql.NotLike(field string, value interface{}) *Clause

* sql.Similar(field string, value string) *Clause

* sql.NotSimilar(field string, value string) *Clause

* sql.In(field string, value ...interface{}) *Clause

* sql.NotIn(field string, value ...interface{}) *Clause

* sql.IsNull(field string) *Clause

* sql.IsNotNull(field string) *Clause

* sql.And(children ...*Clause) *Clause

* sql.Or(children ...*Clause) *Clause

### Query Functions

* sql.FindOne[T](tx *gorm.DB, clause *Clause) (*T, error)

* sql.FindAll[T](tx *gorm.DB, clause *Clause) ([]T, error)

* sql.Count[T](tx *gorm.DB, clause *Clause) (int64, error)

* sql.FindAllComplex[T](tx *gorm.DB, clause *Clause, sort *Sort, pagination *Pagination) ([]T, *Pagination, error)

### Write Functions

* Save[T](tx *gorm.DB, model *T) (*T, error)

* SaveAll[T](tx *gorm.DB, models []T) ([]T, error)

* Delete[T](tx *gorm.DB, model *T) error

* DeleteAll[T](tx *gorm.DB, models []T) error

* DeleteAllByClause[T](tx *gorm.DB, clause *Clause) error

# SQL Package Documentation

The SQL package is a comprehensive Go library that facilitates seamless interaction with relational databases. Leveraging the power of GORM, it provides convenient connectors for MySQL and SQLite databases along with a flexible and expressive clause builder for constructing SQL queries.

## Installation

To integrate the SQL package into your Go project, use the following `go get` command:

```bash
go get -u github.com/metadiv-io/sql
```

## Connectors

### MySQL Connector

Establishes a connection to a MySQL database and returns a GORM database instance.

```
func Connector.MySQL(host, port, username, password, database string) (*gorm.DB, error)
```

### SQLite Connector

Connects to an SQLite database using the specified file path.

```
func Connector.Sqlite(path string) (*gorm.DB, error)
```

### SQLite Memory Connector

Creates an in-memory SQLite database and returns a GORM database instance.

```
func Connector.SqliteMemory() (*gorm.DB, error)
```

## Clause Builder

The SQL package includes a powerful clause builder for constructing SQL queries. The following functions assist in creating clauses:

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

## Query Functions

The SQL package provides versatile query functions:

* sql.FindOne[T](tx *gorm.DB, clause *Clause) (*T, error)

* sql.FindAll[T](tx *gorm.DB, clause *Clause) ([]T, error)

* sql.Count[T](tx *gorm.DB, clause *Clause) (int64, error)

* sql.FindAllComplex[T](tx *gorm.DB, clause *Clause, sort *Sort, pagination *Pagination) ([]T, *Pagination, error)

## Write Functions

For database modification operations, the SQL package offers convenient write functions:

* Save[T](tx *gorm.DB, model *T) (*T, error)

* SaveAll[T](tx *gorm.DB, models []T) ([]T, error)

* Delete[T](tx *gorm.DB, model *T) error

* DeleteAll[T](tx *gorm.DB, models []T) error

* DeleteAllByClause[T](tx *gorm.DB, clause *Clause) error

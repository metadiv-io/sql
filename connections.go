package sql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
MySQL connects to a MySQL database and returns a gorm.DB object.
By default, the silent parameter is false, which means that the
gorm.DB object will log all SQL statements to the console.
*/
func MySQL(host, port, username, password, database string, silent ...bool) (*gorm.DB, error) {
	return gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				username, password, host, port, database)),
		gormConfig(silent...),
	)
}

/*
Sqlite connects to a SQLite database and returns a gorm.DB object.
By default, the silent parameter is false, which means that the
gorm.DB object will log all SQL statements to the console.
*/
func Sqlite(path string, silent ...bool) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(path), gormConfig(silent...))
}

/*
SqliteMem connects to a SQLite database in memory and returns a gorm.DB object.
By default, the silent parameter is false, which means that the
gorm.DB object will log all SQL statements to the console.
*/
func SqliteMem(silent ...bool) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared&busy_timeout=5000"),
		gormConfig(silent...))
}

func gormConfig(silent ...bool) *gorm.Config {
	if len(silent) > 0 && silent[0] {
		return &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}
	return &gorm.Config{}
}

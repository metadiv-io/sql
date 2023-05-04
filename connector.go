package sql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connector = &connector{
	LogLevel: logger.Error,
}

type connector struct {
	LogLevel logger.LogLevel
}

// SetLogLevel sets the log level for the connector.
func (c *connector) SetLogLevel(level logger.LogLevel) {
	c.LogLevel = level
}

// MySQL returns a new MySQL connection.
func (c *connector) MySQL(host, port, username, password, database string) (*gorm.DB, error) {
	return gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)),
		c.gormConfig(),
	)
}

// Sqlite returns a new SQLite connection.
func (c *connector) Sqlite(path string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(path), c.gormConfig())
}

// SqliteMemory returns a new SQLite connection to an in-memory database.
func (c *connector) SqliteMemory() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), c.gormConfig())
}

func (c *connector) gormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(c.LogLevel),
	}
}

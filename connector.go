package sql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect returns a new connector.
func Connect(silent ...bool) *connector {
	var s bool = false
	if len(silent) > 0 {
		s = silent[0]
	}
	return &connector{
		Silent: s,
	}
}

type connector struct {
	Silent bool
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
	return gorm.Open(sqlite.Open("file::memory:?cache=shared&busy_timeout=5000"), c.gormConfig())
}

// SqliteMemoryWithName returns a new SQLite connection to an in-memory database with the given name.
func (c *connector) SqliteMemoryWithName(name string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?mode=memory&cache=shared", name)), c.gormConfig())
}

func (c *connector) gormConfig() *gorm.Config {
	if c.Silent {
		return &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}
	return &gorm.Config{}
}

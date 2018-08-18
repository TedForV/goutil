package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Type is new type for db type, from string
type Type string

// define the all db types
const (
	DBTypeMysql    Type = "mysql"
	DBTypePostgres Type = "Postgres"
	DBTypeSqlite   Type = "Sqlite"
	DBTypeMssql    Type = "mssql"
)

// BaseGorm is a struct store base info for gorm
type BaseGorm struct {
	MaxIdleConn    int
	MaxOpenConn    int
	LifetimeOfConn time.Duration
	ConnStr        string
	DBType         Type
}

var db *gorm.DB

// NewBaseGorm is a func for new base gorm
func NewBaseGorm(connStr string, dbType Type, maxIdleConn, maxOpenConn int, lifeTimeOfconn time.Duration) *BaseGorm {
	return &BaseGorm{
		maxIdleConn,
		maxOpenConn,
		lifeTimeOfconn,
		connStr,
		dbType,
	}

}

// NewConn is a func for new conn for connect db
func (bg *BaseGorm) NewConn() (*gorm.DB, error) {
	var err error
	if db == nil {
		db, err = gorm.Open(string(bg.DBType), bg.ConnStr)
		if err != nil {
			db = nil
			return nil, err
		}
		db.DB().SetMaxIdleConns(bg.MaxIdleConn)
		db.DB().SetMaxOpenConns(bg.MaxOpenConn)
		db.DB().SetConnMaxLifetime(bg.LifetimeOfConn)
		return db, nil
	}

	return db, nil

}

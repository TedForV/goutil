package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DBType string

const (
	DBTypeMysql    DBType = "mysql"
	DBTypePostgres DBType = "Postgres"
	DBTypeSqlite   DBType = "Sqlite"
	DBTypeMssql    DBType = "mssql"
)

type BaseGorm struct {
	DB             *gorm.DB
	MaxIdleConn    int
	MaxOpenConn    int
	LifetimeOfConn time.Duration
	ConnStr        string
}

func NewBaseGorm(connStr string, maxIdleConn int, maxOpenConn int, lifetime time.Duration, dbType DBType) (*BaseGorm, error) {
	db, err := gorm.Open(string(dbType), connStr)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(maxIdleConn)
	db.DB().SetMaxOpenConns(maxOpenConn)
	db.DB().SetConnMaxLifetime(lifetime)
	return &BaseGorm{
		db,
		maxIdleConn,
		maxOpenConn,
		lifetime,
		connStr,
	}, nil
}

func (bg *BaseGorm) NewConn() *gorm.DB {
	return bg.DB.New()
}

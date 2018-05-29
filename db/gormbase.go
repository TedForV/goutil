package db

import (
	"github.com/jinzhu/gorm"
)

type DBType string

const (
	DBTypeMysql    DBType = "mysql"
	DBTypePostgres DBType = "Postgres"
	DBTypeSqlite   DBType = "Sqlite"
	DBTypeMssql    DBType = "mssql"
)

type BaseGorm struct {
	//DB *gorm.DB
	//MaxIdleConn    int
	//MaxOpenConn    int
	//LifetimeOfConn time.Duration
	ConnStr string
	DBType  DBType
}

func NewBaseGorm(connStr string, dbType DBType) *BaseGorm {

	//db.DB().SetMaxIdleConns(maxIdleConn)
	//db.DB().SetMaxOpenConns(maxOpenConn)
	//db.DB().SetConnMaxLifetime(lifetime)
	return &BaseGorm{
		//db,
		//maxIdleConn,
		//maxOpenConn,
		//lifetime,
		connStr,
		dbType,
	}
}

func (bg *BaseGorm) NewConn() (*gorm.DB, error) {
	db, err := gorm.Open(string(bg.DBType), bg.ConnStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

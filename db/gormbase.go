package db

import (
	"github.com/jinzhu/gorm"
)

// DBType is new type for db type, from string
type DBType string

// define the all db types
const (
	DBTypeMysql    DBType = "mysql"
	DBTypePostgres DBType = "Postgres"
	DBTypeSqlite   DBType = "Sqlite"
	DBTypeMssql    DBType = "mssql"
)

// BaseGorm is a struct store base info for gorm
type BaseGorm struct {
	//DB *gorm.DB
	//MaxIdleConn    int
	//MaxOpenConn    int
	//LifetimeOfConn time.Duration
	ConnStr string
	DBType  DBType
}

// NewBaseGorm is a func for new base gorm
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

// NewConn is a func for new conn for connect db
func (bg *BaseGorm) NewConn() (*gorm.DB, error) {
	db, err := gorm.Open(string(bg.DBType), bg.ConnStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

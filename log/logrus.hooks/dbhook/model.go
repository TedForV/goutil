package dbhook

import "time"

type ErrorLog struct {
	Id             int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	ProjectId      int
	ProjectAddress string
	Msg            string
	Trace          string
	CreateAt       time.Time
}

//fit the gorm model, point out the mapping table in db
func (ErrorLog) TableName() string {
	return "error_log"
}

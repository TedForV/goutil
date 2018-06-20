package dbhook

import (
	"github.com/TedForV/goutil/db"
	"github.com/sirupsen/logrus"
)

type DBErrorLogHook struct {
	DB        *db.BaseGorm
	ProjectId int
	Address   string
}

func NewDBErrorLogHook(db *db.BaseGorm, projectId int, address string) *DBErrorLogHook {
	return &DBErrorLogHook{
		DB:        db,
		ProjectId: projectId,
		Address:   address,
	}
}

var triggerLevels = []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}

func (hook *DBErrorLogHook) Levels() []logrus.Level {
	return triggerLevels
}

func (hook *DBErrorLogHook) Fire(entry *logrus.Entry) error {
	return nil
}

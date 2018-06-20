package dbhook

import (
	"github.com/sirupsen/logrus"
)

type DBErrorLogHook struct {
	ProjectId int
	Address   string
}

func NewDBErrorLogHook(projectId int, address string) *DBErrorLogHook {
	return &DBErrorLogHook{
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

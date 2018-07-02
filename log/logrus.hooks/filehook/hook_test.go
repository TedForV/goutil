package filehook

import (
	"github.com/TedForV/goutil/log/logrus.hooks"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFileHook_Fire(t *testing.T) {
	hook := NewFileHook("D:\\Test\\")
	logrus.AddHook(hook)

	logrus.WithField(logrus_hooks.Error_AddInfo_Name, "test").WithField(logrus_hooks.Error_Trace_Name, "test").Debug("debug test")
	logrus.WithField(logrus_hooks.Error_AddInfo_Name, "test").WithField(logrus_hooks.Error_Trace_Name, "test").Info("info test")
	logrus.WithField(logrus_hooks.Error_AddInfo_Name, "test").WithField(logrus_hooks.Error_Trace_Name, "test").Warn("warn test")
	logrus.WithField(logrus_hooks.Error_AddInfo_Name, "test").WithField(logrus_hooks.Error_Trace_Name, "test").Error("error test")
	logrus.WithField(logrus_hooks.Error_AddInfo_Name, "test").WithField(logrus_hooks.Error_Trace_Name, "test").Fatal("fatal test")
	logrus.WithField(logrus_hooks.Error_AddInfo_Name, "test").WithField(logrus_hooks.Error_Trace_Name, "test").Panic("panic test")
}

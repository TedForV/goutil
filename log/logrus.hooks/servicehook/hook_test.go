package servicehook

import (
	"github.com/TedForV/goutil/ms/kit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewErrorLogServiceHook(t *testing.T) {
	hook := NewErrorLogServiceHook(100, 100, "localhost:100", &kit.ETCD3Config{
		Server:        "10.10.10.11:2379",
		DialKeepAlive: time.Second * 3,
		DialTimeout:   time.Second * 3,
	}, "/service/log")
	logrus.AddHook(hook)

	logrus.WithField(Error_Trace_Name, "Trace details...").Error("ms error log test")
}

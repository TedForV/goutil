package servicehook

import (
	"context"
	"testing"
	"time"

	"github.com/TedForV/goutil/log/logrus.hooks"
	"github.com/TedForV/goutil/log/pb"
	"github.com/TedForV/goutil/ms/kit"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func TestNewErrorLogServiceHook(t *testing.T) {
	hook := NewErrorLogServiceHook(100, 100, "gRPC", &kit.ETCD3Config{
		Servers:       []string{"10.10.10.11:2379", "10.10.10.12:2379", "10.10.11.200:2379"},
		DialKeepAlive: time.Second * 3,
		DialTimeout:   time.Second * 3,
	}, "/service/log")
	logrus.AddHook(hook)

	logrus.WithField(logrushooks.ErrorTraceName, "Trace details...").Error("ms error log test")
}

func TestClient(t *testing.T) {
	serviceAddress := "10.10.4.101:9001"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := pb.NewLoggerClient(conn)
	_, err = client.AddErrorLog(context.Background(), &pb.ErrorLog{
		ServiceId:      1,
		ServiceTypeId:  1,
		ProjectAddress: "grpcTest",
		Msg:            "Test",
		Trace:          "Test",
		AdditionalInfo: "test",
	})
	if err != nil {
		t.Error(err)
	}
	//log.Print("Result:%+v", result)
}

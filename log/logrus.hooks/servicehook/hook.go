package servicehook

import (
	"context"
	"github.com/TedForV/goutil/ms/kit"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"rpc_log/pb"
)

const (
	Error_Trace_Name   = "Trace"
	Error_AddInfo_Name = "AdditionalInfo"
)

type ErrorLogServiceHook struct {
	ServiceId        int
	ServiceTypeId    int
	Address          string
	EtcdConf         *kit.ETCD3Config
	LogServicePrefix string
}

func NewErrorLogServiceHook(serviceId int, serviceTypeId int, address string, etcdConfig *kit.ETCD3Config, logServicePrefix string) *ErrorLogServiceHook {
	return &ErrorLogServiceHook{
		ServiceId:        serviceId,
		ServiceTypeId:    serviceTypeId,
		Address:          address,
		EtcdConf:         etcdConfig,
		LogServicePrefix: logServicePrefix,
	}
}

var triggerLevels = []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}

func (hook *ErrorLogServiceHook) Levels() []logrus.Level {
	return triggerLevels
}

func (hook *ErrorLogServiceHook) Fire(entry *logrus.Entry) error {
	if _, ok := kit.GetGrpcBalancer(hook.LogServicePrefix); !ok {
		initLogService(hook.EtcdConf, hook.LogServicePrefix, logFactory)
	}

	defer func() {
		if err := recover(); err != nil {
			//log local
		}
	}()

	data, err := kit.RPC(hook.LogServicePrefix, &pb.ErrorLog{
		ServiceId:      int32(hook.ServiceId),
		ServiceTypeId:  int32(hook.ServiceTypeId),
		ProjectAddress: hook.Address,
		Msg:            entry.Message,
		Trace:          entry.Data[Error_Trace_Name].(string),
		AdditionalInfo: entry.Data[Error_AddInfo_Name].(string),
	})

	if err != nil {
		logrus.Debug(err)
	}
	logrus.Debug(data)

	return nil
}

func initLogService(etcdConfig *kit.ETCD3Config, logServicePrefix string, f sd.Factory) {
	kit.InitialKitGrpc(etcdConfig, logServicePrefix, f)
}

func logFactory(instanceAddress string) (endpoint.Endpoint, io.Closer, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		conn, err := grpc.Dial(instanceAddress, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := pb.NewLoggerClient(conn)
		r := req.(*pb.ErrorLog)
		return client.AddErrorLog(ctx, r)
	}, nil, nil
}

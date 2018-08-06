package servicehook

import (
	"context"
	"github.com/TedForV/goutil/log/logrus.hooks"
	"github.com/TedForV/goutil/log/pb"
	"github.com/TedForV/goutil/ms/kit"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
)

// ErrorLogServiceHook is a hook for sending error log to log service
type ErrorLogServiceHook struct {
	ServiceId        int
	ServiceTypeId    int
	Address          string
	EtcdConf         *kit.ETCD3Config
	LogServicePrefix string
}

// NewErrorLogServiceHook is new method for new ErrorLogServiceHook
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

// Levels is the method must defined in hook
func (hook *ErrorLogServiceHook) Levels() []logrus.Level {
	return triggerLevels
}

// Fire is the method must defined in hook
func (hook *ErrorLogServiceHook) Fire(entry *logrus.Entry) error {
	if _, ok := kit.GetGrpcBalancer(hook.LogServicePrefix); !ok {
		initLogService(hook.EtcdConf, hook.LogServicePrefix, logFactory)
	}

	log := pb.ErrorLog{
		ServiceId:      int32(hook.ServiceId),
		ServiceTypeId:  int32(hook.ServiceTypeId),
		ProjectAddress: hook.Address,
		Msg:            entry.Message,
	}
	if v, ok := entry.Data[logrushooks.ERROR_TRACE_NAME]; ok {
		log.Trace = v.(string)
	}
	if v, ok := entry.Data[logrushooks.ERROR_ADDINFO_NAME]; ok {
		log.AdditionalInfo = v.(string)
	}

	kit.RPC(hook.LogServicePrefix, &log)

	// dead lock
	//if err != nil {
	//	logrus.Info(err)
	//}
	//logrus.Info(data)

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

package kit

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

func NewClient(etcdConfig *ETCD3Config) (etcdv3.Client, error) {
	etcdClient, err := etcdv3.NewClient(context.Background(),
		[]string{
			etcdConfig.Server},
		etcdv3.ClientOptions{
			DialTimeout:   etcdConfig.DialTimeout,
			DialKeepAlive: etcdConfig.DialKeepAlive,
		})
	return etcdClient, err
}

func RegisterService(etcdConfig *ETCD3Config, servicePrefix string, instance string) {
	etcdClient, err := NewClient(etcdConfig)
	if err != nil {
		panic(err)
	}
	register := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   servicePrefix + instance,
		Value: instance,
	}, kitlog.NewNopLogger())
	register.Register()
}

type ETCD3Config struct {
	Server        string
	DialTimeout   time.Duration
	DialKeepAlive time.Duration
}

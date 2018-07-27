package kit

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

// NewClient is a new func for Client
func NewClient(etcdConfig *ETCD3Config) (etcdv3.Client, error) {
	etcdClient, err := etcdv3.NewClient(context.Background(),
		etcdConfig.Servers,
		etcdv3.ClientOptions{
			DialTimeout:   etcdConfig.DialTimeout,
			DialKeepAlive: etcdConfig.DialKeepAlive,
		})
	return etcdClient, err
}

// RegisterService is a func that register the service in etcd
func RegisterService(etcdConfig *ETCD3Config, servicePrefix string, instance string) {
	etcdClient, err := NewClient(etcdConfig)
	if err != nil {
		panic(err)
	}
	register := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   servicePrefix + instance,
		Value: instance,
		TTL:   etcdv3.NewTTLOption(time.Second*1, time.Second*3),
	}, kitlog.NewNopLogger())
	register.Register()
}

// ETCD3Config is the etcd config model
type ETCD3Config struct {
	Servers       []string
	DialTimeout   time.Duration
	DialKeepAlive time.Duration
}

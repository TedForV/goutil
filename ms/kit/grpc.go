package kit

import (
	"context"
	"github.com/go-kit/kit/sd/etcdv3"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"time"
)

var lbs map[string]lb.Balancer

func init() {
	lbs = make(map[string]lb.Balancer)
}

func InitialKitGrpc(etcdConfig ETCD3Config, servicePrefix string, f sd.Factory) {
	etcdClient, err := etcdv3.NewClient(context.Background(),
		[]string{
			etcdConfig.Server},
		etcdv3.ClientOptions{
			DialTimeout:   etcdConfig.DialTimeout,
			DialKeepAlive: etcdConfig.DialKeepAlive,
		})
	if err != nil {
		panic(err)
	}

	instancer, err := etcdv3.NewInstancer(etcdClient, servicePrefix, log.NewNopLogger())

	if err != nil {
		panic(err)
	}

	endpointer := sd.NewEndpointer(instancer, f, log.NewNopLogger())

	balancer := lb.NewRoundRobin(endpointer)

	lbs[servicePrefix] = balancer

}

func GetGrpcBalancer(servicePrefix string) (lb.Balancer, bool) {
	balancer, ok := lbs[servicePrefix]
	return balancer, ok
}

func RPC() {

}

type ETCD3Config struct {
	Server        string
	DialTimeout   time.Duration
	DialKeepAlive time.Duration
}
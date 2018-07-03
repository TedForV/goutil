package kit

import (
	"github.com/go-kit/kit/sd/etcdv3"

	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

var lbs map[string]lb.Balancer

func init() {
	lbs = make(map[string]lb.Balancer)
}

// InitialKitGrpc initial a grpc method in client side for using later
func InitialKitGrpc(etcdConfig *ETCD3Config, servicePrefix string, f sd.Factory) {
	etcdClient, err := NewClient(etcdConfig)

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

// GetGrpcBalancer get fitted balancer for use
func GetGrpcBalancer(servicePrefix string) (lb.Balancer, bool) {
	balancer, ok := lbs[servicePrefix]
	return balancer, ok
}

// RPC is a func to dial the server to do the method
func RPC(servicePrefix string, req interface{}) (interface{}, error) {
	if lber, ok := GetGrpcBalancer(servicePrefix); ok {
		reqEp, err := lber.Endpoint()
		if err != nil {
			return nil, err
		}
		return reqEp(context.Background(), req)
	} else {
		return nil, fmt.Errorf("No such service: %s", servicePrefix)
	}
}

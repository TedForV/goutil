package kit

import (
	"context"
	"fmt"

	"github.com/TedForV/goutil/log/logrus.hooks"
	"github.com/go-kit/kit/sd/etcdv3"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

var lbs map[string]lb.Balancer
var etcdClient etcdv3.Client

func init() {
	lbs = make(map[string]lb.Balancer)
}

// InitialKitGrpc initial a grpc method in client side for using later
func InitialKitGrpc(etcdConfig *ETCD3Config, servicePrefix, methodName string, f sd.Factory) {
	var err error
	if etcdClient == nil {
		etcdClient, err = NewClient(etcdConfig)

		if err != nil {
			logrushooks.RecordLog(etcdConfig, err)
			panic(err)
		}
	}

	instancer, err := etcdv3.NewInstancer(etcdClient, servicePrefix, log.NewNopLogger())

	if err != nil {
		panic(err)
	}

	endpointer := sd.NewEndpointer(instancer, f, log.NewNopLogger())

	balancer := lb.NewRoundRobin(endpointer)
	key := composeBalancerKey(servicePrefix, methodName)
	lbs[key] = balancer

}

// GetGrpcBalancer get fitted balancer for use
func GetGrpcBalancer(servicePrefix string, methodName string) (lb.Balancer, bool) {
	key := composeBalancerKey(servicePrefix, methodName)
	balancer, ok := lbs[key]
	return balancer, ok
}

// RPC is a func to dial the server to do the method
func RPC(servicePrefix string, methodName string, req interface{}) (interface{}, error) {
	if lber, ok := GetGrpcBalancer(servicePrefix, methodName); ok {
		reqEp, err := lber.Endpoint()
		if err != nil {
			return nil, err
		}
		return reqEp(context.Background(), req)
	}

	return nil, fmt.Errorf("No such service: %s", servicePrefix)

}

func composeBalancerKey(servicePrefix, methodName string) string {
	return fmt.Sprintf("%s:%s", servicePrefix, methodName)
}

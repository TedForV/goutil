package gprc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/resolver"

	"google.golang.org/grpc/resolver/manual"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// ResolvedClientConn wrap a clientconn
type ResolvedClientConn struct {
	cc *grpc.ClientConn
	rr *manual.Resolver
}

// FetchEndpoints is a delegate for fetching endpoints
type FetchEndpoints func() []string

const (
	targetTemplate = "%s://%s/%s"
)

// NewResolvedClientConn new a resolvedClientConn
func NewResolvedClientConn(scheme string, endpoints []string, authority string, ticker *time.Ticker, fetch FetchEndpoints, opts ...grpc.DialOption) (*ResolvedClientConn, error) {
	target, err := composeTarget(scheme, endpoints[0], authority)
	if err != nil {
		return nil, err
	}
	rr, err := initialResolver(scheme, endpoints, authority, ticker, fetch)
	if err != nil {
		return nil, err
	}
	r := ResolvedClientConn{
		rr: rr,
	}

	conn, err := grpc.DialContext(context.Background(), target, opts...)

	if err != nil {
		return nil, err
	}
	r.cc = conn
	go updateEndpoints(r.rr, ticker, fetch)
	return &r, nil
}

// GetClientConn get clientconn
func (r *ResolvedClientConn) GetClientConn() *grpc.ClientConn {
	return r.cc
}

func updateEndpoints(rr *manual.Resolver, ticker *time.Ticker, fetch FetchEndpoints) {
	for {
		select {
		case <-ticker.C:
			eps := fetch()
			addrs := epsToAddrs(eps)
			rr.NewAddress(addrs)
		}
	}

}

func initialResolver(scheme string, endpoints []string, authority string, ticker *time.Ticker, fetch FetchEndpoints) (*manual.Resolver, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints is empty")
	}

	rr := manual.NewBuilderWithScheme(scheme)
	initialAddrs(rr, endpoints)
	resolver.Register(rr)
	return rr, nil
}

func initialAddrs(rr *manual.Resolver, endpoints []string) {
	addrs := epsToAddrs(endpoints)
	for _, v := range endpoints {
		addrs = append(addrs, resolver.Address{Addr: v})
	}
	rr.InitialAddrs(addrs)
}

func epsToAddrs(endpoints []string) []resolver.Address {
	addrs := make([]resolver.Address, len(endpoints))
	for _, v := range endpoints {
		addrs = append(addrs, resolver.Address{Addr: v})
	}
	return addrs
}

func composeTarget(scheme, endpoint, authority string) (string, error) {
	if len(scheme) == 0 {
		return "", errors.New("scheme is empty")
	}
	if len(endpoint) == 0 {
		return "", errors.New("endpoint is empty")
	}

	return fmt.Sprintf(targetTemplate, scheme, authority, endpoint), nil

}

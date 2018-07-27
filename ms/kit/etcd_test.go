package kit

import (
	"fmt"
	"testing"
	"time"
)

var etcdConfig = &ETCD3Config{
	Servers:       []string{"10.10.10.11:2379"},
	DialKeepAlive: time.Second * 3,
	DialTimeout:   time.Second * 3,
}

func TestNewClient(t *testing.T) {
	client, err := NewClient(etcdConfig)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(fmt.Sprintf("%+v", client))
	}

}

func TestRegisterService(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	RegisterService(etcdConfig, "/test/test", "localhost:0000")

}

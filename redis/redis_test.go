package redis

import (
	"testing"
	"time"
)

func TestCreateRedisConn(t *testing.T) {
	conn, err := CreateRedisConn("192.168.10.128:6379", time.Second*3, time.Second*10, time.Second*20)
	if err != nil {
		t.Error(err)
	}
	t.Log(conn)
}

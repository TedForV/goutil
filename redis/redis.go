package redis

import (
	redis2 "github.com/gomodule/redigo/redis"
	"time"
)

func CreateRedisConn(url string, connTimeout time.Duration, readTimeout time.Duration, writeTimeout time.Duration) (redis2.Conn, error) {
	return redis2.Dial("tcp", url,
		redis2.DialConnectTimeout(connTimeout),
		redis2.DialReadTimeout(readTimeout),
		redis2.DialWriteTimeout(writeTimeout))
}

package redis

import (
	"fmt"
	"strconv"
	"time"

	redis2 "github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	// ZRANGE command
	ZRANGE = "ZRANGEBYSCORE"
	// ZREVRANGE command
	ZREVRANGE = "ZREVRANGEBYSCORE"
	// MaxInfinite is +inf
	MaxInfinite = "+inf"
	// MinInfinite is -inf
	MinInfinite = "-inf"
)

// CreateRedisConn create a connection to redis
func CreateRedisConn(url string, connTimeout time.Duration, readTimeout time.Duration, writeTimeout time.Duration) (redis2.Conn, error) {
	return redis2.Dial("tcp", url,
		redis2.DialConnectTimeout(connTimeout),
		redis2.DialReadTimeout(readTimeout),
		redis2.DialWriteTimeout(writeTimeout))
}

// SetString set a value to the string key
func SetString(conn *redis2.Conn, key string, data interface{}, expiredSeconds int) error {
	var (
		value string
		err   error
	)

	switch data.(type) {
	case string:
		value = data.(string)
	default:
		value, err = json.MarshalToString(data)
	}
	if err != nil {
		return errors.Wrap(err, "SetString failed.")
	}
	_, err = (*conn).Do("SETEX", key, expiredSeconds, value)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("SET failed. Key:%s Value:%s", key, value))
	}
	return nil
}

// GetString get a value from a string key
func GetString(conn *redis2.Conn, key string, data interface{}) error {
	value, err := redis2.String((*conn).Do("GET", key))
	if err != nil {
		return errors.Wrap(err, "GetString failed.")
	}
	return json.UnmarshalFromString(value, data)
}

// AddSortedSet add a data into a sorted set
func AddSortedSet(conn *redis2.Conn, key string, data interface{}, score int64) error {
	var value string
	var err error
	switch data.(type) {
	case string:
		value = data.(string)
	case int:
		value = strconv.Itoa(data.(int))
	default:
		value, err = json.MarshalToString(data)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("%+v", data))
		}
	}
	_, err = (*conn).Do("ZADD", key, score, value)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("ZADD failed. Key:%s Value:%s Score:%d", key, value, score))
	}
	return nil
}

// GetSortSet get a range data from sorted set key
// if you need sort all items without range, set max = min = 0
func GetSortSet(conn *redis2.Conn, key string, pageNo int, pageRow int, max, min int, isDESC bool) ([]string, error) {
	var maxStr, minStr string
	if max == min && max == 0 {
		maxStr, minStr = MaxInfinite, MinInfinite
	} else {
		maxStr, minStr = strconv.Itoa(max), strconv.Itoa(min)
	}

	var command string
	if isDESC {
		command = ZREVRANGE
	} else {
		command = ZRANGE
	}
	value, err := redis2.Strings((*conn).Do(command, key, maxStr, minStr, "LIMIT", 0, 10))
	if err != nil {
		return value, errors.Wrap(err, fmt.Sprintf("GetSortSet failed. Key: %s pageNo: %d pageRow: %d isDESC: %t", key, pageNo, pageRow, isDESC))
	}
	return value, nil
}

// DeleteSortSetItem delete the item from a sorted set
func DeleteSortSetItem(conn *redis2.Conn, key string, data interface{}) error {
	var value string
	var err error
	switch data.(type) {
	case string:
		value = data.(string)
	case int:
		value = strconv.Itoa(data.(int))
	default:
		value, err = json.MarshalToString(data)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("%+v", data))
		}
	}
	_, err = (*conn).Do("ZREM", key, value)
	return err
}

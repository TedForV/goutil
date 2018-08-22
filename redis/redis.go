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
	// ZRANGEBYSCORE command
	ZRANGEBYSCORE = "ZRANGEBYSCORE"
	// ZREVRANGEBYSCORE command
	ZREVRANGEBYSCORE = "ZREVRANGEBYSCORE"
	// ZCOUNT command
	ZCOUNT = "ZCOUNT"
	// SETEX command
	SETEX = "SETEX"
	// ZADD command
	ZADD = "ZADD"
	// LIMIT command
	LIMIT = "LIMIT"
	// ZREM command
	ZREM = "ZREM"
	// GET command
	GET = "GET"
	// MGET command
	MGET = "MGET"
	// ZRANGE command
	ZRANGE = "ZRANGE"
	// ZCARD command
	ZCARD = "ZCARD"
	// ZREMRANGEBYSCORE command
	ZREMRANGEBYSCORE = "ZREMRANGEBYSCORE"
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
	_, err = (*conn).Do(SETEX, key, expiredSeconds, value)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("SET failed. Key:%s Value:%s", key, value))
	}
	return nil
}

// GetString get a value from a string key
func GetString(conn *redis2.Conn, key string, data interface{}) error {
	value, err := redis2.String((*conn).Do(GET, key))
	if err != nil {
		return errors.Wrap(err, "GetString failed.")
	}
	return json.UnmarshalFromString(value, data)
}

// GetStrings get multi-values
func GetStrings(conn *redis2.Conn, keys []string) (interface{}, error) {
	iKeys := make([]interface{}, len(keys))
	for i, v := range keys {
		iKeys[i] = v
	}
	value, err := redis2.Strings((*conn).Do(MGET, iKeys...))
	if err != nil {
		return nil, errors.Wrap(err, "GetStrings failed.")
	}

	return &value, nil
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
	_, err = (*conn).Do(ZADD, key, score, value)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("ZADD failed. Key:%s Value:%s Score:%d", key, value, score))
	}
	return nil
}

// GetSortSet get a range data from sorted set key
// if you need sort all items without range, set max = min = 0
func GetSortSet(conn *redis2.Conn, key string, pageNo int, pageRow int, max, min int64, isDESC bool) ([]string, error) {
	maxStr, minStr := getRange(max, min)

	var command string
	if isDESC {
		command = ZREVRANGEBYSCORE
	} else {
		command = ZRANGEBYSCORE
	}
	value, err := redis2.Strings((*conn).Do(command, key, maxStr, minStr, LIMIT, (pageNo-1)*pageRow, pageRow))
	if err != nil {
		return value, errors.Wrap(err, fmt.Sprintf("GetSortSet failed. Key: %s pageNo: %d pageRow: %d isDESC: %t", key, pageNo, pageRow, isDESC))
	}
	return value, nil
}

// GetSortedSetCount get sorted set count in range
// if no range limited, set max = min = 0
func GetSortedSetCount(conn *redis2.Conn, key string, max, min int64) (int, error) {
	maxStr, minStr := getRange(max, min)
	return redis2.Int((*conn).Do(ZCOUNT, key, minStr, maxStr))
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
	_, err = (*conn).Do(ZREM, key, value)
	return err
}

// RangeScoreDeleteSortedSet delete a rangable items
func RangeScoreDeleteSortedSet(conn *redis2.Conn, key string, max, min int64) (int, error) {
	maxStr, minStr := getRange(max, min)
	return redis2.Int((*conn).Do(ZREMRANGEBYSCORE, key, minStr, maxStr))
}

func getRange(max, min int64) (string, string) {
	if max == min && max == 0 {
		return MaxInfinite, MinInfinite
	}

	return strconv.FormatInt(max, 10), strconv.FormatInt(min, 10)

}

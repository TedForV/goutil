package redis

import (
	"reflect"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

var testPrefixKey = "test:"

func TestCreateRedisConn(t *testing.T) {
	conn, err := CreateRedisConn("192.168.10.128:6379", time.Second*3, time.Second*10, time.Second*20)
	if err != nil {
		t.Error(err)
	}
	t.Log(conn)
}

func TestSetString(t *testing.T) {
	conn, err := newConn()
	if err != nil {
		t.Error(err)
		return
	}
	key := testPrefixKey + "1"
	input := model{
		ID:     1,
		Name:   "大声的发送的",
		Age:    20,
		Remark: "大声的发现擦拭豆腐干发二娃如果",
	}
	SetString(&conn, key, &input, 10)
	var result model
	GetString(&conn, key, &result)
	assert.Equal(t, input.Name, result.Name)
	assert.Equal(t, input.Age, result.Age)
}

func TestAddSortedSet(t *testing.T) {
	conn, err := newConn()
	if err != nil {
		t.Error(err)
		return
	}
	key := "sortedsettest:1"
	for i := 1; i < 100; i++ {
		AddSortedSet(&conn, key, i, int64(i))
	}
}

// func TestGetSortSet(t *testing.T) {
// 	conn, err := newConn()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	key := "sortedsettest:1"
// 	data, err := GetSortSet(&conn, key, 1, 10, true)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	for i, v := range data {
// 		t.Logf("%d %s", i, v)
// 	}
// }

func TestGetSortSet(t *testing.T) {
	conn, err := newConn()
	if err != nil {
		t.Error(err)
		return
	}
	type args struct {
		conn    *redis2.Conn
		key     string
		pageNo  int
		pageRow int
		max     int
		min     int
		isDESC  bool
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				conn:    &conn,
				key:     "sortedsettest:1",
				pageNo:  1,
				pageRow: 10,
				max:     0,
				min:     0,
				isDESC:  true,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSortSet(tt.args.conn, tt.args.key, tt.args.pageNo, tt.args.pageRow, tt.args.max, tt.args.min, tt.args.isDESC)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSortSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSortSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSortSetItem(t *testing.T) {
	conn, err := newConn()
	if err != nil {
		t.Error(err)
		return
	}
	key := "sortedsettest:1"
	err = DeleteSortSetItem(&conn, key, 95)
	if err != nil {
		t.Error(err)
		return
	}
}

func newConn() (redis.Conn, error) {
	return CreateRedisConn("192.168.10.128:6379", time.Second*3, time.Second*10, time.Second*20)
}

type model struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Age    int32  `json:"age"`
	Remark string `json:"remark"`
}

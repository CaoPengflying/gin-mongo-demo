package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	userKey      = "userdao:uid"
	todayUserKey = "userdao:uid:20201127"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
	rdb = GetRedisClient()
}
func TestSetUserKey(t *testing.T) {
	rdb.SAdd(ctx, userKey, 10000, 10001, 10002, 10003, 10004, 10005)

}
func TestSetTodayUserKey(t *testing.T) {
	rdb.SAdd(ctx, todayUserKey, 10000, 10002, 10003, 10005, 10006, 10007)

}

func TestGetSet(t *testing.T) {
	userIds := rdb.SMembers(ctx, userKey)
	t.Log(userIds)

}

//差集
func TestUnion(t *testing.T) {
	ctx := context.Background()
	diff := rdb.SDiff(ctx, userKey, todayUserKey).Val()
	t.Log(diff)
}

//bitmap统计
//11月签到统计
func TestSetBit(t *testing.T) {
	//用户3000在11月3号签到成功
	client, mock := redismock.NewClientMock()
	mock.ExpectSetBit("uid:sign:3000:202012", 6, 1).SetVal(1)
	mock.ExpectSetBit("uid:sign:3000:202012", 6, 0).SetVal(0)
	mock.ExpectSetBit("uid:sign:3000:202013", 6, 0).SetErr(errors.New("conn fail"))

	tests := []struct {
		Key          string
		Offset       int64
		Value        bool
		ExpectResult error
	}{
		{"uid:sign:3000:202012", 6, true, nil},
		{"uid:sign:3000:202012", 6, false, nil},
		{"uid:sign:3000:202013", 6, false, errors.New("conn fail")},
	}

	for _, test := range tests {
		res := SetBit(ctx, client, test.Key, test.Offset, test.Value)
		assert.Equal(t, test.ExpectResult, res)
	}

}

func TestGetBit(t *testing.T) {
	res := GetBit(ctx, "uid:sign:3000:202012", 6)
	//用户3000在11月3号签到成功
	t.Log(res)
}

func TestBitCount(t *testing.T) {

	bitCount := redis.BitCount{
		0,
		29,
	}
	val := rdb.BitCount(ctx, "uid:sign:3000:202011", &bitCount).Val()
	t.Log(val)
}

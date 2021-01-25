package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

const (
	userKey      = "user:uid"
	todayUserKey = "user:uid:20201127"
)
var rdb *redis.Client
var ctx = context.Background()

func init() {
	rdb = GetRedisClient()
}
func TestSetUserKey(t *testing.T) {
	rdb.SAdd(ctx, userKey, 10000,10001,10002,10003,10004,10005)


}
func TestSetTodayUserKey(t *testing.T) {
	rdb.SAdd(ctx, todayUserKey, 10000,10002,10003,10005,10006,10007)

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
	SetBit(ctx,"uid:sign:3000:202012",6,true)
}

func TestGetBit(t *testing.T) {
	res := GetBit(ctx,"uid:sign:3000:202012",6)
	//用户3000在11月3号签到成功
	t.Log(res)
}

func TestBitCount(t *testing.T) {

	bitCount := redis.BitCount{
		0,
		29,
	}
	val := rdb.BitCount(ctx,"uid:sign:3000:202011",&bitCount).Val()
	t.Log(val)
}

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var client *redis.Client
func init() {
	client = GetRedisClient()
}
//加锁
func Lock(ctx context.Context, key, unique string, expire time.Duration) bool {
	res := client.SetNX(ctx, key, unique, expire)
	return res.Val()
}

//释放锁
func FreeLock(ctx context.Context, key, unique string) (bool, error) {
	value := client.Get(ctx, key)
	if value.Val() == unique {
		res := client.Del(ctx, key)
		if res.Val() == 1 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, ErrNotSameClientLock
	}

}

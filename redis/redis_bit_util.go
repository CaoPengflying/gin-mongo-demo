package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

const BitOneValue = 1
const BitZeroValue = 0

//位图：设置偏移量offset的值
func SetBit(ctx context.Context, key string, offset int64, value bool) {
	client := GetRedisClient()
	if value {
		client.SetBit(ctx, key, offset, BitOneValue)
	} else {
		client.SetBit(ctx, key, offset, BitZeroValue)
	}
}

//位图：获取偏移量offset的值
func GetBit(ctx context.Context, key string, offset int64) bool {
	client := GetRedisClient()
	intCmd := client.GetBit(ctx, key, offset)
	if intCmd.Val() > 0 {
		return true
	} else {
		return false
	}
}

//位图：统计为1的数量
func BitCount(ctx context.Context, key string, start, end int64) int64 {
	client := GetRedisClient()
	bitCount := redis.BitCount{Start: start, End: end}
	resCmd := client.BitCount(ctx, key, &bitCount)
	return resCmd.Val()
}

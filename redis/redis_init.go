package redis

import (
	"gin-mongo-demo/config"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func init() {
	redisClient = config.RedisDbs["redis"]
	if redisClient == nil {
		panic("redis is not found")
	}
}

//返回redis客户端
func GetRedisClient() *redis.Client {
	return redisClient
}

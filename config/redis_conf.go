package config

import (
	"strconv"

	"github.com/go-redis/redis/v8"
)

//RedisConf ...
type RedisConf struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	Auth string `yaml:"password"`
	Db   int    `yaml:"db"`
}

type RedisMap map[string]*redis.Client

// 多redis配置
func GetRedisMap(confMap map[string]RedisConf) RedisMap {
	redisMap := RedisMap{}
	for key, conf := range confMap {
		client := redis.NewClient(&redis.Options{
			Addr:     conf.Addr + ":" + strconv.Itoa(conf.Port),
			Password: conf.Auth,
			DB:       conf.Db,
		})
		redisMap[key] = client
	}
	return redisMap
}

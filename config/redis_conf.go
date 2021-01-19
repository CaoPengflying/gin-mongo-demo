package config

import (
	"github.com/go-redis/redis/v8"
	"strconv"
)

//RedisConf ...
type RedisConf struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	Auth string `yaml:"password"`
	Db   int    `yaml:"db"`
}

type RedisMap map[string]*redis.Client

const (
	Redis_Addr = "81.68.180.162"
	Port       = 6379
	DB         = 5
	Password   = "BYjv49etb8I2I3KI"
)

// go文档说明
// @description 获取redis操作的客户端
func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Redis_Addr + ":" + strconv.Itoa(Port),
		Password: Password,
		DB:       DB,
	})
}

// 多redis配置
func GetRedisMap(confMap map[string]RedisConf) RedisMap {
	redisMap := RedisMap{}
	for key, conf := range confMap {
		client := redis.NewClient(&redis.Options{
			Addr:     conf.Addr + ":" + strconv.Itoa(Port),
			Password: conf.Auth,
			DB:       conf.Db,
		})
		redisMap[key] = client
	}
	return redisMap
}

package config

import (
	"flag"
	"fmt"
	"gin-mongo-demo/middleware/clog"
	"go.uber.org/config"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"
)

var (
	ListenPort int

	AppName string

	RedisDbs RedisMap

	MongoDbs MongoMap

	Consul ConsulConf

	LogConf clog.Config
)

func init() {
	loadConf()

	clog.Init(&LogConf, AppName)
}

func loadConf() {
	loadAppInfo()

	var env string
	flag.StringVar(&env, "env", "dev", "set env")

	testing.Init()
	flag.Parse()

	configDir := getConfDir()
	configName := fmt.Sprintf("application-%s.yaml", env)
	log.Printf("configName=%s", configName)


	configPath := path.Join(configDir, configName)
	opt := config.File(configPath)
	conf, err := config.NewYAML(opt)
	if err != nil {
		log.Printf("read_yaml_fail err=%v", err)
		panic("read_yaml_fail")
	}

	loadRedisConf(conf)

	loadMongoConf(conf)

	loadConsulInfo(conf)

	loadLogConf(conf)

}

//loadLogConf 加载日志配置信息
func loadLogConf(conf *config.YAML) {
	err := conf.Get("clog").Populate(&LogConf)
	if err != nil {
		panic("file_log_config_is_not_found")
	}
}

// loadConsulInfo 加载注册中心
func loadConsulInfo(conf *config.YAML) {
	err := conf.Get("consul").Populate(&Consul)
	if err != nil {
		panic("consul_config_is_not_found")
	}
}

// loadAppInfo 加载应用信息
func loadAppInfo() {
	configDir := getConfDir()
	configName := fmt.Sprintf("application.yaml")
	configPath := path.Join(configDir, configName)
	opt := config.File(configPath)
	conf, err := config.NewYAML(opt)

	if err != nil {
		log.Printf("read_yaml_fail err=%v", err)
		panic("application.yaml is not found")
	}

	ListenPort, err = strconv.Atoi(conf.Get("application").Get("port").String())
	if err != nil {
		ListenPort = 52308
	}

	AppName = conf.Get("application").Get("name").String()
}

//loadRedisConf 加载Redis服务
func loadRedisConf(conf *config.YAML) {
	defaultRedisConf := RedisConf{}
	err := conf.Get("redis").Populate(&defaultRedisConf)
	if err != nil {
		panic("redis_config_is_not_found")
	}
	redisMap := map[string]RedisConf{}
	redisMap["redis"] = defaultRedisConf

	RedisDbs = GetRedisMap(redisMap)
}

// loadMongoConf 加载mongo服务
func loadMongoConf(conf *config.YAML) {
	defaultMongoConf := MongoConf{}
	err := conf.Get("mongo").Populate(&defaultMongoConf)
	if err != nil {
		panic("mongo_config_is_not_found")
	}

	mongoMap := map[string]MongoConf{}
	mongoMap["mongo"] = defaultMongoConf

	MongoDbs = GetMongoMap(mongoMap)
}

func getConfDir() string {
	dir := "gin-mongo-demo"
	for i := 0; i < 5; i++ {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			break
		}
		dir = filepath.Join("..", dir)
	}

	return dir
}

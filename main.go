package main

import (
	"fmt"
	"gin-mongo-demo/middleware/clog"
	"gin-mongo-demo/middleware/trace"

	"gin-mongo-demo/config"
	"gin-mongo-demo/controller"
	"gin-mongo-demo/controller/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	//consulReg := consul.NewRegistry(
	//	registry.Addrs(clog.Consul.Addr),
	//)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(trace.Trace())
	InitRoute(engine)

	server := web.NewService(
		web.Address(fmt.Sprintf(":%d", config.ListenPort)),
		web.Handler(engine),
		//web.Registry(consulReg),
		web.Name(config.AppName),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	clog.Info("event=start log_conf=%v", config.LogConf.FileConfig.LogFilePath)
	err := server.Run()

	if err != nil {
		clog.Error("event=server_start_err err=%v", err)
	}

}

func InitRoute(g *gin.Engine) {
	controller.Include(user.InitRouter)

	controller.Init(g)
}

package main

import (
	"fmt"

	"gin-mongo-demo/config"
	"gin-mongo-demo/controller"
	"gin-mongo-demo/controller/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func init() {
	fmt.Println("init ")

}
func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs(config.Consul.Addr),
	)

	engine := gin.New()

	InitRoute(engine)

	server := web.NewService(
		web.Address(fmt.Sprintf(":%d", config.ListenPort)),
		web.Handler(engine),
		web.Registry(consulReg),
		web.Name(config.AppName),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	err := server.Run()

	if err != nil {
		fmt.Printf("event=server_start_err err=%s", err)
	}
}

func InitRoute(g *gin.Engine) {
	controller.Include(user.InitRouter)

	controller.Init(g)
}

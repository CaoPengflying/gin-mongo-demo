package main

import "C"
import (
	"fmt"
	"gin-mongo-demo/dao"
	"gin-mongo-demo/entity"
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
		registry.Addrs("127.0.0.1:8500"),
	)
	r := gin.Default()
	r.GET("hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "hello"})
	})
	r.POST("getByName", func(context *gin.Context) {
		user := entity.User{}
		err := context.Bind(&user)
		if err != nil {
			panic(err)
		}
		context.JSON(200, gin.H{"message": dao.GetByName(user.Name)})
	})
	r.POST("/createUser", func(context *gin.Context) {
		user := entity.User{}
		err := context.Bind(&user)
		if err != nil {
			panic(err)
		}
		dao.Insert(user)
	})
	server := web.NewService(
		web.Address(":8001"),
		web.Handler(r),
		web.Registry(consulReg),
		web.Name("micro-demo"),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//server.Init()
	err := server.Run()
	if err != nil {
		fmt.Println(err)
	}
}

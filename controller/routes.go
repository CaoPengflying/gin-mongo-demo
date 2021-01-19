package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Option func(*gin.Engine)

var options = make([]Option, 0, 8)

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init ...
func Init(g *gin.Engine)  {
	for _, opt := range options {
		opt(g)
	}

	g.Handle("GET", "/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "SUCCESS")
	})

	g.Handle("HEAD", "ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

}

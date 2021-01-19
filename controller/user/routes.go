package user

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	group := g.Group("user")

	group.POST("create_user", CreateUser)
}

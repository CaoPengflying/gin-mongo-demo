package user

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	group := g.Group("userdao")

	group.POST("create_user", CreateUser)
}

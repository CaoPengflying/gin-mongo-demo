package user

import (
	"gin-mongo-demo/middleware/clog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserSignInfoReq struct {
	UserNo string `json:"user_no"`
}

func GetUserSignInfo(ctx *gin.Context) {
	var req GetUserSignInfoReq
	err := ctx.Bind(&req)
	if err != nil {
		clog.InfoC(ctx, "event=parse_create_user_req_fail err=%v", err)
		return
	}
	clog.InfoC(ctx, "create_user")

	ctx.JSON(http.StatusOK, "create userdao success")

}

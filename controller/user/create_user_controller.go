package user

import (
	"gin-mongo-demo/middleware/clog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type CreateUserReq struct {
	Id     bson.ObjectId `json:"id"`
	Name   string        `json:"name"`
	UserNo string        `json:"user_no"`
	Age    int           `json:"age"`
	OrgNo  string        `json:"org_no"`
}

func CreateUser(ctx *gin.Context) {
	var req CreateUserReq
	err := ctx.Bind(&req)
	if err != nil {
		clog.InfoC(ctx, "event=parse_create_user_req_fail err=%v", err)
		return
	}
	clog.InfoC(ctx, "create_user")


	ctx.JSON(http.StatusOK, "create userdao success")

}

package user

import (
	"log"
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
		log.Printf("event=parse_create_user_req_fail err=%v", err)
		return
	}

	ctx.JSON(http.StatusOK, "create user success")

}

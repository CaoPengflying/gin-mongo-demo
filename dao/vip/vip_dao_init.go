package vip

import (
	"gin-mongo-demo/config"
	"gin-mongo-demo/constants"
	"gopkg.in/mgo.v2"
)

const (
	ColVip = "vip"
)

var userSession *mgo.Session

func init() {
	userSession := config.MongoDbs["mongo"]
	if userSession == nil {
		panic("mongo is not found")
	}
}

func GetVipSession() (*mgo.Session,*mgo.Collection) {
	s := userSession.Copy()

	c := s.DB(constants.DbEggCnode).C(ColVip)

	return s,c
}

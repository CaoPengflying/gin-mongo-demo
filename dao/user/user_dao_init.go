package user

import (
	"gin-mongo-demo/config"
	"gin-mongo-demo/constants"
	"gopkg.in/mgo.v2"
)

const (
	ColUser = "user"
)

var userSession *mgo.Session

func init() {
	userSession = config.MongoDbs["mongo"]
	if userSession == nil {
		panic("mongo is not found")
	}
}

func GetUserSession() (*mgo.Session,*mgo.Collection) {
	s := userSession.Copy()
	c := s.DB(constants.DbEggCnode).C(ColUser)

	return s,c
}

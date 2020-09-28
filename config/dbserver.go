package config

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	URL = "mongodb://81.68.180.162:27017,81.68.180.162:27018,81.68.180.162:27019?maxPoolSize=50"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		session, err := mgo.Dial(URL)
		if err != nil {
			log.Println("mongo connect error", err)
			panic(err)
		}
		session.SetMode(mgo.SecondaryPreferred, true)
	}
	return session.Copy()
}

func init() {
	session, _ = mgo.Dial(URL)
}

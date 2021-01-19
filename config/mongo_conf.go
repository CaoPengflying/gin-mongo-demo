package config

import (
	"gopkg.in/mgo.v2"

	"log"
)

const (
	URL = "mongodb://81.68.180.162:27017,81.68.180.162:27018,81.68.180.162:27019?maxPoolSize=50"
)

type MongoMap map[string]*mgo.Session

type MongoConf struct {
	Url string `yaml:"url"`
}
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

func GetMongoMap(confMap map[string]MongoConf) MongoMap {
	mongoMap := MongoMap{}
	for key, conf := range confMap {
		session, err :=mgo.Dial(conf.Url)
		if err != nil {
			log.Println("mongo connect error", err)
			panic(err)
		}
		session.SetMode(mgo.SecondaryPreferred, true)
		mongoMap[key] = session
	}
	return mongoMap
}

func init() {
	session, _ = mgo.Dial(URL)
}



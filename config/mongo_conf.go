package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

type MongoMap map[string]*mgo.Session

type MongoConf struct {
	Url string `yaml:"url"`
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



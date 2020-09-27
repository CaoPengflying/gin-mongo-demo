package config

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	URL = "mongodb://81.68.180.162:27017"
)

type Collection struct {
	DB string
	string
}

func GetSession() *mgo.Session {
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Println("mongo connect error", err)
		panic(err)
	}
	return session
}


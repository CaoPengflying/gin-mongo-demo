package vip

import (
	"gin-mongo-demo/entity"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//类型嵌套

/**
插入一个文档
*/
func InsertVip(vip entity.Vip) {
	session,c := GetVipSession()
	defer session.Close()

	err := c.Insert(vip)
	if err != nil {
		log.Fatal(err)
	}
}

func GetByName(name string) *entity.Vip {
	session,c := GetVipSession()
	defer session.Close()

	vip := entity.Vip{}
	query := bson.M{"name": name}
	err := c.Find(query).One(&vip)
	if err != nil {
		panic(err)
	}
	return &vip
}

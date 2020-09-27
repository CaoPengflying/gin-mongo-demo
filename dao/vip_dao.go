package dao

import (
	"gin-mongo-demo/config"
	"gin-mongo-demo/entity"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//类型嵌套

/**
插入一个文档
*/
func InsertVip(vip entity.Vip) {
	session := config.GetSession()
	defer session.Close()
	c := session.DB("egg_cnode").C("vip")
	err := c.Insert(vip)
	if err != nil {
		log.Fatal(err)
	}
}

func GetVipByName(name string) *entity.Vip {
	session := config.GetSession()
	defer session.Close()
	c := session.DB("egg_cnode").C("vip")
	vip := entity.Vip{}
	query := bson.M{"name": name}
	err := c.Find(query).One(&vip)
	if err != nil {
		panic(err)
	}
	return &vip
}

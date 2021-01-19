package user

import (
	"gin-mongo-demo/constants"
	"gin-mongo-demo/entity"
	"gopkg.in/mgo.v2/bson"

)

/**
插入一个文档
*/
func Insert(u entity.User) {
	session,c := GetUserSession()
	defer session.Close()

	err := c.Insert(u)
	if err != nil {
		panic("error")
	}
}

func UpdateByName(user *entity.User) {
	session,c := GetUserSession()
	defer session.Close()

	query := bson.M{entity.UserName: user.Name}
	err := c.Update(query, user)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateById(id interface{}, data map[string]interface{}) error {
	session,c := GetUserSession()
	defer session.Close()

	update := bson.M{
		"$set": bson.M(data),
	}

	err := c.UpdateId(id, update)
	return err
}

/**
根据主键删除文档
*/
func DeleteById(id bson.ObjectId) error {
	session,c := GetUserSession()
	defer session.Close()

	return c.RemoveId(id)
}



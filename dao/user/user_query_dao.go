package user

import (
	"gin-mongo-demo/constants"
	"gin-mongo-demo/entity"
	"gopkg.in/mgo.v2/bson"
)

/**
根据某一字段查询一个文档
*/
func GetByName(name string) *entity.User {
	session,c := GetUserSession()
	defer session.Close()

	user := entity.User{}
	query := bson.M{entity.UserName: name}
	err := c.Find(query).One(&user)
	if err != nil {
		panic(err)
	}

	return &user
}

/**
根据某一字段查询文档切片
*/
func ListByOrgNo(orgNo string) []entity.User {
	session,c := GetUserSession()
	defer session.Close()

	var userList []entity.User
	query := bson.M{entity.UserOrgNo: orgNo}
	err := c.Find(query).All(&userList)
	if err != nil {
		panic(err)
	}

	return userList
}

/**
组合查询文档
*/
func GetByOrgNoOrName(orgNo string, name string) []entity.User {
	session,c := GetUserSession()
	defer session.Close()

	var userList []entity.User
	query := bson.M{"$or": []bson.M{{entity.UserOrgNo: orgNo}, {entity.UserName: name}}}
	err := c.Find(query).All(&userList)
	if err != nil {
		panic(err)
	}

	return userList
}

/**
in 查询文档
*/
func ListByOrgNos(orgNos []string) []entity.User {
	session,c := GetUserSession()
	defer session.Close()

	var userList []entity.User
	query := bson.M{entity.UserOrgNo: bson.M{"$in": orgNos}}
	err := c.Find(query).All(&userList)
	if err != nil {
		panic(err)
	}

	return userList
}

/**
根据大于条件查询文档
*/
func ListByGtAge(age int) []entity.User {
	session,c := GetUserSession()
	defer session.Close()

	var userList []entity.User
	query := bson.M{entity.UserAge: bson.M{"$gt": age}}
	err := c.Find(query).All(&userList)
	if err != nil {
		panic(err)
	}
	return userList
}

/**

 */
func ListByNotEq(name string) []entity.User {
	session,c := GetUserSession()
	defer session.Close()


	var userList []entity.User
	query := bson.M{entity.UserName: bson.M{"$ne": name}}
	err := c.Find(query).Limit(constants.PageOffset).All(&userList)
	if err != nil {
		panic(err)
	}

	return userList
}
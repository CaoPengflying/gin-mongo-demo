package userdao

import (
	"context"
	"gin-mongo-demo/constants"
	"gin-mongo-demo/entity"
	"gin-mongo-demo/middleware/clog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserQuery struct {
}

// @description 根据姓名查询用户
// @auth caopengfei 2021/1/27
// @param name string
// @return user,error
func (u *UserQuery) GetByName(context context.Context, name string) (*entity.User, error) {
	session, c := GetUserSession()
	defer session.Close()

	user := entity.User{}
	query := bson.M{entity.UserName: name}
	err := c.Find(query).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		clog.ErrorC(context, "event=get_user_by_name_fail err=%v name=%s", err, name)
		return nil, err
	}

	return &user, nil
}

/**
根据某一字段查询文档切片
*/
func (u *UserQuery) ListByOrgNo(orgNo string) []entity.User {
	session, c := GetUserSession()
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
func (u *UserQuery) GetByOrgNoOrName(orgNo string, name string) []entity.User {
	session, c := GetUserSession()
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
func (u *UserQuery) ListByOrgNos(orgNos []string) []entity.User {
	session, c := GetUserSession()
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
func (u *UserQuery) ListByGtAge(age int) []entity.User {
	session, c := GetUserSession()
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
func (u *UserQuery) ListByNotEq(name string) []entity.User {
	session, c := GetUserSession()
	defer session.Close()

	var userList []entity.User
	query := bson.M{entity.UserName: bson.M{"$ne": name}}
	err := c.Find(query).Limit(constants.PageOffset).All(&userList)
	if err != nil {
		panic(err)
	}

	return userList
}

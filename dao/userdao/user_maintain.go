package userdao

import (
	"gin-mongo-demo/entity"
	"gin-mongo-demo/middleware/clog"
	"gopkg.in/mgo.v2/bson"
)

type UserMaintain struct {
}

// @description 插入用户
// @auth caopengfei 2021/1/27
// @param user
// @return
func (um *UserMaintain) Insert(u entity.User) {
	session, c := GetUserSession()
	defer session.Close()

	err := c.Insert(u)
	if err != nil {
		panic("error")
	}
}

// @description 根据姓名修改用户信息
// @auth caopengfei 2021/1/27
// @param
// @return
func (um *UserMaintain) UpdateByName(user *entity.User) {
	session, c := GetUserSession()
	defer session.Close()

	query := bson.M{entity.UserName: user.Name}
	err := c.Update(query, user)
	if err != nil {
		clog.Error("event=update_by_name_fail err=%v", err)
	}
}

// @description 根据id修改用户信息
// @auth caopengfei 2021/1/27
// @param
// @return
func (um *UserMaintain) UpdateById(id interface{}, data map[string]interface{}) error {
	session, c := GetUserSession()
	defer session.Close()

	update := bson.M{
		"$set": bson.M(data),
	}

	err := c.UpdateId(id, update)
	return err
}

// @description 根据id删除用户
// @auth caopengfei 2021/1/27
// @param
// @return
func (um *UserMaintain) DeleteById(id bson.ObjectId) error {
	session, c := GetUserSession()
	defer session.Close()

	return c.RemoveId(id)
}

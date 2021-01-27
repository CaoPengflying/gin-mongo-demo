package userdao

import (
	"gin-mongo-demo/entity"
	"gopkg.in/mgo.v2/bson"
)

type IUserMaintain interface {
	Insert(u entity.User)
	UpdateByName(user *entity.User)
	UpdateById(id interface{}, data map[string]interface{}) error
	DeleteById(id bson.ObjectId) error
}

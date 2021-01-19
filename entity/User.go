package entity

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Name   string        `bson:"name"`
	UserNo string        `bson:"user_no"`
	Age    int           `bson:"age"`
	OrgNo  string        `bson:"org_no"`
	a      string
}

const (
	UserName  = "name"
	UserNo    = "user_no"
	UserAge   = "age"
	UserOrgNo = "org_no"
)

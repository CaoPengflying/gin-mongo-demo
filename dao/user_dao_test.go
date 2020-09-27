package dao

import (
	"gin-mongo-demo/entity"
	"testing"
)

func TestInsert(t *testing.T) {
	user := entity.User{
		Name:   "003",
		UserNo: "0003",
		Age:    27,
		OrgNo:  "A003",
	}
	Insert(user)
}

func TestGetByName(t *testing.T) {
	name := "002"
	user := GetByName(name)
	t.Log(user)
}

func TestGetByOrgNo(t *testing.T) {
	orgNo := "A001"
	userList := ListByOrgNo(orgNo)
	t.Log(userList)
}

func TestGetByOrgNoAndName(t *testing.T) {
	orgNo := "A001"
	name := "test_d1"
	list := GetByOrgNoOrName(orgNo, name)
	t.Log(list)
}

func TestListByOrgNos(t *testing.T) {
	orgNos := []string{"A001", "A002"}
	userList := ListByOrgNos(orgNos)
	t.Log(userList)

}

func TestListByGtAge(t *testing.T) {
	age := 20
	userList := ListByGtAge(age)
	t.Log(userList)
}

func TestListByNotEq(t *testing.T) {
	name := "cpf"
	userList := ListByNotEq(name)
	t.Log(userList)
}

func TestUpdateById(t *testing.T) {
	name := "002"
	user := GetByName(name)
	user.UserNo = "0004"
	UpdateByName(user)
}
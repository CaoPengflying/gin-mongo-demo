package userdao

import (
	"context"
	"testing"
)

var userQuery = UserQuery{}

func TestGetByName(t *testing.T) {
	name := "002"
	user, _ := userQuery.GetByName(context.Background(), name)
	t.Log(user)
}

func TestGetByOrgNo(t *testing.T) {
	orgNo := "A001"
	userList := userQuery.ListByOrgNo(orgNo)
	t.Log(userList)
}

func TestGetByOrgNoAndName(t *testing.T) {
	orgNo := "A001"
	name := "test_d1"
	list := userQuery.GetByOrgNoOrName(orgNo, name)
	t.Log(list)
}

func TestListByOrgNos(t *testing.T) {
	orgNos := []string{"A001", "A002"}
	userList := userQuery.ListByOrgNos(orgNos)
	t.Log(userList)

}

func TestListByGtAge(t *testing.T) {
	age := 20
	userList := userQuery.ListByGtAge(age)
	t.Log(userList)
}

func TestListByNotEq(t *testing.T) {
	name := "cpf"
	userList := userQuery.ListByNotEq(name)
	t.Log(userList)
}

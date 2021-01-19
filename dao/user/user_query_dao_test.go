package user

import (
	"testing"
)


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


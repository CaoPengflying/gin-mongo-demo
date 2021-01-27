package userdao

import (
	"context"
	"gin-mongo-demo/entity"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var userMaintain = UserMaintain{}

func TestRead(t *testing.T) {
	for true {
		age := 20
		userList := userQuery.ListByGtAge(age)
		t.Log(userList)
	}
}

func TestWrite(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		name := "000000" + strconv.Itoa(i)
		user := entity.User{
			Name:   name[len(name)-6:],
			UserNo: name[len(name)-6:],
			Age:    27,
			OrgNo:  "A" + name[len(name)-6:],
		}
		userMaintain.Insert(user)
		t.Log(name)
	}
}

func TestInsert(t *testing.T) {
	user := entity.User{
		Name:   "004",
		UserNo: "0003",
		Age:    27,
		OrgNo:  "A003",
	}
	userMaintain.Insert(user)
}

func TestInsertList(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(i)
		name := "000000" + strconv.Itoa(i)
		user := entity.User{
			Name:   name[len(name)-6:],
			UserNo: name[len(name)-6:],
			Age:    27,
			OrgNo:  "A" + name[len(name)-6:],
		}
		userMaintain.Insert(user)
	}

}

func TestUpdateByName(t *testing.T) {
	name := "003"
	//userdao := GetByName(name)
	//userdao.UserNo = "0004"
	user := &entity.User{
		UserNo: "0002",
		Name:   name,
	}
	userMaintain.UpdateByName(user)
}

func TestUpdateById(t *testing.T) {
	u1, _ := userQuery.GetByName(context.Background(), "004")
	//userdao := GetByName(name)
	//userdao.UserNo = "0004"
	user := map[string]interface{}{}
	user["user_no"] = "0004"
	err := userMaintain.UpdateById(u1, user)
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteById(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"000000", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := userQuery.GetByName(context.Background(), tt.name)
			err := userMaintain.DeleteById(user.Id)
			require.Equalf(t, tt.wantErr, err != nil, "delete by id error = %v", err, tt.wantErr)
		})
	}
}

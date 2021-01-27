package service

import (
	"context"
	"gin-mongo-demo/entity"
	mock_userdao "gin-mongo-demo/mock/daomock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserSignInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_userdao.NewMockIUserQuery(ctrl)

	user1 := entity.User{
		Name: "cpf",
		Age:  20,
	}

	user2 := entity.User{
		Name: "cpf",
		Age:  23,
	}

	m.EXPECT().GetByName(context.Background(), "cpf").Return(user1, nil)
	m.EXPECT().GetByName(context.Background(), "zzc").Return(user2, nil)

	tests := []struct {
		name         string
		expectResult bool
	}{
		{"cpf", true}, {"zzc", false},
	}

	for _, test := range tests {
		flag, err := GetUserSignInfo(context.Background(), test.name)

		assert.Nil(t, err)
		assert.Equal(t, flag, test.expectResult)
	}

}

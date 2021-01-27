package service

import (
	"context"
	"gin-mongo-demo/dao/userdao"
	"gin-mongo-demo/redis"
)

func GetUserSignInfo(ctx context.Context, userName string) (bool, error) {
	userQuery := userdao.UserQuery{}
	user, err := userQuery.GetByName(ctx, userName)
	if err != nil {
		return false, err
	}

	signFlag := redis.GetBit(ctx, user.UserNo, int64(user.Age))

	return signFlag, nil
}

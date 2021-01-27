// @Description query user info interface
// @Author caopengfei 2021/1/27 11:15
package userdao

import (
	"context"

	"gin-mongo-demo/entity"
)

type IUserQuery interface {
	GetByName(context context.Context, name string) (*entity.User, error)
	ListByOrgNo(orgNo string) []entity.User
	GetByOrgNoOrName(orgNo string, name string) []entity.User
	ListByOrgNos(orgNos []string) []entity.User
	ListByGtAge(age int) []entity.User
	ListByNotEq(name string) []entity.User
}

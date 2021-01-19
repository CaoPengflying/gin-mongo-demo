package vip

import (
	"gin-mongo-demo/entity"
	"testing"
)

func TestInsertVip(t *testing.T) {
	integral := entity.Integral{
		LeftInte: 100,
		SumInte:  100,
	}
	changeFlows := []entity.ChangeFlow{{ChangeInfo: "1to2"}, {ChangeInfo: "2to3"}}

	areas := []entity.Area{{Province: "河北", City: "唐山", Address: "迁安"}, {Province: "浙江", City: "杭州", Address: "江干"}}
	vip := entity.Vip{
		Name:        "cpf",
		Integral:    integral,
		ChangeFlows: changeFlows,
		Areas:       areas,
	}
	InsertVip(vip)
}

func TestGetVipByName(t *testing.T) {
	vip := GetByName("cpf")
	t.Log(vip)
}


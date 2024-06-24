package ckhouse

import (
	"encoding/json"

	"github.com/moke-game/game/services/analytics/names"
)

type PetAct struct {
	Base
	UID        string `json:"uid"`         //玩家uid
	PetId      int64  `json:"pet_id"`      //宠物蛋ID
	ActId      int8   `json:"act_id"`      //1:孵化开始,2:孵化取消3:孵化成功,4:喂养蛋,5:洗练宠物
	ItemId     int32  `json:"item_id"`     //物品ID
	PartId     int32  `json:"part_id"`     //部位ID
	PartVal    int32  `json:"part_val"`    //部位值
	CreateTime string `json:"create_time"` //开始时间 Y-m-d H:i:s
}

func (g *PetAct) GetEventName() names.EventName {
	return "pet_act"
}

func (g *OrderPay) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

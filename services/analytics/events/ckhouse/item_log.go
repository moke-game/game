package ckhouse

import (
	"encoding/json"

	"github.com/moke-game/game/services/analytics/names"
)

type ItemLog struct {
	Base
	ID          string `json:"id"`
	UID         string `json:"uid"`          //
	ItemId      int32  `json:"item_id"`      // '物品Id',
	ItemCnt     int32  `json:"item_cnt"`     // '物品数量',
	ItemType    int32  `json:"item_type"`    // '物品类型',
	ItemQuality int32  `json:"item_quality"` // '物品品质',
	Act         int8   `json:"act"`          // '1:增加 2:减少',
	Source      string `json:"source"`       // '物品来源',
	CreateTime  string `json:"create_time"`  //开始时间 Y-m-d H:i:s
}

func (g *ItemLog) GetEventName() names.EventName {
	return "item_log"
}

func (g *ItemLog) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

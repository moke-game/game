package ckhouse

import (
	"encoding/json"

	"github.com/moke-game/game/services/analytics/names"
)

type HeroUpgrade struct {
	Base
	UID        string `json:"uid"`         //玩家uid
	HeroId     int32  `json:"hero_id"`     //英雄ID
	HeroLv     int32  `json:"hero_lv"`     //英雄升级后的等级
	CreateTime string `json:"create_time"` //开始时间 Y-m-d H:i:s
}

func (g *HeroUpgrade) GetEventName() names.EventName {
	return "hero_level_up_log"
}

func (g *HeroUpgrade) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

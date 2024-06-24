package ckhouse

import (
	"encoding/json"

	"github.com/moke-game/game/services/analytics/names"
)

type BattleUserLog struct {
	Base
	UID          string `json:"uid"`            //玩家uid
	RoomId       string `json:"room_id"`        //房间ID
	PlayId       int32  `json:"play_id"`        //玩法ID
	Cup          int32  `json:"cup"`            //获取奖杯数
	Win          int32  `json:"win"`            //是否胜利
	Mvp          int32  `json:"mvp"`            //是否mvp
	Again        int32  `json:"again"`          //是否点击再次进入
	Reconnect    int32  `json:"reconnect"`      //是否重连
	RobotTake    int32  `json:"robot_take"`     //是否机器人托管
	MapId        int32  `json:"map_id"`         //地图ID
	HeroId       int32  `json:"hero_id"`        //使用的英雄ID
	HeroLv       int32  `json:"hero_lv"`        //使用英雄的等级
	SkinId       int32  `json:"skin_id"`        //英雄使用的皮肤ID
	CampId       int32  `json:"camp_id"`        //玩家阵营ID
	CampKillCnt  int32  `json:"camp_kill_cnt"`  //玩家阵营击败人数
	CampDeathCnt int32  `json:"camp_death_cnt"` //玩家阵营死亡人数
	KillCnt      int32  `json:"kill_cnt"`       //击杀人数
	DeathCnt     int32  `json:"death_cnt"`      //死亡次数
	DamageCnt    int32  `json:"damage_cnt"`     //输出伤害
	BedamageCnt  int32  `json:"bedamage_cnt"`   //承受伤害
	GoalCnt      int32  `json:"goal_cnt"`       //进球数
	PetPart1Id   int32  `json:"pet_part1_id"`   //宠物部件1ID
	PetPart1Lv   int32  `json:"pet_part1_lv"`   //宠物部件1品质
	PetPart2Id   int32  `json:"pet_part2_id"`   //宠物部件2ID
	PetPart2Lv   int32  `json:"pet_part2_lv"`   //宠物部件2品质
	PetPart3Id   int32  `json:"pet_part3_id"`   //宠物部件3ID
	PetPart3Lv   int32  `json:"pet_part3_lv"`   //宠物部件3品质
	PetPart4Id   int32  `json:"pet_part4_id"`   //宠物部件4ID
	PetPart4Lv   int32  `json:"pet_part4_lv"`   //宠物部件4品质
	BeginTime    string `json:"begin_time"`     //战斗开始时间 Y-m-d H:i:s
	EndTime      string `json:"end_time"`       //战斗结束时间
}

func (g *BattleUserLog) GetEventName() names.EventName {
	return "battle_user_log"
}

func (g *BattleUserLog) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

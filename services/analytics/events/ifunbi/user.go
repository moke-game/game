package ifunbi

import (
	"github.com/moke-game/game/services/analytics/names"
)

type LoginEvent struct {
	CurrentGold    string `json:"current_gold"`     //当前金币数量	当前金币数量	Double
	CurrentDiamond string `json:"current_diamond"`  //当前钻石数量	当前星星数量	Double
	LastLoginTime  string `json:"last_login_time"`  //最后登录时间		时间戳
	RegisterTime   string `json:"register_time"`    //账号注册时间		时间戳		该用户的第一条包含账户 ID 的数据（含事件和用户属性数据）入库时，该条数据#time 字段的时间
	FirstLoginTime string `json:"first_login_time"` //首次登录时间		时间戳
	LoginMode      string `json:"login_mode"`       //登录方式	区分访客登录，Facebook，Google，apple账号登录，Email等
	Tel            string `json:"tel"`              //用户手机号
}

func (g *LoginEvent) Name() names.EventName {
	return names.BI_LOGIN
}

type RegisterEvent struct {
	CurrentGold    string `json:"current_gold"`     //当前金币数量	当前金币数量	Double
	CurrentDiamond string `json:"current_diamond"`  //当前钻石数量	当前星星数量	Double
	RegisterTime   string `json:"register_time"`    //账号注册时间		时间戳		该用户的第一条包含账户 ID 的数据（含事件和用户属性数据）入库时，该条数据#time 字段的时间
	FirstLoginTime string `json:"first_login_time"` //首次登录时间		时间戳
	LoginMode      string `json:"login_mode"`       //登录方式	区分访客登录，Facebook，Google，apple账号登录，Email等
}

func (g *RegisterEvent) Name() names.EventName {
	return names.BI_Register
}

type LogoutEvent struct {
	LastLoginTime string `json:"last_login_time"` //首次登录时间		时间戳
	Duration      string `json:"duration"`        //在线时长 单位：秒，玩家本次从登陆到退出的游戏在线时长
}

func (g *LogoutEvent) Name() names.EventName {
	return names.BI_Logout
}

type GuideEvent struct {
	GuideStep string `json:"guide_step"` //当前所在教程步数	项目组需要提供新手引导步骤
	GuideName string `json:"guide_name"` //当前所在教程名
}

func (g *GuideEvent) Name() names.EventName {
	return names.BI_Guide
}

type OnlineCntEvent struct {
	TimeStamp   string `json:"timeStamp"`    //记录时间点 毫秒级时间
	OnlineCount string `json:"online_count"` //实时在线玩家数量
	Rids        string `json:"rids"`         //如果在线人数不能区分渠道，平台，国家，区服时，需要传目前在线的用户ID的集合 "rids":"[1032755, 1028781]"
}

func (g *OnlineCntEvent) Name() names.EventName {
	return names.BI_OnlineCnt
}

package names

type EventName string

func (e EventName) String() string {
	return string(e)
}

const (
	GameMoneyChangedEvent        EventName = "GameMoneyChangedEvent"        //代币变更 不包括现金
	GameRubyChangedEvent         EventName = "GameRubyChangedEvent"         //钻石变更
	GamePropertyChangedEvent     EventName = "GamePropertyChangedEvent"     //道具变更
	GameShopPropertyChangedEvent EventName = "GameShopPropertyChangedEvent" //商城道具变更
	GamePlayerRechargeEvent      EventName = "GamePlayerRechargeEvent"      //充值记录
	GameOnlineAmountEvent        EventName = "GameOnlineAmountEvent"        //在线人员
	GameOnlineDurationEvent      EventName = "GameOnlineDurationEvent"      //用户退出
	GamePlayerLoginRecordEvent   EventName = "GamePlayerLoginRecordEvent"   //用户登录
	GameTaskMovedEvent           EventName = "GameTaskMovedEvent"           //任务变更
	GameCreateRoleEvent          EventName = "GameCreateRoleEvent"          //创建校色 = 注册事件
	BI_LOGIN                     EventName = "login"                        //用户登录
	BI_Register                  EventName = "register"                     //用户注册
	BI_Logout                    EventName = "exit"                         //用户退出
	BI_Guide                     EventName = "guide"                        //新手引导
	BI_OnlineCnt                 EventName = "onlinecnt"                    //在线人数
	BI_Payment                   EventName = "payment"                      //充值
	BI_Resource                  EventName = "resource_change"              //资源变动
)

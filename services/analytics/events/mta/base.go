package mta

import (
	pb "github.com/moke-game/platform/api/gen/analytics"
)

type Base struct {
	Uuid         string `json:"uuid"` //事件唯一标识 消息的uniq_id 注意是32位的，不带“-”中划线,事件可能重复上报，但是平台需要防止重复统计,e.g. 123e4567e89b12d3a456426655440000
	Category     string `json:"ctgr"` //事件类别 如：ActivatedEvent
	Label        string `json:"labl"` //事件类别
	CreatedAt    string `json:"crat"` //事件发生时间,格式符合ISO 8601, 如：2019-01-14T11:04:46.500+0800
	ClientId     string `json:"clid"` //客户端的唯一标识 appid 由UPS用户支付中心分配
	Umid         string `json:"umid"` //APP安装实例唯一编号 设备ID 必填，umid获取机制：游戏客户端通过SDK提供的API获取，针对GMS类型的数据，游戏客户端传给游戏后端，游戏后端写log。注意：同一个设备打的MTA所有事件中的umid一致。安卓获取方法：Plugins.getData().getUmid()，参考文档；IOS获取方法：[CJSDKDataGA getGameUmid]，参考文档；
	GlobalUserId string `json:"gusd"` //全局唯一用户ID uid 用户登录后就必填。为了防止可能切换登录帐号，需要确定每次事件的所有者，此为系统的全局唯一用户ID（用户中心 全局唯一用户ID。针对聚合包，则为聚合账号ID；针对包，则为全局用户ID）
	UserId       string `json:"usid"` //登录用户的唯一标识 uid 用户登录后就必填。为了防止可能切换登录帐号，需要确定每次事件的所有者，此为系统的用户ID（SDK用户ID或聚合用户ID或UPS用户ID）
	GameIp       string `json:"gmip"` //客户端玩家IP 服务端自动获取，不需要SDK或者APP传递
	ChannelId    string `json:"cchd"` //渠道ID 必须和聚合的渠道ID一一对应，渠道ID由聚合系统来定。如果子事件中有渠道ID字段，则跟这个一样，只需填充这个即可，其他不填
	//-------- 基本上各个时间都有的参数
	GameId         string `json:"gmid"` //游戏唯一标识 游戏ID，有BOSS系统分配
	ChannelName    string `json:"clnm"` //渠道名称，写渠道ID即可，渠道ID必须和聚合的渠道ID一一对应，渠道ID由聚合系统来定
	GamePlatformId string `json:"gpid"` //大区id，来源于BOSS系统，必须按照BOSS系统的大区规则
	GameServerId   string `json:"gsid"` //区服id，来源于BOSS系统，必须按照BOSS系统的区服规则
	RoleId         string `json:"rlid"` //角色ID
	RoleName       string `json:"rlnm"` //角色昵称
	Level          string `json:"lvl"`  //等级
}

func (m *Base) Platform() pb.DeliveryType {
	return pb.DeliveryType_Local
}

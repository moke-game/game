package mta

import (
	"encoding/json"

	pb "github.com/moke-game/platform/api/gen/analytics"

	"github.com/moke-game/game/services/analytics/names"
)

type GameOnlineAmountEvent struct {
	Uuid         string `json:"uuid"` //事件唯一标识 消息的uniq_id 注意是32位的，不带“-”中划线,事件可能重复上报，但是平台需要防止重复统计,e.g. 123e4567e89b12d3a456426655440000
	Category     string `json:"ctgr"` //事件类别 如：ActivatedEvent
	Label        string `json:"abl"`  //事件类别
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
	GameServerName string `json:"gsnm"` //区服名称
	Online         string `json:"onli"` //当前在线人数(在线日志每1分钟记录一次，必须是1分钟记录一次)
}

func (m *GameOnlineAmountEvent) Platform() pb.DeliveryType {
	return pb.DeliveryType_Local
}

func (g *GameOnlineAmountEvent) GetEventName() names.EventName {
	name := names.GameOnlineAmountEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GameOnlineAmountEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

type GameOnlineDurationEvent struct {
	Base
	LoginTime     string `json:"lgti"` //登录时间,格式为：2019-01-14T11:04:46.500+0800
	LoginOut      string `json:"lgou"` //登出时间,格式为：2019-01-14T11:04:46.500+0800
	TimeLong      string `json:"telg"` //本次登录登出在线总时长，单位：min（分钟）
	ChannelUserId string `json:"cuid"` //渠道用户ID，游戏研发接入聚合系统时登录验证返回的sdkUserID，即为渠道用户ID
}

func (g *GameOnlineDurationEvent) GetEventName() names.EventName {
	name := names.GameOnlineDurationEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GameOnlineDurationEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

type GamePlayerLoginRecordEvent struct {
	Base
	Idfa          string `json:"idfa"` //设备IDFA，IDFA只有IOS设备能获取，故IOS设备必填；安卓设备填空字符串
	Caid          string `json:"caid"` //设备CAID，CAID只有IOS设备能获取
	Imei          string `json:"imei"` //设备IMEI 设备IMEI，IMEI只有安卓设备能获取，必填；IOS设备填空字符串
	GameMac       string `json:"gmac"` //玩家手持设备为IOS的，获取此数据
	GameSn        string `json:"gmsn"` //game_sn 手持设备的产品序列号，必须获取
	VipLevel      string `json:"vpvl"` //VIP等级
	Diamond       string `json:"dmnd"` //钻石数量或者元宝数量，即为游戏内能交易的货币
	GameCoin      string `json:"gmcn"` //游戏币数量
	Guide         string `json:"gude"` //新手引导状态存放
	OperateType   string `json:"optp"` //1：登录，2：正常退出登录，3：背景运行退出登录，4：杀进程退出登录
	LoginTime     string `json:"lgtm"` //登录时间,格式如：2019-01-14T11:04:46.500+0800
	SessionId     string `json:"ssid"` //当前用户登录登出的唯一会话ID，长度必须是32位。一个会话标识只能对应唯一的实名用户，一个实名用户可以拥有多个会话标识；同一用户单次游戏会话中，上下线动作必须使用同一会话标识上报。备注：会话标识仅标识一次用户会话，生命周期仅为一次上线和与之匹配的一次下线，不会对生命周期之外的任何业务有任何影响
	BattleValue   string `json:"btvl"` //战力
	ChannelUserId string `json:"cuid"` //渠道用户ID。针对聚合包（华为、小米、vivo等，即接入聚合SDK），渠道用户ID就是聚合系统时登录验证返回的sdkUserID，即为渠道用户ID，这些在接入聚合时都会返回给研发，研发需要存入；针对包（接入安卓SDK和IOS SDK的包），渠道用户ID存入账号ID即可。
	IsFirst       string `json:"isft"` //是否为第一次，在登录表中用来标记是否为第1次登录（账号注册）, 0:否, 1:是
	Zxpi          string `json:"zxpi"` //中宣部实名认证唯一ID
	UserType      string `json:"ustp"` //用户类型：0-游客；1-非游客。默认为1
}

func (g *GamePlayerLoginRecordEvent) GetEventName() names.EventName {
	name := names.GamePlayerLoginRecordEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GamePlayerLoginRecordEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

type GameTaskMovedEvent struct {
	Base
	TaskId    string `json:"tsid"` //任务ID，保证游戏内的任务ID唯一，建议是整数
	TaskName  string `json:"tsnm"` //任务名称
	TaskLevel string `json:"tsll"` //任务枚举，即，主线任务  支线任务  活动任务  等类型
	TaskType  string `json:"tstp"` //任务种类，枚举，研发童鞋提供给技术
	Status    string `json:"sts"`  //领取结果，枚举：1-领取任务；2-完成任务
}

func (g *GameTaskMovedEvent) GetEventName() names.EventName {
	name := names.GameTaskMovedEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GameTaskMovedEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

// 充值日志
type GamePlayerRechargeEvent struct {
	Base
	Pays            string `json:"pays"` //充值方式为“支付宝”“微信支付”“易宝支付”“快钱”等，填写英文或拼音简称
	GoodsId         string `json:"gdid"` //商品ID
	IsFirstPay      string `json:"isfp"` //是否首充:是为1,否为0
	Amount          string `json:"amnt"` //充值金额，单位：分。
	Currency        string `json:"cncy"` //货币单位：默认人民币CNY，采用国际货币代码标准。针对海外游戏，如果非人民币的话，需要写明对应的货币单位
	GameCurrency    string `json:"gmcy"` //钻石数(或游戏币)记录真实获得钻石数(游戏币)（如充值30元送300钻石，则记录600钻石）
	PayDatee        string `json:"pydt"` //充值成功时间,格式为：2019-01-14T11:04:46.500+0800
	PlatformOrderId string `json:"poid"` //订单号
	ChannelOrderId  string `json:"coid"` //渠道订单号
}

func (g *GamePlayerRechargeEvent) GetEventName() names.EventName {
	name := names.GamePlayerRechargeEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GamePlayerRechargeEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

// 注册事件
type GameCreateRoleEvent struct {
	Base
	Idfa          string `json:"idfa"` //设备IDFA，IDFA只有IOS设备能获取，故IOS设备必填；安卓设备填空字符串
	Caid          string `json:"caid"` //设备CAID, CAID只有IOS设备能获取
	Imei          string `json:"imei"` //设备IMEI，IMEI只有安卓设备能获取，必填；IOS设备填空字符串
	GameMac       string `json:"gmac"` //玩家手持设备为IOS的，获取此数据
	GameSn        string `json:"gmsn"` //手持设备的产品序列号，必须获取
	SessionId     string `json:"ssid"` //与渠道接口对接握手会话信息（信息认证、登录令牌等）
	ChannelUserId string `json:"cuid"` //渠道用户ID，针对聚合包（华为、小米、vivo等，即接入聚合SDK），渠道用户ID就是聚合系统时登录验证返回的sdkUserID，即为渠道用户ID，这些在接入聚合时都会返回给研发，研发需要存入；针对包（接入安卓SDK和IOS SDK的包），渠道用户ID存入账号ID即可。
	UserType      string `json:"ustp"` //用户类型：0-游客；1-非游客。默认为1
}

func (g *GameCreateRoleEvent) GetEventName() names.EventName {
	name := names.GameCreateRoleEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GameCreateRoleEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

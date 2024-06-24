package ifunbi

import (
	"encoding/json"

	pb "github.com/moke-game/platform/api/gen/analytics"

	"github.com/moke-game/game/services/analytics/names"
)

type IEvent interface {
	Name() names.EventName
}

type Event struct {
	DeviceId       string `json:"device_id"`        //设备id		String		设备ID +“#”+ 手机卡ID，设备ID指的是android IMEI，IOS idfa。定义设备id+手机卡ID为一台唯一设备，变 动一个即为新设备。例如：ea14858065fba5ff9b2d68fc4bb55248#244f5ad638bf0939，其中 “ea14858065fba5ff9b2d68fc4bb55248”是设备ID，“244f5ad638bf0939”是手机卡ID。							客户端
	AppId          string `json:"app_id"`           //应用id		String		应用ID,该字段由数据平台接入时分配，用于区分不同游戏的数据，可联系数据平台负责人							数据部
	OSPlatform     string `json:"platform"`         //平台		Int		平台号，-1 为 pc、0 为 unity 或其他 IDE、1 为 Android、2 为 IOS、3 为 WinPhone、4 为 H5 Web 等
	GameArea       string `json:"game_area"`        //区服		String		游戏区服，不分区服时，设置为 global，分区服时设置为区服的代称。例如：S-116
	ApkVersion     string `json:"apk_version"`      //版本		String		应用包版本，例如：1.1.0，1.2.0							客户端
	Ts             string `json:"ts"`               //毫秒级时间		时间戳		毫秒级时间
	Ip             string `json:"ip"`               //ip		String		客户端公网IP地址，例如：182.150.46.53							客户端
	Event          string `json:"event"`            //事件名		String		事件名，如登陆事件，该值为login
	Source         string `json:"source"`           //数据来源		Int		区分数据来源，如1为客户端上报，2为服务器上报，3第三方上报
	AccountId      string `json:"account_id"`       //账号id		Int		账号ID，游戏服务端用户唯一标识，例如：655745
	DistinctId     string `json:"distinct_id"`      //访客id		Int		accountId与distinctId至少要传入一个，如果所有事件都是在用户登录状态下触发的，则只传入account_id是可行的，如果有事件是在未登录状态（包括注册前）触发的，则建议填入访客ID
	UserName       string `json:"user_name"`        //角色昵称		String
	MachineCode    string `json:"machine_code"`     //机器码		String		ios 取的uuid 安卓取的imei（ifun slot项目）							客户端
	FirstLoginTime string `json:"first_login_time"` //首次登录时间		时间戳
	RegisterTime   string `json:"register_time"`    //账号注册时间		时间戳		该用户的第一条包含账户 ID 的数据（含事件和用户属性数据）入库时，该条数据#time 字段的时间
	EventData      string `json:"event_data"`       //事件数据		String		根据不同的事件类型传输不同格式的数据
	CountryCode    string `json:"country_code"`     //国家代码		String
	Country        string `json:"country"`          //国家		String		用户所在国家，根据 IP 地址生成							客户端
	Province       string `json:"province"`         //省份		String		用户所在省份，根据 IP 地址生成							客户端
	City           string `json:"city"`             //城市		string		用户所在城市，根据 IP 地址生成							客户端
	Brand          string `json:"brand"`            //手机型号		String		手机型号，例如：samsung，oppo，vivo							客户端
	Network        string `json:"network"`          //网络类型		String		网络类型，例如：Wifi							客户端
	Language       string `json:"language"`         //系统语言		String		系统语言，例如：中文简体							客户端
	data           any
	//Others      string `json:"others"`       //其他		String		备用字段
	//BaseEventLog string `json:"base_event_log"` //客户端上报原始数据		String		客户端上报原始数据							客户端
	//LastLoginTime  string `json:"last_login_time"`  //最后登录时间		时间戳
	//RegisterTime   string `json:"register_time"`    //账号注册时间		时间戳		该用户的第一条包含账户 ID 的数据（含事件和用户属性数据）入库时，该条数据#time 字段的时间
	//CurrentGold    string `json:"current_gold"`     //当前金币数量	当前金币数量	Double
	//CurrentDiamond string `json:"current_diamond"`  //当前钻石数量	当前星星数量	Double
	//ChannelId    string `json:"channel_id"`   //渠道		String		如Google，AppStore，taptap、小米应用商城，华为应用商城，应用宝等 如果有马甲包需区分:Google-马甲包，Appstore-马甲包							客户端
	//FirstPayTime   string `json:"first_pay_time"`   //首次付费时间		时间戳
	//PayTimes       string `json:"pay_times"`        //累积充值次数		Int
	//Idfa            string `json:"idfa"`             //苹果设备标识		String
	//Udid            string `json:"udid"`             //安卓设备标识		String
	//ActiveTime     string `json:"active_time"`      //激活时间		时间戳		该用户的第一条数据（含事件和用户属性数据）入库时，该条数据#time 字段的时间
	//CurrentResource string `json:"current_resource"` //当前资源		String		适用于游戏资源多的项目，可以列表形式上报列表：[{resource_id:资源ID（金币，钻石等...）,resource_num:资源数量（数值）},{resource_id:"金币",resource_num:123}]
	//Manufacturer string `json:"manufacturer"` //制造商		String		手机生产厂商，例如：MacBookPro11,3							客户端
	//OsVersion    string `json:"os_version"`   //操作系统版本		String		手机操作系统版本，例如：Android OS 9							客户端
	//Operator     string `json:"operator"`     //运营商		String		运营商，例如：中国移动、联通、电信等
	//PackageId  string `json:"package_id"`  //分包		String		渠道包ID，例如: lyjlzz0001，因部分游戏存在渠道分包的现象，故采用此字段来区分不同的渠道分包							客户端
	//SessionId    string `json:"session_id"`   //会话ID		String		会话ID
	//TotalPay       string `json:"total_pay"`        //累计充值总额	充值总额	Double
}

func (m *Event) Platform() pb.DeliveryType {
	return pb.DeliveryType_SDK
}

func (g *Event) GetEventName() names.EventName {
	return names.EventName(g.Event)
}

func (g *Event) ToJson() (data []byte, err error) {
	if data, err = json.Marshal(g.data); err != nil {
		return
	}
	g.EventData = string(data)
	return json.Marshal(g)
}

func (g *Event) Data(data any) {
	g.data = data
}

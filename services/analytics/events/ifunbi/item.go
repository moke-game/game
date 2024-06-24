package ifunbi

import (
	"github.com/moke-game/game/services/analytics/names"
)

type PaymentEvent struct {
	OrderId      string `json:"order_id"`       //订单id		字符串
	ThirdOrderId string `json:"third_order_id"` //第三方订单id	第三方订单号	字符串
	PayScene     string `json:"pay_scene"`      //充值场景	充值场景，在机台内购买就用机台id，大厅内购买就用大厅位置，其他场景也用可区分的字符串来区分开	字符串
	PayLevel     string `json:"pay_1_level"`    //支付1级分类	1级分类，如商城、通行证、首充礼包、破产充值……详见游戏内设定	字符串
	GoodsId      string `json:"goods_id"`       //商品ID	项目组需要提供对应商品对照表（必给）	字符串
	GoodsName    string `json:"goods_name"`     //商品名称		字符串
	PayMethod    string `json:"pay_method"`     //支付方式	Apple，Google，支付宝，微信支付等	字符串
	PayGroup     string `json:"pay_group"`      //支付分组	Eg：实际付费，沙盒模式付费	字符串
	PayType      string `json:"pay_type"`       //支付币种	实际币种符号+实际金额，举例：玩家实际上花费6人民币买了一个礼包，需要在pay_type报RMB 6这个属性，如果是100新加坡币，就报新加坡币符号+实际金额；	字符串
	PayAmount    string `json:"pay_amount"`     //支付金额	服务器提前配置好的美金金额，比如服务器配置的礼包为0.99美金，这里就报0.99美金	数值
	GetGoods     string `json:"get_goods"`      //获得物品		列表
	IsFirstPay   string `json:"is_first_pay"`   //是否首购		布尔
}

func (g *PaymentEvent) Name() names.EventName {
	return names.BI_Payment
}

type ResourceEvent struct {
	ResourceId   string `json:"resource_id"`   //资源ID	如金币，钻石，道具等对应的物品id	字符串
	ResourceName string `json:"resource_name"` //资源名称	gold，diamond等	字符串
	ChangeType   string `json:"change_type"`   //变动类型	获取、消耗	字符串
	ChangeNum    string `json:"change_num"`    //变动数量	获取为正，消耗为负数	数值
	ChangeBefore string `json:"change_before"` //变化前数量		数值
	ChangeAfter  string `json:"change_after"`  //变化后数量		数值
	ChangeReason string `json:"change_reason"` //变动原因	通用货币消耗与获得需特别说明	字符串
}

func (g *ResourceEvent) Name() names.EventName {
	return names.BI_Resource
}

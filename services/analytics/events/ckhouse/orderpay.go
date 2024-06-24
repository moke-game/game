package ckhouse

import (
	"encoding/json"

	"github.com/moke-game/game/services/analytics/names"
)

type OrderPay struct {
	Base
	UID          string // '用户ID',
	OrderId      string // '平台订单ID',
	TransId      string // '渠道订单ID',
	Source       int32  // '平台 0:u3d 1:ios 2:android',
	PayType      int32  // '支付方式 -1:免费 0:货币 >0:代币',
	ActId        int32  // '活动表ID',
	ShopId       int32  // 'shop表ID ',
	ShopType     int32  // '商品类型 1，钻石充值 >1:其它礼包',
	ShopActType  int32  // '商城标签',
	ItemBoxId    int32  // '道具礼包',
	PurchaseId   int32  // 'purchase表ID',
	PurchaseName string // '渠道商品名称'
	RechargeId   string // '渠道档位名称',
	RechargeType int32  // '消费处理类型',
	Price        int32  // '价格',
	SandBox      bool   // '是否沙盒',
	IsFirst      bool   // '是否首次',
	CheckStatus  int32  // '验单状态',
	FinishStatus int32  // '发货状态',
	FinishTime   string // '发货时间',
	CreateTime   string // '创建时间',
}

func (g *OrderPay) GetEventName() names.EventName {
	return "order_pay"
}

func (g *PetAct) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

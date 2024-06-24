package mta

import (
	"encoding/json"

	"github.com/moke-game/game/services/analytics/names"
)

// 货币变动(产出|消耗)
type GameMoneyChangedEvent struct {
	Base
	MoneyType       string `json:"mytp"` //货币类型（写英文或者数字）。枚举对应需研发提供
	MoneyName       string `json:"mynm"` //货币名称
	HavingAmount    string `json:"hvmt"` //变化数量，当类型sts=get，为获得数量；当sts=del，为消耗数量
	RemainingAmount string `json:"rnmt"` //剩余数量
	ChangeType      string `json:"chtp"` //变化类型：get-获得；del-消耗
	ChangeWay       string `json:"chwy"` //来源途径：云购\充值\拍卖行\商店等（写英文或者数字）。枚举对应需研发提供
}

func (g *GameMoneyChangedEvent) GetEventName() names.EventName {
	name := names.GameMoneyChangedEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GameMoneyChangedEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

// 钻石变动(产出|消耗)
type GameRubyChangedEvent struct {
	Base
	MoneyType       string `json:"mytp"` //货币类型（写英文或者数字）。枚举对应需研发提供
	MoneyName       string `json:"mynm"` //货币名称
	HavingAmount    string `json:"hvmt"` //变化数量，当类型sts=get，为获得数量；当sts=del，为消耗数量
	RemainingAmount string `json:"rnmt"` //剩余数量
	ChangeType      string `json:"chtp"` //变化类型：get-获得；del-消耗
	ChangeWay       string `json:"chwy"` //来源途径：云购\充值\拍卖行\商店等（写英文或者数字）。枚举对应需研发提供
}

func (g *GameRubyChangedEvent) GetEventName() names.EventName {
	name := names.GameRubyChangedEvent
	g.Category = string(name)
	g.Label = g.Category
	return name

}

func (g *GameRubyChangedEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

// 道具变动(产出|消耗)
type GamePropertyChangedEvent struct {
	Base
	PropId          string `json:"ppid"` //道具ID
	PropName        string `json:"ppnm"` //道具名称
	HavingAmount    string `json:"hvmt"` //变化数量，当类型sts=get，为获得数量；当sts=del，为消耗数量
	RemainingAmount string `json:"rnmt"` //剩余数量
	Status          string `json:"sts"`  //类型：get-获得；del-消耗
	Way             string `json:"way"`  //来源途径：云购\充值\拍卖行\商店等（写英文或者数字）。枚举对应需研发提供
}

func (g *GamePropertyChangedEvent) GetEventName() names.EventName {
	name := names.GamePropertyChangedEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GamePropertyChangedEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

// 商城道具变动(产出|消耗)
type GameShopPropertyChangedEvent struct {
	Base
	ProductId       string `json:"prid"` //道具ID
	PropId          string `json:"ppid"` //道具ID
	PropName        string `json:"ppnm"` //道具名称
	HavingAmount    string `json:"hvmt"` //变化数量，当类型sts=get，为获得数量；当sts=del，为消耗数量
	RemainingAmount string `json:"rnmt"` //剩余数量
	Status          string `json:"sts"`  //类型：get-获得；del-消耗
	Way             string `json:"way"`  //商店种类：（写英文或者数字）。枚举对应需研发提供，商店如商城商店、荣誉商店等商店类型
	MoneyType       string `json:"mytp"` //货币类型（写英文或者数字）。枚举对应需研发提供
	MoneyAmount     string `json:"mymt"` //货币数量
}

func (g *GameShopPropertyChangedEvent) GetEventName() names.EventName {
	name := names.GameShopPropertyChangedEvent
	g.Category = string(name)
	g.Label = g.Category
	return name
}

func (g *GameShopPropertyChangedEvent) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

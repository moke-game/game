package ckhouse

import (
	"encoding/json"

	"github.com/moke-game/game/services/analytics/names"
)

type UserBase struct {
	Base
	UID         string `json:"uid"`
	OpenId      string `json:"open_id"`
	DeviceId    string `json:"device_id"`
	CreateTime  string `json:"create_time"`
	CreateDay   string `json:"create_day"`
	IP          string `json:"ip"` //ip		String
	System      string `json:"system"`
	CountryCode string `json:"country_code"` //国家代码		String
	Country     string `json:"country"`      //国家		String		用户所在国家，根据 IP 地址生成							客户端
	Province    string `json:"province"`     //省份		String		用户所在省份，根据 IP 地址生成							客户端
	City        string `json:"city"`         //城市		string		用户所在城市，根据 IP 地址生成							客户端
	Brand       string `json:"brand"`        //手机型号		String		手机型号，例如：samsung，oppo，vivo							客户端
	Network     string `json:"network"`      //网络类型		String		网络类型，例如：Wifi							客户端
	Language    string `json:"language"`     //系统语言		String		系统语言，例如：中文简体							客户端
	//Version     string `json:"version"`      //游戏版本
}

type UserLogin struct {
	UserBase
	IsGuest    int8 `json:"is_guest"`
	IsRegister int8 `json:"is_register"`
}

func (g *UserLogin) GetEventName() names.EventName {
	return "user_login"
}

func (g *UserLogin) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

type UserReg struct {
	UserBase
	IsBind  int8 `json:"is_bind"`
	IsGuest int8 `json:"is_guest"`
}

func (g *UserReg) GetEventName() names.EventName {
	return "user_reg"
}

func (g *UserReg) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

type UserDeviceReg struct {
	UserBase
	IsGuest int8 `json:"is_guest"`
}

func (g *UserDeviceReg) GetEventName() names.EventName {
	return "user_device_reg"
}

func (g *UserDeviceReg) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

type UserExit struct {
	UserBase
	LoginTime   string `json:"login_time"`   //登录时间
	LoginSecond int32  `json:"login_second"` //登录持续秒数
}

func (g *UserExit) GetEventName() names.EventName {
	return "user_exit"
}

func (g *UserExit) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

type UserOnlineCnt struct {
	Base
	Pod        string `json:"pod"`
	OnlineCnt  int    `json:"online_cnt"`
	CreateDay  string `json:"create_day"`
	CreateTime string `json:"create_time"` //开始时间 Y-m-d H:i:s
}

func (g *UserOnlineCnt) GetEventName() names.EventName {
	return "user_online_cnt"
}

func (g *UserOnlineCnt) ToJson() (data []byte, err error) {
	return json.Marshal(g)
}

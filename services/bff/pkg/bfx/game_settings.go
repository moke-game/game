package bfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type GameSettingParams struct {
	fx.In

	GameUrl      string `name:"GameUrl"`
	HostName     string `name:"HostName"`
	IFunSDKAppId string `name:"IFunSDKAppId"`
	IFunSDKUrl   string `name:"IFunSDKUrl"`
	BossGmHost   string `name:"boss_gm_host"`
	BossGmGameId string `name:"boss_gm_game_id"`
}

type GameSettingsResult struct {
	fx.Out

	GameUrl      string `name:"GameUrl" envconfig:"GAME_URL" default:"localhost:8081"`
	HostName     string `name:"HostName" envconfig:"HOST_NAME" default:"bff"`
	IFunSDKAppId string `name:"IFunSDKAppId" envconfig:"IFUN_SDK_APPID" default:"10008"`
	// product url: https://account.ifunservice.com/ums/
	IFunSDKUrl   string `name:"IFunSDKUrl" envconfig:"IFUN_SDK_URL" default:"http://prod-game.gameloyo.com/ums"`
	BossGmHost   string `name:"boss_gm_host" envconfig:"BOSS_GM_HOST" default:"http://192.168.90.20:8383"`
	BossGmGameId string `name:"boss_gm_game_id" envconfig:"BOSS_GM_GAME_ID" default:"fr"`
}

func (g *GameSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out GameSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)

package rfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/utility"
)

type RoomSettingParams struct {
	fx.In

	RoomUrl       string `name:"RoomUrl"`
	BattleRoomMax int32  `name:"BattleRoomMax"`
	RoomHostName  string `name:"Hostname"`
}

type RoomSettingsResult struct {
	fx.Out

	RoomUrl       string `name:"RoomUrl" envconfig:"ROOM_URL" default:"localhost:8888"`
	BattleRoomMax int32  `name:"BattleRoomMax" envconfig:"BATTLE_ROOM_Max" default:"10"`
	RoomHostName  string `name:"Hostname"envconfig:"HOSTNAME" default:"room"`
}

func (g *RoomSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out RoomSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)

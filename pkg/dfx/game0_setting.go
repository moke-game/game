package dfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/utility"
)

// SettingsParams  you can customize it as your need
type SettingsParams struct {
	fx.In

	GameUrl string `name:"GameUrl"`
	DbName  string `name:"DbName"`
}

type SettingsResult struct {
	fx.Out

	GameUrl string `name:"GameUrl" envconfig:"GAME_URL" default:"localhost:8081"`
	DbName  string `name:"DbName" envconfig:"DB_NAME" default:"game"`
}

func (g *SettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

// SettingsModule  config your app settings
var SettingsModule = fx.Provide(
	func() (out SettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)

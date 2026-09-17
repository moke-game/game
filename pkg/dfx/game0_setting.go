package dfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

// SettingsParams is injected game config.
type SettingsParams struct {
	fx.In

	GameUrl string `name:"GameUrl"`
	DbName  string `name:"DbName"`
}

// SettingsResult is loaded from GAME_URL / DB_NAME.
type SettingsResult struct {
	fx.Out

	GameUrl string `name:"GameUrl" envconfig:"GAME_URL" default:"localhost:8081"`
	DbName  string `name:"DbName" envconfig:"DB_NAME" default:"game"`
}

// LoadFromEnv fills settings from the process environment.
func (g *SettingsResult) LoadFromEnv() error {
	return utility.Load(g)
}

// SettingsModule provides game settings from the environment.
var SettingsModule = ProvideFromEnv[SettingsResult]()

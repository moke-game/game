package dfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
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

// SettingsModule provides game settings from the environment.
var SettingsModule = platformfx.ProvideFromEnv[SettingsResult]()

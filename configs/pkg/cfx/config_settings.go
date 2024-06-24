package cfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type ConfigSettingParams struct {
	fx.In
	ConfigPath     string `name:"ConfigPath"`
	WordsPath      string `name:"WordsPath"`
	NavMeshPath    string `name:"NavMeshPath"`
	MapPath        string `name:"MapPath"`
	MatchGroupSize int32  `name:"MatchGroupSize"`
	AIPath         string `name:"AIPath"`
	GooglePayConf  string `name:"GooglePayConf"`
}

type ConfigSettingsResult struct {
	fx.Out

	ConfigPath     string `name:"ConfigPath" envconfig:"CONFIG_PATH" default:"./configs/data"`
	WordsPath      string `name:"WordsPath" envconfig:"WORDS_PATH" default:"./configs/wordsfilter"`
	NavMeshPath    string `name:"NavMeshPath" envconfig:"NAVMESH_PATH" default:"./configs/maps"`
	MapPath        string `name:"MapPath" envconfig:"MAP_PATH" default:"./configs/maps"`
	MatchGroupSize int32  `name:"MatchGroupSize" envconfig:"Match_Group_Size" default:"3"`
	AIPath         string `name:"AIPath" envconfig:"AI_PATH" default:"./configs/ai"`
	GooglePayConf  string `name:"GooglePayConf" envconfig:"GOOGLE_PAY_CONF" default:"./configs/iap/googlepay/finalrumble-101-5dbc5dc04183.json"`
}

func (g *ConfigSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out ConfigSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)

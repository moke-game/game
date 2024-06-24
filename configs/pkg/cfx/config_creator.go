package cfx

import (
	"encoding/json"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	cfg "github.com/moke-game/game/configs/code"
	"github.com/moke-game/game/configs/customs"
	wordsfilter2 "github.com/moke-game/game/services/common/wordsfilter"
)

type ConfigsResult struct {
	fx.Out
	*cfg.Tables
	Nickname    *customs.Nicknames
	HeroLevelUp *customs.CharacterLevelUp
	PlayRewards *customs.PlayRewards
	RankConfigs *customs.RankConfigs
	DropLibrary *customs.DropLibrary
}

type ConfigsParams struct {
	fx.In
	*cfg.Tables
	Nickname    *customs.Nicknames
	HeroLevelUp *customs.CharacterLevelUp
	PlayRewards *customs.PlayRewards
	RankConfigs *customs.RankConfigs
	DropLibrary *customs.DropLibrary
}

func createLoader(dir string) cfg.JsonLoader {
	return func(file string) ([]map[string]interface{}, error) {
		fPath := path.Join(dir, file+".json")
		if bytes, err := os.ReadFile(fPath); err != nil {
			return nil, errors.Wrap(err, "read file error, file: "+fPath)
		} else {
			jsonData := make([]map[string]interface{}, 0)
			if err = json.Unmarshal(bytes, &jsonData); err != nil {
				return nil, errors.Wrap(err, "unmarshal json error, file: "+fPath)
			}
			return jsonData, nil
		}
	}
}

func (cf *ConfigsResult) loadWords(path string) error {
	viper.SetConfigName("words")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	words := make([]string, 0)
	if err := viper.UnmarshalKey("words", &words); err != nil {
		return err
	}
	wordsfilter2.Load(words)
	return nil
}

func (cf *ConfigsResult) Execute(path string, wPath string, navPath string, mapPath string, aiPath string) (err error) {
	table, err := cfg.NewTables(createLoader(path))
	if err != nil {
		return err
	}
	cf.Tables = table
	cf.customize()
	if err := cf.loadWords(wPath); err != nil {
		return err
	}

	return nil
}

func (cf *ConfigsResult) customize() {
	cf.Nickname = customs.CreateNickNames()
	cf.Nickname.Init(cf.Tables.TblNickName)
	cf.HeroLevelUp = customs.CreateCharacterLevelUp()
	cf.HeroLevelUp.Init(cf.Tables.TblCharacterLvup)
	cf.PlayRewards = customs.CreatePlayRewards()
	cf.PlayRewards.Init(cf.Tables.TblPlaySelect)

	cf.RankConfigs = customs.CreateRankConfigs()
	cf.RankConfigs.Init(cf.Tables.TblRankConfig)

	cf.DropLibrary = customs.CreateDropLibrary()
	cf.DropLibrary.Init(cf.Tables.TblDropLibrary, cf.TblDropGroup)
}

var ConfigsCreator = fx.Provide(
	func(
		l *zap.Logger,
		s ConfigSettingParams,
	) (out ConfigsResult, err error) {
		err = out.Execute(s.ConfigPath, s.WordsPath, s.NavMeshPath, s.MapPath, s.AIPath)
		return
	},
)

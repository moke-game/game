package gmfx

import (
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
)

type GmSettingParams struct {
	fx.In
	Name     string `name:"GmName"`
	GmUrl    string `name:"GmUrl"`
	AESKey   string `name:"AESKey"`
	MysqlURL string `name:"MysqlURL"`
}

type GmSettingResult struct {
	fx.Out
	Name     string `name:"GmName" envconfig:"GM_NAME" default:"gm"`
	GmUrl    string `name:"GmUrl" envconfig:"GM_URL" default:"localhost:8081"`
	AESKey   string `name:"AESKey" envconfig:"AES_KEY" default:"CTeGahnbQWfAr5hW"`
	MysqlURL string `name:"MysqlURL" envconfig:"GM_MYSQL_URL" default:""` //root:123qweasd@tcp(192.168.110.67:3306)/fr-admin?charset=utf8mb4&parseTime=true
}

func (l *GmSettingResult) LoadFromEnv() (err error) {
	err = utility.Load(l)
	return
}

var GmSettingsModule = fx.Provide(
	func() (out GmSettingResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)

var GmMysqlModule = fx.Provide(
	func(params GmSettingParams) (out ofx.GormDriverResult, err error) {
		if params.MysqlURL != "" {
			out.Dialector = mysql.New(mysql.Config{DSN: params.MysqlURL})
		}
		return
	},
)

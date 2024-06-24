package gmfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/moke-game/game/api/gen/gm"
)

type GmClientParams struct {
	fx.In

	GmClient pb.GMServiceClient `name:"GmClient"`
}

type GmClientResult struct {
	fx.Out

	GmClient pb.GMServiceClient `name:"GmClient"`
}

func NewGmClient(host string, sSetting sfx.SecuritySettingsParams) (pb.GMServiceClient, error) {
	if sSetting.MTLSEnable {
		if conn, err := tools.DialWithSecurity(
			host,
			sSetting.ClientCert,
			sSetting.ClientKey,
			sSetting.ServerName,
			sSetting.ServerCaCert,
		); err != nil {
			return nil, err
		} else {
			return pb.NewGMServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewGMServiceClient(conn), nil
		}
	}
}

var GmClientModule = fx.Provide(
	func(
		setting GmSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out GmClientResult, err error) {
		if cli, e := NewGmClient(setting.GmUrl, sSetting); e != nil {
			err = e
		} else {
			out.GmClient = cli
		}
		return
	},
)

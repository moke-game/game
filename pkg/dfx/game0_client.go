package dfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb "github.com/moke-game/game/api/gen/game0/api"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
)

type DemoClientParams struct {
	fx.In

	Game0Client pb.Game0ServiceClient `name:"Game0Client"`
}

type DemoClientResult struct {
	fx.Out

	Game0Client pb.Game0ServiceClient `name:"Game0Client"`
}

func NewDemoClient(host string, logger *zap.Logger, sSetting sfx.SecuritySettingsParams) (pb.Game0ServiceClient, error) {
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
			return pb.NewGame0ServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewGame0ServiceClient(conn), nil
		}
	}
}

// Game0ClientModule  you can inject it to other module
var Game0ClientModule = fx.Provide(
	func(
		setting SettingsParams,
		sSetting sfx.SecuritySettingsParams,
		logger *zap.Logger,
	) (out DemoClientResult, err error) {
		if cli, e := NewDemoClient(setting.GameUrl, logger, sSetting); e != nil {
			err = e
		} else {
			out.Game0Client = cli
		}
		return
	},
)

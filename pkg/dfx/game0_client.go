package dfx

import (
	"go.uber.org/fx"

	pb "github.com/moke-game/game/api/gen/game0/api"

	"github.com/gstones/moke-kit/server/pkg/sfx"
)

// DemoClientParams is the injected game0 gRPC client.
type DemoClientParams struct {
	fx.In

	Game0Client pb.Game0ServiceClient `name:"Game0Client"`
}

// DemoClientResult provides a game0 gRPC client.
type DemoClientResult struct {
	fx.Out

	Game0Client pb.Game0ServiceClient `name:"Game0Client"`
}

// NewDemoClient dials the game service with the process TLS settings.
func NewDemoClient(host string, sSetting sfx.SecuritySettingsParams) (pb.Game0ServiceClient, error) {
	return NewClient(host, sSetting, pb.NewGame0ServiceClient)
}

// Game0ClientModule provides an outbound Game0Service client.
var Game0ClientModule = fx.Provide(
	func(setting SettingsParams, sSetting sfx.SecuritySettingsParams) (out DemoClientResult, err error) {
		cli, err := NewDemoClient(setting.GameUrl, sSetting)
		if err != nil {
			return out, err
		}
		out.Game0Client = cli
		return
	},
)

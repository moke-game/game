package dfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/moke-game/platform/pkg/platformfx"

	pb "github.com/moke-game/game/api/gen/game0/api"
)

// Game0ClientParams is the injected game0 gRPC client.
type Game0ClientParams struct {
	fx.In

	Game0Client pb.Game0ServiceClient `name:"Game0Client"`
}

// Game0ClientResult provides a game0 gRPC client.
type Game0ClientResult struct {
	fx.Out

	Game0Client pb.Game0ServiceClient `name:"Game0Client"`
}

// NewGame0Client dials the game service with the process TLS settings.
func NewGame0Client(host string, sec sfx.SecuritySettingsParams) (pb.Game0ServiceClient, error) {
	return platformfx.NewClient(host, sec, pb.NewGame0ServiceClient)
}

// Game0ClientModule provides an outbound Game0Service client.
var Game0ClientModule = fx.Provide(
	func(setting SettingsParams, sec sfx.SecuritySettingsParams) (out Game0ClientResult, err error) {
		cli, err := NewGame0Client(setting.GameUrl, sec)
		if err != nil {
			return out, err
		}
		out.Game0Client = cli
		return
	},
)

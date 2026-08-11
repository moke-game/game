package modules

import (
	"go.uber.org/fx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"

	"github.com/moke-game/game/internal/services/game0"
	"github.com/moke-game/game/pkg/dfx"
)

// GrpcModule starts game gRPC with platform AuthMiddlewareModule by default.
var GrpcModule = fx.Module("grpcService",
	dfx.SettingsModule,
	auth.AuthMiddlewareModule,
	game0.ServiceModule,
)

// HttpModule starts game gRPC + HTTP gateway with AuthMiddlewareModule.
var HttpModule = fx.Module("httpService",
	dfx.SettingsModule,
	auth.AuthMiddlewareModule,
	game0.ServiceModule,
)

// TcpModule starts game TCP (zinx) service.
var TcpModule = fx.Module("tcpService",
	dfx.SettingsModule,
	game0.ServiceModule,
)

// AllModule starts all game transports.
// Pair with platform auth.AuthMiddlewareModule (thin) or auth.AuthAllModule
// (aggregate) in main — do not embed auth here to avoid duplicate providers.
var AllModule = fx.Module("allService",
	dfx.SettingsModule,
	game0.ServiceModule,
)

// GrpcClientModule provides a game gRPC client.
var GrpcClientModule = fx.Module("grpcClient",
	dfx.SettingsModule,
	dfx.Game0ClientModule,
)

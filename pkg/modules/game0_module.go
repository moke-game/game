package modules

import (
	"go.uber.org/fx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"

	"github.com/moke-game/game/internal/services/game0"
	"github.com/moke-game/game/pkg/dfx"
)

// GrpcModule starts game gRPC only, with platform AuthMiddlewareModule.
var GrpcModule = fx.Module("grpcService",
	dfx.SettingsModule,
	auth.AuthMiddlewareModule,
	game0.ServiceInstance,
	game0.GrpcService,
)

// HttpModule starts game gRPC + HTTP gateway with AuthMiddlewareModule.
var HttpModule = fx.Module("httpService",
	dfx.SettingsModule,
	auth.AuthMiddlewareModule,
	game0.ServiceInstance,
	game0.GrpcService,
	game0.HttpService,
)

// TcpModule starts game TCP (zinx) only.
// Note: zinx does not use gRPC AuthMiddleware; protect at network edge if needed.
var TcpModule = fx.Module("tcpService",
	dfx.SettingsModule,
	game0.ServiceInstance,
	game0.TcpService,
)

// AllModule starts all game transports (shared Service instance).
// Pair with platform auth.AuthMiddlewareModule (thin) or auth.AuthAllModule
// (aggregate) in main — do not embed auth here to avoid duplicate providers.
var AllModule = fx.Module("allService",
	dfx.SettingsModule,
	game0.ServiceInstance,
	game0.GrpcService,
	game0.HttpService,
	game0.TcpService,
)

// GrpcClientModule provides a game gRPC client.
var GrpcClientModule = fx.Module("grpcClient",
	dfx.SettingsModule,
	dfx.Game0ClientModule,
)

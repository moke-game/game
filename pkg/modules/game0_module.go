package modules

import (
	"go.uber.org/fx"

	"github.com/moke-game/game/internal/services/game0"
	"github.com/moke-game/game/pkg/dfx"
)

// GrpcModule starts game gRPC only.
// Auth is not included — pair auth.AuthMiddlewareModule or auth.AuthAllModule in main.
var GrpcModule = fx.Module("grpcService",
	dfx.SettingsModule,
	game0.ServiceInstance,
	game0.GrpcService,
)

// HttpModule starts game gRPC + HTTP gateway.
// Auth is not included — pair auth in main the same way as GrpcModule / AllModule.
var HttpModule = fx.Module("httpService",
	dfx.SettingsModule,
	game0.ServiceInstance,
	game0.GrpcService,
	game0.HttpService,
)

// TcpModule starts game TCP (zinx) only.
//
// WARNING: zinx does not use gRPC AuthMiddleware. Callers can spoof uid.
// Opt-in for local experiments only; keep off the public network.
var TcpModule = fx.Module("tcpService",
	dfx.SettingsModule,
	game0.ServiceInstance,
	game0.TcpService,
)

// AllModule starts authenticated game transports (gRPC + HTTP gateway).
// TCP is intentionally omitted — use TcpModule / AllWithTCPModule only for
// trusted local demos. Pair with auth.AuthMiddlewareModule (thin) or
// auth.AuthAllModule (aggregate) in main.
var AllModule = fx.Module("allService",
	dfx.SettingsModule,
	game0.ServiceInstance,
	game0.GrpcService,
	game0.HttpService,
)

// AllWithTCPModule is AllModule plus unauthenticated TCP (zinx).
// Do not use on public networks.
var AllWithTCPModule = fx.Module("allServiceWithTcp",
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

package modules

import (
	"go.uber.org/fx"

	"github.com/moke-game/game/internal/services/game0"
	"github.com/moke-game/game/pkg/dfx"
)

// gameCore is the shared game instance + settings. Transport modules add
// gRPC / HTTP / TCP on top. Auth is never included here — pair it in main
// via Platform (aggregate) or PlatformClients (thin).
var gameCore = fx.Options(
	dfx.SettingsModule,
	game0.ServiceInstance,
)

// GrpcModule starts game gRPC only.
var GrpcModule = fx.Module("grpcService",
	gameCore,
	game0.GrpcService,
)

// HttpModule starts game gRPC + HTTP gateway.
var HttpModule = fx.Module("httpService",
	gameCore,
	game0.GrpcService,
	game0.HttpService,
)

// TcpModule starts game TCP (zinx) only.
//
// WARNING: zinx does not use gRPC AuthMiddleware. Callers can spoof uid.
// Opt-in for local experiments only; keep off the public network.
var TcpModule = fx.Module("tcpService",
	gameCore,
	game0.TcpService,
)

// AllModule starts game gRPC + HTTP gateway (TCP omitted by default).
// Pair with Platform (AuthAllModule) or PlatformClients (AuthMiddlewareModule).
var AllModule = fx.Module("allService",
	gameCore,
	game0.GrpcService,
	game0.HttpService,
)

// AllWithTCPModule is AllModule plus unauthenticated TCP (zinx).
// Do not use on public networks.
var AllWithTCPModule = fx.Module("allServiceWithTcp",
	gameCore,
	game0.GrpcService,
	game0.HttpService,
	game0.TcpService,
)

// GrpcClientModule provides a game gRPC client.
var GrpcClientModule = fx.Module("grpcClient",
	dfx.SettingsModule,
	dfx.Game0ClientModule,
)

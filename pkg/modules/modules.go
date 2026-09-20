package modules

import (
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"go.uber.org/fx"

	analytics "github.com/moke-game/platform/services/analytics/pkg/module"
	auth "github.com/moke-game/platform/services/auth/pkg/module"
	buddy "github.com/moke-game/platform/services/buddy/pkg/module"
	chat "github.com/moke-game/platform/services/chat/pkg/module"
	knapsack "github.com/moke-game/platform/services/knapsack/pkg/module"
	leaderboard "github.com/moke-game/platform/services/leaderboard/pkg/module"
	mail "github.com/moke-game/platform/services/mail/pkg/module"
	matchmaking "github.com/moke-game/platform/services/matchmaking/pkg/module"
	party "github.com/moke-game/platform/services/party/pkg/module"
	profile "github.com/moke-game/platform/services/profile/pkg/module"

	"github.com/moke-game/game/internal/services/game0"
	"github.com/moke-game/game/pkg/dfx"
)

// LEGO pieces for fxmain.Main(...).
//
//	fxmain.Main(modules.Aggregate) // local/dev: game + in-process platform
//	fxmain.Main(modules.Thin)      // prod-like: game + remote platform clients
//
// Swap a brick instead of rewriting main:
//
//	fxmain.Main(modules.Infra, modules.GrpcModule, modules.PlatformClients)

// Infra is what fxmain.Main does not enable by default but this game uses:
// NATS + in-process local MQ (Watch/Hi publish both) and Redis ICache.
var Infra = fx.Options(
	mfx.NatsModule,
	mfx.LocalModule,
	ofx.RedisCacheModule,
)

// gameCore is the shared game instance + settings. Transport modules add
// gRPC / HTTP / TCP on top. Auth is never included here — pair it via
// Platform (aggregate) or PlatformClients (thin).
var gameCore = fx.Options(
	dfx.SettingsModule,
	game0.ServiceInstance,
)

func withTransports(name string, transports ...fx.Option) fx.Option {
	return fx.Module(name, append([]fx.Option{gameCore}, transports...)...)
}

// GrpcModule starts game gRPC only.
var GrpcModule = withTransports("grpcService", game0.GrpcService)

// HttpModule starts game gRPC + HTTP gateway.
var HttpModule = withTransports("httpService", game0.GrpcService, game0.HttpService)

// TcpModule starts game TCP (zinx) only.
//
// WARNING: zinx does not use gRPC AuthMiddleware. Callers can spoof uid.
// Opt-in for local experiments only; keep off the public network.
var TcpModule = withTransports("tcpService", game0.TcpService)

// AllModule is the default public surface: gRPC + HTTP (no TCP).
// Pair with Platform (AuthAllModule) or PlatformClients (AuthMiddlewareModule).
var AllModule = HttpModule

// AllWithTCPModule is AllModule plus unauthenticated TCP (zinx).
// Do not use on public networks.
var AllWithTCPModule = withTransports(
	"allServiceWithTcp",
	game0.GrpcService,
	game0.HttpService,
	game0.TcpService,
)

// GrpcClientModule provides a game gRPC client.
var GrpcClientModule = fx.Module("grpcClient",
	dfx.SettingsModule,
	dfx.Game0ClientModule,
)

// Platform hosts in-process platform services (local/dev aggregate).
// AuthAllModule starts AuthService and public AuthMiddleware.
var Platform = fx.Options(
	auth.AuthAllModule,
	analytics.AnalyticsModule,
	profile.ProfileModule,
	knapsack.KnapsackModule,
	mail.MailModule,
	party.PartyModule,
	buddy.BuddyModule,
	leaderboard.LeaderboardModule,
	chat.ChatModule,
	matchmaking.MatchmakingModule,
)

// PlatformClients talks to a remote platform (prod-like thin).
// AuthMiddlewareModule validates tokens via AUTH_URL — must be remote, not PORT.
var PlatformClients = fx.Options(
	auth.AuthMiddlewareModule,
	profile.ProfileClientModule,
	mail.MailClientModule,
	knapsack.KnapsackClientModule,
	party.PartyClientModule,
	buddy.BuddyClientModule,
	leaderboard.LeaderboardClientPublic,
	chat.ChatClientModule,
)

// Aggregate is the local/dev process: game transports + in-process platform.
var Aggregate = fx.Options(Infra, AllModule, Platform)

// Thin is the game-only process: game transports + remote platform clients.
var Thin = fx.Options(Infra, AllModule, PlatformClients)

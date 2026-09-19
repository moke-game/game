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
)

// Infra is what fxmain.Main does not enable by default but this game uses:
// NATS + in-process local MQ (Watch/Hi publish both) and Redis ICache.
var Infra = fx.Options(
	mfx.NatsModule,
	mfx.LocalModule,
	ofx.RedisCacheModule,
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
//
//	fxmain.Main(modules.Aggregate)
var Aggregate = fx.Options(Infra, AllModule, Platform)

// Thin is the game-only process: game transports + remote platform clients.
//
//	fxmain.Main(modules.Thin)
var Thin = fx.Options(Infra, AllModule, PlatformClients)

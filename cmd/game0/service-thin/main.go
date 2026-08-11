package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
	buddy "github.com/moke-game/platform/services/buddy/pkg/module"
	chat "github.com/moke-game/platform/services/chat/pkg/module"
	knapsack "github.com/moke-game/platform/services/knapsack/pkg/module"
	leaderboard "github.com/moke-game/platform/services/leaderboard/pkg/module"
	mail "github.com/moke-game/platform/services/mail/pkg/module"
	party "github.com/moke-game/platform/services/party/pkg/module"
	profile "github.com/moke-game/platform/services/profile/pkg/module"

	"github.com/moke-game/game/pkg/modules"
)

// Thin topology: game-only process talking to remote platform over gRPC clients.
// Set AUTH_URL (and other *_URL envs) to the remote platform endpoints.
// AuthMiddlewareModule validates tokens via the remote AuthService.
func main() {
	fxmain.Main(
		mfx.NatsModule,
		mfx.LocalModule,
		ofx.RedisCacheModule,

		// game public API
		modules.AllModule,

		// remote auth middleware + platform clients
		auth.AuthMiddlewareModule,
		profile.ProfileClientModule,
		mail.MailClientModule,
		knapsack.KnapsackClientModule,
		party.PartyClientModule,
		buddy.BuddyClientModule,
		leaderboard.LeaderboardClientPublic,
		chat.ChatClientModule,
	)
}

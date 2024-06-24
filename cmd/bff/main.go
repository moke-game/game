package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

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

	iap "github.com/gstones/moke-kit/3rd/iap/pkg/module"
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"

	configs "github.com/moke-game/game/configs/pkg/module"
	bff "github.com/moke-game/game/services/bff/pkg/module"
)

func initEnvs() {
	err := os.Setenv("APP_NAME", "bff")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("TIMEOUT", "15")
	if err != nil {
		panic(err)
	}
}
func initPprof() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
}

func main() {
	initEnvs()
	initPprof()
	fxmain.Main(
		//clients
		auth.AuthMiddlewareModule,
		profile.ProfileClientModule,
		knapsack.KnapsackClientModule,
		chat.ChatClientModule,
		party.PartyClientModule,
		matchmaking.MatchClientModule,
		buddy.BuddyClientModule,
		mail.MailAllClientModule,
		analytics.AnalyticsClientModule,
		leaderboard.LeaderboardClientAll,

		// bff
		bff.BffModule,
		configs.ConfigsModule,

		mfx.NatsModule,
		iap.IAPModule,
	)
}

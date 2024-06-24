package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gstones/moke-kit/3rd/agones/pkg/agonesfx"
	iap "github.com/gstones/moke-kit/3rd/iap/pkg/module"
	"github.com/gstones/moke-kit/fxmain"
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
	match "github.com/moke-game/platform/services/matchmaking/pkg/module"
	party "github.com/moke-game/platform/services/party/pkg/module"
	profile "github.com/moke-game/platform/services/profile/pkg/module"

	configs "github.com/moke-game/game/configs/pkg/module"
	bff "github.com/moke-game/game/services/bff/pkg/module"
	gm "github.com/moke-game/game/services/gm/pkg/module"
	room "github.com/moke-game/game/services/room/pkg/module"
)

func initEnvs() {
	err := os.Setenv("APP_NAME", "game")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("TIMEOUT", "200")
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
		// bff
		bff.BffModule,
		configs.ConfigsModule,

		// platform services
		auth.AuthAllModule,
		profile.ProfileAllModule,
		knapsack.KnapsackAllModule,
		room.RoomModule,
		chat.ChatAllModule,
		party.PartyAllModule,
		match.MatchAllModule,
		buddy.BuddyAllModule,
		mail.MailAllModule,
		gm.GmAllModule,
		analytics.AnalyticsAllModule,
		leaderboard.LeaderboardAll,
		// infrastructures
		mfx.NatsModule,
		mfx.LocalModule,
		ofx.RedisCacheModule,
		fx.NopLogger,
		agonesfx.AgonesSDKModule,
		iap.IAPModule,
	)
}

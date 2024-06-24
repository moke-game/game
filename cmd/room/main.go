package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	agones "github.com/gstones/moke-kit/3rd/agones/pkg/module"
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"go.uber.org/fx"

	analytics "github.com/moke-game/platform/services/analytics/pkg/module"
	auth "github.com/moke-game/platform/services/auth/pkg/module"
	knapsack "github.com/moke-game/platform/services/knapsack/pkg/module"
	leaderboard "github.com/moke-game/platform/services/leaderboard/pkg/module"
	matchmaking "github.com/moke-game/platform/services/matchmaking/pkg/module"
	party "github.com/moke-game/platform/services/party/pkg/module"
	profile "github.com/moke-game/platform/services/profile/pkg/module"

	configs "github.com/moke-game/game/configs/pkg/module"
	room "github.com/moke-game/game/services/room/pkg/module"
)

func initEnvs() {
	err := os.Setenv("APP_NAME", "room")
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
		mfx.NatsModule,
		agones.AgonesSDKModule,
		configs.ConfigsModule,
		room.RoomModule,
		// platform services
		auth.AuthMiddlewareModule,
		profile.ProfileClientModule,
		knapsack.KnapsackClientModule,
		fx.NopLogger,
		matchmaking.MatchClientModule,
		party.PartyClientModule,
		leaderboard.LeaderboardClientPrivate,
		analytics.AnalyticsClientModule,
	)
}

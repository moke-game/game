package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	match "github.com/moke-game/platform/services/matchmaking/pkg/module"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
	chat "github.com/moke-game/platform/services/chat/pkg/module"
	knapsack "github.com/moke-game/platform/services/knapsack/pkg/module"
	mail "github.com/moke-game/platform/services/mail/pkg/module"
	profile "github.com/moke-game/platform/services/profile/pkg/module"

	gm "github.com/moke-game/game/services/gm/pkg/module"
)

func main() {
	fxmain.Main(
		profile.ProfileClientModule,
		chat.ChatPrivateClientModule,
		auth.AuthClientModule,
		mail.MailClientPrivateModule,
		knapsack.KnapsackClientModule,
		gm.GmModule,
		match.MatchClientModule,

		// infrastructures
		mfx.NatsModule,
		ofx.RedisCacheModule,
	)
}

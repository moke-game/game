package module

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/moke-game/game/services/bff/internal"
	"github.com/moke-game/game/services/bff/internal/routers/buddy"
	"github.com/moke-game/game/services/bff/internal/routers/chat"
	"github.com/moke-game/game/services/bff/internal/routers/gm"
	"github.com/moke-game/game/services/bff/internal/routers/knapsack"
	"github.com/moke-game/game/services/bff/internal/routers/leaderboard"
	"github.com/moke-game/game/services/bff/internal/routers/login"
	"github.com/moke-game/game/services/bff/internal/routers/mail"
	"github.com/moke-game/game/services/bff/internal/routers/matchmaking"
	"github.com/moke-game/game/services/bff/internal/routers/party"
	"github.com/moke-game/game/services/bff/internal/routers/profile"
	"github.com/moke-game/game/services/bff/pkg/bfx"
	"github.com/moke-game/game/services/client/interactive"
)

// BffModule  backend for frontend service module
var BffModule = fx.Module("bff", fx.Options(
	bfx.SettingsModule,
	internal.Module,
	RoutersModule,
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("bff")
	}),
))

var RoutersModule = fx.Module("routers", fx.Options(
	login.RouterLogin,
	profile.RouterProfile,
	knapsack.RouterKnapsack,
	chat.RouterChat,
	party.RouterParty,
	matchmaking.RouterMatchmaking,
	buddy.RouterBuddy,
	mail.RouterMail,
	leaderboard.RouterLeaderboard,
	gm.RouterGM,
))

func NewClientShell(url string) (*interactive.Client, error) {
	c := &interactive.Client{}
	err := c.Init(url)
	if err != nil {
		return nil, err
	}
	return c, nil
}

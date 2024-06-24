package module

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/moke-game/game/services/room/internal"
	"github.com/moke-game/game/services/room/pkg/rfx"
)

// RoomModule  backend for frontend service module
var RoomModule = fx.Module("room", fx.Options(
	rfx.SettingsModule,
	internal.Module,
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("room")
	}),
))

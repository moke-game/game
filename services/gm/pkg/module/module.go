package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/game/services/gm/internal/service"
	"github.com/moke-game/game/services/gm/pkg/gmfx"
)

// GmModule Provides gm service
var GmModule = fx.Module("gm",
	service.GmService,
	gmfx.GmSettingsModule,
	gmfx.GmMysqlModule,
)

// GmClientModule Provides gm client for grpc
var GmClientModule = fx.Module("gm_client",
	gmfx.GmClientModule,
	gmfx.GmSettingsModule,
)

// GmAllModule Provides client, service for gm
var GmAllModule = fx.Module("gm_all",
	service.GmService,
	gmfx.GmClientModule,
	gmfx.GmSettingsModule,
	gmfx.GmMysqlModule,
)

package module

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	mfx2 "github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/utility"

	"github.com/moke-game/game/configs/pkg/cfx"
)

// ConfigsModule config  module
var ConfigsModule = fx.Module("configs", fx.Options(
	cfx.SettingsModule,
	cfx.ConfigsCreator,
	ConfigsGlobalModule,
))
var ConfigsGlobal *cfx.ConfigsParams
var DeploymentGlobal utility.Deployments

var ConfigsGlobalModule = fx.Invoke(
	func(
		l *zap.Logger,
		p cfx.ConfigsParams,
		aParams mfx2.AppParams,
	) {
		ConfigsGlobal = &p
		DeploymentGlobal = utility.ParseDeployments(aParams.Deployment)
	},
)

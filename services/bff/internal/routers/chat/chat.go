package chat

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/moke-game/platform/api/gen/chat"
	profile "github.com/moke-game/platform/api/gen/profile"
	cfx2 "github.com/moke-game/platform/services/chat/pkg/cfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	"github.com/moke-game/game/configs/pkg/cfx"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
)

type Router struct {
	logger  *zap.Logger
	cClient chat.ChatServiceClient
	pClient profile.ProfileServiceClient
	configs cfx.ConfigsParams
	mq      miface.MessageQueue
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_CHATMessage, r.chatMessage)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_ChatGetEmoji, r.chatGetEmoji)
}

func CreateRouter(
	logger *zap.Logger,
	cClient chat.ChatServiceClient,
	pClient profile.ProfileServiceClient,
	configs cfx.ConfigsParams,
	mq miface.MessageQueue,
) (*Router, error) {
	return &Router{
		logger:  logger,
		cClient: cClient,
		pClient: pClient,
		configs: configs,
		mq:      mq,
	}, nil
}

var RouterChat = fx.Provide(
	func(
		logger *zap.Logger,
		cParams cfx2.ChatClientParams,
		pParams pfx.ProfileClientParams,
		setting bfx.GameSettingParams,
		configs cfx.ConfigsParams,
		mqParams mfx.MessageQueueParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			cParams.ChatClient,
			pParams.ProfileClient,
			configs,
			mqParams.MessageQueue,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

package login

import (
	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/mq/miface"
	mfx2 "github.com/gstones/moke-kit/mq/pkg/mfx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	apb "github.com/moke-game/platform/api/gen/auth"
	"github.com/moke-game/platform/api/gen/buddy"
	"github.com/moke-game/platform/api/gen/chat"
	kpb "github.com/moke-game/platform/api/gen/knapsack"
	ptpb "github.com/moke-game/platform/api/gen/party"
	ppb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/auth/pkg/afx"
	bfx2 "github.com/moke-game/platform/services/buddy/pkg/bfx"
	"github.com/moke-game/platform/services/chat/pkg/cfx"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
	"github.com/moke-game/platform/services/party/pkg/ptfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
)

type Router struct {
	logger         *zap.Logger
	mq             miface.MessageQueue
	aClient        apb.AuthServiceClient
	pClient        ppb.ProfileServiceClient
	pPrivateClient ppb.ProfilePrivateServiceClient
	cClient        chat.ChatServiceClient
	bClient        buddy.BuddyServiceClient
	ptClient       ptpb.PartyServiceClient
	kClient        kpb.KnapsackServiceClient
	appId          string
}

// Register register handlers
// 注册处理器
func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_Auth, r.handleAuth)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_Heartbeat, r.handleHeartbeat)
}

func CreateRouter(
	logger *zap.Logger,
	authClient apb.AuthServiceClient,
	profileClient ppb.ProfileServiceClient,
	pPrivateClient ppb.ProfilePrivateServiceClient,
	cClient chat.ChatServiceClient,
	bClient buddy.BuddyServiceClient,
	ptClient ptpb.PartyServiceClient,
	kClient kpb.KnapsackServiceClient,
	name string,
	mq miface.MessageQueue,
) (*Router, error) {
	return &Router{
		logger:         logger,
		aClient:        authClient,
		appId:          name,
		pClient:        profileClient,
		pPrivateClient: pPrivateClient,
		cClient:        cClient,
		bClient:        bClient,
		ptClient:       ptClient,
		kClient:        kClient,
		mq:             mq,
	}, nil
}

var RouterLogin = fx.Provide(
	func(
		logger *zap.Logger,
		aParams afx.AuthClientParams,
		pParams pfx.ProfileClientParams,
		cParams cfx.ChatClientParams,
		bParams bfx2.BuddyClientParams,
		partyParams ptfx.PartyClientParams,
		kParams kfx.KnapsackClientParams,
		hParams bfx2.BuddyClientParams,
		setting bfx.GameSettingParams,
		appParams mfx.AppParams,
		mqParams mfx2.MessageQueueParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			aParams.AuthClient,
			pParams.ProfileClient,
			pParams.ProfilePrivateClient,
			cParams.ChatClient,
			bParams.BuddyClient,
			partyParams.PartyClient,
			kParams.KnapsackClient,
			appParams.Deployment,
			mqParams.MessageQueue,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

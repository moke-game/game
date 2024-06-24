package buddy

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	pb2 "github.com/moke-game/platform/api/gen/buddy"
	pb4 "github.com/moke-game/platform/api/gen/party"
	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/buddy/pkg/bfx"
	"github.com/moke-game/platform/services/party/pkg/ptfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"
)

type Router struct {
	logger   *zap.Logger
	pClient  pb.ProfileServiceClient
	bClient  pb2.BuddyServiceClient
	ptClient pb4.PartyServiceClient
	mq       miface.MessageQueue
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendGet, r.handleFriendGet)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendAdd, r.handleFriendAdd)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendAgree, r.handleFriendAgree)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendRefuse, r.handleFriendRefuse)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendDelete, r.handleFriendDelete)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendNickName, r.handleFriendNickName)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendDeleteBlack, r.handleFriendDeleteBlack)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendGetPlayerInfoByCode, r.handleFriendGetPlayerInfoByCode)
}

func CreateRouter(
	logger *zap.Logger,
	pClient pb.ProfileServiceClient,
	bClient pb2.BuddyServiceClient,
	ptClient pb4.PartyServiceClient,
	mq miface.MessageQueue,
) (*Router, error) {
	return &Router{
		logger:   logger,
		pClient:  pClient,
		bClient:  bClient,
		ptClient: ptClient,
		mq:       mq,
	}, nil
}

var RouterBuddy = fx.Provide(
	func(
		logger *zap.Logger,
		pParams pfx.ProfileClientParams,
		bParams bfx.BuddyClientParams,
		ptParams ptfx.PartyClientParams,
		mqParams mfx.MessageQueueParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			pParams.ProfileClient,
			bParams.BuddyClient,
			ptParams.PartyClient,
			mqParams.MessageQueue,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

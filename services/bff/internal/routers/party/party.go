package party

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/moke-game/platform/api/gen/buddy"

	bfx2 "github.com/moke-game/platform/services/buddy/pkg/bfx"

	match "github.com/moke-game/platform/api/gen/matchmaking"
	ptpb "github.com/moke-game/platform/api/gen/party"
	ppb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/matchmaking/pkg/matchfx"
	"github.com/moke-game/platform/services/party/pkg/ptfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	"github.com/moke-game/game/configs/pkg/cfx"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
)

type Router struct {
	logger   *zap.Logger
	ptClient ptpb.PartyServiceClient
	pClient  ppb.ProfilePrivateServiceClient
	mClient  match.MatchServiceClient
	bClient  buddy.BuddyServiceClient
	configs  cfx.ConfigsParams
	mq       miface.MessageQueue
	redis    *redis.Client
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_GetPartyInfo, r.getPartyInfo)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_JoinParty, r.joinParty)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_LeaveParty, r.leaveParty)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_KickParty, r.kickParty)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_READY_PARTY, r.readyParty)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_CANCEL_READY_PARTY, r.cancelReadyParty)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_FriendPartyInfo, r.friendPartyInfo)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_CreateParty, r.createParty)
}

func CreateRouter(
	logger *zap.Logger,
	ptClient ptpb.PartyServiceClient,
	pClient ppb.ProfilePrivateServiceClient,
	mClient match.MatchServiceClient,
	bClient buddy.BuddyServiceClient,
	configs cfx.ConfigsParams,
	mq miface.MessageQueue,
	redis *redis.Client,
) (*Router, error) {
	return &Router{
		logger:   logger,
		configs:  configs,
		mq:       mq,
		ptClient: ptClient,
		pClient:  pClient,
		mClient:  mClient,
		bClient:  bClient,
		redis:    redis,
	}, nil
}

var RouterParty = fx.Provide(
	func(
		logger *zap.Logger,
		ptParams ptfx.PartyClientParams,
		pParams pfx.ProfileClientParams,
		mParams matchfx.MatchClientParams,
		bParams bfx2.BuddyClientParams,
		setting bfx.GameSettingParams,
		configs cfx.ConfigsParams,
		mqParams mfx.MessageQueueParams,
		redisParams ofx.RedisParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			ptParams.PartyClient,
			pParams.ProfilePrivateClient,
			mParams.MatchClient,
			bParams.BuddyClient,
			configs,
			mqParams.MessageQueue,
			redisParams.Redis,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

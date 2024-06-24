package matchmaking

import (
	mfx2 "github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/utility"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb4 "github.com/moke-game/platform/api/gen/party"
	"github.com/moke-game/platform/services/party/pkg/ptfx"

	"github.com/moke-game/platform/services/matchmaking/pkg/matchfx"

	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	pb2 "github.com/moke-game/platform/api/gen/matchmaking"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/configs/pkg/cfx"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
)

type Router struct {
	logger      *zap.Logger
	pClient     pb.ProfileServiceClient
	mClient     pb2.MatchServiceClient
	partyClient pb4.PartyServiceClient
	mq          miface.MessageQueue
	configs     cfx.ConfigsParams
	deployments utility.Deployments
	redisCli    *redis.Client
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_MatchingSingleStart, r.handleMatchSingleStart)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_MatchingCancel, r.handleMatchCancel)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_MatchingStatus, r.handleMatchStatus)
}

func CreateRouter(
	logger *zap.Logger,
	pClient pb.ProfileServiceClient,
	mClient pb2.MatchServiceClient,
	partyClient pb4.PartyServiceClient,
	mq miface.MessageQueue,
	redisCli *redis.Client,
	configs cfx.ConfigsParams,
	deployment string,
) (*Router, error) {
	return &Router{
		logger:      logger,
		pClient:     pClient,
		mClient:     mClient,
		partyClient: partyClient,
		mq:          mq,
		redisCli:    redisCli,
		configs:     configs,
		deployments: utility.ParseDeployments(deployment),
	}, nil
}

var RouterMatchmaking = fx.Provide(
	func(
		logger *zap.Logger,
		pParams pfx.ProfileClientParams,
		setting bfx.GameSettingParams,
		mParams matchfx.MatchClientParams,
		partyParams ptfx.PartyClientParams,
		mqParams mfx.MessageQueueParams,
		redisParams ofx.RedisParams,
		configs cfx.ConfigsParams,
		appParams mfx2.AppParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			pParams.ProfileClient,
			mParams.MatchClient,
			partyParams.PartyClient,
			mqParams.MessageQueue,
			redisParams.Redis,
			configs,
			appParams.Deployment,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

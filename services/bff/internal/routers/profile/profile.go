package profile

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"

	"github.com/moke-game/platform/api/gen/leaderboard"
	"github.com/moke-game/platform/api/gen/mail"
	party "github.com/moke-game/platform/api/gen/party"
	"github.com/moke-game/platform/services/leaderboard/pkg/lbfx"
	"github.com/moke-game/platform/services/mail/pkg/mailfx"
	"github.com/moke-game/platform/services/party/pkg/ptfx"

	"go.uber.org/fx"
	"go.uber.org/zap"

	kpb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/matchmaking"
	ppb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
	"github.com/moke-game/platform/services/matchmaking/pkg/matchfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/configs/pkg/cfx"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
)

type Router struct {
	logger         *zap.Logger
	pClient        ppb.ProfileServiceClient
	kClient        kpb.KnapsackServiceClient
	mmClient       matchmaking.MatchServiceClient
	ptClient       party.PartyServiceClient
	configs        cfx.ConfigsParams
	lbClient       leaderboard.LeaderboardServiceClient
	mq             miface.MessageQueue
	mailPrivClient mail.MailPrivateServiceClient
	ppClient       ppb.ProfilePrivateServiceClient
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_NewPlayer, r.handleNewPlayer)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_SimpleInfo, r.handleSimpleInfo)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_PlayerRename, r.handleRename)
}

func CreateRouter(
	logger *zap.Logger,
	pClient ppb.ProfileServiceClient,
	kClient kpb.KnapsackServiceClient,
	mmClient matchmaking.MatchServiceClient,
	ptClient party.PartyServiceClient,
	configs cfx.ConfigsParams,
	kPrivateClient kpb.KnapsackPrivateServiceClient,
	lbClient leaderboard.LeaderboardServiceClient,
	mq miface.MessageQueue,
	mailPrivClient mail.MailPrivateServiceClient,
	ppClient ppb.ProfilePrivateServiceClient,
) (*Router, error) {
	return &Router{
		logger:         logger,
		pClient:        pClient,
		kClient:        kClient,
		ptClient:       ptClient,
		configs:        configs,
		mmClient:       mmClient,
		lbClient:       lbClient,
		mq:             mq,
		mailPrivClient: mailPrivClient,
		ppClient:       ppClient,
	}, nil
}

var RouterProfile = fx.Provide(
	func(
		logger *zap.Logger,
		pParams pfx.ProfileClientParams,
		kParams kfx.KnapsackClientParams,
		mmParams matchfx.MatchClientParams,
		partyParams ptfx.PartyClientParams,
		configs cfx.ConfigsParams,
		lbParams lbfx.LeaderboardClientParams,
		mqParams mfx.MessageQueueParams,
		mailPrivParams mailfx.MailClientPrivateParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			pParams.ProfileClient,
			kParams.KnapsackClient,
			mmParams.MatchClient,
			partyParams.PartyClient,
			configs,
			kParams.KnapsackPrivateClient,
			lbParams.Client,
			mqParams.MessageQueue,
			mailPrivParams.MailClient,
			pParams.ProfilePrivateClient,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

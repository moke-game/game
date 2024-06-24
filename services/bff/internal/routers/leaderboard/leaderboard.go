package leaderboard

import (
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	knapsack "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/leaderboard"
	"github.com/moke-game/platform/api/gen/mail"
	profile "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
	"github.com/moke-game/platform/services/leaderboard/pkg/lbfx"
	"github.com/moke-game/platform/services/mail/pkg/mailfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/configs/pkg/module"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	leaderboard2 "github.com/moke-game/game/services/common/leaderboard"
)

type Router struct {
	logger               *zap.Logger
	lbClient             leaderboard.LeaderboardServiceClient
	lbPrivateClient      leaderboard.LeaderboardPrivateServiceClient
	profileClient        profile.ProfileServiceClient
	profilePrivateClient profile.ProfilePrivateServiceClient
	mailClient           mail.MailPrivateServiceClient
	knapsackClient       knapsack.KnapsackServiceClient
}

func CreateRouter(
	logger *zap.Logger,
	lbClient leaderboard.LeaderboardServiceClient,
	lbPrivateClient leaderboard.LeaderboardPrivateServiceClient,
	profilePrivateClient profile.ProfileServiceClient,
	mailClient mail.MailPrivateServiceClient,
	knapsackClient knapsack.KnapsackServiceClient,
) (*Router, error) {
	return &Router{
		logger:          logger,
		lbClient:        lbClient,
		lbPrivateClient: lbPrivateClient,
		mailClient:      mailClient,
		profileClient:   profilePrivateClient,
		knapsackClient:  knapsackClient,
	}, nil
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_GetLeaderboard, r.getLeaderboard)
	r.trySettlement()
}

// 排行榜结算定时器
func (r *Router) trySettlement() {
	for k := range module.ConfigsGlobal.RankConfigs.Leaderboards {
		go func(id int32) {
			for {
				if period := leaderboard2.MakeLeaderboardPeriod(id); period == "" {
					r.logger.Error("make leaderboard period error", zap.Int32("id", id))
					return
				} else if duration, err := leaderboard2.CalculateLeftDuration(id); err != nil {
					r.logger.Error("calculate left duration failed", zap.Int32("id", id), zap.Error(err))
					return
				} else {
					timer := time.NewTimer(duration)
					<-timer.C
					if err := r.trySettlementRewards(period, id); err != nil {
						r.logger.Error("try settlement rewards failed", zap.Error(err))
					}
				}
			}
		}(k)
	}
}

var RouterLeaderboard = fx.Provide(
	func(
		logger *zap.Logger,
		lbParams lbfx.LeaderboardClientParams,
		lbPrivateParams lbfx.LeaderboardClientPrivateParams,
		profileParams pfx.ProfileClientParams,
		mailParams mailfx.MailClientPrivateParams,
		knapsackParams kfx.KnapsackClientParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			lbParams.Client,
			lbPrivateParams.Client,
			profileParams.ProfileClient,
			mailParams.MailClient,
			knapsackParams.KnapsackClient,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

package knapsack

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/configs/pkg/cfx"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
)

type Router struct {
	logger  *zap.Logger
	kClient pb.KnapsackServiceClient
	configs cfx.ConfigsParams
	mq      miface.MessageQueue
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_WatchingKnapsack, r.watchingKnapsack)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_DiamondExchangeItem, r.diamondExchangeItem)
}

func CreateRouter(
	logger *zap.Logger,
	kClient pb.KnapsackServiceClient,
	configs cfx.ConfigsParams,
	mq miface.MessageQueue,
) (*Router, error) {
	return &Router{
		logger:  logger,
		kClient: kClient,
		configs: configs,
		mq:      mq,
	}, nil
}

var RouterKnapsack = fx.Provide(
	func(
		logger *zap.Logger,
		kParams kfx.KnapsackClientParams,
		setting bfx.GameSettingParams,
		configs cfx.ConfigsParams,
		mqParams mfx.MessageQueueParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			kParams.KnapsackClient,
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

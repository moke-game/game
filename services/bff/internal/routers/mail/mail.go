package mail

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/configs/pkg/cfx"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
	pb2 "github.com/moke-game/platform/api/gen/knapsack"
	pb "github.com/moke-game/platform/api/gen/mail"
	pb3 "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
	"github.com/moke-game/platform/services/mail/pkg/mailfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"
)

type Router struct {
	logger  *zap.Logger
	mClient pb.MailServiceClient
	kClient pb2.KnapsackServiceClient
	pClient pb3.ProfileServiceClient
	configs cfx.ConfigsParams
	mq      miface.MessageQueue
}

func (r *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_ReadMail, r.readMail)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_DeleteReadMail, r.deleteReadMail)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_GetMailRewards, r.getMailRewards)
	register.RegisterHandler(cpb.C2S_EVENT_C2S_WatchMail, r.watchMail)
}

func CreateRouter(
	logger *zap.Logger,
	kClient pb.MailServiceClient,
	configs cfx.ConfigsParams,
	mq miface.MessageQueue,
	ks pb2.KnapsackServiceClient,
	pClient pb3.ProfileServiceClient,
) (*Router, error) {
	return &Router{
		logger:  logger,
		mClient: kClient,
		configs: configs,
		mq:      mq,
		kClient: ks,
		pClient: pClient,
	}, nil
}

var RouterMail = fx.Provide(
	func(
		logger *zap.Logger,
		kParams mailfx.MailClientParams,
		setting bfx.GameSettingParams,
		configs cfx.ConfigsParams,
		mqParams mfx.MessageQueueParams,
		ksParams kfx.KnapsackClientParams,
		params pfx.ProfileClientParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			kParams.MailClient,
			configs,
			mqParams.MessageQueue,
			ksParams.KnapsackClient,
			params.ProfileClient,
		); err != nil {
			err = e
		} else {
			out.Router = r
		}
		return
	},
)

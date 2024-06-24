package gm

import (
	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/mq/miface"
	mfx2 "github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/utility"
	"github.com/gstones/zinx/ziface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	apb "github.com/moke-game/platform/api/gen/auth"
	"github.com/moke-game/platform/api/gen/buddy"
	"github.com/moke-game/platform/api/gen/chat"
	kpb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/leaderboard"
	ptpb "github.com/moke-game/platform/api/gen/party"
	ppb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/auth/pkg/afx"
	bfx2 "github.com/moke-game/platform/services/buddy/pkg/bfx"
	"github.com/moke-game/platform/services/chat/pkg/cfx"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
	"github.com/moke-game/platform/services/leaderboard/pkg/lbfx"
	"github.com/moke-game/platform/services/party/pkg/ptfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	cpb "github.com/moke-game/game/api/gen/common"
	cfgfx "github.com/moke-game/game/configs/pkg/cfx"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
)

type cmd func(request ziface.IRequest, args ...string) error

type Router struct {
	logger          *zap.Logger
	mq              miface.MessageQueue
	aClient         apb.AuthServiceClient
	pClient         ppb.ProfileServiceClient
	privateClient   ppb.ProfilePrivateServiceClient
	cClient         chat.ChatServiceClient
	bClient         buddy.BuddyServiceClient
	ptClient        ptpb.PartyServiceClient
	kpClient        kpb.KnapsackPrivateServiceClient
	lbPrivateClient leaderboard.LeaderboardPrivateServiceClient
	configs         cfgfx.ConfigsParams
	appId           string
	commands        map[string]cmd
	deployments     utility.Deployments
	redisCli        *redis.Client
}

func (c *Router) Register(register biface.IRegister) {
	register.RegisterHandler(cpb.C2S_EVENT_C2S_BffGMCommand, c.handleGM)
	c.initHandler()
}

func (c *Router) initHandler() {
	c.RegisterCommand("addCurrency", c.addCurrency)            // 添加货币
	c.RegisterCommand("addLeaderboard", c.addLeaderboards)     // 添加排行榜假数据
	c.RegisterCommand("updateScore", c.updateSelfScore)        // 更新排行榜分数
	c.RegisterCommand("clearLeaderboard", c.clearLeaderboards) // 清空排行榜排名
	c.RegisterCommand("feat", c.unlockFeat)                    //解锁功能
	c.RegisterCommand("additem", c.addItem)                    //添加道具
}

func (c *Router) RegisterCommand(command string, cmd cmd) {
	c.commands[command] = cmd
}

func CreateRouter(
	logger *zap.Logger,
	authClient apb.AuthServiceClient,
	profileClient ppb.ProfileServiceClient,
	profilePrivateClient ppb.ProfilePrivateServiceClient,
	cClient chat.ChatServiceClient,
	bClient buddy.BuddyServiceClient,
	ptClient ptpb.PartyServiceClient,
	kClient kpb.KnapsackPrivateServiceClient,
	lbClient leaderboard.LeaderboardPrivateServiceClient,
	configs cfgfx.ConfigsParams,
	name string,
	mq miface.MessageQueue,
	redisCli *redis.Client,
) (*Router, error) {
	return &Router{
		logger:          logger,
		aClient:         authClient,
		appId:           name,
		pClient:         profileClient,
		cClient:         cClient,
		bClient:         bClient,
		ptClient:        ptClient,
		kpClient:        kClient,
		lbPrivateClient: lbClient,
		configs:         configs,
		mq:              mq,
		commands:        make(map[string]cmd),
		privateClient:   profilePrivateClient,
		deployments:     utility.ParseDeployments(name),
		redisCli:        redisCli,
	}, nil

}

var RouterGM = fx.Provide(
	func(
		logger *zap.Logger,
		aParams afx.AuthClientParams,
		pParams pfx.ProfileClientParams,
		cParams cfx.ChatClientParams,
		bParams bfx2.BuddyClientParams,
		partyParams ptfx.PartyClientParams,
		kParams kfx.KnapsackClientParams,
		hParams bfx2.BuddyClientParams,
		lbParams lbfx.LeaderboardClientPrivateParams,
		configs cfgfx.ConfigsParams,
		setting bfx.GameSettingParams,
		appParams mfx.AppParams,
		mqParams mfx2.MessageQueueParams,
		redisParams ofx.RedisParams,
	) (out routers.RouterResult, err error) {
		if r, e := CreateRouter(
			logger,
			aParams.AuthClient,
			pParams.ProfileClient,
			pParams.ProfilePrivateClient,
			cParams.ChatClient,
			bParams.BuddyClient,
			partyParams.PartyClient,
			kParams.KnapsackPrivateClient,
			lbParams.Client,
			configs,
			appParams.Deployment,
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

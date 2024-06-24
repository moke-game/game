package service

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"

	mfx2 "github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"

	pb3 "github.com/moke-game/platform/api/gen/auth"
	"github.com/moke-game/platform/api/gen/chat"
	kpb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/mail"
	pb2 "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/auth/pkg/afx"
	"github.com/moke-game/platform/services/chat/pkg/cfx"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
	"github.com/moke-game/platform/services/mail/pkg/mailfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	pb "github.com/moke-game/game/api/gen/gm"
	"github.com/moke-game/game/services/gm/errors"
	"github.com/moke-game/game/services/gm/internal/db"
	"github.com/moke-game/game/services/gm/pkg/gmfx"
)

type Service struct {
	utility.WithoutAuth
	logger     *zap.Logger
	mq         miface.MessageQueue
	appId      string
	deployment string
	profileCLi pb2.ProfilePrivateServiceClient
	authCli    pb3.AuthServiceClient
	mailCli    mail.MailPrivateServiceClient
	chatCli    chat.ChatPrivateServiceClient
	knapsack   kpb.KnapsackPrivateServiceClient
	aesKey     string
	url        string
	db         *db.Database
	gormDb     *gorm.DB
	redisCli   *redis.Client
}

func (s *Service) GetBlockedUserInfo(_ context.Context, request *pb.GetBlockedUserInfoRequest) (*pb.GetBlockedUserInfoResponse, error) {
	res := &pb.GetBlockedUserInfoResponse{}
	if info, err := s.db.GetBlockedList(request.GetUid()); err != nil {
		s.logger.Error("get blockList failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else {
		if v, ok := info[request.GetUid()]; ok {
			res.Detail = v
		}
		return res, nil
	}
}

func (s *Service) RegisterWithGatewayServer(server siface.IGatewayServer) error {
	return pb.RegisterGMServiceHandlerFromEndpoint(
		context.Background(), server.GatewayRuntimeMux(), s.url, server.GatewayOption(),
	)
}

func (s *Service) WatchGM(request *pb.WatchGMRequest, server pb.GMService_WatchGMServer) error {
	if request.Uid == "" {
		s.logger.Error("uid is empty", zap.String("uid", request.Uid))
		return errors.ErrClientParamFailure
	}
	topic := makeBlockedListTopic(request.Uid)
	_, err := s.mq.Subscribe(server.Context(), topic, func(msg miface.Message, err error) common.ConsumptionCode {
		if err != nil {
			s.logger.Error("subscribe block list topic failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		bannedInfo := &pb.BannedInfo{}
		if err := json.Unmarshal(msg.Data(), bannedInfo); err != nil {
			s.logger.Error("unmarshal message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		if err := server.Send(&pb.WatchGMResponse{
			Detail: bannedInfo,
		}); err != nil {
			s.logger.Error("send message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		return common.ConsumeAck
	})
	if err != nil {
		s.logger.Error("subscribe block list topic failed", zap.Error(err))
		return errors.ErrGeneralFailure
	}
	<-server.Context().Done()
	return nil
}

func NewService(
	l *zap.Logger,
	mq miface.MessageQueue,
	deployment string,
	appId string,
	profileCLi pb2.ProfilePrivateServiceClient,
	authCli pb3.AuthServiceClient,
	chatClient chat.ChatPrivateServiceClient,
	mailClient mail.MailPrivateServiceClient,
	knapsack kpb.KnapsackPrivateServiceClient,
	redis *redis.Client,
	aesKey string,
	url string,
	gromDb *gorm.DB,
) (result *Service, err error) {
	result = &Service{
		logger:     l,
		mq:         mq,
		appId:      appId,
		deployment: deployment,
		profileCLi: profileCLi,
		aesKey:     aesKey,
		url:        url,
		authCli:    authCli,
		mailCli:    mailClient,
		chatCli:    chatClient,
		knapsack:   knapsack,
		db:         db.OpenDatabase(l, redis),
		gormDb:     gromDb,
		redisCli:   redis,
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterGMServiceServer(server.GrpcServer(), s)
	return nil
}

var GmService = fx.Provide(
	func(
		l *zap.Logger,
		setting gmfx.GmSettingParams,
		mqParams mfx.MessageQueueParams,
		aParams mfx2.AppParams,
		pParams pfx.ProfileClientParams,
		authParams afx.AuthClientParams,
		chatParams cfx.ChatPrivateClientParams,
		mailParams mailfx.MailClientPrivateParams,
		ksparams kfx.KnapsackClientParams,
		redisParams ofx.RedisParams,
		gormParams ofx.GormParams,
	) (out sfx.GrpcServiceResult, gt sfx.GatewayServiceResult, err error) {
		if s, err := NewService(
			l,
			mqParams.MessageQueue,
			aParams.Deployment,
			aParams.AppId,
			pParams.ProfilePrivateClient,
			authParams.AuthClient,
			chatParams.ChatClient,
			mailParams.MailClient,
			ksparams.KnapsackPrivateClient,
			redisParams.Redis,
			setting.AESKey,
			setting.GmUrl,
			gormParams.GormDB,
		); err != nil {
			return out, gt, err
		} else {
			out.GrpcService = s
			gt.GatewayService = s
		}
		return
	},
)

package game0

import (
	"context"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"
	"github.com/gstones/zinx/ziface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	pb2 "open-match.dev/open-match/pkg/pb"

	pb "github.com/moke-game/game/api/gen/game0/api"
	"github.com/moke-game/game/internal/services/game0/db_nosql"
	"github.com/moke-game/game/internal/services/game0/domain"
	"github.com/moke-game/game/pkg/dfx"
)

type Service struct {
	utility.WithoutAuth
	logger      *zap.Logger
	gameHandler *domain.Game
}

func (s *Service) Run(request *pb2.RunRequest, server pb2.MatchFunction_RunServer) error {
	s.logger.Info("Run", zap.Any("request", request))

	return nil
}

// ---------------- grpc ----------------

func (s *Service) Watch(request *pb.WatchRequest, server pb.Game0Service_WatchServer) error {
	topic := request.GetTopic()
	s.logger.Info("Watch", zap.String("topic", topic))

	if err := s.gameHandler.Watch(
		server.Context(),
		topic,
		func(message string) error {
			if err := server.Send(&pb.WatchResponse{
				Message: message,
			}); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return err
	}

	return nil
}

func (s *Service) Hi(_ context.Context, request *pb.HiRequest) (*pb.HiResponse, error) {
	message := request.GetMessage()
	s.logger.Info("Hi", zap.String("message", message))

	if err := s.gameHandler.Hi(request.GetUid(), request.GetTopic(), request.GetMessage()); err != nil {
		return nil, err
	}
	return &pb.HiResponse{
		Message: "response:  " + message,
	}, nil

}
func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterGame0ServiceServer(server.GrpcServer(), s)
	pb2.RegisterMatchFunctionServer(server.GrpcServer(), s)
	return nil
}

// ---------------- gateway ----------------

func (s *Service) RegisterWithGatewayServer(server siface.IGatewayServer) error {
	return pb.RegisterGame0ServiceHandlerFromEndpoint(
		context.Background(),
		server.GatewayRuntimeMux(),
		server.Endpoint(),
		server.GatewayOption(),
	)
}

//---------------- zinx ----------------

func (s *Service) PreHandle(_ ziface.IRequest) {

}

func (s *Service) Handle(request ziface.IRequest) {
	switch request.GetMsgID() {
	case 1:
		req := &pb.HiRequest{}
		if err := proto.Unmarshal(request.GetData(), req); err != nil {
			s.logger.Error("unmarshal request data error", zap.Error(err))
		} else {
			if err := s.gameHandler.Hi(req.GetUid(), req.GetTopic(), req.GetMessage()); err != nil {
				s.logger.Error("Hi error", zap.Error(err))
			}
		}
	case 2:
		req := &pb.WatchRequest{}
		if err := proto.Unmarshal(request.GetData(), req); err != nil {
			s.logger.Error("unmarshal request data error", zap.Error(err))
		} else {
			if err := s.gameHandler.Watch(
				request.GetConnection().Context(),
				req.GetTopic(),
				func(message string) error {
					resp := &pb.WatchResponse{
						Message: message,
					}
					if data, err := proto.Marshal(resp); err != nil {
						return err
					} else if err := request.GetConnection().SendMsg(2, data); err != nil {
						return err
					}
					return nil
				}); err != nil {
				s.logger.Error("Watch error", zap.Error(err))
			}
		}
	}
}

func (s *Service) PostHandle(_ ziface.IRequest) {

}

func (s *Service) RegisterWithServer(server siface.IZinxServer) {
	server.ZinxServer().AddRouter(1, s)
	server.ZinxServer().AddRouter(2, s)
}

func NewService(
	logger *zap.Logger,
	coll diface.ICollection,
	mq miface.MessageQueue,
	redis *redis.Client,
) (result *Service, err error) {
	handler := domain.NewGame(
		logger,
		db_nosql.OpenDatabase(logger, coll),
		mq,
		redis,
	)

	result = &Service{
		logger:      logger,
		gameHandler: handler,
	}
	return
}

var GrpcService = fx.Provide(
	func(
		l *zap.Logger,
		dProvider ofx.DocumentStoreParams,
		setting dfx.SettingsParams,
		mqParams mfx.MessageQueueParams,
		redisClient ofx.RedisParams,
	) (out sfx.GrpcServiceResult, err error) {
		if coll, err := dProvider.DriverProvider.OpenDbDriver(setting.DbName); err != nil {
			return out, err
		} else if s, err := NewService(
			l,
			coll,
			mqParams.MessageQueue,
			redisClient.Redis,
		); err != nil {
			return out, err
		} else {
			out.GrpcService = s
		}
		return
	},
)

var HttpService = fx.Provide(
	func(
		l *zap.Logger,
		dProvider ofx.DocumentStoreParams,
		setting dfx.SettingsParams,
		mqParams mfx.MessageQueueParams,
		redisClient ofx.RedisParams,
	) (out sfx.GatewayServiceResult, err error) {
		if coll, err := dProvider.DriverProvider.OpenDbDriver(setting.DbName); err != nil {
			return out, err
		} else if s, err := NewService(
			l,
			coll,
			mqParams.MessageQueue,
			redisClient.Redis,
		); err != nil {
			return out, err
		} else {
			out.GatewayService = s
		}
		return
	},
)

var TcpService = fx.Provide(
	func(
		l *zap.Logger,
		dProvider ofx.DocumentStoreParams,
		setting dfx.SettingsParams,
		mqParams mfx.MessageQueueParams,
		redisClient ofx.RedisParams,
	) (out sfx.ZinxServiceResult, err error) {
		if coll, err := dProvider.DriverProvider.OpenDbDriver(setting.DbName); err != nil {
			return out, err
		} else if s, err := NewService(
			l,
			coll,
			mqParams.MessageQueue,
			redisClient.Redis,
		); err != nil {
			return out, err
		} else {
			out.ZinxService = s
		}
		return
	},
)

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/gstones/moke-kit/3rd/agones/aiface"
	"github.com/gstones/moke-kit/3rd/agones/pkg/agonesfx"
	common3 "github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/znet"

	"github.com/moke-game/game/services/common/constants"

	apb "github.com/moke-game/platform/api/gen/auth"
	kpb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/leaderboard"
	match "github.com/moke-game/platform/api/gen/matchmaking"
	ppb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/auth/pkg/afx"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
	"github.com/moke-game/platform/services/leaderboard/pkg/lbfx"
	"github.com/moke-game/platform/services/matchmaking/pkg/matchfx"
	data2 "github.com/moke-game/platform/services/matchmaking/pkg/module/data"
	"github.com/moke-game/platform/services/party/pkg/ptfx"
	"github.com/moke-game/platform/services/profile/pkg/pfx"

	cpb "github.com/moke-game/game/api/gen/common"
	roompb "github.com/moke-game/game/api/gen/room"
	"github.com/moke-game/game/configs/pkg/cfx"
	configs "github.com/moke-game/game/configs/pkg/module"
	common2 "github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/notification"
	"github.com/moke-game/game/services/common/pack"
	"github.com/moke-game/game/services/room/internal/room"
	"github.com/moke-game/game/services/room/internal/room/riface"
	"github.com/moke-game/game/services/room/pkg/rfx"
)

type Service struct {
	znet.BaseRouter
	logger       *zap.Logger
	roomMgr      *RoomMgr
	agones       *Agones
	asyncHandler map[uint32]riface.IHandler

	aClient        apb.AuthServiceClient
	pClient        ppb.ProfileServiceClient
	pPrivateClient ppb.ProfilePrivateServiceClient
	kClient        kpb.KnapsackPrivateServiceClient
	matchClient    match.MatchServiceClient
	lbClient       leaderboard.LeaderboardPrivateServiceClient
	rClient        *redis.Client
	configs        cfx.ConfigsParams
	setting        rfx.RoomSettingParams
	mq             miface.MessageQueue
}

func (s *Service) RegisterWithServer(server siface.IZinxServer) {
	for v := range cpb.C2S_EVENT_name {
		if (v >= 20000 && v < 60000) || v == 90000 {
			if v == 20002 {
				continue
			}
			server.ZinxServer().AddRouter(uint32(v), s)
		}
	}
	server.ZinxServer().SetPacket(pack.NewDataPack())
	server.ZinxServer().SetDecoder(pack.NewLTEV_Decoder())

	if err := s.agones.ready(); err != nil {
		s.logger.Error("agones ready failed", zap.Error(err))
	}
	if err := s.agones.watchGameServer(); err != nil {
		s.logger.Error("watch game server failed", zap.Error(err))
	}
}

func (s *Service) logRequest(request ziface.IRequest) {
	if request.GetMsgID() == uint32(cpb.C2S_EVENT_C2S_RoomHeartbeat) ||
		request.GetMsgID() == uint32(cpb.C2S_EVENT_C2S_Move) ||
		request.GetMsgID() == uint32(cpb.C2S_EVENT_C2S_MovePrepare) ||
		request.GetMsgID() == uint32(cpb.C2S_EVENT_C2S_StopMoving) {
		return
	}
	logger := s.logger
	if uid, err := request.GetConnection().GetProperty(constants.ConnUid); err == nil {
		logger = s.logger.With(zap.String("uid", uid.(string)))
	}
	if roomId, err := request.GetConnection().GetProperty(constants.ConnRoomId); err == nil {
		logger = s.logger.With(zap.String("roomId", roomId.(string)))
	}
	if name, ok := cpb.C2S_EVENT_name[int32(request.GetMsgID())]; ok {
		logger.Info(
			"room request",
			zap.String("msg", name),
			zap.Any("data", request.GetData()),
		)
	} else {
		logger.Info(
			"room request",
			zap.Uint32("msgID", request.GetMsgID()),
			zap.Any("data", request.GetData()),
		)
	}
}
func (s *Service) recover() {
	if !configs.DeploymentGlobal.IsProd() {
		return
	}
	if r := recover(); r != nil {
		s.logger.Error("panic",
			zap.Any("recover", r),
			zap.String("stack", string(debug.Stack())),
		)
	}
}

func (s *Service) Handle(request ziface.IRequest) {
	defer s.recover()
	s.logRequest(request)
	if request.GetMsgID() == uint32(cpb.C2S_EVENT_C2S_EnterRoom) {
		if err := s.enterRoom(request); err != nil {
			errCode := cpb.ERRORCODE_COMMON_ERROR
			s.logger.Error("enter room failed", zap.Error(err))
			if err := common2.SendResponse(
				request.GetConnection(),
				uint32(cpb.S2C_EVENT_S2C_EnterRoom),
				errCode,
				nil,
			); err != nil {
				s.logger.Error("send response failed", zap.Error(err))
			}
			return
		}
	}
	if uid, err := request.GetConnection().GetProperty(constants.ConnUid); err != nil {
		s.logger.Error("get uid failed", zap.Error(err))
		s.kickOut(request, false)
	} else if roomId, err := request.GetConnection().GetProperty(constants.ConnRoomId); err != nil {
		s.logger.Error("get roomId failed", zap.Error(err))
		s.kickOut(request, false)
	} else if r, err := s.roomMgr.LoadRoom(roomId.(string)); err != nil {
		s.logger.Error("load room failed", zap.Error(err))
		if err := common2.SendResponse(
			request.GetConnection(),
			request.GetMsgID()+1,
			cpb.ERRORCODE_SUCCESS,
			nil,
		); err != nil {
			s.logger.Error("send response failed", zap.Error(err))
		}
	} else {
		r.Receive(uid.(string), request)
	}
}

func (s *Service) kickOut(request ziface.IRequest, back2Hall bool) {
	s.logger.Warn("kick out")
	reason := roompb.KickOutReason_KICK_OUT_REASON_STATE_ERROR
	if back2Hall {
		reason = roompb.KickOutReason_KICK_OUT_REASON_BACK_TO_HALL
	}
	_ = common2.SendNotify(request.GetConnection(), uint32(cpb.S2C_EVENT_NTF_KickOut), &roompb.NtfKickOut{
		Reason: reason,
	})
}

func (s *Service) enterRoom(request ziface.IRequest) error {
	req := &roompb.C2SEnterRoom{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return err
	}
	if req.RoomId == "" {
		return fmt.Errorf("invalid room id:%s", req.RoomId)
	}
	uid, _, err := s.checkToken(req.Token)
	if err != nil {
		return err
	}
	result := &data2.MatchResult{}
	if len(req.RoomToken) > 0 {
		if _, data, err := s.checkToken(req.RoomToken); err != nil {
			return err
		} else {
			if err = json.Unmarshal(data, result); err != nil {
				return err
			}
		}
	}
	allowCreat := result.IsFirstEnter
	if result.PlayId == 0 {
		if !s.agones.CheckAndDeleteReserve(uid) {
			return fmt.Errorf("player %s not reserved or reserved timeout", uid)
		}
		allowCreat = true
	}
	if r, err := s.roomMgr.LoadOrCreateRoom(
		req.RoomId,
		func() (riface.IRoom, error) {
			return room.CreateRoom(
				req.RoomId,
				s.logger,
				s.pClient,
				s.pPrivateClient,
				s.kClient,
				s.matchClient,
				s.lbClient,
				s.configs,
				s.rClient,
				s.mq,
				result,
				s.setting,
			)
		},
		allowCreat,
	); err != nil {
		return err
	} else {
		if req.CountryCode != "" {
			request.GetConnection().SetProperty(constants.ConnCountryCode, req.CountryCode)
		}
		request.GetConnection().SetProperty(constants.ConnToken, req.Token)
		request.GetConnection().SetProperty(constants.ConnUid, uid)
		request.GetConnection().SetProperty(constants.ConnRoomId, req.RoomId)
		go func(uid string) {
			s.subNotification(request.GetConnection(), uid)
			<-request.GetConnection().Context().Done()
			if err := s.agones.DeletePlayer(uid); err != nil {
				s.logger.Warn("Agones delete player failed", zap.String("roomID", r.RoomId()), zap.Error(err))
			}

			//r.Receive(uid, common.CreateExitMsg())
		}(uid)
	}
	return nil
}

func (s *Service) subNotification(connection ziface.IConnection, uid string) {
	topic := notification.MakeGamePrivateNotifyTopic(uid)
	_, err := s.mq.Subscribe(connection.Context(), topic, func(msg miface.Message, err error) common3.ConsumptionCode {
		if err != nil {
			s.logger.Error("subscribe buddy topic failed", zap.Error(err))
			return common3.ConsumeNackPersistentFailure
		}
		var event notification.NotifyEvent
		if err := json.Unmarshal(msg.Data(), &event); err != nil {
			s.logger.Error("unmarshal notify event failed", zap.Error(err))
			return common3.ConsumeNackPersistentFailure
		}
		if event == notification.NotifyEventLoginOverride {
			connection.Stop()
		}
		return common3.ConsumeAck
	})
	if err != nil {
		s.logger.Error("subscribe buddy topic failed", zap.Error(err))
	}
}

func (s *Service) checkToken(token string) (string, []byte, error) {
	if resp, err := s.aClient.ValidateToken(context.Background(), &apb.ValidateTokenRequest{
		AccessToken: token,
	}); err != nil {
		return "", nil, err
	} else {
		return resp.GetUid(), resp.CustomData, nil
	}
}

func NewService(
	l *zap.Logger,
	aClient apb.AuthServiceClient,
	pClient ppb.ProfileServiceClient,
	pPrivateClient ppb.ProfilePrivateServiceClient,
	kClient kpb.KnapsackPrivateServiceClient,
	matchClient match.MatchServiceClient,
	lb leaderboard.LeaderboardPrivateServiceClient,
	mq miface.MessageQueue,
	rClient *redis.Client,
	agones aiface.IAgones,
	configs cfx.ConfigsParams,
	setting rfx.RoomSettingParams,
) (result *Service, err error) {
	rm := NewRoomMgr(l, agones)
	ag, err := NewAgones(agones, l, rm)
	if err != nil {
		return nil, err
	}

	result = &Service{
		logger:         l,
		aClient:        aClient,
		pClient:        pClient,
		pPrivateClient: pPrivateClient,
		kClient:        kClient,
		rClient:        rClient,
		matchClient:    matchClient,
		lbClient:       lb,
		roomMgr:        rm,
		configs:        configs,
		agones:         ag,
		setting:        setting,
		mq:             mq,
		asyncHandler:   make(map[uint32]riface.IHandler),
	}
	return
}

var Module = fx.Provide(
	func(
		l *zap.Logger,
		aParams afx.AuthClientParams,
		pParams pfx.ProfileClientParams,
		kParams kfx.KnapsackClientParams,
		matchParams matchfx.MatchClientParams,
		partyParams ptfx.PartyClientParams,
		lbParams lbfx.LeaderboardClientPrivateParams,
		redisParams ofx.RedisParams,
		configs cfx.ConfigsParams,
		agParams agonesfx.SDKParams,
		setting rfx.RoomSettingParams,
		mqParams mfx.MessageQueueParams,
	) (out sfx.ZinxServiceResult, err error) {
		if svc, e := NewService(
			l,
			aParams.AuthClient,
			pParams.ProfileClient,
			pParams.ProfilePrivateClient,
			kParams.KnapsackPrivateClient,
			matchParams.MatchClient,
			lbParams.Client,
			mqParams.MessageQueue,
			redisParams.Redis,
			agParams.SDK,
			configs,
			setting,
		); e != nil {
			err = e
		} else {
			out.ZinxService = svc
		}
		return
	},
)

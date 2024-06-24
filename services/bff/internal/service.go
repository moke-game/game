package internal

import (
	"context"
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/znet"

	"github.com/moke-game/game/services/common/constants"

	"github.com/gstones/moke-kit/orm/nosql/key"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"

	configs "github.com/moke-game/game/configs/pkg/module"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/bff/internal/biface"
	"github.com/moke-game/game/services/bff/internal/routers"
	"github.com/moke-game/game/services/bff/pkg/bfx"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/pack"
)

const (
	TickTime = time.Minute
)

type Service struct {
	znet.BaseRouter
	logger    *zap.Logger
	url       string
	hostName  string
	routers   []biface.IRouter
	registers *Registers
	redis     *redis.Client
	msgIdHash map[int32]int32
	meter     metric.Meter
}

func (s *Service) RegisterWithServer(server siface.IZinxServer) {
	s.hashEventIds()
	for _, v := range s.routers {
		v.Register(s.registers)
	}
	for k := range s.registers.GetHandlers() {
		server.ZinxServer().AddRouter(uint32(k), s)
	}

	server.ZinxServer().SetOnConnStart(s.onConnStart)
	server.ZinxServer().SetOnConnStop(s.onConnStop)
	server.ZinxServer().SetPacket(pack.NewDataPack())
	server.ZinxServer().SetDecoder(pack.NewLTEV_Decoder())
	s.UpdateOnline(server)
}

func (s *Service) onConnStart(connection ziface.IConnection) {
	name := fmt.Sprintf("bff.online.counts.%s", s.hostName)
	if counter, err := s.meter.Int64UpDownCounter(
		name,
		metric.WithDescription("bff online player counts"),
		metric.WithUnit("{online}"),
	); err != nil {
		s.logger.Error("Int64UpDownCounter err", zap.Error(err))
	} else {
		counter.Add(connection.Context(), 1)
	}
	s.logger.Info("onConnStart", zap.String("remoteAddr", connection.RemoteAddr().String()))
}

func (s *Service) onConnStop(connection ziface.IConnection) {
	name := fmt.Sprintf("bff.online.counts.%s", s.hostName)
	if counter, err := s.meter.Int64UpDownCounter(
		name,
		metric.WithDescription("bff online player counts"),
		metric.WithUnit("{online}"),
	); err != nil {
		s.logger.Error("Int64UpDownCounter err", zap.Error(err))
	} else {
		counter.Add(connection.Context(), -1)
	}
	uid := ""
	if u, err := connection.GetProperty(constants.ConnUid); err == nil {
		uid = u.(string)
	}

	s.logger.Info("onConnStop",
		zap.String("remoteAddr", connection.RemoteAddr().String()),
		zap.String("uid", uid),
	)
}

func (s *Service) UpdateOnline(server siface.IZinxServer) {
	go func() {
		ticker := time.NewTicker(TickTime)
		for {
			<-ticker.C
			online := server.ZinxServer().GetConnMgr().Len()
			sec := time.Now().UTC().Unix()
			hour := sec / (60 * 60)
			min := sec / 60
			onlineKey, err := MakeOnlineKey(strconv.FormatInt(hour, 10))
			if err != nil {
				s.logger.Error("UpdateOnline err", zap.Error(err))
				continue
			}
			serverKey := s.hostName + "_" + strconv.FormatInt(min, 10)
			if res := s.redis.HSet(context.TODO(), onlineKey.String(), serverKey, online); res.Err() != nil {
				s.logger.Error("UpdateOnline HSet err", zap.Error(res.Err()))
				continue
			}
			s.redis.Expire(context.TODO(), onlineKey.String(), time.Hour*24*15)
		}
	}()

}

func MakeOnlineKey(hour string) (key.Key, error) {
	return key.NewKeyFromParts("online", hour)
}

func (s *Service) hashEventIds() {
	for k, v := range cpb.C2S_EVENT_name {
		name, ok := strings.CutPrefix(v, "C2S_")
		if !ok {
			s.logger.Warn("invalid event name", zap.String("name", v))
			continue
		}
		if name != "" {
			s2cKey := fmt.Sprintf("S2C_%s", name)
			if id, ok := cpb.S2C_EVENT_value[s2cKey]; ok {
				s.msgIdHash[k] = id
			}
		}
	}
	// chttodo 暂时处理
	s.msgIdHash[20002] = 35028
}

func (s *Service) logRequest(uid string, request ziface.IRequest) {
	if request.GetMsgID() == uint32(cpb.C2S_EVENT_C2S_Heartbeat) {
		return
	}
	if name, ok := cpb.C2S_EVENT_name[int32(request.GetMsgID())]; ok {
		s.logger.Info(
			"bff request",
			zap.String("uid", uid),
			zap.String("msg", name),
			zap.Any("data", request.GetData()),
		)
	} else {
		s.logger.Info(
			"bff request",
			zap.String("uid", uid),
			zap.Uint32("msgID", request.GetMsgID()),
			zap.Any("data", request.GetData()),
		)
	}
}

func (s *Service) logResponse(uid string, respId int32, code cpb.ERRORCODE, msg proto.Message) {
	if respId == int32(cpb.S2C_EVENT_S2C_Heartbeat) {
		return
	}
	if name, ok := cpb.S2C_EVENT_name[respId]; ok {
		s.logger.Info(
			"bff response",
			zap.String("uid", uid),
			zap.String("msgID", name),
			zap.Any("data", msg),
			zap.Any("code", code),
		)
	} else {
		s.logger.Info(
			"bff response",
			zap.String("uid", uid),
			zap.Uint32("msgID", uint32(respId)),
			zap.Any("data", msg),
			zap.Any("code", code),
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
	request.GetConnection().GetConnID()
	uid := ""
	if u, err := request.GetConnection().GetProperty(constants.ConnUid); err == nil {
		uid = u.(string)
	}

	eventId := cpb.C2S_EVENT(request.GetMsgID())
	if handle := s.registers.GetHandler(eventId); handle == nil {
		s.logger.Error("can not find handler", zap.String("msgID", eventId.String()))
	} else {
		s.logRequest(uid, request)
		resp, code := handle(request.GetConnection().Context(), request)
		if resp == nil && code == cpb.ERRORCODE_SUCCESS {
			s.logger.Warn("handler return nil response", zap.String("msgID", eventId.String()))
			return
		} else if respId, ok := s.msgIdHash[int32(request.GetMsgID())]; !ok {
			s.logger.Error("can not find response msg id", zap.String("msgID", eventId.String()))
		} else if err := common.SendResponse(request.GetConnection(), uint32(respId), code, resp); err != nil {
			s.logger.Error("send response failed", zap.Error(err))
		} else {
			s.logResponse(uid, respId, code, resp)
		}
	}
}

func NewService(
	l *zap.Logger,
	url string,
	rs []biface.IRouter,
	hostName string,
	redisClient *redis.Client,
) (result *Service, err error) {
	result = &Service{
		logger:    l,
		url:       url,
		hostName:  hostName,
		routers:   rs,
		registers: NewRegisters(),
		msgIdHash: make(map[int32]int32),
		redis:     redisClient,
		meter:     otel.Meter(hostName),
	}
	return
}

var Module = fx.Provide(
	func(
		l *zap.Logger,
		setting bfx.GameSettingParams,
		routers routers.RouterParams,
		redisParams ofx.RedisParams,
	) (out sfx.ZinxServiceResult, err error) {
		hostName := setting.HostName
		if len(hostName) == 0 || hostName == "bff" {
			hostName += strconv.FormatInt(time.Now().Unix(), 10)
		}
		if svc, e := NewService(
			l,
			setting.GameUrl,
			routers.Routers,
			hostName,
			redisParams.Redis,
		); e != nil {
			err = e
		} else {
			out.ZinxService = svc
		}
		return
	},
)

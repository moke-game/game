package room

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/zpack"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	kpb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/leaderboard"
	match "github.com/moke-game/platform/api/gen/matchmaking"
	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/matchmaking/pkg/module/data"

	"github.com/moke-game/game/services/room/pkg/rfx"

	cpb "github.com/moke-game/game/api/gen/common"
	_ "github.com/moke-game/game/configs/code"
	"github.com/moke-game/game/configs/pkg/cfx"
	configs "github.com/moke-game/game/configs/pkg/module"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/room/internal/room/riface"
)

const (
	FrameRate      = 20                                  // 帧率
	TickerInterval = 1000 * time.Millisecond / FrameRate //基于上面定义的帧率推算每帧的间隔
	ReadBuffCap    = 1024
)

type RecMsg struct {
	uid string
	msg ziface.IRequest
}

type Room struct {
	logger   *zap.Logger
	roomId   string
	mapId    int32
	playId   int32
	handlers map[uint32]riface.IHandler

	pClient        pb.ProfileServiceClient
	pPrivateClient pb.ProfilePrivateServiceClient
	knapsackClient kpb.KnapsackPrivateServiceClient
	matchClient    match.MatchServiceClient
	lbClient       leaderboard.LeaderboardPrivateServiceClient
	rClient        *redis.Client
	configs        cfx.ConfigsParams
	msgSender      *MsgSender
	msgCh          chan *RecMsg
	mq             miface.MessageQueue
	isRunning      *atomic.Bool
	isOver         bool
	duration       time.Duration
}

func (r *Room) Receive(uid string, message ziface.IRequest) {
	r.msgCh <- &RecMsg{uid: uid, msg: message}
}

func (r *Room) MapId() int32 {
	return r.mapId
}

func NewRoom(
	roomId string,
	logger *zap.Logger,
	pClient pb.ProfileServiceClient,
	pPrivateClient pb.ProfilePrivateServiceClient,
	knapsackClient kpb.KnapsackPrivateServiceClient,
	matchClient match.MatchServiceClient,
	configs cfx.ConfigsParams,
	rClient *redis.Client,
	lbClient leaderboard.LeaderboardPrivateServiceClient,
	mq miface.MessageQueue,
	setting rfx.RoomSettingParams,
) (*Room, error) {
	return &Room{
		roomId:         roomId,
		logger:         logger,
		pClient:        pClient,
		pPrivateClient: pPrivateClient,
		rClient:        rClient,
		knapsackClient: knapsackClient,
		matchClient:    matchClient,
		configs:        configs,
		isRunning:      atomic.NewBool(false),
		lbClient:       lbClient,
		mq:             mq,
	}, nil
}

func (r *Room) Init(result *data.MatchResult) error {
	playId := result.PlayId
	r.playId = playId
	r.msgCh = make(chan *RecMsg, ReadBuffCap)
	if hub, err := CreateMsgHub(r.logger); err != nil {
		r.logger.Error("create msg hub error", zap.Error(err))
		return err
	} else {
		r.msgSender = hub
	}
	r.registerHandlers()
	return nil
}

func (r *Room) registerHandlers() {
	r.handlers = make(map[uint32]riface.IHandler)
	r.handlers[uint32(cpb.C2S_EVENT_C2S_EnterRoom)] = r.enterRoom
	r.handlers[uint32(cpb.C2S_EVENT_C2S_LeaveRoom)] = r.exitRoom
	r.handlers[uint32(cpb.C2S_EVENT_C2S_Move)] = r.move
}

func (r *Room) recover() {
	if !configs.DeploymentGlobal.IsProd() {
		return
	}
	if e := recover(); e != nil {
		r.logger.Error("panic",
			zap.Any("recover", e),
			zap.String("stack", string(debug.Stack())),
		)
	}
}

func (r *Room) Handle(uid string, request ziface.IRequest) {
	defer r.recover()
	if h, ok := r.handlers[request.GetMsgID()]; !ok {
		r.logger.Error(
			"room handle can not found msg id",
			zap.String("roomId", r.roomId),
			zap.Uint32("msgId", request.GetMsgID()),
		)
	} else {
		requestId := request.GetMsgID()
		if resp, errCode := h(uid, request); resp == nil {
			return
		} else if msgName, ok := cpb.S2C_EVENT_name[int32(requestId+1)]; !ok {
			r.logger.Error("room handle can not found response msg id", zap.Uint32("msgId", requestId+1))
		} else if err := common.SendResponse(request.GetConnection(), requestId+1, errCode, resp); err != nil {
			if err.Error() == "Connection closed when send buff msg" {
				r.logger.Info("client connection already closed", zap.String("uid", uid))
			} else {
				r.logger.Error("room send response failed", zap.Error(err))
			}
		} else if requestId != uint32(cpb.C2S_EVENT_C2S_RoomHeartbeat) &&
			requestId != uint32(cpb.C2S_EVENT_C2S_Move) &&
			requestId != uint32(cpb.C2S_EVENT_C2S_MovePrepare) &&
			requestId != uint32(cpb.C2S_EVENT_C2S_StopMoving) {
			r.logger.Info(
				"room send response",
				zap.String("roomId", r.roomId),
				zap.String("uid", uid),
				zap.String("msgName", msgName),
				zap.Any("data", resp),
				zap.Int32("errorCode", int32(errCode)),
			)
		}
	}
}

func (r *Room) Run() error {
	if r.isRunning.Load() {
		return fmt.Errorf("room:%s is already running", r.roomId)
	}
	ticker := time.NewTicker(TickerInterval)
	defer ticker.Stop()
	defer r.Destroy()
	lastTime := time.Now()
	for {
		now := <-ticker.C
		duration := now.Sub(lastTime)
		lastTime = now
		if r.CheckCountdown(duration) {
			return nil
		}
		if r.Update(duration) {
			r.GameEnd(duration)
		}
		r.RoomHeart()
	}
}
func (r *Room) GameEnd(duration time.Duration) {
	r.isOver = true
	r.duration = 0
}

func (r *Room) RoomHeart() {
}

func (r *Room) CheckCountdown(duration time.Duration) bool {
	if r.isOver {
		r.duration -= duration
		if r.duration <= 0 {
			r.logger.Debug("room is end, destroy it", zap.String("roomId", r.roomId))
			return true
		}
	}
	return false
}

func (r *Room) Destroy() {

}

func (r *Room) RoomId() string {
	return r.roomId
}

func (r *Room) Broadcast(id cpb.S2C_EVENT, msg proto.Message) {
	if d, err := proto.Marshal(msg); err != nil {
		r.logger.Error("room broadcast marshal failed", zap.Error(err))
	} else {
		pack := zpack.NewMsgPackage(uint32(id), d)
		r.msgSender.BroadcastExclude(pack)
	}
}

func (r *Room) Update(dt time.Duration) bool {
	r.consumeMsg()
	// TODO game update with dt
	return false
}

func (r *Room) consumeMsg() {
	if len(r.msgCh) <= 0 {
		return
	}
	msgRec, _, _, _ := lo.Buffer(r.msgCh, len(r.msgCh))
	for _, v := range msgRec {
		r.Handle(v.uid, v.msg)
	}
}

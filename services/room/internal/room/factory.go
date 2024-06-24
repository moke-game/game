package room

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	kpb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/leaderboard"
	match "github.com/moke-game/platform/api/gen/matchmaking"

	"github.com/moke-game/game/services/room/pkg/rfx"

	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/matchmaking/pkg/module/data"

	"github.com/moke-game/game/configs/pkg/cfx"
	"github.com/moke-game/game/services/room/internal/room/riface"
)

func CreateRoom(
	roomId string,
	logger *zap.Logger,
	pClient pb.ProfileServiceClient,
	pPrivateClient pb.ProfilePrivateServiceClient,
	kClient kpb.KnapsackPrivateServiceClient,
	matchClient match.MatchServiceClient,
	lbClient leaderboard.LeaderboardPrivateServiceClient,
	configs cfx.ConfigsParams,
	rClient *redis.Client,
	mq miface.MessageQueue,
	result *data.MatchResult,
	setting rfx.RoomSettingParams,
) (riface.IRoom, error) {
	if room, err := NewRoom(
		roomId,
		logger,
		pClient,
		pPrivateClient,
		kClient,
		matchClient,
		configs,
		rClient,
		lbClient,
		mq,
		setting,
	); err != nil {
		return nil, err
	} else if err := room.Init(result); err != nil {
		return nil, err
	} else {
		return room, nil
	}
}

func CreateMsgHub(logger *zap.Logger) (*MsgSender, error) {
	msg := &MsgSender{
		logger: logger,
	}
	if err := msg.Init(); err != nil {
		return nil, err
	}
	return msg, nil
}

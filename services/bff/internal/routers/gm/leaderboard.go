package gm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"

	"github.com/moke-game/game/configs/pkg/module"
	leaderboard2 "github.com/moke-game/game/services/common/leaderboard"
	"github.com/moke-game/platform/api/gen/leaderboard"
)

// addLeaderboards .
func (c *Router) addLeaderboards(request ziface.IRequest, args ...string) error {
	if len(args) < 2 {
		c.logger.Error("args len error", zap.Strings("args", args))
		return fmt.Errorf("args:%v len error", args)
	}
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		c.logger.Error("id error", zap.Error(err))
		return fmt.Errorf("id:%s error:%v", args[0], err)
	}

	num, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		c.logger.Error("num error", zap.Error(err))
		return fmt.Errorf("num:%s error:%v", args[1], err)
	}
	conf := module.ConfigsGlobal.RankConfigs.GetRankInfo(int32(id))
	if conf == nil {
		c.logger.Error("rank config not found", zap.Int32("id", int32(id)))
		return fmt.Errorf("rank config not found:%d", int32(id))
	}
	conf2 := module.ConfigsGlobal.TblRankConfig.Get(conf.ID)
	if conf2 == nil {
		c.logger.Error("rank config not found", zap.Int32("id", conf.ID))
		return fmt.Errorf("rank config not found:%d", conf.ID)
	}

	scores := make(map[string]float64)
	if conf2.WorldRank == 1 {
		// random country codes with scores
		for i := 1; i < int(num)+1; i++ {
			code := randomdata.Country(randomdata.TwoCharCountry)
			scores[code] = float64(i)
		}
	} else {
		for i := 1; i < int(num)+1; i++ {
			us := fmt.Sprintf("%d", i)
			scores[us] = float64(i)
		}
	}

	period := leaderboard2.MakeLeaderboardPeriod(int32(id))
	_, err = c.lbPrivateClient.UpdateScore(context.Background(), &leaderboard.UpdateScoreRequest{
		Id:     period,
		Scores: scores,
	})
	if err != nil {
		c.logger.Error("update score error", zap.Error(err))
		return fmt.Errorf("update score error:%v", err)
	}

	return nil
}

func (c *Router) updateSelfScore(request ziface.IRequest, args ...string) error {
	if len(args) < 3 {
		c.logger.Error("args len error", zap.Strings("args", args))
		return fmt.Errorf("args:%v len error", args)
	}
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		c.logger.Error("id error", zap.Error(err))
		return fmt.Errorf("id:%s error:%v", args[0], err)
	}

	uid := args[1]
	updateType, err := strconv.Atoi(args[2])
	if err != nil {
		c.logger.Error("updateType error", zap.Error(err))
		return fmt.Errorf("updateType:%s error:%v", args[2], err)
	}

	score, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		c.logger.Error("score error", zap.Error(err))
		return fmt.Errorf("score:%s error:%v", args[3], err)
	}

	period := leaderboard2.MakeLeaderboardPeriod(int32(id))
	_, err = c.lbPrivateClient.UpdateScore(context.Background(), &leaderboard.UpdateScoreRequest{
		UpdateType: leaderboard.UpdateScoreRequest_UpdateType(updateType),
		Id:         period,
		Scores:     map[string]float64{uid: score},
	})
	if err != nil {
		c.logger.Error("update score error", zap.Error(err))
		return err
	}

	return nil
}

func (c *Router) clearLeaderboards(request ziface.IRequest, args ...string) error {
	if len(args) < 1 {
		c.logger.Error("args len error", zap.Strings("args", args))
		return fmt.Errorf("args:%v len error", args)
	}
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		c.logger.Error("id error", zap.Error(err))
		return fmt.Errorf("id:%s error:%v", args[0], err)
	}
	period := leaderboard2.MakeLeaderboardPeriod(int32(id))
	if _, err := c.lbPrivateClient.ClearLeaderboard(context.Background(), &leaderboard.ClearLeaderboardRequest{
		Id: period,
	}); err != nil {
		c.logger.Error("clear leaderboard error", zap.Error(err))
		return err
	}
	return nil
}

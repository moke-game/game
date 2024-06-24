package leaderboard

import (
	"fmt"
	"time"

	"github.com/moke-game/game/configs/pkg/module"
)

func CalculateLeftDuration(id int32) (time.Duration, error) {
	conf := module.ConfigsGlobal.RankConfigs.GetRankInfo(id)
	if conf == nil {
		return time.Duration(0), fmt.Errorf("rank config not found")
	}
	now := time.Now().Unix()
	if conf.EndTime > 0 && now >= conf.EndTime {
		return time.Duration(0), fmt.Errorf("rank config end time is over")
	}
	mod := time.Duration(conf.Period) * time.Hour
	duration := time.Duration(now-conf.BeginTime) * time.Second
	left := int64(duration.Seconds()) % int64(mod.Seconds())
	return mod - time.Duration(left)*time.Second, nil
}

func MakeLeaderboardPeriod(id int32) string {
	conf := module.ConfigsGlobal.RankConfigs.GetRankInfo(id)
	if conf == nil {
		return ""
	}
	now := time.Now().Unix()
	if conf.EndTime > 0 && now >= conf.EndTime {
		now = conf.EndTime - 100
	}

	duration := now - conf.BeginTime
	hour := duration / 3600

	period := hour / int64(conf.Period)
	return fmt.Sprintf("%d_%d_%d", id, conf.ID, period)
}

// CalculateOpenHours 计算当前期数排行榜已经开始多少小时
func CalculateOpenHours(id int32) (int32, error) {
	conf := module.ConfigsGlobal.RankConfigs.GetRankInfo(id)
	if conf == nil {
		return 0, fmt.Errorf("rank config not found")
	}
	now := time.Now().Unix()
	if conf.BeginTime <= 0 || now < conf.BeginTime {
		return 0, fmt.Errorf("rank not start")
	}
	left := int32((now - conf.BeginTime) / (60 * 60))
	hours := left % conf.Period
	return hours, nil
}

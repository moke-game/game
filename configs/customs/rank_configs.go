package customs

import (
	"time"

	cfg "github.com/moke-game/game/configs/code"
)

type RankInfo struct {
	ID        int32
	BeginTime int64
	EndTime   int64
	Period    int32
	IsCountry bool
}

type RankConfigs struct {
	Leaderboards    map[int32][]*RankInfo
	LeaderboardsIds map[int32]int32
}

func CreateRankConfigs() *RankConfigs {
	return &RankConfigs{
		Leaderboards:    make(map[int32][]*RankInfo),
		LeaderboardsIds: make(map[int32]int32),
	}
}

func (n *RankConfigs) Init(tbl *cfg.TblRankConfig) {
	for _, v := range tbl.GetDataList() {
		btime, err := time.ParseInLocation("2006-01-02 15:04:05", v.BeginTime, time.Local)
		if err != nil {
			panic(err)
		}
		eUnix := int64(0)
		if v.EndTime != "" {
			etime, err := time.ParseInLocation("2006-01-02 15:04:05", v.EndTime, time.Local)
			if err != nil {
				panic(err)
			}
			eUnix = etime.Unix()
		}
		if _, ok := n.Leaderboards[v.RankContent]; !ok {
			n.Leaderboards[v.RankContent] = make([]*RankInfo, 0)
		}
		info := &RankInfo{
			ID:        v.ID,
			BeginTime: btime.Unix(),
			EndTime:   eUnix,
			Period:    v.BillingCycle,
			IsCountry: v.WorldRank == 1,
		}
		n.Leaderboards[v.RankContent] = append(n.Leaderboards[v.RankContent], info)
	}
}

func (n *RankConfigs) GetRankInfo(tp int32) *RankInfo {
	ranks, ok := n.Leaderboards[tp]
	if !ok {
		return nil
	}

	for _, v := range ranks {
		now := time.Now().Unix()
		if v.BeginTime <= now && (v.EndTime == 0 || now <= v.EndTime) {
			return v
		}
	}
	return ranks[len(ranks)-1]
}

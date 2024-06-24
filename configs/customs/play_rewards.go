package customs

import cfg "github.com/moke-game/game/configs/code"

type PlayRewards struct {
	rewards map[int32]map[int32][]*cfg.BattleSpecialReward
}

func CreatePlayRewards() *PlayRewards {
	return &PlayRewards{
		rewards: make(map[int32]map[int32][]*cfg.BattleSpecialReward),
	}
}

func (pr *PlayRewards) Init(tbl *cfg.TblPlaySelect) {
	rewards := make(map[int32][]*cfg.BattleSpecialReward)
	for _, data := range tbl.GetDataList() {
		for _, v := range data.Rewards3 {
			rewards[v.MapId] = append(rewards[v.MapId], v)
		}
		pr.rewards[data.ID] = rewards
	}
}

func (pr *PlayRewards) GetRewards(id, mapId, index int32) int32 {
	if rewards, ok := pr.rewards[id]; ok {
		if rds, ok := rewards[mapId]; ok {
			if index >= int32(len(rds)) {
				return 0
			}
			return rds[index].DropId
		}
	}
	return 0
}

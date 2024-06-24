package msg_transfer

import (
	"strconv"

	"github.com/moke-game/platform/api/gen/leaderboard"
	profile "github.com/moke-game/platform/api/gen/profile"

	bff "github.com/moke-game/game/api/gen/bff"
)

func TransferLeaderboard(entries []*leaderboard.LeaderboardEntry, infos map[string]*profile.ProfileBasic) []*bff.LeaderboardEntry {
	res := make([]*bff.LeaderboardEntry, 0)
	for _, entry := range entries {
		country := ""
		uidInt64 := int64(0)
		if uid, err := strconv.ParseInt(entry.Uid, 10, 64); err != nil {
			country = entry.Uid
		} else {
			uidInt64 = uid
		}

		ele := &bff.LeaderboardEntry{
			Uid:     uidInt64,
			Score:   int32(entry.Score),
			Country: country,
			Star:    int32(entry.Star),
		}
		if basic, ok := infos[entry.Uid]; ok {
			ele.Nickname = basic.GetNickname()
			ele.Avatar = basic.GetAvatar()
		}
		res = append(res, ele)
	}
	return res
}

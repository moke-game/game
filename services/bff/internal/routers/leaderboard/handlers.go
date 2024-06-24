package leaderboard

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/moke-game/platform/api/gen/leaderboard"
	"github.com/moke-game/platform/api/gen/mail"
	profile "github.com/moke-game/platform/api/gen/profile"

	"github.com/moke-game/game/services/common/constants"

	"github.com/gstones/zinx/ziface"

	bffpb "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	cfg "github.com/moke-game/game/configs/code"
	"github.com/moke-game/game/configs/pkg/module"
	"github.com/moke-game/game/services/common"
	leaderboard2 "github.com/moke-game/game/services/common/leaderboard"
	"github.com/moke-game/game/services/common/msg_transfer"
)

func (r *Router) getLeaderboard(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SGetLeaderboard{}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_PROTO_UNMARSHAL_ERROR
	}
	conf := module.ConfigsGlobal.RankConfigs.GetRankInfo(req.Id)
	if conf == nil {
		r.logger.Error("rank config not found", zap.Int32("period", req.Id))
		return nil, cpb.ERRORCODE_CONFIG_NOT_FOUND
	}
	country := ""
	size := int32(100)
	if conf.IsCountry {
		if code, err := request.GetConnection().GetProperty(constants.ConnCountryCode); err == nil {
			country = code.(string)
		}
		if country == "" {
			country = "UNKNOWN"
		}
		size = 10
	}

	if period := leaderboard2.MakeLeaderboardPeriod(req.Id); period == "" {
		r.logger.Error("have no period can use", zap.Int32("rank content", req.Id))
		return nil, cpb.ERRORCODE_COMMON_ERROR
	} else if resp, err := r.lbClient.GetLeaderboard(ctx, &leaderboard.GetLeaderboardRequest{
		Id:       period,
		PageSize: size,
	}); err != nil {
		r.logger.Error("get leaderboard failed", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else if len(resp.Entries) <= 0 {
		return &bffpb.S2CGetLeaderboard{
			Id: req.Id,
		}, cpb.ERRORCODE_SUCCESS
	} else if selfRank, err := r.lbClient.GetRank(ctx, &leaderboard.GetRankRequest{Id: period, Country: country}); err != nil {
		r.logger.Error("get self rank failed", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		basics := make(map[string]*profile.ProfileBasic)
		if !conf.IsCountry {
			basics = r.getProfileBasicsInfo(ctx, resp.Entries)
		}
		entries := msg_transfer.TransferLeaderboard(resp.Entries, basics)
		return &bffpb.S2CGetLeaderboard{
			Entries: entries,
			SelfEntry: &bffpb.SelfEntry{
				Rank:    int32(selfRank.Rank),
				Score:   int32(selfRank.Score),
				Country: country,
			},
			Id: req.Id,
		}, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) getProfileBasicsInfo(ctx context.Context, entries []*leaderboard.LeaderboardEntry) map[string]*profile.ProfileBasic {
	uids := make([]string, 0)
	for _, v := range entries {
		uids = append(uids, v.Uid)
	}
	resp, err := r.profilePrivateClient.GetProfileBasics(ctx, &profile.GetProfileBasicsRequest{
		Uid: uids,
	})
	if err != nil {
		r.logger.Error("get profile basics failed", zap.Error(err))
		return nil
	}
	return resp.Basics

}

func (r *Router) trySettlementRewards(period string, id int32) error {
	if resp, err := r.lbPrivateClient.ExpireLeaderboard(
		context.Background(),
		&leaderboard.ExpireLeaderboardRequest{
			Id:  period,
			Num: 5000,
		},
	); err != nil {
		return err
	} else if resp.IsDeleted {
		conf := module.ConfigsGlobal.TblRankConfig.Get(id)
		go r.sendLeaderboardRewards(resp.Entries, conf)
	}
	return nil
}

func (r *Router) sendLeaderboardRewards(entries []*leaderboard.LeaderboardEntry, cfg *cfg.RankConfig) {
	if cfg == nil || len(cfg.RankReward) <= 0 {
		return
	}
	rewards, ranks := r.makeRewards(entries, cfg)
	for k, v1 := range rewards {
		if _, err := r.mailClient.SendMail(
			context.Background(),
			&mail.SendMailRequest{
				SendType: mail.SendMailRequest_ROLE,
				RoleIds:  []string{k},
				Mail: &mail.Mail{
					TemplateId:   cfg.MailID,
					TemplateArgs: []string{ranks[k]},
					Rewards:      v1,
				},
			},
		); err != nil {
			r.logger.Error(
				"send leaderboard mail rewards failed",
				zap.Error(err),
				zap.String("uid", k),
				zap.Any("rewards", v1),
			)
		}
	}
}

func (r *Router) makeRewards(
	entries []*leaderboard.LeaderboardEntry,
	cfg *cfg.RankConfig,
) (map[string][]*mail.MailReward, map[string]string) {
	rewards := make(map[string][]*mail.MailReward)
	ranks := make(map[string]string)
	for i, v := range entries {
		index := getIndexByRank(int32(i+1), cfg.RankNum)
		if index == -1 {
			continue
		}
		ranks[v.Uid] = strconv.Itoa(i + 1)
		if index >= len(cfg.RankReward) {
			r.logger.Warn(
				"rank reward index out of range",
				zap.Int32("index", int32(i+1)),
				zap.Int32("rank", cfg.ID),
				zap.Any("rankNum", cfg.RankNum),
				zap.Any("rankReward", cfg.RankReward),
			)
			continue
		}
		rewardBox := cfg.RankReward[index]
		rbConfig := module.ConfigsGlobal.TblItemBox.Get(rewardBox)
		items := msg_transfer.Configs2MailItems(rbConfig.Item)
		rewards[v.Uid] = items
	}
	return rewards, ranks

}

func getIndexByRank(rank int32, rankRange []int32) int {
	for i, v := range rankRange {
		if rank <= v {
			return i
		}
	}
	return -1
}

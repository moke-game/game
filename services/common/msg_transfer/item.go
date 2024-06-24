package msg_transfer

import (
	pb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/api/gen/mail"

	bffpb "github.com/moke-game/game/api/gen/bff"
	cfg "github.com/moke-game/game/configs/code"
)

func Items2Knapsack(items *bffpb.Items) map[int64]*pb.Item {
	res := make(map[int64]*pb.Item)
	for k, v := range items.Items {
		res[k] = Item2Knapsack(v)
	}
	return res
}

func Item2Knapsack(item *bffpb.Item) *pb.Item {
	return &pb.Item{
		Id:     item.Id,
		Type:   item.ConfigId,
		Expire: item.Expired,
		Num:    item.Num,
	}
}

func Configs2MailItems(rewards []*cfg.ItemReward) []*mail.MailReward {
	var mailItems []*mail.MailReward
	for _, item := range rewards {
		mailItems = append(mailItems, &mail.MailReward{Id: int64(item.Id), Num: item.Num, Type: item.Id})
	}
	return mailItems
}

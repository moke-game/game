package knapsack

import (
	"context"
	"fmt"
	"math"

	common2 "github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/gstones/zinx/ziface"

	"github.com/moke-game/game/services/common/constants"

	kpb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/services/knapsack/changes"

	bpb "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common"
)

func (r *Router) watchingKnapsack(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bpb.C2SWatchingKnapsack{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}

	uid, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		r.logger.Error("get uid from context err")
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if resp, err := r.kClient.GetKnapsack(ctx, &kpb.GetKnapsackRequest{}); err != nil {
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		topic := changes.CreateTopic(uid.(string))
		if _, err := r.mq.Subscribe(ctx, topic, func(msg miface.Message, err error) common2.ConsumptionCode {
			if msg, err := changes.UnPack(msg.Data()); err != nil {
				r.logger.Error("unpack knapsack changes failed", zap.Error(err))
			} else {
				watchResp := r.toWatchModifyResp(ctx, request, msg)
				if err := common.SendResponse(
					request.GetConnection(),
					uint32(cpb.S2C_EVENT_S2C_WatchingKnapsack),
					0,
					watchResp,
				); err != nil {
					if err.Error() == "connection closed when send buff msg" {
						r.logger.Warn("connection closed when send knapsack changes buff msg")
					} else {
						r.logger.Error("send knapsack changes failed", zap.Error(err))
					}
				} else {
					r.logger.Info(
						"watch knapsack changes success",
						zap.String("uid", uid.(string)),
						zap.String("changes", watchResp.String()),
					)
				}
				if watchResp.Knapsack.Features != nil {
					request.GetConnection().SetProperty(constants.ConnFeatures, watchResp.Knapsack.Features)
				}
			}
			return common2.ConsumeAck
		}); err != nil {
			r.logger.Error("subscribe knapsack changes failed", zap.Error(err))
			return nil, cpb.ERRORCODE_COMMON_ERROR
		}
		msg := r.toWatchResp(resp.Knapsack)
		request.GetConnection().SetProperty(constants.ConnFeatures, resp.Knapsack.Features)
		return msg, cpb.ERRORCODE_SUCCESS
	}
}
func (r *Router) toWatchModifyResp(ctx context.Context, request ziface.IRequest, msg *kpb.KnapsackModify) *bpb.S2CWatchingKnapsack {
	items := make(map[int64]*bpb.Item)
	for k, v := range msg.Knapsack.GetItems() {
		items[k] = &bpb.Item{
			Id:       v.Id,
			Num:      v.Num,
			Expired:  v.Expire,
			ConfigId: v.Type,
		}
	}
	return &bpb.S2CWatchingKnapsack{
		Knapsack: &bpb.MKnapsack{
			Items:    items,
			Features: msg.Knapsack.Features,
		},
	}
}

func (r *Router) toWatchResp(msg *kpb.Knapsack) *bpb.S2CWatchingKnapsack {
	items := make(map[int64]*bpb.Item)
	for k, v := range msg.GetItems() {
		items[k] = &bpb.Item{
			Id:       v.Id,
			Num:      v.Num,
			Expired:  v.Expire,
			ConfigId: v.Type,
		}
	}
	return &bpb.S2CWatchingKnapsack{
		Knapsack: &bpb.MKnapsack{
			Items:    items,
			Features: msg.Features,
		},
	}
}

func (r *Router) diamondExchangeItem(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bpb.C2SDiamondExchangeItem{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if remove, add, err := r.calculateExchangeDiamondPrice(req.Item); err != nil {
		return nil, cpb.ERRORCODE_CONFIG_NOT_FOUND
	} else if _, err := r.kClient.RemoveThenAddItem(ctx, &kpb.RemoveThenAddItemRequest{
		RemoveItems: map[int64]*kpb.Item{
			remove.Id: remove,
		},
		AddItems: add,
	}); err != nil {
		r.logger.Error("diamondExchangeItem RemoveThenAddItem err", zap.Any("removes items", remove), zap.Error(err))
		return nil, cpb.ERRORCODE_NOT_ENOUGH
	}
	return &bpb.S2CDiamondExchangeItem{}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) calculateExchangeDiamondPrice(items *bpb.Items) (*kpb.Item, map[int64]*kpb.Item, error) {
	cost := &kpb.Item{
		Id:   int64(bpb.ItemType_kDiamond),
		Type: int32(bpb.ItemType_kDiamond),
	}

	exchanged := map[int64]*kpb.Item{}
	for _, v := range items.GetItems() {
		conf := r.configs.TblItem.Get(v.ConfigId)
		if conf == nil {
			return nil, nil, fmt.Errorf("calculateExchangeGoldPrice GetItemByID id %d not found", v.ConfigId)
		}
		if v.Num <= 0 {
			continue
		}
		if conf.DiamondExchange <= 0 {
			return nil, nil, fmt.Errorf("item:%d can not exchange by diamond", v.ConfigId)
		}
		// 万分比
		diamondPercent := v.Num * conf.DiamondExchange
		if diamondPercent < 10000 {
			diamondPercent = 10000
		}
		dNum := math.Ceil(float64(diamondPercent) / 10000)
		cost.Num += int32(dNum)

		if _, ok := exchanged[int64(v.ConfigId)]; ok {
			exchanged[int64(v.ConfigId)].Num += v.Num
		} else {
			ex := &kpb.Item{
				Id:   int64(v.ConfigId),
				Type: v.ConfigId,
				Num:  v.Num,
			}
			exchanged[int64(v.ConfigId)] = ex
		}
	}
	return cost, exchanged, nil
}

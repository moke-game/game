package gm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"

	kpb "github.com/moke-game/platform/api/gen/knapsack"
)

// addCurrency .
func (c *Router) addCurrency(request ziface.IRequest, args ...string) error {
	if len(args) < 3 {
		return fmt.Errorf("args:%v error", args)
	}
	uid := args[0]
	currencyType := args[1]
	currencyCount := args[2]

	cType, _ := strconv.ParseInt(currencyType, 10, 64)
	cCount, _ := strconv.ParseInt(currencyCount, 10, 64)
	item := &kpb.Item{
		Id:   cType,
		Type: int32(cType),
		Num:  int32(cCount),
	}
	_, err := c.kpClient.AddItem(context.TODO(), &kpb.AddItemPrivateRequest{
		Uid:    uid,
		Items:  map[int64]*kpb.Item{cType: item},
		Source: "gm",
	})
	if err != nil {
		c.logger.Error("add currency err", zap.String("uid", uid), zap.Error(err))
		return err
	}
	return nil
}

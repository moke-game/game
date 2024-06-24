package gm

import (
	"context"
	"fmt"

	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"

	kpb "github.com/moke-game/platform/api/gen/knapsack"

	"github.com/moke-game/game/configs/pkg/module"
)

// addLeaderboards .
func (c *Router) unlockFeat(request ziface.IRequest, args ...string) error {
	if len(args) < 1 {
		c.logger.Error("args len error", zap.Strings("args", args))
		return fmt.Errorf("args:%v len error", args)
	}
	uid := args[0]

	features := make(map[int32]bool)
	for _, v := range module.ConfigsGlobal.TblFeature.GetDataList() {
		if v.DefaultSystem == 1 {
			continue
		}
		features[v.ID] = true
	}

	if _, err := c.kpClient.AddItem(context.Background(), &kpb.AddItemPrivateRequest{
		Features: features,
		Uid:      uid,
	}); err != nil {
		return err
	}
	return nil
}

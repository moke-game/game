package analytics

import (
	"context"

	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/pkg/global"

	"github.com/moke-game/game/services/common"

	pb "github.com/moke-game/platform/api/gen/analytics"

	"github.com/moke-game/game/services/analytics/aiface"
)

func Analytics(events []aiface.IEvent, logger *zap.Logger) {
	aEvents := &pb.AnalyticsEvents{}
	for _, v := range events {
		v.GetEventName() //事件名称赋值  不能省略
		if data, err := v.ToJson(); err == nil {
			aEvents.Events = append(aEvents.Events, &pb.Event{
				DeliverTo:  v.Platform(),
				Event:      v.GetEventName().String(),
				Properties: data,
			})
			logger.Info("---------Event:"+v.GetEventName().String(), zap.String("content", string(data)))
		} else {
			logger.Error("analytics event json fail", zap.Error(err), zap.Any("name", v.GetEventName()), zap.Any("event", v))
		}
	}
	go func() {
		defer func() {
			common.Cover(logger)
		}()
		if _, err := global.GetAnalyticsSender().Analytics(context.Background(), aEvents); err != nil {
			logger.Error("analytics event send fail", zap.Error(err), zap.Any("events", aEvents))
		}
	}()
}

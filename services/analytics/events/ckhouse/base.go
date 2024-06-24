package ckhouse

import (
	pb "github.com/moke-game/platform/api/gen/analytics"
)

type Base struct {
}

func (m *Base) Platform() pb.DeliveryType {
	return pb.DeliveryType_ClickHouse
}

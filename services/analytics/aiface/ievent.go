package aiface

import (
	pb "github.com/moke-game/platform/api/gen/analytics"

	"github.com/moke-game/game/services/analytics/names"
)

type IEvent interface {
	// GetEventName returns the event name of the event.
	GetEventName() names.EventName
	// ToJson returns the json data of the event.
	ToJson() (data []byte, err error)
	// Platform
	Platform() pb.DeliveryType
}

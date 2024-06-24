package riface

import (
	"github.com/gstones/zinx/ziface"

	"github.com/moke-game/platform/services/matchmaking/pkg/module/data"
)

type RoomCreator func() (IRoom, error)

type IMessage interface {
	GetMsgId() uint32 // Gets the ID of the message(获取消息ID)
	GetMsgData() []byte
}

type IRoom interface {
	Init(result *data.MatchResult) error
	Run() error
	RoomId() string
	Receive(uid string, message ziface.IRequest)
}

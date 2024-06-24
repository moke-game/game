package riface

import (
	"github.com/gstones/zinx/ziface"
	"google.golang.org/protobuf/proto"

	cpb "github.com/moke-game/game/api/gen/common"
)

type IHandler func(uid string, request ziface.IRequest) (proto.Message, cpb.ERRORCODE)

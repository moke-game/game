package biface

import (
	"context"

	"github.com/gstones/zinx/ziface"
	"google.golang.org/protobuf/proto"

	cpb "github.com/moke-game/game/api/gen/common"
)

type IHandler func(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE)

type IRegister interface {
	RegisterHandler(event cpb.C2S_EVENT, handler IHandler)
	RegisterHandlers(start cpb.C2S_EVENT, end cpb.C2S_EVENT, handler IHandler)
}

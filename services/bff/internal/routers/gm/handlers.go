package gm

import (
	"context"
	"fmt"
	"strings"

	"github.com/gstones/zinx/ziface"
	"google.golang.org/protobuf/proto"

	bffpb "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/configs/pkg/module"
)

func (c *Router) handleGM(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	if module.DeploymentGlobal.IsProd() {
		return nil, cpb.ERRORCODE_FEATURE_NOT_OPEN
	}
	req := &bffpb.C2SBffGMCommand{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_PROTO_UNMARSHAL_ERROR
	}
	commands := strings.Split(req.Command, " ")
	if len(commands) < 1 {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	var err error
	if h, ok := c.commands[commands[0]]; ok {
		err = h(request, commands[1:]...)
	} else {
		err = fmt.Errorf("gm command:%s not found", commands[0])
	}
	resp := &bffpb.S2CBffGMCommand{}
	if err != nil {
		resp.Error = err.Error()
	}
	return resp, cpb.ERRORCODE_SUCCESS
}

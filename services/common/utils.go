package common

import (
	"context"
	"fmt"
	"runtime"

	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"github.com/moke-game/game/services/common/constants"

	cpb "github.com/moke-game/game/api/gen/common"
)

const (
	GROUP_SIZE             = 3
	RENAME_CONFIG_NICK_MAX = 116 //修改昵称最大长度
	GROUP_REFUSE_TIME      = 300 //队伍拒绝再次申请秒数
)

func MakeAuthCtxOut(ctx context.Context, connection ziface.IConnection) context.Context {
	if params := MakeAuthParams(connection); params != nil {
		return MakeCtxOutWithParams(ctx, params)
	} else {
		return ctx
	}
}

func MakeAuthParams(connection ziface.IConnection) map[string]string {
	if token, err := connection.GetProperty(constants.ConnToken); err != nil {
		return nil
	} else {
		return map[string]string{
			"authorization": fmt.Sprintf("%s %v", "bearer", token),
		}
	}
}

func MakeCtxOutWithParams(ctx context.Context, params map[string]string) context.Context {
	md := metadata.Pairs()
	for k, v := range params {
		md.Append(k, v)
	}
	return mm.MD(md).ToOutgoing(ctx)
}

func SendResponse(connect ziface.IConnection, msgId uint32, code cpb.ERRORCODE, msg proto.Message) error {
	if connect == nil {
		return nil
	} else if data, err := proto.Marshal(msg); err != nil {
		return err
	} else if data, err = proto.Marshal(&cpb.Response{
		Code: code,
		Data: data,
	}); err != nil {
		return err
	} else if err := connect.SendBuffMsg(msgId, data); err != nil {
		return err
	}
	return nil
}

func SendNotify(connect ziface.IConnection, msgId uint32, msg proto.Message) error {
	if connect == nil {
		return fmt.Errorf("connection is nil")
	} else if data, err := proto.Marshal(msg); err != nil {
		return err
	} else if err := connect.SendBuffMsg(msgId, data); err != nil {
		return err
	}
	return nil
}

func Cover(logger *zap.Logger) {
	if r := recover(); r != nil {
		buf := make([]byte, 2048)
		l := runtime.Stack(buf, false)
		if logger != nil {
			logger.Error("Panic Cover", zap.Any("error", r), zap.Any("stack", string(buf[:l])))
		} else {
			fmt.Printf("Panic Cover - error:%v stack:%v", r, string(buf[:l]))
		}
	}
}

package chat

import (
	"context"
	"strconv"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/gstones/zinx/ziface"

	"github.com/moke-game/game/services/common/constants"

	"github.com/moke-game/platform/api/gen/chat"
	ppb "github.com/moke-game/platform/api/gen/profile"

	"github.com/moke-game/game/services/common"

	bffpb "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	configs "github.com/moke-game/game/configs/pkg/module"
)

func (r *Router) chatMessage(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SCHATMessage{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	uid, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		r.logger.Error("get uid from context err")
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	pro, err := r.pClient.GetProfile(ctx, &ppb.GetProfileRequest{Uid: uid.(string)})
	if err != nil {
		r.logger.Error("get profile msg error", zap.Error(err))
		return nil, cpb.ERRORCODE_PLAYER_NOT_EXIST
	}
	dest := &chat.Destination{
		Channel: int32(req.ChatType),
		Id:      strconv.FormatInt(req.ChatInfo.ReceiveUid, 10),
	}
	msg := &chat.ChatMessage_Message{
		ProfileId: uid.(string),
		Nickname:  pro.Profile.Nickname,
		Avatar:    pro.Profile.Avatar,
		Content:   req.ChatInfo.Content,
		Emoji:     req.ChatInfo.EmojiId,
		Timestamp: time.Now().Unix(),
	}

	var cStream chat.ChatService_ChatClient
	if out, err := request.GetConnection().GetProperty("chat"); err != nil {
		r.logger.Error("get chat stream err", zap.Error(err))
		return nil, cpb.ERRORCODE_COMMON_ERROR
	} else {
		cStream = out.(chat.ChatService_ChatClient)
	}
	if err := cStream.Send(&chat.ChatRequest{
		Kind: &chat.ChatRequest_Message{
			Message: &chat.ChatMessage{
				Destination: dest,
				Message:     msg,
			},
		},
	}); err != nil {
		r.logger.Error("send chat message err", zap.Error(err))
	}
	return nil, cpb.ERRORCODE_SUCCESS
}

func (r *Router) chatGetEmoji(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	cfg := configs.ConfigsGlobal.TblChatEmoji.GetDataMap()
	resp := &bffpb.S2CChatGetEmoji{}
	for i := range cfg {
		resp.EmojiId = append(resp.EmojiId, i)
	}
	return resp, cpb.ERRORCODE_SUCCESS
}

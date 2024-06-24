package mail

import (
	"context"
	"strconv"

	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	bffpb "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/msg_transfer"
	pb2 "github.com/moke-game/platform/api/gen/knapsack"
	pb "github.com/moke-game/platform/api/gen/mail"
	pb3 "github.com/moke-game/platform/api/gen/profile"
)

func (r *Router) watchMail(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if err := r.watchMailResponse(request, ctx); err != nil {
		r.logger.Error("watch mail response err", zap.Error(err))
		return &bffpb.S2CWatchMail{}, cpb.ERRORCODE_WATCH_ERROR
	}
	return nil, cpb.ERRORCODE_SUCCESS
}

func (r *Router) readMail(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SReadMail{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}

	updates := make(map[int64]pb.MailStatus)
	if len(req.GetIds()) <= 0 {
		updates[0] = pb.MailStatus_READ
	} else {
		for _, v := range req.GetIds() {
			updates[v] = pb.MailStatus_READ
		}
	}

	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if _, err := r.mClient.UpdateMail(ctx, &pb.UpdateMailRequest{
		Updates: updates,
	}); err != nil {
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		return &bffpb.S2CReadMail{}, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) deleteReadMail(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SDeleteReadMail{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}

	updates := make(map[int64]pb.MailStatus)
	if len(req.GetIds()) <= 0 {
		updates[0] = pb.MailStatus_DELETED
	} else {
		for _, v := range req.GetIds() {
			updates[v] = pb.MailStatus_DELETED
		}
	}

	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if _, err := r.mClient.UpdateMail(ctx, &pb.UpdateMailRequest{
		Updates: updates,
	}); err != nil {
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		return &bffpb.S2CDeleteReadMail{}, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) getMailRewards(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SGetMailRewards{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}

	updates := make(map[int64]pb.MailStatus)
	if len(req.GetIds()) <= 0 {
		updates[0] = pb.MailStatus_REWARDED
	} else {
		for _, v := range req.GetIds() {
			updates[v] = pb.MailStatus_REWARDED
		}
	}

	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if resp, err := r.mClient.UpdateMail(ctx, &pb.UpdateMailRequest{
		Updates: updates,
	}); err != nil {
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		rewards := &bffpb.Items{
			Items: make(map[int64]*bffpb.Item),
		}
		bffRewards := &bffpb.Items{
			Items: make(map[int64]*bffpb.Item),
		}
		for _, v := range resp.Rewards {
			item := &bffpb.Item{
				Id:       v.Id,
				Num:      v.Num,
				Expired:  v.Expire,
				ConfigId: v.Type,
			}
			bffRewards.Items[v.Id] = item
			rewards.Items[v.Id] = item
		}
		if len(rewards.Items) > 0 {
			add := msg_transfer.Items2Knapsack(rewards)
			if _, err := r.kClient.AddItem(ctx, &pb2.AddItemRequest{Items: add}); err != nil {
				return nil, cpb.ERRORCODE_RPC_ERROR
			}
		}
		return &bffpb.S2CGetMailRewards{
			Rewards: bffRewards,
		}, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) watchMailResponse(request ziface.IRequest, ctx context.Context) error {
	req := &bffpb.C2SWatchMail{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return err
	}

	profile, err := r.pClient.GetProfile(ctx, &pb3.GetProfileRequest{})
	if err != nil {
		return err
	}
	if cli, err := r.mClient.Watch(ctx, &pb.WatchMailRequest{
		Language: req.Language,
		//Channel:      strconv.Itoa(int(profile.GetProfile().Channel)),
		RegisterTime: profile.GetProfile().RegisterTime,
	}); err != nil {
		return err
	} else {
		go func() {
			for {
				if resp, err := cli.Recv(); err != nil {
					if status.Code(err) == codes.Canceled {
						r.logger.Info("chat connection closed")
						return
					}
					r.logger.Error("watch mail err", zap.Error(err))
					return
				} else {
					changes := r.makeMailChanges(resp.Mails, req.Language)
					r.logger.Info("watched mails changed", zap.Any("changes", changes))
					if err := common.SendNotify(request.GetConnection(), uint32(cpb.S2C_EVENT_NTF_MailChange), changes); err != nil {
						r.logger.Error("send mail change err", zap.Error(err))
					}
				}
			}
		}()
	}
	return nil
}

func (r *Router) makeMailChanges(mails map[int64]*pb.Mail, language string) *bffpb.NtfMailChange {
	ntf := &bffpb.NtfMailChange{
		Mails: make(map[int64]*bffpb.Mail),
	}
	for k, m := range mails {
		sta := bffpb.MailStatus(m.Status)
		rewards := &bffpb.Items{
			Items: make(map[int64]*bffpb.Item),
		}
		for _, v1 := range m.GetRewards() {
			rewards.Items[v1.GetId()] = &bffpb.Item{
				Id:       v1.Id,
				ConfigId: v1.Type,
				Num:      v1.Num,
				Expired:  v1.Expire,
			}
		}
		body, ok := m.Body[language]
		if !ok {
			body = m.Body["en"]
		}
		title, ok := m.Title[language]
		if !ok {
			title = m.Title["en"]
		}

		ntf.Mails[k] = &bffpb.Mail{
			Id:       m.Id,
			From:     m.From,
			Body:     body,
			Date:     m.Date,
			Expire:   m.ExpireAt,
			Status:   sta,
			Rewards:  rewards,
			Title:    title,
			Template: strconv.Itoa(int(m.TemplateId)),
			Params:   m.TemplateArgs,
		}
	}
	r.logger.Info("watched mails changed", zap.Any("changes", ntf))
	return ntf
}

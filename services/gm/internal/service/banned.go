package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/moke-game/game/services/common/notification"
	"github.com/moke-game/game/services/gm/errors"
	"github.com/moke-game/game/services/gm/internal/common"

	pb2 "github.com/moke-game/platform/api/gen/auth"
	"github.com/moke-game/platform/api/gen/chat"
	"github.com/moke-game/platform/api/gen/mail"
	profile "github.com/moke-game/platform/api/gen/profile"

	pb "github.com/moke-game/game/api/gen/gm"
)

func (s *Service) KickOffline(_ context.Context, request *pb.KickOfflineRequest) (*pb.KickOfflineResponse, error) {
	if request.Type == 2 && request.Val != "" {
		topic := notification.MakeGamePrivateNotifyTopic(request.Val)
		if err := s.mq.Publish(topic, miface.WithJSON(notification.NotifyEventKickOffline)); err != nil {
			s.logger.Error("publish kick offline failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
	} else if request.Val == "" {
		topic := notification.MakeGamePublicNotifyTopic()
		if err := s.mq.Publish(topic, miface.WithJSON(notification.NotifyEventKickOffline)); err != nil {
			s.logger.Error("publish kick offline failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
	} else {
		s.logger.Error("kick offline failed", zap.Int32("type", request.Type), zap.String("val", request.Val))
		return nil, fmt.Errorf("invalid request")
	}
	return &pb.KickOfflineResponse{}, nil
}

func (s *Service) PlayerBannedList(
	ctx context.Context,
	request *pb.PlayerBannedListRequest,
) (*pb.PlayerBannedListResponse, error) {
	platform := ""
	if request.GetPlatformId() != "0" {
		platform = request.GetPlatformId()
	}
	profileLst := make(map[string]*profile.Profile)
	req := makeGetProfileReq(request.Type, request.Val, platform, request.Page, request.PageSize)
	if resp, err := s.profileCLi.GetProfilePrivate(ctx, req); err != nil {
		return nil, err
	} else {
		for _, v := range resp.Profiles {
			profileLst[v.GetUid()] = v
		}
	}
	uids := make([]string, 0)
	for k := range profileLst {
		uids = append(uids, k)
	}
	bl, err := s.db.GetBlockedList(uids...)
	if err != nil {
		s.logger.Error("get blockList failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	res := make([]*pb.PlayerBannedInfo, 0)
	for k, v := range profileLst {
		banInfo := common.TransProfile2PlayerBanInfo(v)
		if info, ok := bl[k]; ok {
			banInfo.BannedType = info.BannedType
			banInfo.LockTime = info.BannedTime
			banInfo.UnlockTime = info.UnsealTime
		}
		res = append(res, banInfo)
	}

	return &pb.PlayerBannedListResponse{
		TotalCount:     int32(len(res)),
		BannedInfoList: res,
	}, nil

}

func (s *Service) PlayerBanned(
	ctx context.Context,
	request *pb.PlayerBannedRequest,
) (*pb.PlayerBannedResponse, error) {
	reqMsg := &pb.BannedReqMsg{}
	if data, e := common.CBCDecrypt([]byte(s.aesKey), request.Data); e != nil {
		s.logger.Error("CBCDecrypt failed", zap.Error(e), zap.String("data", request.Data))
		return nil, errors.ErrGeneralFailure
	} else if e := protojson.Unmarshal([]byte(data), reqMsg); e != nil {
		s.logger.Error("unmarshal data failed", zap.Error(e), zap.String("data", data))
		return nil, errors.ErrGeneralFailure
	}
	now := time.Now()
	var duration time.Duration
	end := int64(0)
	if reqMsg.BannedDuration > 0 {
		duration = time.Duration(reqMsg.BannedDuration) * time.Hour
		end = now.Add(duration).Unix()
	}
	banInfo := &pb.BannedInfo{
		BannedType:   reqMsg.BannedType,
		BannedTime:   now.Unix(),
		UnsealTime:   end,
		BannedReason: reqMsg.BannedReason,
	}
	switch reqMsg.BannedType {
	case -1:
		if info, err := s.db.GetBlockedById(reqMsg.GetRoleId()); err != nil {
			s.logger.Error("get blockList failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		} else if info == nil {
			s.logger.Error("blockList not found", zap.String("uid", reqMsg.GetRoleId()))
			return nil, errors.ErrGeneralFailure
		} else {
			if info.BannedType == 1 {
				if _, err := s.chatCli.AddBlocked(ctx, &chat.AddBlockedRequest{
					ProfileId: reqMsg.GetRoleId(),
				}); err != nil {
					s.logger.Error("chat add blocked failed", zap.Error(err))
					return nil, errors.ErrGeneralFailure
				} else if err := s.db.RemoveBlockedList(reqMsg.GetRoleId()); err != nil {
					s.logger.Error("add blockList failed", zap.Error(err))
					return nil, errors.ErrGeneralFailure
				}
			} else if info.BannedType == 0 {
				if _, err := s.authCli.AddBlocked(ctx, &pb2.BlockListRequest{
					Uid: reqMsg.GetRoleId(),
				}); err != nil {
					s.logger.Error("auth add blockList failed", zap.Error(err))
					return nil, errors.ErrGeneralFailure
				} else if err := s.db.RemoveBlockedList(reqMsg.GetRoleId()); err != nil {
					s.logger.Error("add blockList failed", zap.Error(err))
					return nil, errors.ErrGeneralFailure
				}
			}
		}
	case 1:
		if _, err := s.chatCli.AddBlocked(ctx, &chat.AddBlockedRequest{
			ProfileId: reqMsg.GetRoleId(),
			Duration:  int64(duration.Seconds()),
			IsBlocked: true,
		}); err != nil {
			s.logger.Error("chat add blocked failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		} else if err := s.db.AddBlockedList(reqMsg.RoleId, banInfo); err != nil {
			s.logger.Error("add blockList failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}

	case 0:
		if _, err := s.authCli.AddBlocked(ctx, &pb2.BlockListRequest{
			Uid:      reqMsg.GetRoleId(),
			IsBlock:  true,
			Duration: int64(duration.Seconds()),
		}); err != nil {
			s.logger.Error("auth add blockList failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		} else if err := s.db.AddBlockedList(reqMsg.RoleId, banInfo); err != nil {
			s.logger.Error("add blockList failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
		topic := makeBlockedListTopic(reqMsg.GetRoleId())
		if err := s.mq.Publish(topic, miface.WithJSON(banInfo)); err != nil {
			s.logger.Error("publish block list failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
	case 2:
		if _, err := s.mailCli.SendMail(ctx, &mail.SendMailRequest{
			SendType: mail.SendMailRequest_ROLE,
			RoleIds:  []string{reqMsg.GetRoleId()},
			Mail: &mail.Mail{
				Date:     now.Unix(),
				ExpireAt: end,
				Title:    map[string]string{"en": "Mail warning", "zh": "警告"},
				Body:     map[string]string{"en": reqMsg.GetRemark()},
				From:     "System",
			},
		}); err != nil {
			s.logger.Error("send mail failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
	}
	return &pb.PlayerBannedResponse{
		Status: "ok",
		Code:   "1",
		Info:   "banned success",
		BannedInfo: &pb.BannedInfo{
			BannedTime: now.Unix(),
			UnsealTime: end,
		},
	}, nil
}

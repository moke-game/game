package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/moke-game/game/api/gen/gm"
	"github.com/moke-game/game/services/gm/errors"
	"github.com/moke-game/game/services/gm/internal/common"
	pb3 "github.com/moke-game/platform/api/gen/knapsack"
	pb2 "github.com/moke-game/platform/api/gen/profile"
)

// GetPlayerInfoList 1-角色名(模糊查询)； 2-角色ID；3-账号ID；4-渠道用户ID；5-聚合订单号；6-角色名(精确查询)。
// 注意：当type=1或2时，如果平台ID和区服ID传的都是0，则代表查询游戏内符合当前条件的所有角色；当type=3或4或5时，
// 平台ID传空字符串和，区服ID传空数组；当type=6时，平台ID和区服ID必有值且不为0
func (s *Service) GetPlayerInfoList(
	ctx context.Context,
	request *pb.GetPlayerInfoListRequest,
) (*pb.GetPlayerInfoListResponse, error) {
	platform := ""
	if request.GetPlatformId() != "0" {
		platform = request.GetPlatformId()
	}
	req := makeGetProfileReq(request.Type, request.Val, platform, request.Page, request.PageSize)
	if resp, err := s.profileCLi.GetProfilePrivate(ctx, req); err != nil {
		return nil, err
	} else {
		infoLst := make([]*pb.PlayerInfo, 0)
		for _, v := range resp.Profiles {
			if info := common.TransProfile2PlayerInfo(v); info != nil {
				infoLst = append(infoLst, info)
			}
		}
		return &pb.GetPlayerInfoListResponse{PlayerInfoList: infoLst}, nil
	}
}

func (s *Service) GetPlayerInfo(ctx context.Context, request *pb.GetPlayerInfoRequest) (*pb.GetPlayerInfoResponse, error) {
	platform := ""
	if request.GetPlatformId() != "0" {
		platform = request.GetPlatformId()
	}

	req := &pb2.GetProfilePrivateRequest{
		PlatformId: platform,
		Kind: &pb2.GetProfilePrivateRequest_Uids_{
			Uids: &pb2.GetProfilePrivateRequest_Uids{
				Uid: []string{request.RoleId},
			},
		},
	}

	if resp, err := s.profileCLi.GetProfilePrivate(ctx, req); err != nil {
		return nil, errors.ErrProfileNotFound
	} else if len(resp.Profiles) <= 0 {
		return nil, errors.ErrProfileNotFound
	} else if knapsack, err := s.knapsack.GetKnapsack(ctx, &pb3.GetKnapsackRequest{
		Uid: request.RoleId,
	}); err != nil {
		s.logger.Error("get knapsack failed", zap.Error(err))
		return nil, errors.ErrKnapsackNotFound
	} else {
		info := common.TransProfile2ProfileDetail(resp.Profiles[0], knapsack.GetKnapsack())
		info.GoodsList = common.TransKnapsack2Goods(knapsack.GetKnapsack())
		return info, nil
	}
}

func (s *Service) QueryUserList(ctx context.Context, request *pb.QueryUserListRequest) (*pb.QueryUserListResponse, error) {
	reqMsg := &pb.QueryUserMsg{}
	if data, e := common.CBCDecrypt([]byte(s.aesKey), request.Data); e != nil {
		s.logger.Error("CBCDecrypt failed", zap.Error(e), zap.String("data", request.Data))
		return nil, e
	} else if e := protojson.Unmarshal([]byte(data), reqMsg); e != nil {
		s.logger.Error("unmarshal data failed", zap.Error(e), zap.String("data", data))
		return nil, e
	}
	platform := ""
	if reqMsg.GetPlatformId() != "0" {
		platform = reqMsg.GetPlatformId()
	}

	req := &pb2.GetProfilePrivateRequest{
		PlatformId: platform,
		ChannelId:  reqMsg.ChannelId,
	}
	if reqMsg.UserId != "" {
		req.Kind = &pb2.GetProfilePrivateRequest_Uids_{
			Uids: &pb2.GetProfilePrivateRequest_Uids{
				Uid: []string{reqMsg.UserId},
			},
		}
	} else if reqMsg.RoleName != "" {
		req.Kind = &pb2.GetProfilePrivateRequest_Name_{
			Name: &pb2.GetProfilePrivateRequest_Name{
				Name:     reqMsg.RoleName,
				IsRegexp: true,
				Page:     reqMsg.Page,
				PageSize: reqMsg.PageSize,
			},
		}
	}

	if resp, err := s.profileCLi.GetProfilePrivate(ctx, req); err != nil {
		return nil, err
	} else {
		infoLst := make([]*pb.PlayerInfo, 0)
		for _, v := range resp.Profiles {
			if info := common.TransProfile2PlayerInfo(v); info != nil {
				infoLst = append(infoLst, info)
			}
		}
		return &pb.QueryUserListResponse{UserList: infoLst}, nil
	}

}

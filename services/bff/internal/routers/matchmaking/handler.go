package matchmaking

import (
	"context"

	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/constants"
	"github.com/moke-game/platform/api/gen/matchmaking"
	pb "github.com/moke-game/platform/api/gen/party"
)

const (
	MatchReady  = 1
	MatchMaking = 2
)

func (r *Router) handleMatchSingleStart(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	return nil, cpb.ERRORCODE_COMMON_ERROR
}

func (r *Router) handleMatchCancel(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bff.C2SMatchingCancel{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	iUid, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	uid := iUid.(string)
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	//组队匹配中也通过此消息发送取消匹配的操作，只有队长有权限取消匹配，取消匹配后除队长外，其他队友状态仍处于准备状态
	if _, err = r.mClient.MatchCancel(ctx, &matchmaking.MatchCancelRequest{
		ProfileId: uid,
	}); err != nil {
		r.logger.Error("MatchCancel err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	partyReq := &pb.GetPartyRequest{}
	partyReq.Kind = &pb.GetPartyRequest_Uid{
		Uid: uid,
	}
	if party, err := r.partyClient.GetParty(ctx, partyReq); err == nil {
		if party.Party != nil && len(party.Party.Members) > 0 {
			for mid := range party.Party.Members {
				if mid == uid {
					continue
				}
			}
		}
	}

	return &bff.S2CMatchingCancel{}, cpb.ERRORCODE_SUCCESS

}

func (r *Router) handleMatchStatus(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	iUid, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil || iUid == nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	uid := iUid.(string)
	resp := &bff.S2CMatchingStatus{}
	statusResp, err := r.mClient.MatchStatus(ctx, &matchmaking.MatchStatusRequest{ProfileId: uid})
	if err != nil {
		r.logger.Error("MatchStatus err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	if statusResp != nil && statusResp.MatchTime > 0 {
		resp.Status = MatchMaking
		return resp, cpb.ERRORCODE_SUCCESS
	}

	ptReq := &pb.GetPartyRequest{}
	ptReq.Kind = &pb.GetPartyRequest_Uid{
		Uid: uid,
	}
	if partyResp, err := r.partyClient.GetParty(ctx, ptReq); err == nil {
		for _, member := range partyResp.Party.Members {
			if member.Uid != uid {
				continue
			}
			if member.Status == int32(bff.PartyMemberStatus_STATUS_READY) {
				resp.Status = MatchReady
				return resp, cpb.ERRORCODE_SUCCESS
			}
			break
		}
	}
	return resp, cpb.ERRORCODE_SUCCESS
}

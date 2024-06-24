package party

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/gstones/zinx/ziface"

	"github.com/moke-game/game/services/common/constants"

	"github.com/moke-game/platform/api/gen/buddy"
	ptpb "github.com/moke-game/platform/api/gen/party"
	ppb "github.com/moke-game/platform/api/gen/profile"

	bffpb "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/msg_transfer"
)

func (r *Router) getPartyInfo(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	uid, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	reqMsg := &bffpb.C2SGetPartyInfo{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	ptReq := &ptpb.GetPartyRequest{}
	if reqMsg.GetPartyId() != "" {
		ptReq.Kind = &ptpb.GetPartyRequest_Pid{
			Pid: reqMsg.GetPartyId(),
		}
	} else if reqMsg.GetMemberId() != 0 {
		ptReq.Kind = &ptpb.GetPartyRequest_Uid{
			Uid: strconv.FormatInt(reqMsg.GetMemberId(), 10),
		}
	} else {
		ptReq.Kind = &ptpb.GetPartyRequest_Uid{
			Uid: uid.(string),
		}
	}
	if resp, err := r.ptClient.GetParty(ctx, ptReq); err != nil {
		r.logger.Error("get party error", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else if resp == nil || resp.Party == nil {
		return nil, cpb.ERRORCODE_SUCCESS
	} else if msg, err := msg_transfer.Party(resp.Party); err != nil {
		r.logger.Error("party changes transfer msg error", zap.Error(err))
		return nil, cpb.ERRORCODE_COMMON_ERROR
	} else {
		if proResp, err := r.pClient.GetProfileBasics(ctx, &ppb.GetProfileBasicsRequest{}); err != nil {
			r.logger.Error("get profile error", zap.Error(err))
			return nil, cpb.ERRORCODE_RPC_ERROR
		} else if basic, ok := proResp.Basics[uid.(string)]; !ok || basic == nil {
			r.logger.Error("get profile error", zap.Error(err))
			return nil, cpb.ERRORCODE_PLAYER_NOT_EXIST
		} else if err := r.watchParty(
			request.GetConnection(),
			basic,
			msg.PartyId,
			false,
			msg.PlayId,
			false,
			true,
		); err != nil {
			r.logger.Error("watch party error", zap.Error(err))
			return nil, cpb.ERRORCODE_RPC_ERROR
		}
		if property, err := request.GetConnection().GetProperty(IdPropertyKey); err != nil || property == nil {
			request.GetConnection().SetProperty(IdPropertyKey, msg.PartyId)
		}
		resp := &bffpb.S2CGetPartyInfo{
			PartyInfo: msg,
		}
		return resp, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) friendPartyInfo(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	buddyResp, err := r.bClient.GetBuddies(ctx, &buddy.GetBuddyRequest{})
	if err != nil {
		r.logger.Error("buddies client getBuddies err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	party := make([]*bffpb.PartyInfo, 0)
	for _, b := range buddyResp.Buddies.Buddies {
		ptReq := &ptpb.GetPartyRequest{
			Kind: &ptpb.GetPartyRequest_Uid{
				Uid: b.Uid,
			},
		}
		if resp, err := r.ptClient.GetParty(ctx, ptReq); err != nil {
			//r.logger.Error("get party error", zap.Error(err))
			continue
		} else if msg, err := msg_transfer.Party(resp.Party); err != nil {
			//r.logger.Error("party changes transfer msg error", zap.Error(err))
			continue
		} else {
			if len(msg.PartyMembers) < common.GROUP_SIZE {
				party = append(party, msg)
			}
		}
	}
	return &bffpb.S2CFriendPartyInfo{Info: party}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) createParty(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	uidProp, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	uidStr := uidProp.(string)
	reqMsg := &bffpb.C2SCreateParty{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	if resp, err := r.pClient.GetProfileBasics(ctx, &ppb.GetProfileBasicsRequest{}); err != nil {
		r.logger.Error("get profile error", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else if basic, ok := resp.Basics[uidStr]; !ok || basic == nil {
		r.logger.Error("get profile error", zap.Error(err))
		return nil, cpb.ERRORCODE_PLAYER_NOT_EXIST
	} else if err := r.watchParty(
		request.GetConnection(),
		basic,
		"",
		true,
		reqMsg.PlayId,
		false,
		false,
	); err != nil {
		r.logger.Error("watch party error", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	return nil, cpb.ERRORCODE_SUCCESS
}

func (r *Router) joinParty(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	uid, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	reqMsg := &bffpb.C2SJoinParty{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	} else if reqMsg.GetPartyId() == "" {
		r.logger.Error("party id is empty")
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	ptReq := &ptpb.GetPartyRequest{}
	ptReq.Kind = &ptpb.GetPartyRequest_Pid{
		Pid: strings.ToLower(reqMsg.PartyId),
	}
	if resp, err := r.ptClient.GetParty(ctx, ptReq); err != nil {
		r.logger.Error("get party error", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else if msg, err := msg_transfer.Party(resp.Party); err != nil {
		r.logger.Error("party changes transfer msg error", zap.Error(err))
		return nil, cpb.ERRORCODE_COMMON_ERROR
	} else if len(msg.PartyMembers) >= common.GROUP_SIZE {
		return &bffpb.S2CJoinParty{
			RetCode: int32(cpb.ERRORCODE_GROUP_MAX),
			PartyId: reqMsg.PartyId,
		}, cpb.ERRORCODE_SUCCESS
	} else {
		//检测是否在队伍黑名单中
		nowTime := time.Now().UTC().Unix()
		if tim, ok := resp.Party.Party.Refuse[uid.(string)]; ok {
			if nowTime-tim <= common.GROUP_REFUSE_TIME {
				return &bffpb.S2CJoinParty{
					RetCode: int32(cpb.ERRORCODE_GROUP_REFUSE_JOIN),
					PartyId: reqMsg.PartyId,
				}, cpb.ERRORCODE_SUCCESS
			}
		}
		//检测是否在双方黑名单中
		verifyResp, err := r.bClient.VerifyBlocked(ctx, &buddy.VerifyBlockedRequest{UidSelf: uid.(string), UidOther: resp.Party.Party.Owner})
		if err != nil {
			r.logger.Error("VerifyBlocked error", zap.Error(err))
		}
		if verifyResp.IsBlocked {
			return &bffpb.S2CJoinParty{
				RetCode: int32(cpb.ERRORCODE_GROUP_REFUSE_JOIN),
				PartyId: reqMsg.PartyId,
			}, cpb.ERRORCODE_SUCCESS
		}
	}
	return &bffpb.S2CJoinParty{
		RetCode: int32(cpb.ERRORCODE_SUCCESS),
		PartyId: reqMsg.PartyId,
	}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) leaveParty(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if res, err := r.ptClient.LeaveParty(ctx, &ptpb.LeavePartyRequest{}); err != nil {
		r.logger.Error("leave party error", zap.Error(err), zap.Any("resp", res))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else if err := destroyParty(request.GetConnection()); err != nil {
		r.logger.Error("destroy party error", zap.Error(err))
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	return &bffpb.S2CLeaveParty{}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) kickParty(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	reqMsg := &bffpb.C2SKickParty{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	partyId := ""
	if idProp, err := request.GetConnection().GetProperty(IdPropertyKey); err != nil {
		r.logger.Error("get id error", zap.Error(err))
		return nil, cpb.ERRORCODE_COMMON_ERROR
	} else if id, ok := idProp.(string); !ok {
		r.logger.Error("id is not string", zap.Any("id", idProp))
	} else {
		partyId = id
	}
	uidStr := strconv.FormatInt(reqMsg.GetMemberId(), 10)
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if resp, err := r.ptClient.KickOut(ctx, &ptpb.KickOutRequest{
		PartyId: partyId,
		Uid:     uidStr,
	}); err != nil {
		r.logger.Error("kick out error", zap.Error(err), zap.Any("resp", resp))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	return &bffpb.S2CKickParty{}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) readyParty(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	if err := r.updatePartyMemberStatus(ctx, request, bffpb.PartyMemberStatus_STATUS_READY); err != nil {
		r.logger.Error("update party member status error", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		partyId := ""
		if idProp, err := request.GetConnection().GetProperty(IdPropertyKey); err != nil {
			r.logger.Error("get IdPropertyKey error", zap.Error(err))
		} else if id, ok := idProp.(string); !ok {
			r.logger.Error("id is not string", zap.Any("id", idProp))
		} else {
			partyId = id
		}
		ptReq := &ptpb.GetPartyRequest{}
		ptReq.Kind = &ptpb.GetPartyRequest_Pid{
			Pid: partyId,
		}

	}
	return &bffpb.S2CReadyParty{}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) cancelReadyParty(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	if err := r.updatePartyMemberStatus(ctx, request, bffpb.PartyMemberStatus_STATUS_NOT_READY); err != nil {
		r.logger.Error("update party member status error", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	return &bffpb.S2CCancelReadyParty{}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) updatePartyMemberStatus(
	ctx context.Context,
	request ziface.IRequest,
	status bffpb.PartyMemberStatus,
) error {
	partyId := ""
	if idProp, err := request.GetConnection().GetProperty(IdPropertyKey); err != nil {
		return err
	} else if id, ok := idProp.(string); !ok {
		return fmt.Errorf("id is not string, id: %v", idProp)
	} else {
		partyId = id
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if _, err := r.ptClient.UpdateMember(ctx, &ptpb.UpdateMemberRequest{
		PartyId: partyId,
		Member: &ptpb.Member{
			Status: int32(status),
		},
	}); err != nil {
		return err
	}
	return nil
}

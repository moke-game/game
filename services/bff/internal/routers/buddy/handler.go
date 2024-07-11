package buddy

import (
	"context"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/gstones/zinx/ziface"

	"github.com/moke-game/game/services/common/constants"

	pb "github.com/moke-game/platform/api/gen/buddy"
	pb4 "github.com/moke-game/platform/api/gen/party"
	pb2 "github.com/moke-game/platform/api/gen/profile"

	bffpb "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common"
)

func (r *Router) handleFriendGet(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SFriendGet{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	selfUid, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	resp := &bffpb.S2CFriendGet{}
	if buddy, err := r.bClient.GetBuddies(ctx, &pb.GetBuddyRequest{}); err != nil {
		r.logger.Error("handleFriendGet getBuddies err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		ids := make(map[string]int64)
		idArr := make([]string, 0)
		remakes := make(map[string]string)
		teammateMap := make(map[string]*pb2.ProfileBasic)
		switch req.ListType {
		case 1: //好友列表
			for _, bdy := range buddy.Buddies.Buddies {
				ids[bdy.Uid] = bdy.ActTime
				remakes[bdy.Uid] = bdy.Remark
				idArr = append(idArr, bdy.Uid)
			}
		case 2: //申请列表
			for _, inviter := range buddy.Buddies.Inviters {
				ids[inviter.Uid] = inviter.ReqTime
				remakes[inviter.Uid] = ""
				idArr = append(idArr, inviter.Uid)
			}
		case 3: //黑名单列表
			for _, block := range buddy.Buddies.Blocked {
				ids[block.Uid] = block.AddTime
				remakes[block.Uid] = ""
				idArr = append(idArr, block.Uid)
			}
		}
		friends := make([]*bffpb.FriendInfo, 0)
		inviteeStatus, er := r.pClient.GetProfileStatus(ctx, &pb2.GetProfileStatusRequest{Uid: idArr})
		if er != nil {
			r.logger.Error("GetProfileStatus err", zap.Error(er))
		}
		if inviteeStatus == nil || inviteeStatus.Status == nil {
			inviteeStatus = &pb2.GetProfileStatusResponse{Status: make(map[string]int32)}
		}
		for _, uid := range idArr {
			tim, ok := ids[uid]
			if !ok {
				tim = time.Now().UTC().Unix()
			}
			nickName := ""
			if remake, ok := remakes[uid]; ok {
				nickName = remake
			}
			if playerInfo := r.getSimplePlayerInfo(uid, inviteeStatus, ctx); playerInfo != nil {
				friend := &bffpb.FriendInfo{
					Info:    playerInfo,
					Remarks: nickName,
					Tim:     tim,
				}
				friends = append(friends, friend)
			} else {
				if basic, ok := teammateMap[uid]; ok {
					uidInt, err := strconv.ParseInt(basic.Uid, 10, 64)
					if err != nil {
						r.logger.Error("handleFriendGet ParseInt error", zap.Error(err))
						continue
					}
					player := &bffpb.PlayerSimpleInfo{
						Uid:    uidInt,
						Name:   basic.Nickname,
						Head:   basic.Avatar,
						Level:  1,
						HeroId: basic.HeroId,
						Online: false,
					}
					friend := &bffpb.FriendInfo{
						Info:    player,
						Remarks: nickName,
						Tim:     tim,
					}
					friends = append(friends, friend)
				}
			}
		}
		resp.ListType = req.ListType
		resp.Infos = friends
		resp.MyFriendCode = selfUid.(string)
		return resp, cpb.ERRORCODE_SUCCESS
	}

}

func (r *Router) getSimplePlayerInfo(uid string, inviteeStatus *pb2.GetProfileStatusResponse, ctx context.Context) *bffpb.PlayerSimpleInfo {
	id, _ := strconv.ParseInt(uid, 10, 64)
	if prep, err := r.pClient.GetProfile(ctx, &pb2.GetProfileRequest{
		Uid: uid,
	}); err == nil {
		party, _ := r.ptClient.GetParty(ctx, &pb4.GetPartyRequest{Kind: &pb4.GetPartyRequest_Uid{Uid: uid}})
		partyId := "0"
		var state int32 = 0 //0=正常 1=战斗中 2=组队中
		var onlineStatus int32 = 0
		if status, ok := inviteeStatus.Status[uid]; ok {
			onlineStatus = status
		}
		if onlineStatus == int32(pb2.ProfileStatus_BATTLE) {
			state = 1
		}
		if party != nil && party.Party != nil && party.Party.Party != nil {
			partyId = party.Party.Party.Id
			state = 2
		}
		playerInfo := &bffpb.PlayerSimpleInfo{
			Uid:           id,
			Name:          prep.Profile.Nickname,
			Head:          prep.Profile.Avatar,
			Level:         1,
			HeroId:        prep.Profile.HeroId,
			Online:        onlineStatus > 0,
			State:         state,
			PartyId:       partyId,
			LastLoginTime: int32(time.Now().Unix()), //chttodo 最后登录时间暂时使用当前时间
		}
		return playerInfo
	}
	return nil
}

func (r *Router) handleFriendGetPlayerInfoByCode(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	req := &bffpb.C2SFriendGetPlayerInfoByCode{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	_, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	resp := &bffpb.S2CFriendGetPlayerInfoByCode{}
	if prep, err := r.pClient.GetProfile(ctx, &pb2.GetProfileRequest{
		Uid: req.FriendCode,
	}); err == nil {
		id, _ := strconv.ParseInt(prep.Profile.Uid, 10, 64)
		state := int32(0)
		partyId := ""
		partyReq := &pb4.GetPartyRequest{
			Kind: &pb4.GetPartyRequest_Uid{Uid: prep.Profile.Uid},
		}
		if partyRes, err := r.ptClient.GetParty(ctx, partyReq); err == nil {
			if partyRes != nil && partyRes.Party != nil && partyRes.Party.Party != nil {
				state = int32(2)
				partyId = partyRes.Party.Party.Id
			}
		}
		playerInfo := &bffpb.PlayerSimpleInfo{
			Uid:           id,
			Name:          prep.Profile.Nickname,
			Head:          prep.Profile.Avatar,
			Level:         1,
			HeroId:        prep.Profile.HeroId,
			Online:        prep.Profile.OnlineStatus == 1,
			State:         state,
			PartyId:       partyId,
			LastLoginTime: int32(time.Now().Unix()),
		}
		resp.Infos = playerInfo
		return resp, cpb.ERRORCODE_SUCCESS
	}
	return resp, cpb.ERRORCODE_PLAYER_NOT_EXIST
}

func (r *Router) handleFriendAdd(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SFriendAdd{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	_, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	addUidStr := strconv.FormatInt(req.AddUid, 10)
	reqAdd := &pb.AddBuddyRequest{}
	reqAdd.Uid = append(reqAdd.Uid, addUidStr)
	if _, err := r.bClient.AddBuddy(ctx, reqAdd); err != nil {
		r.logger.Error("buddy client addBuddy err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	return &bffpb.S2CFriendAdd{}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) handleFriendAgree(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SFriendAgree{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	_, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	reqReply := &pb.ReplyAddBuddyRequest{}
	uid := strconv.FormatInt(req.Uid, 10)
	reqReply.Uid = append(reqReply.Uid, uid)
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if rep, err := r.bClient.ReplyAddBuddy(ctx, reqReply); err != nil {
		r.logger.Error("buddy client replyAddBuddy err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		if len(rep.Failed) > 0 {
			return nil, cpb.ERRORCODE_COMMON_ERROR
		}
		selfProfile, err := r.pClient.GetProfile(ctx, &pb2.GetProfileRequest{})
		if err != nil {
			return nil, cpb.ERRORCODE_PLAYER_NOT_EXIST
		}
		//通知申请好友的玩家
		selfState := int32(0)
		selfPartyId := ""
		partyReq := &pb4.GetPartyRequest{
			Kind: &pb4.GetPartyRequest_Uid{Uid: selfProfile.Profile.Uid},
		}
		if partyRes, err := r.ptClient.GetParty(ctx, partyReq); err == nil {
			if partyRes != nil && partyRes.Party != nil && partyRes.Party.Party != nil {
				selfState = int32(2)
				selfPartyId = partyRes.Party.Party.Id
			}
		}
		resp := &bffpb.S2CFriendAgree{}
		selfUid, _ := strconv.ParseInt(selfProfile.Profile.Uid, 10, 64)
		//chttodo 部分信息固定
		friend := &bffpb.PlayerSimpleInfo{
			Uid:           selfUid,
			Name:          selfProfile.Profile.Nickname,
			Head:          selfProfile.Profile.Avatar,
			HeroId:        selfProfile.Profile.HeroId,
			Level:         1,
			Online:        true,
			State:         selfState,
			PartyId:       selfPartyId,
			LastLoginTime: int32(time.Now().Unix() / 60), //分钟
		}
		resp.FriendInfo = friend
		//if code := r.pushMqMsg(resp, int32(cpb.S2C_EVENT_S2C_FriendAgree), uid); code != cpb.ERRORCODE_SUCCESS {
		//return rep, code
		//}

		if prep, err := r.pClient.GetProfile(ctx, &pb2.GetProfileRequest{
			Uid: uid,
		}); err != nil {
			return nil, cpb.ERRORCODE_COMMON_ERROR
		} else {
			otherState := int32(0)
			otherPartyId := ""
			otherPartyReq := &pb4.GetPartyRequest{
				Kind: &pb4.GetPartyRequest_Uid{Uid: strconv.FormatInt(req.Uid, 10)},
			}
			if partyRes, err := r.ptClient.GetParty(ctx, otherPartyReq); err == nil {
				if partyRes != nil && partyRes.Party != nil && partyRes.Party.Party != nil {
					otherState = int32(2)
					otherPartyId = partyRes.Party.Party.Id
				}
			}
			//chttodo 部分信息固定
			friend := &bffpb.PlayerSimpleInfo{
				Uid:           req.Uid,
				Name:          prep.Profile.Nickname,
				Head:          prep.Profile.Avatar,
				HeroId:        prep.Profile.HeroId,
				Level:         1,
				Online:        prep.Profile.OnlineStatus == 1,
				State:         otherState,
				PartyId:       otherPartyId,
				LastLoginTime: int32(time.Now().Unix() / 60), //分钟
			}
			resp.FriendInfo = friend
		}
		return resp, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) handleFriendRefuse(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SFriendRefuse{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	_, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	reqRefuse := &pb.RefuseBuddyRequest{}
	for _, uid := range req.Uids {
		reqRefuse.Uid = append(reqRefuse.Uid, strconv.FormatInt(uid, 10))
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if _, err := r.bClient.RefuseBuddy(ctx, reqRefuse); err != nil {
		r.logger.Error("buddy client RefuseBuddy err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		resp := &bffpb.S2CFriendRefuse{}
		resp.Uids = req.Uids
		return resp, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) handleFriendDelete(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SFriendDelete{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	//selfUid, err := request.GetConnection().GetProperty(constants.ConnUid)
	//if err != nil {
	//	return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	//}
	uid := strconv.FormatInt(req.DelUid, 10)
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	//删除同时拉黑
	if req.IsBlacklist {
		reqBlocked := &pb.ProfileIds{
			ProfileIds: []*pb.ProfileId{{
				ProfileId: uid,
			}},
		}
		if _, err := r.bClient.AddBlockedProfiles(ctx, reqBlocked); err != nil {
			r.logger.Error("buddy client addBlockedProfiles err", zap.Error(err))
			return nil, cpb.ERRORCODE_RPC_ERROR
		}

	} else {
		reqRem := &pb.RemoveBuddyRequest{
			Uid: uid,
		}
		if _, err := r.bClient.RemoveBuddy(ctx, reqRem); err != nil {
			r.logger.Error("buddy client removeBuddy err", zap.Error(err))
			return nil, cpb.ERRORCODE_RPC_ERROR
		}
	}
	//通知被删除的玩家
	//id, _ := strconv.ParseInt(selfUid.(string), 10, 64)
	//remNotice := &bffpb.S2CFriendDelete{
	//	DelUid: id,
	//}
	//if code := r.pushMqMsg(remNotice, int32(cpb.S2C_EVENT_S2C_FriendDelete), uid); code != cpb.ERRORCODE_SUCCESS {
	//	//return nil, code
	//}

	resp := &bffpb.S2CFriendDelete{
		DelUid:      req.DelUid,
		IsBlacklist: req.IsBlacklist,
	}
	return resp, cpb.ERRORCODE_SUCCESS
}

func (r *Router) handleFriendNickName(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SFriendNickName{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	_, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	uid := strconv.FormatInt(req.Uid, 10)
	remarkReq := &pb.RemarkRequest{
		Uid:    uid,
		Remark: req.NickName,
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if _, err := r.bClient.Remark(ctx, remarkReq); err != nil {
		r.logger.Error("buddy client remark err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	resp := &bffpb.S2CFriendNickName{
		Uid:      req.Uid,
		NickName: req.NickName,
	}
	return resp, cpb.ERRORCODE_SUCCESS
}

func (r *Router) handleFriendDeleteBlack(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bffpb.C2SFriendDeleteBlack{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	_, err := request.GetConnection().GetProperty(constants.ConnUid)
	if err != nil {
		return nil, cpb.ERRORCODE_PLAYER_NOT_LOGIN
	}
	uid := strconv.FormatInt(req.Uid, 10)
	deleteReq := &pb.ProfileIds{
		ProfileIds: []*pb.ProfileId{{
			ProfileId: uid,
		}},
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if _, err := r.bClient.RemoveBlockedProfiles(ctx, deleteReq); err != nil {
		r.logger.Error("buddy client removeBlockedProfiles err", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	resp := &bffpb.S2CFriendDeleteBlack{
		Uid: req.Uid,
	}
	return resp, cpb.ERRORCODE_SUCCESS
}
func (r *Router) getErrorCode(err string) cpb.ERRORCODE {
	sts := strings.Split(err, "=")
	str := strings.TrimSpace(sts[len(sts)-1])
	switch str {
	case "ErrBuddyAlreadyAdded":
		return cpb.ERRORCODE_FRIEND_IS_REPEAT
	case "ErrBuddyAlreadyRequested":
		return cpb.ERRORCODE_FRIEND_ALREADY_APPLY
	case "ErrBuddyAlreadyInYourRequestList":
		return cpb.ERRORCODE_FRIEND_ALREADY_APPLY
	case "ErrBuddiesNotFound":
		return cpb.ERRORCODE_FRIEND_NOT_EXIST
	case "ErrInviterNotFound":
		return cpb.ERRORCODE_FRIEND_NOT_IN_REQUEST
	case "ErrSelfBuddiesTopLimit":
		return cpb.ERRORCODE_FRIEND_MY_COUNT_MAX
	case "ErrSelfInviterTopLimit":
		return cpb.ERRORCODE_FRIEND_MY_COUNT_MAX
	case "ErrTargetInviterTopLimit":
		return cpb.ERRORCODE_FRIEND_YOU_COUNT_MAX
	case "ErrTargetBuddiesTopLimit":
		return cpb.ERRORCODE_FRIEND_YOU_COUNT_MAX
	case "ErrCanNotAddSelf":
		return cpb.ERRORCODE_FRIEND_YOU_COUNT_MAX
	case "ErrInTargetBlockedList":
		return cpb.ERRORCODE_FRIEND_IN_YOU_BLACKLIST
	case "ErrInSelfBlockedList":
		return cpb.ERRORCODE_FRIEND_IN_MY_BLACKLIST
	case "ErrBlockedNumExceed":
		//return cpb.
	}
	return cpb.ERRORCODE_FRIEND_NOT_IN_REQUEST
}

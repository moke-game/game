package profile

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/gstones/zinx/ziface"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	errors2 "github.com/moke-game/platform/services/profile/errors"

	pb2 "github.com/moke-game/platform/api/gen/knapsack"
	pb "github.com/moke-game/platform/api/gen/profile"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	configs "github.com/moke-game/game/configs/pkg/module"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/wordsfilter"
)

func (r *Router) handleNewPlayer(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bff.C2SNewPlayer{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}

	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	nameMaxLen := r.configs.TblGlobalConfig.Get(116).Num
	if !configs.DeploymentGlobal.IsProd() {
		nameMaxLen = 10000
	}

	heroId := req.HeroId

	unlockConfig := configs.ConfigsGlobal.TblCharacterUnlock.Get(heroId)
	if unlockConfig == nil {
		return nil, cpb.ERRORCODE_PLAYER_HERO_ILLEGAL
	}

	//heroId := int32(10010)
	if len(req.GetName()) == 0 {
		r.logger.Error("name empty")
		return nil, cpb.ERRORCODE_PLAYER_NAME_LEN_ILLEGAL
	} else if len(req.GetName()) > int(nameMaxLen) {
		r.logger.Error("name too long", zap.String("name", req.GetName()))
		return nil, cpb.ERRORCODE_PLAYER_NAME_LEN_ILLEGAL
	} else if wordsfilter.IsSensitive(req.GetName()) {
		r.logger.Error("name sensitive", zap.String("name", req.GetName()))
		return nil, cpb.ERRORCODE_PLAYER_NAME_ILLEGAL
	} else if h := r.configs.TblCharacter.Get(req.HeroId); h == nil {
		r.logger.Error("hero id illegal", zap.Int32("heroId", req.HeroId))
		return nil, cpb.ERRORCODE_PLAYER_HERO_ILLEGAL
	} else if rProfile, err := r.pClient.CreateProfile(ctx, &pb.CreateProfileRequest{
		Profile: &pb.Profile{
			Nickname: req.GetName(),
			HeroId:   heroId,
			Avatar:   fmt.Sprintf("%d", heroId),
		},
	}); err != nil {
		r.logger.Error("create profile error", zap.Error(err), zap.Any("resp", rProfile))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else if resp, err := r.kClient.AddItem(ctx, &pb2.AddItemRequest{
		Items: r.makeDefaultItems(),
	}); err != nil {
		r.logger.Error("add item error", zap.Error(err), zap.Any("resp", resp))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		info := &bff.PlayerSimpleInfo{
			Uid:          cast.ToInt64(rProfile.Profile.Uid),
			Name:         rProfile.Profile.Nickname,
			Head:         rProfile.Profile.Avatar,
			State:        rProfile.Profile.OnlineStatus,
			HeroId:       heroId,
			PetProfileId: rProfile.Profile.PetProfileId,
		}
		return &bff.S2CNewPlayer{
			SimpleInfo: info,
		}, cpb.ERRORCODE_SUCCESS
	}
}

func (r *Router) makeDefaultItems() map[int64]*pb2.Item {
	items := make(map[int64]*pb2.Item)
	for k, v := range r.configs.TblItem.GetDataMap() {
		if v.Default > 0 {
			itemId := int64(k)
			expire := int64(0)
			items[itemId] = &pb2.Item{
				Id:     itemId,
				Num:    v.Default,
				Type:   v.ID,
				Expire: expire,
			}
		}
	}
	return items
}
func (r *Router) handleSimpleInfo(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bff.C2SSimpleInfo{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	uid := strconv.FormatInt(req.GetUid(), 10)
	if resp, err := r.pClient.GetProfile(ctx, &pb.GetProfileRequest{
		Uid: uid,
	}); err != nil {
		if errors.Is(err, errors2.ErrNotFound) {
			return nil, cpb.ERRORCODE_NO_USER
		}
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		info := &bff.PlayerSimpleInfo{
			Uid:    req.GetUid(),
			Name:   resp.Profile.Nickname,
			HeroId: resp.Profile.HeroId,
			Head:   resp.Profile.Avatar,
			Online: resp.Profile.OnlineStatus == 1,
		}
		return &bff.S2CSimpleInfo{
			Info: &bff.PlayerInfo{
				SimpleInfo: info,
			},
		}, cpb.ERRORCODE_SUCCESS
	}
}
func (r *Router) handleRename(ctx context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bff.C2SPlayerRename{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	if profileResp, err := r.pClient.GetProfile(ctx, &pb.GetProfileRequest{}); err != nil {
		r.logger.Error("handleRename GetProfile error", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	} else {
		renameCfg := configs.ConfigsGlobal.TblGlobalConfig.Get(common.RENAME_CONFIG_NICK_MAX)
		if len(req.Name) > int(renameCfg.Num) {
			return nil, cpb.ERRORCODE_RENAME_LENGTH_MAX
		}
		if isSensitive := wordsfilter.IsSensitive(req.Name); isSensitive {
			return nil, cpb.ERRORCODE_RENAME_NOT_ALLOW
		}
		res := &bff.S2CPlayerRename{
			Name: req.Name,
		}
		if profileResp.Profile.Nickname == req.Name {
			return res, cpb.ERRORCODE_SUCCESS
		}
		profileResp.Profile.Nickname = req.Name

		_, err := r.pClient.UpdateProfile(ctx, &pb.UpdateProfileRequest{Profile: profileResp.Profile})
		if err != nil {
			r.logger.Error("handleRename UpdateProfile error", zap.Error(err))
			return nil, cpb.ERRORCODE_RENAME_REPEAT
		}
		return res, cpb.ERRORCODE_SUCCESS
	}
}

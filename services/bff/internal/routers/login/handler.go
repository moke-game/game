package login

import (
	"context"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/constants"
	apb "github.com/moke-game/platform/api/gen/auth"
	ppb "github.com/moke-game/platform/api/gen/profile"
)

// 3rd login auth verify
// 第三方登陆验证
func (r *Router) umsCheck(req *bff.C2SAuth) (string, error) {
	//只有生产环境走ums验证
	//if configs.DeploymentGlobal.IsProd() {
	// check ums token
	// 第三方登陆校验
	//}
	//测试环境直接返回
	return req.Openid, nil
}

func (r *Router) handleAuth(ctx context.Context, request ziface.IRequest) (mess proto.Message, code cpb.ERRORCODE) {
	var req = &bff.C2SAuth{}
	//获取请求数据
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		r.logger.Error("unmarshal request fail", zap.Error(err))
		return nil, cpb.ERRORCODE_PROTO_UNMARSHAL_ERROR
	}

	// 第三方登陆校验
	uid3Rd, err := r.umsCheck(req)
	if err != nil {
		r.logger.Error("ums check fail", zap.Error(err))
		return nil, cpb.ERRORCODE_COMMON_ERROR
	}

	//申请token
	aResp, err := r.aClient.Authenticate(ctx, &apb.AuthenticateRequest{
		AppId: r.appId,
		Id:    uid3Rd,
	})
	if err != nil {
		r.logger.Error("authenticate fail", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}

	//写入connection属性
	r.setAttrs(request, aResp)
	ctx = common.MakeAuthCtxOut(ctx, request.GetConnection())
	resp, err := r.pClient.IsProfileExist(ctx, &ppb.IsProfileExistRequest{
		Uid: aResp.Uid,
	})
	if err != nil {
		r.logger.Error("check profile exist fail", zap.Error(err))
		return nil, cpb.ERRORCODE_RPC_ERROR
	}
	if !resp.Exist {
		//create profile
		profile := r.initProfile(aResp.Uid)
		if _, err := r.pClient.CreateProfile(ctx, &ppb.CreateProfileRequest{
			Profile: profile,
		}); err != nil {
			r.logger.Error("create profile fail", zap.Error(err))
			return nil, cpb.ERRORCODE_RPC_ERROR
		}
	}
	return &bff.S2CAuth{}, cpb.ERRORCODE_SUCCESS
}

func (r *Router) initProfile(uid string) *ppb.Profile {
	return &ppb.Profile{
		Uid:      uid,
		Nickname: random.RandString(6),
	}
}
func (r *Router) setAttrs(request ziface.IRequest, aResp *apb.AuthenticateResponse) {
	request.GetConnection().SetProperty(constants.ConnUid, aResp.Uid)
	request.GetConnection().SetProperty(constants.ConnToken, aResp.GetAccessToken())
}

func (r *Router) handleHeartbeat(_ context.Context, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	req := &bff.C2SHeartbeat{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, cpb.ERRORCODE_COMMON_ERROR
	} else {
		return &bff.S2CHeartbeat{
			SysTime: int32(time.Now().Unix()),
			Params:  req.Params,
		}, cpb.ERRORCODE_SUCCESS
	}
}

package service

import pb "github.com/moke-game/platform/api/gen/profile"

func makeGetProfileReq(tp, val, platform string, page, pageSize int32) *pb.GetProfilePrivateRequest {
	req := &pb.GetProfilePrivateRequest{
		PlatformId: platform,
	}
	switch tp {
	case "1": // 角色名(模糊查询)
		req.Kind = &pb.GetProfilePrivateRequest_Name_{
			Name: &pb.GetProfilePrivateRequest_Name{
				Name:     val,
				IsRegexp: true,
				Page:     page,
				PageSize: pageSize,
			},
		}
	case "2": // 角色ID
		req.Kind = &pb.GetProfilePrivateRequest_Uids_{
			Uids: &pb.GetProfilePrivateRequest_Uids{
				Uid: []string{val},
			},
		}
	case "3": // 账号ID
		req.Kind = &pb.GetProfilePrivateRequest_Account{
			Account: val,
		}
	case "6": // 角色名(精确查询)
		req.Kind = &pb.GetProfilePrivateRequest_Name_{
			Name: &pb.GetProfilePrivateRequest_Name{
				Name: val,
			},
		}
	}
	return req
}

package party

import (
	"context"
	"strconv"
	"strings"

	"github.com/gstones/moke-kit/orm/nosql/key"
	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ptpb "github.com/moke-game/platform/api/gen/party"
	ppb "github.com/moke-game/platform/api/gen/profile"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common"
	"github.com/moke-game/game/services/common/msg_transfer"
)

func (r *Router) watchParty(
	connect ziface.IConnection,
	profile *ppb.ProfileBasic,
	partyId string,
	isCreate bool,
	playId int32,
	isPublish bool,
	isReWatch bool,
) error {
	// 关闭之前的stream
	if err := destroyParty(connect); err != nil {
		return err
	}
	// 开启新的stream
	ctx := common.MakeAuthCtxOut(connect.Context(), connect)
	member := msg_transfer.ProfileToPartyMember(profile)
	member.Status = int32(bff.PartyMemberStatus_STATUS_NOT_READY)
	party := &ptpb.PartySetting{}
	if isCreate {
		party = createDefaultParty(playId)
	}
	ctx, cancel := context.WithCancel(ctx)
	if stream, err := r.ptClient.JoinParty(ctx, &ptpb.JoinPartyRequest{
		Id:        partyId,
		Member:    member,
		IsCreate:  isCreate,
		Party:     party,
		IsPublish: isPublish,
	}); err != nil {
		cancel()
		return err
	} else {
		go func() {
			defer cancel()
			isResp := false
			if isReWatch {
				isResp = true
			}
			for {
				msgRecv, err := stream.Recv()
				if err != nil {
					if status.Code(err) == codes.Canceled {
						r.logger.Info("party connection closed")
						return
					}
					r.logger.Error("party changes receive msg error", zap.Error(err))
					return
				}
				if !isResp {
					if msg, err := msg_transfer.Party(msgRecv.Party); err != nil {
						r.logger.Error("party changes transfer msg error", zap.Error(err))
					} else {
						ntf := &bff.NtfPartyUpdate{
							PartyInfo: msg,
						}
						if err := common.SendNotify(
							connect,
							uint32(cpb.S2C_EVENT_NTF_PartyUpdate),
							ntf,
						); err != nil {
							r.logger.Error("party changes send msg error", zap.Error(err))
							return
						}
						connect.SetProperty(IdPropertyKey, msg.PartyId)
					}
				} else {
					r.noticePartyChanges(profile.Uid, connect, ctx, msgRecv.Party)
				}
				isResp = true
			}
		}()
		connect.SetProperty(StreamCancel, cancel)
		go func() {
			<-connect.Context().Done()
			if _, err := r.ptClient.LeaveParty(ctx, &ptpb.LeavePartyRequest{}); err != nil {
				r.logger.Error("leave party error", zap.Error(err))
			}
			cancel()
		}()
	}
	return nil
}
func (r *Router) noticePartyChanges(uid string, connect ziface.IConnection, ctx context.Context, changes *ptpb.PartyInfo) {
	if changes.GetParty() != nil {
		notice := &bff.NtfPartyInfo{}
		if changes.GetParty().GetType() != 0 {
			notice.PlayId = changes.GetParty().GetType()
		}
		if changes.GetParty().GetOwner() != "" {
			if uid, err := strconv.ParseInt(changes.GetParty().GetOwner(), 10, 64); err != nil {
				r.logger.Error("parse owner error", zap.Error(err))
			} else {
				notice.Owner = uid
			}
		}
		if err := common.SendNotify(
			connect,
			uint32(cpb.S2C_EVENT_NTF_PartyInfo),
			notice,
		); err != nil {
			r.logger.Error("send party map error", zap.Error(err))
		}
	}

	statusNotice := &bff.NtfPartyMemberStatus{}
	statusNotice.Status = make(map[int64]bff.PartyMemberStatus)
	leaveNotice := &bff.NtfPartyMemberLeave{}
	joinNotice := &bff.NtfPartyMemberJoin{}
	joinNotice.Members = make([]*bff.PartyMember, 0)
	heroNotice := &bff.NtfPartyMemberHero{}
	heroNotice.Heros = make(map[int64]int32)
	for k, v := range changes.GetMembers() {
		mid, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			continue
		}
		if v.Nickname != "" {
			if mem, err := msg_transfer.PartyMember(v); err != nil {
				r.logger.Error("transfer party member error", zap.Error(err))
			} else {
				joinNotice.Members = append(joinNotice.Members, mem)
			}
		} else if v.GetStatus() != 0 {
			statusNotice.Status[mid] = bff.PartyMemberStatus(v.GetStatus())
		} else if v.GetIsLeave() {
			leaveNotice.MemberIds = append(leaveNotice.MemberIds, mid)
			leaveNotice.LeaveReason = v.LeaveReason
			if strings.Compare(k, uid) == 0 {
				if err := destroyParty(connect); err != nil {
					r.logger.Error("destroy party error", zap.Error(err))
				}
			}
		} else if v.GetHeroId() != 0 {
			heroNotice.Heros[mid] = v.GetHeroId()
		}
	}
	if len(statusNotice.Status) > 0 {
		if err := common.SendNotify(
			connect,
			uint32(cpb.S2C_EVENT_NTF_PartyMemberStatus),
			statusNotice,
		); err != nil {
			r.logger.Error("send party member status error", zap.Error(err))
		}
	}

	if len(joinNotice.Members) > 0 {
		r.logger.Info("party member join", zap.Any("msg", joinNotice))
		if err := common.SendNotify(
			connect,
			uint32(cpb.S2C_EVENT_NTF_PartyMemberJoin),
			joinNotice,
		); err != nil {
			r.logger.Error("send party member join error", zap.Error(err))
		}
	}

	if len(leaveNotice.MemberIds) > 0 {
		if err := common.SendNotify(
			connect,
			uint32(cpb.S2C_EVENT_NTF_PartyMemberLeave),
			leaveNotice,
		); err != nil {
			r.logger.Error("send party member leave error", zap.Error(err))
		}
	}

	if len(heroNotice.Heros) > 0 {
		if err := common.SendNotify(
			connect,
			uint32(cpb.S2C_EVENT_NTF_PartyMemberHero),
			heroNotice,
		); err != nil {
			r.logger.Error("send party member hero error", zap.Error(err))
		}
	}
}

func createDefaultParty(tp int32) *ptpb.PartySetting {
	return &ptpb.PartySetting{
		Type:      tp,
		MaxMember: 3,
	}
}

func destroyParty(connect ziface.IConnection) error {
	if party, err := connect.GetProperty(StreamCancel); err == nil {
		if cancel, ok := party.(context.CancelFunc); ok {
			cancel()
		}
		connect.RemoveProperty(StreamCancel)
		connect.RemoveProperty(IdPropertyKey)
	}
	return nil
}

func (r *Router) GetRoomPlayerAmount(basic *ppb.ProfileBasic) (int, error) {
	hostName := basic.RoomHostname
	if hostName == "" {
		hostName = "room"
	}
	roomPlayerKey, err := key.NewKeyFromParts("battle", "room", "player", hostName, basic.RoomId)
	if err != nil {
		return 0, err
	}
	result, err := r.redis.Get(context.TODO(), roomPlayerKey.String()).Int()
	if err != nil {
		return 0, err
	}
	return result, err
}

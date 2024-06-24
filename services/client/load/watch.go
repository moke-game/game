package load

import (
	"io"
	"net"

	"github.com/gstones/zinx/ziface"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	roompb "github.com/moke-game/game/api/gen/room"
	"github.com/moke-game/game/services/client/common"
)

func (r *Robot) watchResponse(conn net.Conn) {
	for {
		resp, err := common.RecvMsg(conn)

		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, net.ErrClosed) {
				return
			}
			r.shell.PrintErrln(err)
			return
		}
		r.handleResponse(resp)
	}
}

func (r *Robot) handleResponse(msg ziface.IMessage) {
	data := msg.GetData()
	if msg.GetMsgID() > 2000 {
		resp := &cpb.Response{}
		if err := proto.Unmarshal(msg.GetData(), resp); err != nil {
			r.shell.PrintErrln(err)
			return
		}
		if resp.Code != 0 {
			r.shell.Printf("receive err resp: %v \r\n", resp)
			return
		}
		data = resp.Data
	}

	switch msg.GetMsgID() {
	case uint32(cpb.S2C_EVENT_S2C_Auth):
		resp := &bff.S2CAuth{}
		r.printResp(data, resp)
		r.token = resp.Token
		r.uid = resp.Uid
		if err := r.unlockFeats(); err != nil {
			r.shell.PrintErrln(err)
		} else if err := r.unlockHero(); err != nil {
			r.shell.PrintErrln(err)
		} else if err := r.chooseHero(); err != nil {
			r.shell.PrintErrln(err)
		} else if err := r.getMissions(); err != nil {
			r.shell.PrintErrln(err)
		}
	case uint32(cpb.S2C_EVENT_S2C_NewPlayer):
		r.printResp(data, &bff.S2CNewPlayer{})
	case uint32(cpb.S2C_EVENT_S2C_SimpleInfo):
		r.printResp(data, &bff.S2CSimpleInfo{})
	case uint32(cpb.S2C_EVENT_S2C_GetRoomInfo):
		resp := &bff.S2CGetRoomInfo{}
		r.printResp(data, resp)
		if err := r.enterRoom(resp); err != nil {
			r.shell.PrintErrln(err)
		}
	case uint32(cpb.S2C_EVENT_S2C_EnterRoom):
		resp := &roompb.S2CEnterRoom{}
		r.printResp(data, resp)
		r.mapID = resp.MapId
		for _, v := range resp.Players {
			if v.Uid == r.uid {
				r.pos = v.GetPos()
				r.speed = float64(v.GetSpeed()) / 10000.0
				break
			}
		}
		if err := r.enterScene(); err != nil {
			r.shell.PrintErrln(err)
		}
	//case uint32(cpb.S2C_EVENT_S2C_MatchingGroupReady):
	//	resp := &bff.S2CMatchingGroupReady{}
	//	r.printResp(data, resp)

	case uint32(cpb.S2C_EVENT_S2C_EnterScene):
		r.printResp(data, &roompb.S2CEnterScene{})
	case uint32(cpb.S2C_EVENT_S2C_WatchingKnapsack):
		r.printResp(data, &bff.S2CWatchingKnapsack{})
	case uint32(cpb.S2C_EVENT_S2C_DiamondExchangeItem):
		r.printResp(data, &bff.S2CDiamondExchangeItem{})
	case uint32(cpb.S2C_EVENT_S2C_CHATReceiveWorldMessage):
		r.printResp(data, &bff.S2CCHATReceiveWorldMessage{})
	case uint32(cpb.S2C_EVENT_S2C_CHATReceivePlayerMessage):
		r.printResp(data, &bff.S2CCHATReceivePlayerMessage{})
	case uint32(cpb.S2C_EVENT_S2C_CHATReceiveTeamMessage):
		r.printResp(data, &bff.S2CCHATReceiveTeamMessage{})
	case uint32(cpb.S2C_EVENT_S2C_JoinParty):
		r.printResp(data, &bff.S2CJoinParty{})
	case uint32(cpb.S2C_EVENT_S2C_LeaveParty):
		r.printResp(data, &bff.S2CLeaveParty{})
	case uint32(cpb.S2C_EVENT_S2C_KickParty):
		r.printResp(data, &bff.S2CKickParty{})
	case uint32(cpb.S2C_EVENT_S2C_READY_PARTY):
		r.printResp(data, &bff.S2CReadyParty{})
	case uint32(cpb.S2C_EVENT_S2C_CANCEL_READY_PARTY):
		r.printResp(data, &bff.S2CCancelReadyParty{})
	case uint32(cpb.S2C_EVENT_NTF_PartyInfo):
		r.printResp(data, &bff.NtfPartyInfo{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberStatus):
		r.printResp(data, &bff.NtfPartyMemberStatus{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberHero):
		r.printResp(data, &bff.NtfPartyMemberHero{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberJoin):
		r.printResp(data, &bff.NtfPartyMemberJoin{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberLeave):
		r.printResp(data, &bff.NtfPartyMemberLeave{})
	case uint32(cpb.S2C_EVENT_NTF_RoomSkills):
		r.printResp(data, &roompb.NtfRoomSkills{})
	case uint32(cpb.S2C_EVENT_NTF_RoomHits):
		r.printResp(data, &roompb.NtfRoomHits{})
	case uint32(cpb.S2C_EVENT_NTF_RoomBuffEffect):
		r.printResp(data, &roompb.NtfRoomBuffEffect{})
	case uint32(cpb.S2C_EVENT_NTF_RoomBuffKnockback):
		r.printResp(data, &roompb.NtfRoomBuffKnockback{})
	case uint32(cpb.S2C_EVENT_NTF_UnitEnterEyeshot):
		r.printResp(data, &roompb.NtfUnitEnterEyeshot{})
	case uint32(cpb.S2C_EVENT_NTF_PlayerEnterEyeshot):
		r.printResp(data, &roompb.NtfPlayerEnterEyeshot{})
	case uint32(cpb.S2C_EVENT_NTF_UnitLeaveEyeshot):
		r.printResp(data, &roompb.NtfUnitLeaveEyeshot{})
	case uint32(cpb.S2C_EVENT_NTF_UnitMove):
		r.printResp(data, &roompb.NtfUnitMove{})
	case uint32(cpb.S2C_EVENT_NTF_UnitMovementBatch):
		resp := &roompb.NtfUnitMovementBatch{}
		r.printResp(data, resp)
		for _, m := range resp.Movements {
			if f := m.GetFlash(); f != nil {
				if f.UnitId == r.uid {
					r.OnStop(f.GetPos())
				}
			} else if s := m.GetStop(); s != nil {
				if s.UnitId == r.uid {
					r.OnStop(s.GetPos())
				}
			}
		}
	case uint32(cpb.S2C_EVENT_NTF_AttributeUpdate):
		resp := &roompb.NtfAttributeUpdate{}
		r.printResp(data, resp)
		if resp.UnitId == r.uid {
			if spd, ok := resp.Changes[5]; ok {
				r.speed = float64(spd) / 10000.0
			}
		}
	}

}

func (r *Robot) printResp(data []byte, pMsg proto.Message) {
	if err := proto.Unmarshal(data, pMsg); err != nil {
		r.shell.PrintErrln(err)
	}
}

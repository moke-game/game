package room

import (
	"github.com/gstones/zinx/ziface"
	"google.golang.org/protobuf/proto"

	cpb "github.com/moke-game/game/api/gen/common"
	roompb "github.com/moke-game/game/api/gen/room"
)

func (r *Room) enterRoom(uid string, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	return &roompb.S2CEnterRoom{
		RoomId: r.RoomId(),
		MapId:  r.mapId,
	}, cpb.ERRORCODE_SUCCESS
}
func (r *Room) exitRoom(uid string, _ ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	return &roompb.S2CLeaveRoom{}, cpb.ERRORCODE_SUCCESS

}

func (r *Room) move(uid string, request ziface.IRequest) (proto.Message, cpb.ERRORCODE) {
	return &roompb.S2CMove{}, cpb.ERRORCODE_SUCCESS
}

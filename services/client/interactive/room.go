package interactive

import (
	"net"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	roompb "github.com/moke-game/game/api/gen/room"
	"github.com/moke-game/game/services/client/common"
)

type Room struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewRoom(conn net.Conn) *Room {
	l := &Room{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Room) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "room",
		Help:    "room interactive shell",
		Aliases: []string{"r"},
	}
	l.initSubCmd()
}
func (l *Room) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Room) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "get",
		Help:    "get room info",
		Aliases: []string{"g"},
		Func:    l.get,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "move",
		Help:    "room move",
		Aliases: []string{"m"},
		Func:    l.move,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "skill",
		Help:    "use skill",
		Aliases: []string{"s"},
		Func:    l.useSkill,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "enter",
		Help:    "enter room",
		Aliases: []string{"en"},
		Func:    l.enter,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "exit",
		Help:    "exit room",
		Aliases: []string{"ex"},
		Func:    l.exit,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "enterscene",
		Aliases: []string{"es"},
		Func:    l.enterScene,
		Help:    "enter scene",
	})
}

func (l *Room) move(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &roompb.C2SMove{
		Current: &roompb.Vector{
			X: 1505550,
			Y: 0,
			Z: 1979753,
		},
		Direct: 1,
		Dest:   &roompb.Vector{},
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_Move, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "move send success: %v \n", req.String())
}

func (l *Room) enter(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	token := slogger.ReadLine(c, "token: ")
	roomIdIn := slogger.ReadLine(c, "roomId: ")
	//roomId, err := strconv.ParseInt(roomIdIn, 10, 64)
	//if err != nil {
	//	slogger.Warn(c, err)
	//	return
	//}
	req := &roompb.C2SEnterRoom{
		Token:  token,
		RoomId: roomIdIn,
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_EnterRoom, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "enter send success: %v \n", req.String())
}

func (l *Room) exit(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &roompb.S2CLeaveRoom{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_LeaveRoom, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "exit send success: %v \n", req.String())
}

func (l *Room) enterScene(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &roompb.C2SEnterScene{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_EnterScene, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "enter scene success!")
}

func (l *Room) useSkill(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	idIn := slogger.ReadLine(c, "skill id: ")
	id, err := strconv.Atoi(idIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := &roompb.C2SUseSkill{
		Skill: &roompb.Skill{
			Cid: int32(id),
			SelfPos: &roompb.Vector{
				X: 0,
				Y: 0,
				Z: 0,
			},
			TargetPos: &roompb.Vector{
				X: 1,
				Y: 0,
				Z: 1,
			},
		},
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_UseSkill, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "use skill send success: %v \n", req.String())
}

func (l *Room) get(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SGetRoomInfo{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_GetRoomInfo, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get room info send success: %v \n", req.String())
}

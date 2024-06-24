package interactive

import (
	"net"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Party struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewParty(conn net.Conn) *Party {
	l := &Party{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Party) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "party",
		Help:    "party interactive shell",
		Aliases: []string{"pt"},
	}
	l.initSubCmd()
}
func (l *Party) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Party) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "join",
		Help:    "join party",
		Aliases: []string{"j"},
		Func:    l.joinParty,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "leave",
		Help:    "leave party",
		Aliases: []string{"l"},
		Func:    l.leaveParty,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "kick",
		Help:    "kick party",
		Aliases: []string{"k"},
		Func:    l.kickParty,
	})

	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "ready",
		Help:    "ready party member",
		Aliases: []string{"r"},
		Func:    l.ReadyParty,
	})

	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "cancel",
		Help:    "cancel ready party member",
		Aliases: []string{"c"},
		Func:    l.CancelReadyParty,
	})

	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "map",
		Help:    "choose map",
		Aliases: []string{"m"},
		Func:    l.chooseMap,
	})
}
func (l *Party) joinParty(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	partyId := slogger.ReadLine(c, "party id: ")
	req := &bff.C2SJoinParty{
		PartyId: partyId,
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_JoinParty, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "join party send success: %v \n", req.String())
}

func (l *Party) leaveParty(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SLeaveParty{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_LeaveParty, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "leave party send success: %v \n", req.String())
}

func (l *Party) kickParty(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	playerIn := slogger.ReadLine(c, "player id: ")
	playerId, err := strconv.Atoi(playerIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := &bff.C2SKickParty{
		MemberId: int64(playerId),
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_KickParty, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "kick party send success: %v \n", req.String())
}

func (l *Party) ReadyParty(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SReadyParty{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_READY_PARTY, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "ready party  send success: %v \n", req.String())
}

func (l *Party) CancelReadyParty(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SCancelReadyParty{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_CANCEL_READY_PARTY, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "cancel ready party send success: %v \n", req.String())
}

func (l *Party) chooseMap(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	mapIn := slogger.ReadLine(c, "map id: ")
	mapId, err := strconv.Atoi(mapIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := &bff.C2SChoosePartyMap{
		PlayId: int32(mapId),
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_ChoosePartyMap, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "choose map send success: %v \n", req.String())
}

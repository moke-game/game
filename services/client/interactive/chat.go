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

type Chat struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewChat(conn net.Conn) *Chat {
	l := &Chat{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Chat) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "chat",
		Help:    "chat interactive shell",
		Aliases: []string{"c"},
	}
	l.initSubCmd()
}
func (l *Chat) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Chat) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "msg",
		Help:    "send msg",
		Aliases: []string{"m"},
		Func:    l.sendMsg,
	})

}
func (l *Chat) sendMsg(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	msg := slogger.ReadLine(c, "msg: ")
	cIn := slogger.ReadLine(c, "channel(0:world,1:player,2:team): ")

	channel, err := strconv.Atoi(cIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	target := 0
	if channel != 0 {
		targetIn := slogger.ReadLine(c, "target id(team id or player uid): ")
		target, err = strconv.Atoi(targetIn)
		if err != nil {
			slogger.Warn(c, err)
			return
		}
	}

	req := &bff.C2SCHATMessage{
		ChatType: bff.ChatType(channel),
		ChatInfo: &bff.ChatInfo{
			ReceiveUid: int64(target),
			Content:    msg,
		},
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_CHATMessage, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}

	slogger.Infof(c, "send msg send success: %v \n", req.String())
}

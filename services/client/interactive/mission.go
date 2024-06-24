package interactive

import (
	"net"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Mission struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewMission(conn net.Conn) *Mission {
	l := &Mission{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Mission) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "mission",
		Help:    "mission interactive shell",
		Aliases: []string{"ms"},
	}
	l.initSubCmd()
}
func (l *Mission) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Mission) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "get",
		Help:    "get missions",
		Aliases: []string{"g"},
		Func:    l.getMissions,
	})

	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "touch",
		Help:    "touch missions",
		Aliases: []string{"t"},
		Func:    l.touchMissions,
	})
}

func (l *Mission) touchMissions(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2STouchMission{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_TouchMission, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "touch missions send success: %v \n", req.String())
}

func (l *Mission) getMissions(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SGetMission{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_GetMission, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get missions send success: %v \n", req.String())
}

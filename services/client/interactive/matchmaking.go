package interactive

import (
	"net"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Match struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewMatch(conn net.Conn) *Match {
	m := &Match{
		conn: conn,
	}
	m.initShell()
	return m
}

func (m *Match) initShell() {
	m.cmd = &ishell.Cmd{
		Name:    "match",
		Help:    "matchmaking interactive shell",
		Aliases: []string{"mm"},
	}
	m.initSubCmd()
}
func (m *Match) GetCmd() *ishell.Cmd {
	return m.cmd
}

func (m *Match) initSubCmd() {
	m.cmd.AddCmd(&ishell.Cmd{
		Name:    "single",
		Help:    "single match ready",
		Aliases: []string{"s"},
		Func:    m.singleMath,
	})
	m.cmd.AddCmd(&ishell.Cmd{
		Name:    "cancel",
		Help:    "single match cancel",
		Aliases: []string{"c"},
		Func:    m.Cancel,
	})
}

func (m *Match) singleMath(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SMatchingSingleStart{}
	err := common.SendMsg(m.conn, cpb.C2S_EVENT_C2S_MatchingSingleStart, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "single match send success: %v \n", req.String())
}

func (m *Match) Cancel(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SMatchingCancel{}
	err := common.SendMsg(m.conn, cpb.C2S_EVENT_C2S_MatchingCancel, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "single match cancel send success: %v \n", req.String())
}

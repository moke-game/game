package interactive

import (
	"net"

	"github.com/Pallinder/go-randomdata"
	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Login struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewLogin(conn net.Conn) *Login {
	l := &Login{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Login) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "login",
		Help:    "login interactive shell",
		Aliases: []string{"l"},
	}
	l.initSubCmd()
}
func (l *Login) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Login) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "auth",
		Help:    "auth",
		Aliases: []string{"a"},
		Func:    l.auth,
	})
	// logout
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "logout",
		Help:    "logout",
		Aliases: []string{"o"},
		Func:    l.logout,
	})
}

func (l *Login) auth(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	openId := slogger.ReadLine(c, "account: ")

	req := &bff.C2SAuth{
		Token:       openId,
		Openid:      openId,
		MachineCode: openId,
		CountryCode: randomdata.Country(randomdata.TwoCharCountry),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_Auth, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "auth send success: %v \n", req.String())
}

func (l *Login) logout(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SLogoff{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_Logoff, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "logout send success: %v  \n", req.String())
}

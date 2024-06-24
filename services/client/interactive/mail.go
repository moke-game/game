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

type Mail struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewMail(conn net.Conn) *Mail {
	l := &Mail{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Mail) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "mail",
		Help:    "Mail interactive shell",
		Aliases: []string{"m"},
	}
	l.initSubCmd()
}
func (l *Mail) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Mail) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "reward",
		Help:    "get mail rewards",
		Aliases: []string{"r"},
		Func:    l.getMailRewards,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "read",
		Help:    "read mail",
		Aliases: []string{"read"},
		Func:    l.readMails,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "delete",
		Help:    "delete read mails",
		Aliases: []string{"d"},
		Func:    l.deleteReadMails,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "watch",
		Help:    "watch mail changes",
		Aliases: []string{"w"},
		Func:    l.watchMail,
	})
}

func (l *Mail) getMailRewards(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	ids := make([]int64, 0)
	idIn := slogger.ReadLine(c, "mail id: ")
	mid, err := strconv.Atoi(idIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	if mid > 0 {
		ids = append(ids, int64(mid))
	}
	req := &bff.C2SGetMailRewards{
		Ids: ids,
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_GetMailRewards, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get mail rewards send success: %v \n", req.String())
}

func (l *Mail) readMails(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	ids := make([]int64, 0)
	idIn := slogger.ReadLine(c, "mail id: ")
	mid, err := strconv.Atoi(idIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	if mid > 0 {
		ids = append(ids, int64(mid))
	}

	req := &bff.C2SReadMail{
		Ids: ids,
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_ReadMail, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "read mails send success: %v \n", req.String())
}

func (l *Mail) deleteReadMails(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	ids := make([]int64, 0)
	idIn := slogger.ReadLine(c, "mail id: ")
	mid, err := strconv.Atoi(idIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	if mid > 0 {
		ids = append(ids, int64(mid))
	}

	req := &bff.C2SDeleteReadMail{
		Ids: ids,
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_DeleteReadMail, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "delete mails send success: %v \n", req.String())
}

func (l *Mail) watchMail(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SWatchMail{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_WatchMail, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "watch mail send success: %v \n", req.String())
}

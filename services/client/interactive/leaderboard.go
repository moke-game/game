package interactive

import (
	"fmt"
	"net"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/pkg/errors"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Leaderboard struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewLeaderboard(conn net.Conn) *Leaderboard {
	m := &Leaderboard{
		conn: conn,
	}
	m.initShell()
	return m
}

func (m *Leaderboard) initShell() {
	m.cmd = &ishell.Cmd{
		Name:    "leaderboard",
		Help:    "leaderboard interactive shell",
		Aliases: []string{"lb"},
	}
	m.initSubCmd()
}
func (m *Leaderboard) GetCmd() *ishell.Cmd {
	return m.cmd
}

func (m *Leaderboard) initSubCmd() {
	m.cmd.AddCmd(&ishell.Cmd{
		Name:    "get",
		Help:    "get leaderboard",
		Aliases: []string{"g"},
		Func:    m.getLeaderboard,
	})

	m.cmd.AddCmd(&ishell.Cmd{
		Name:    "add",
		Help:    "add members",
		Aliases: []string{"add"},
		Func:    m.addMembers,
	})

	m.cmd.AddCmd(&ishell.Cmd{
		Name:    "update",
		Help:    "update self score",
		Aliases: []string{"u"},
		Func:    m.updateSelfScore,
	})

	m.cmd.AddCmd(&ishell.Cmd{
		Name:    "clear",
		Help:    "clear leaderboard",
		Aliases: []string{"c"},
		Func:    m.clearLeaderboard,
	})

	m.cmd.AddCmd(&ishell.Cmd{
		Name:    "star",
		Help:    "star leaderboard",
		Aliases: []string{"star"},
		Func:    m.starLeaderboard,
	})

}

func (m *Leaderboard) getLeaderboard(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	id := 1
	idIn := slogger.ReadLine(c, "leaderboard id (default:1): ")
	if idIn != "" {
		id, _ = strconv.Atoi(idIn)
	}
	req := &bff.C2SGetLeaderboard{
		Id: int32(id),
	}
	err := common.SendMsg(m.conn, cpb.C2S_EVENT_C2S_GetLeaderboard, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get leaderboard send success: %v \n", req.String())
}

func (m *Leaderboard) addMembers(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	id := 1
	idIn := slogger.ReadLine(c, "leaderboard id (default:1): ")
	if idIn != "" {
		id, _ = strconv.Atoi(idIn)
	}

	counts := 100
	countsIn := slogger.ReadLine(c, "leaderboard counts (default:100): ")
	if countsIn != "" {
		counts, _ = strconv.Atoi(countsIn)
	}

	cmd := fmt.Sprintf("addLeaderboard %d %d", id, counts)
	if err := common.SendMsg(m.conn, cpb.C2S_EVENT_C2S_BffGMCommand, &bff.C2SBffGMCommand{
		Command: cmd,
	}); err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "add leaderboard send success: %s \n", cmd)
}

func (m *Leaderboard) updateSelfScore(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	id := 1
	idIn := slogger.ReadLine(c, "leaderboard id (default:1): ")
	if idIn != "" {
		id, _ = strconv.Atoi(idIn)
	}

	uid := slogger.ReadLine(c, "uid: ")
	score := slogger.ReadLine(c, "score: ")
	updateTypeIn := slogger.ReadLine(c, "updateType (default:ADD,1:GT,2:LT): ")
	updateType := 0
	if updateTypeIn != "" {
		updateType, _ = strconv.Atoi(updateTypeIn)
	}
	cmd := fmt.Sprintf("updateScore %d %s %d %s", id, uid, updateType, score)
	if err := common.SendMsg(m.conn, cpb.C2S_EVENT_C2S_BffGMCommand, &bff.C2SBffGMCommand{
		Command: cmd,
	}); err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "update score send success: %s \n", cmd)
}

func (m *Leaderboard) clearLeaderboard(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	id := "1"
	idIn := slogger.ReadLine(c, "leaderboard id (default:1): ")
	if idIn != "" {
		id = idIn
	}

	cmd := fmt.Sprintf("clearLeaderboard %s", id)
	if err := common.SendMsg(m.conn, cpb.C2S_EVENT_C2S_BffGMCommand, &bff.C2SBffGMCommand{
		Command: cmd,
	}); err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "clear leaderboard send success: %s \n", cmd)
}

func (m *Leaderboard) starLeaderboard(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	id := 1
	idIn := slogger.ReadLine(c, "leaderboard id (default:1): ")
	if idIn != "" {
		id, _ = strconv.Atoi(idIn)
	}
	uid := slogger.ReadLine(c, "uid: ")
	if uid == "" {
		slogger.Warn(c, errors.New("uid is empty"))
		return
	}
	uid64, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		slogger.Warn(c, err)
		return
	}

	if err := common.SendMsg(m.conn, cpb.C2S_EVENT_C2S_StarLeaderboard, &bff.C2SStarLeaderboard{
		Id:  int32(id),
		Uid: uid64,
	}); err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "star leaderboard send success \n")
}

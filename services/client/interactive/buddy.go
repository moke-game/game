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

type Buddy struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewBuddy(conn net.Conn) *Buddy {
	b := &Buddy{
		conn: conn,
	}
	b.initShell()
	return b
}

func (b *Buddy) initShell() {
	b.cmd = &ishell.Cmd{
		Name:    "buddy",
		Help:    "buddy interactive shell",
		Aliases: []string{"bdy"},
	}
	b.initSubCmd()
}
func (b *Buddy) GetCmd() *ishell.Cmd {
	return b.cmd
}

func (b *Buddy) initSubCmd() {
	b.cmd.AddCmd(&ishell.Cmd{
		Name:    "friends",
		Help:    "get friend",
		Aliases: []string{"f"},
		Func:    b.getFriends,
	})
	b.cmd.AddCmd(&ishell.Cmd{
		Name:    "apply",
		Help:    "apply friend",
		Aliases: []string{"ap"},
		Func:    b.addFriend,
	})
	b.cmd.AddCmd(&ishell.Cmd{
		Name:    "agree",
		Help:    "agree friend apply",
		Aliases: []string{"ag"},
		Func:    b.agreeFriend,
	})
	b.cmd.AddCmd(&ishell.Cmd{
		Name:    "refuse",
		Help:    "refuse friend apply",
		Aliases: []string{"re"},
		Func:    b.refuseFriend,
	})
	b.cmd.AddCmd(&ishell.Cmd{
		Name:    "delete",
		Help:    "delete friend",
		Aliases: []string{"del"},
		Func:    b.deleteFriend,
	})
	b.cmd.AddCmd(&ishell.Cmd{
		Name:    "remark",
		Help:    "remark friend nickname",
		Aliases: []string{"rk"},
		Func:    b.remarkFriend,
	})
	b.cmd.AddCmd(&ishell.Cmd{
		Name:    "delBlack",
		Help:    "delete black list",
		Aliases: []string{"delB"},
		Func:    b.deleteBlack,
	})
}

func (b *Buddy) getFriends(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SFriendGet{}
	err := common.SendMsg(b.conn, cpb.C2S_EVENT_C2S_FriendGet, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get friends send success: %v \n", req.String())
}

func (b *Buddy) addFriend(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SFriendAdd{}
	addUid := slogger.ReadLine(c, "uid: ")
	uid, _ := strconv.ParseInt(addUid, 10, 64)
	req.AddUid = uid
	err := common.SendMsg(b.conn, cpb.C2S_EVENT_C2S_FriendAdd, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "add friend send success: %v \n", req.String())
}

func (b *Buddy) agreeFriend(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	agreeUid := slogger.ReadLine(c, "uid: ")
	uid, _ := strconv.ParseInt(agreeUid, 10, 64)
	req := &bff.C2SFriendAgree{}
	req.Uid = uid
	err := common.SendMsg(b.conn, cpb.C2S_EVENT_C2S_FriendAgree, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "agree friend send success: %v \n", req.String())
}

func (b *Buddy) refuseFriend(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	agreeUid := slogger.ReadLine(c, "uid: ")
	addBlack := slogger.ReadLine(c, "addBlack: ")
	uid, _ := strconv.ParseInt(agreeUid, 10, 64)
	isBlack, _ := strconv.ParseBool(addBlack)
	req := &bff.C2SFriendRefuse{}
	req.Uids = append(req.Uids, uid)
	req.IsBlacklist = isBlack

	err := common.SendMsg(b.conn, cpb.C2S_EVENT_C2S_FriendRefuse, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "refuse friend send success: %v \n", req.String())
}

func (b *Buddy) deleteFriend(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	delUid := slogger.ReadLine(c, "uid: ")
	addBlack := slogger.ReadLine(c, "addBlack: ")
	uid, _ := strconv.ParseInt(delUid, 10, 64)
	isBlack, _ := strconv.ParseBool(addBlack)
	req := &bff.C2SFriendDelete{}
	req.DelUid = uid
	req.IsBlacklist = isBlack
	err := common.SendMsg(b.conn, cpb.C2S_EVENT_C2S_FriendDelete, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "delete delete send success: %v \n", req.String())
}

func (b *Buddy) remarkFriend(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	fUid := slogger.ReadLine(c, "uid: ")
	remark := slogger.ReadLine(c, "remark: ")
	uid, _ := strconv.ParseInt(fUid, 10, 64)
	req := &bff.C2SFriendNickName{}
	req.Uid = uid
	req.NickName = remark

	err := common.SendMsg(b.conn, cpb.C2S_EVENT_C2S_FriendNickName, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "remark friend send success: %v \n", req.String())
}

func (b *Buddy) deleteBlack(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	fUid := slogger.ReadLine(c, "uid: ")
	uid, _ := strconv.ParseInt(fUid, 10, 64)
	req := &bff.C2SFriendDeleteBlack{}
	req.Uid = uid
	err := common.SendMsg(b.conn, cpb.C2S_EVENT_C2S_FriendDeleteBlack, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "delete black send success: %v \n", req.String())
}

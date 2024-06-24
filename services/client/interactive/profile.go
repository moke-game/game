package interactive

import (
	"fmt"
	"net"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/spf13/cast"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Profile struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewProfile(conn net.Conn) *Profile {
	l := &Profile{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Profile) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "profile",
		Help:    "profile interactive shell",
		Aliases: []string{"p"},
	}
	l.initSubCmd()
}
func (l *Profile) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Profile) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "new",
		Help:    "new profile",
		Aliases: []string{"n"},
		Func:    l.newProfile,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "get",
		Help:    "get profile",
		Aliases: []string{"g"},
		Func:    l.getProfile,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "ch",
		Help:    "choose hero",
		Aliases: []string{"n"},
		Func:    l.chooseHero,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "playerRoom",
		Help:    "get player room info",
		Aliases: []string{"pr"},
		Func:    l.getPlayerRoom,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "guide",
		Help:    "guide step",
		Aliases: []string{"gd"},
		Func:    l.guideStep,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "cup",
		Help:    "add cup",
		Aliases: []string{"cup"},
		Func:    l.addCup,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "feat",
		Help:    "unlock all feat",
		Aliases: []string{"fe"},
		Func:    l.unlockFeats,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "hero",
		Help:    "unlock all hero",
		Aliases: []string{"h"},
		Func:    l.unlockHeros,
	})
}

func (l *Profile) guideStep(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SNewPlayerGuide{
		Step: cast.ToInt32(slogger.ReadLine(c, "step: ")),
	}
	if err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_NewPlayerGuide, req); err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Profile) newProfile(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	nickname := slogger.ReadLine(c, "nickname: ")
	heroIdIn := slogger.ReadLine(c, "heroId: ")
	heroId, err := strconv.Atoi(heroIdIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := &bff.C2SNewPlayer{
		Name:   nickname,
		HeroId: int32(heroId),
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_NewPlayer, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "new profile send success: %v \n", req.String())
}

func (l *Profile) getProfile(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	uidIn := slogger.ReadLine(c, "uid: ")
	heroId, err := strconv.Atoi(uidIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := &bff.C2SSimpleInfo{
		Uid: int64(heroId),
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_SimpleInfo, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get profile send success: %v \n", req.String())
}

func (l *Profile) chooseHero(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	uidIn := slogger.ReadLine(c, "heroId: ")
	heroId, err := strconv.Atoi(uidIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := &bff.C2SChooseHero{
		HeroId: int32(heroId),
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_ChooseHero, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get profile send success: %v \n", req.String())
}

func (l *Profile) addCup(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	uid := slogger.ReadLine(c, "uid:")
	heroId := slogger.ReadLine(c, "heroId (default:10002):")
	if heroId == "" {
		heroId = "10002"
	}
	cup := slogger.ReadLine(c, "cup (default:100):")
	if cup == "" {
		cup = "100"
	}

	cmd := fmt.Sprintf("addHeroCup %s %s %s", uid, heroId, cup)
	if err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_BffGMCommand, &bff.C2SBffGMCommand{
		Command: cmd,
	}); err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "add cup send success: %s \n", cmd)
}

func (l *Profile) unlockFeats(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	uid := slogger.ReadLine(c, "uid:")
	cmd := fmt.Sprintf("feat %s", uid)
	if err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_BffGMCommand, &bff.C2SBffGMCommand{
		Command: cmd,
	}); err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "unlock feats send success: %s \n", cmd)
}

func (l *Profile) unlockHeros(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	uid := slogger.ReadLine(c, "uid:")
	cmd := fmt.Sprintf("hero %s", uid)
	if err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_BffGMCommand, &bff.C2SBffGMCommand{
		Command: cmd,
	}); err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "unlock heros send success: %s \n", cmd)
}

func (l *Profile) getPlayerRoom(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	uid := slogger.ReadLine(c, "uid:")

	req := &bff.C2SGetPlayerRoomInfo{
		Uid: int64(cast.ToInt(uid)),
	}
	if err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_GetPlayerRoomInfo, req); err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "get player room send success: %v \n", req.String())
}

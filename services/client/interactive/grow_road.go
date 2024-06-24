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

type GrowRoad struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewGrowRoad(conn net.Conn) *GrowRoad {
	gr := &GrowRoad{
		conn: conn,
	}
	gr.initShell()
	return gr
}

func (gr *GrowRoad) initShell() {
	gr.cmd = &ishell.Cmd{
		Name:    "growRoad",
		Help:    "GrowRoad interactive shell",
		Aliases: []string{"gr"},
	}
	gr.initSubCmd()
}
func (gr *GrowRoad) GetCmd() *ishell.Cmd {
	return gr.cmd
}

func (gr *GrowRoad) initSubCmd() {
	gr.cmd.AddCmd(&ishell.Cmd{
		Name:    "info",
		Help:    "get grow road info",
		Aliases: []string{"i"},
		Func:    gr.getGrowRoadInfo,
	})

	gr.cmd.AddCmd(&ishell.Cmd{
		Name:    "reward",
		Help:    "get grow road reward",
		Aliases: []string{"r"},
		Func:    gr.getGrowRoadReward,
	})

}

func (gr *GrowRoad) getGrowRoadInfo(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	req := &bff.C2SGetCupInfo{}
	err := common.SendMsg(gr.conn, cpb.C2S_EVENT_C2S_GetCupInfo, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get grow road info send success: %v \n", req.String())
}

func (gr *GrowRoad) getGrowRoadReward(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	idIn := slogger.ReadLine(c, "input reward id: ")

	id, err := strconv.ParseInt(idIn, 10, 32)
	if err != nil {
		slogger.Warn(c, err)
		return
	}

	req := &bff.C2SGetCupReward{
		Id: int32(id),
	}
	err = common.SendMsg(gr.conn, cpb.C2S_EVENT_C2S_GetCupReward, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "get grow road reward send success: %v \n", req.String())
}

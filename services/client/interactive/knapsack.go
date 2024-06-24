package interactive

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Knapsack struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewKnapsack(conn net.Conn) *Knapsack {
	l := &Knapsack{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Knapsack) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "knapsack",
		Help:    "knapsack interactive shell",
		Aliases: []string{"k"},
	}
	l.initSubCmd()
}
func (l *Knapsack) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Knapsack) initSubCmd() {
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "watch",
		Help:    "watch knapsack",
		Aliases: []string{"w"},
		Func:    l.watchKnapsack,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "exchange",
		Help:    "exchange items",
		Aliases: []string{"e"},
		Func:    l.exchangeItem,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "add",
		Help:    "add items",
		Aliases: []string{"a"},
		Func:    l.addItem,
	})
}
func (l *Knapsack) watchKnapsack(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SWatchingKnapsack{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_WatchingKnapsack, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "watch knapsack send success: %v \n", req.String())
}

func (l *Knapsack) exchangeItem(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	itemIdIn := slogger.ReadLine(c, "item id: ")
	itemIds := strings.Split(itemIdIn, ",")
	items := make(map[int64]*bff.Item)
	for _, itemId := range itemIds {
		itemId, err := strconv.Atoi(itemId)
		if err != nil {
			slogger.Warn(c, err)
			return
		}
		numIn := slogger.ReadLine(c, "num: ")
		num, err := strconv.Atoi(numIn)
		if err != nil {
			slogger.Warn(c, err)
			return
		}
		items[int64(itemId)] = &bff.Item{
			Id:       int64(itemId),
			ConfigId: int32(itemId),
			Num:      int32(num),
		}
	}

	req := &bff.C2SDiamondExchangeItem{
		Item: &bff.Items{
			Items: items,
		},
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_DiamondExchangeItem, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "exchange item send success: %v \n", req.String())
}

func (l *Knapsack) addItem(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	uid := slogger.ReadLine(c, "uid: ")
	itemIdIn := slogger.ReadLine(c, "item id: ")
	itemId, err := strconv.Atoi(itemIdIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	numIn := slogger.ReadLine(c, "num: ")
	num, err := strconv.Atoi(numIn)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	cmd := fmt.Sprintf("addCurrency %s %d %d", uid, itemId, num)
	req := &bff.C2SBffGMCommand{
		Command: cmd,
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_BffGMCommand, req)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	slogger.Infof(c, "add item send success: %v \n", req.String())
}

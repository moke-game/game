package interactive

import (
	"net"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/spf13/cast"

	"github.com/gstones/moke-kit/logging/slogger"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/client/common"
)

type Shop struct {
	cmd  *ishell.Cmd
	conn net.Conn
}

func NewShop(conn net.Conn) *Shop {
	l := &Shop{
		conn: conn,
	}
	l.initShell()
	return l
}

func (l *Shop) initShell() {
	l.cmd = &ishell.Cmd{
		Name:    "shop",
		Help:    "shop interactive shell",
		Aliases: []string{"s"},
	}
	l.initSubCmd()
}
func (l *Shop) GetCmd() *ishell.Cmd {
	return l.cmd
}

func (l *Shop) initSubCmd() {

	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "check",
		Help:    "check items is can purchase",
		Aliases: []string{"c"},
		Func:    l.check,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "pay",
		Help:    "pay items",
		Aliases: []string{"p"},
		Func:    l.pay,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "buy",
		Help:    "shop buy items",
		Aliases: []string{"b"},
		Func:    l.buy,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "products",
		Help:    "shop products",
		Aliases: []string{"d"},
		Func:    l.products,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "firsts",
		Help:    "first pay infos",
		Aliases: []string{"f"},
		Func:    l.firsts,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name:    "firstday",
		Help:    "first pay day",
		Aliases: []string{"n"},
		Func:    l.firstday,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "day7infos",
		Help: "day7 act infos",
		Func: l.day7infos,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "day7coin",
		Help: "day7 act coin reward",
		Func: l.day7coin,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "day7buy",
		Help: "day7 act buy",
		Func: l.day7buy,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "mcards",
		Help: "month card list",
		Func: l.mcards,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "mcarddone",
		Help: "month card done",
		Func: l.mcarddone,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "trigger",
		Help: "trigger act",
		Func: l.trigger,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "cdkey",
		Help: "cdkey act",
		Func: l.cdkey,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "signinfo",
		Help: "signinfo act",
		Func: l.signActInfo,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "signbuy",
		Help: "sign buy day",
		Func: l.signBuyDay,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "signdone",
		Help: "sign day done",
		Func: l.signDayDone,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "unlockSkin",
		Help: "unlockSkin",
		Func: l.unlockSkin,
	})
	l.cmd.AddCmd(&ishell.Cmd{
		Name: "selectSkin",
		Help: "selectSkin",
		Func: l.selectSkin,
	})
}

func (l *Shop) selectSkin(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SSelectSkin{
		SkinId: cast.ToInt32(slogger.ReadLine(c, "skinId: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_SelectSkin, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) unlockSkin(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SUnlockSkin{
		SkinId: cast.ToInt32(slogger.ReadLine(c, "skinId: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_UnlockSkin, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) signDayDone(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SSignDayDone{
		Day:   cast.ToInt32(slogger.ReadLine(c, "day: ")),
		ActId: 10001,
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_SignDayDone, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) signBuyDay(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SSignBuyDay{
		Day:   cast.ToInt32(slogger.ReadLine(c, "day: ")),
		ActId: 10001,
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_SignBuyDay, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) signActInfo(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SSignActInfo{
		ActId: 10001,
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_SignActInfo, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) cdkey(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SCDKeyReward{
		CdKey: slogger.ReadLine(c, "cdkey: "),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_CDKeyReward, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) trigger(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SWatchTriggerAct{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_WatchTriggerAct, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) mcarddone(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SMonthCardActDone{
		ActId: cast.ToInt32(slogger.ReadLine(c, "act_id: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_MonthCardActDone, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) mcards(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SMonthCardActInfo{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_MonthCardActInfo, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) day7infos(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SDay7ActInfos{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_Day7ActInfos, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) day7coin(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SDay7ActReward{
		Lv: cast.ToInt32(slogger.ReadLine(c, "lv: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_Day7ActReward, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) day7buy(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SDay7ActBuy{
		Day:       cast.ToInt32(slogger.ReadLine(c, "day: ")),
		ItemBoxId: cast.ToInt32(slogger.ReadLine(c, "itembox id: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_Day7ActBuy, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "send success: %v \n", req.String())
}

func (l *Shop) check(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SCanPurchase{
		PurchaseId: cast.ToInt32(slogger.ReadLine(c, "purchase_id: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_CanPurchase, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "check items is can purchase send success: %v \n", req.String())
}

func (l *Shop) products(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SShopProducts{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_ShopProducts, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "shop products send success: %v \n", req.String())
}

func (l *Shop) buy(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SShopBuy{
		ShopId: cast.ToInt32(slogger.ReadLine(c, "shop_id: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_ShopBuy, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "shop buy send success: %v \n", req.String())
}

func (l *Shop) pay(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	transId := slogger.ReadLine(c, "transaction_id: ")
	platform := slogger.ReadLine(c, "platform(1:appstore,2:googlePlay): ")
	receipt := slogger.ReadLine(c, "receipt: ")
	pId, err := strconv.ParseInt(platform, 10, 64)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := &bff.C2SPurchase{
		TransactionId: transId,
		Platform:      bff.Platform(pId),
		PurchaseId:    cast.ToInt32(slogger.ReadLine(c, "product_id(int): ")),
		Receipt:       receipt,
	}
	err = common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_Purchase, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "purchase items send success: %v \n", req.String())
}

func (l *Shop) firsts(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SFirstPayActInfo{}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_FirstPayActInfo, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "first pay infos send success: %v \n", req.String())
}
func (l *Shop) firstday(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	req := &bff.C2SFirstPayActDone{
		ActId: cast.ToInt32(slogger.ReadLine(c, "ActId: ")),
		Day:   cast.ToInt32(slogger.ReadLine(c, "Day: ")),
	}
	err := common.SendMsg(l.conn, cpb.C2S_EVENT_C2S_FirstPayActDone, req)
	if err != nil {
		slogger.Warn(c, err)
	}
	slogger.Infof(c, "shop products send success: %v \n", req.String())
}

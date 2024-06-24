package interactive

import (
	"net"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/zinx/ziface"
	"google.golang.org/protobuf/proto"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	roompb "github.com/moke-game/game/api/gen/room"
	"github.com/moke-game/game/services/client/common"
)

type Client struct {
	conn  net.Conn
	shell *ishell.Shell
}

func (c *Client) Init(url string) error {

	c.shell = ishell.New()
	if conn, err := net.Dial("tcp", url); err != nil {
		c.shell.Println("client start err, exit!", err)
		return err
	} else {
		c.conn = conn
	}
	c.initShells(c.shell)
	c.shell.Println("game client interactive shell ready")
	return nil

	c.shell = ishell.New()
	if conn, err := net.Dial("tcp", url); err != nil {
		c.shell.Println("client start err, exit!", err)
		return err
	} else {
		c.conn = conn
	}
	c.initShells(c.shell)
	c.shell.Println("game client interactive shell ready")
	//ca := x509.NewCertPool()
	//caBytes, err := os.ReadFile("../../configs/tls-server/ca.crt") // "../../configs/x509/ca_cert.pem
	//if err != nil {
	//	c.shell.Println("read ca cert err, exit!", err)
	//	return err
	//}
	////c.shell.Printf("read ca cert ok:%s ", caBytes)
	//if ok := ca.AppendCertsFromPEM(caBytes); !ok {
	//	c.shell.Println("append cert err, exit!")
	//	return nil
	//}
	//if conn, err := tls.Dial("tcp", url, &tls.Config{
	//	//InsecureSkipVerify: true,
	//	RootCAs: ca,
	//}); err != nil {
	//	c.shell.Println("client start err, exit!", err)
	//	return err
	//} else {
	//	c.conn = conn
	//}
	c.initShells(c.shell)
	c.shell.Println("game client interactive shell ready")
	return nil
}

func (c *Client) Run() {
	c.shell.Interrupt(func(cxt *ishell.Context, count int, input string) {
		if count >= 2 {
			cxt.Stop()
		}
		if count == 1 {
			c.Close()
			slogger.Info(cxt, "interrupted, press again to exit")
		}
	})
	c.heartbeat()
	c.watchResponse()
	c.shell.Run()
}

func (c *Client) heartbeat() {
	go func() {
		for {
			if err := common.SendMsg(c.conn, cpb.C2S_EVENT_C2S_Heartbeat, &bff.C2SHeartbeat{}); err != nil {
				c.shell.Println("send heartbeat err, exit!", err)
				c.Close()
				return
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func (c *Client) watchResponse() {
	go func() {
		for {
			resp, err := common.RecvMsg(c.conn)
			if err != nil {
				c.shell.Println("receive msg err, exit!", err)
				c.Close()
				return
			}
			c.handleResponse(resp)
		}
	}()
}

func (c *Client) handleResponse(msg ziface.IMessage) {
	data := msg.GetData()
	if msg.GetMsgID() > 2000 {
		resp := &cpb.Response{}
		if err := proto.Unmarshal(msg.GetData(), resp); err != nil {
			c.shell.Printf("unmarshal msg:%d err:%v , exit! \n", msg.GetMsgID(), err)
			return
		}
		if resp.Code != 0 {
			c.shell.Printf("receive msg:%d, errcode:%s \n", msg.GetMsgID(), resp.Code)
			return
		}
		data = resp.Data
	}

	switch msg.GetMsgID() {
	case uint32(cpb.S2C_EVENT_S2C_Auth):
		c.printResp(data, &bff.S2CAuth{})
	case uint32(cpb.S2C_EVENT_S2C_NewPlayer):
		c.printResp(data, &bff.S2CNewPlayer{})
	case uint32(cpb.S2C_EVENT_S2C_SimpleInfo):
		c.printResp(data, &bff.S2CSimpleInfo{})
	case uint32(cpb.S2C_EVENT_S2C_EnterRoom):
		c.printResp(data, &roompb.S2CEnterRoom{})
	case uint32(cpb.S2C_EVENT_S2C_WatchingKnapsack):
		c.printResp(data, &bff.S2CWatchingKnapsack{})
	case uint32(cpb.S2C_EVENT_S2C_DiamondExchangeItem):
		c.printResp(data, &bff.S2CDiamondExchangeItem{})
	case uint32(cpb.S2C_EVENT_S2C_CHATReceiveWorldMessage):
		c.printResp(data, &bff.S2CCHATReceiveWorldMessage{})
	case uint32(cpb.S2C_EVENT_S2C_CHATReceivePlayerMessage):
		c.printResp(data, &bff.S2CCHATReceivePlayerMessage{})
	case uint32(cpb.S2C_EVENT_S2C_CHATReceiveTeamMessage):
		c.printResp(data, &bff.S2CCHATReceiveTeamMessage{})
	case uint32(cpb.S2C_EVENT_S2C_JoinParty):
		c.printResp(data, &bff.S2CJoinParty{})
	case uint32(cpb.S2C_EVENT_S2C_LeaveParty):
		c.printResp(data, &bff.S2CLeaveParty{})
	case uint32(cpb.S2C_EVENT_S2C_KickParty):
		c.printResp(data, &bff.S2CKickParty{})
	case uint32(cpb.S2C_EVENT_S2C_READY_PARTY):
		c.printResp(data, &bff.S2CReadyParty{})
	case uint32(cpb.S2C_EVENT_S2C_CANCEL_READY_PARTY):
		c.printResp(data, &bff.S2CCancelReadyParty{})
	case uint32(cpb.S2C_EVENT_NTF_PartyInfo):
		c.printResp(data, &bff.NtfPartyInfo{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberStatus):
		c.printResp(data, &bff.NtfPartyMemberStatus{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberHero):
		c.printResp(data, &bff.NtfPartyMemberHero{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberJoin):
		c.printResp(data, &bff.NtfPartyMemberJoin{})
	case uint32(cpb.S2C_EVENT_NTF_PartyMemberLeave):
		c.printResp(data, &bff.NtfPartyMemberLeave{})
	case uint32(cpb.S2C_EVENT_NTF_RoomSkills):
		c.printResp(data, &roompb.NtfRoomSkills{})
	case uint32(cpb.S2C_EVENT_NTF_RoomHits):
		c.printResp(data, &roompb.NtfRoomHits{})
	case uint32(cpb.S2C_EVENT_NTF_RoomBuffEffect):
		c.printResp(data, &roompb.NtfRoomBuffEffect{})
	case uint32(cpb.S2C_EVENT_NTF_RoomBuffKnockback):
		c.printResp(data, &roompb.NtfRoomBuffKnockback{})
	case uint32(cpb.S2C_EVENT_NTF_UnitEnterEyeshot):
		c.printResp(data, &roompb.NtfUnitEnterEyeshot{})
	case uint32(cpb.S2C_EVENT_NTF_PlayerEnterEyeshot):
		c.printResp(data, &roompb.NtfPlayerEnterEyeshot{})
	case uint32(cpb.S2C_EVENT_NTF_UnitLeaveEyeshot):
		c.printResp(data, &roompb.NtfUnitLeaveEyeshot{})
	case uint32(cpb.S2C_EVENT_NTF_UnitMove):
		c.printResp(data, &roompb.NtfUnitMove{})
	case uint32(cpb.S2C_EVENT_NTF_UnitStopMoving):
		c.printResp(data, &roompb.NtfUnitStopMoving{})
	case uint32(cpb.S2C_EVENT_S2C_GetRoomInfo):
		c.printResp(data, &bff.S2CGetRoomInfo{})
	case uint32(cpb.S2C_EVENT_NTF_MailChange):
		c.printResp(data, &bff.NtfMailChange{})
	case uint32(cpb.S2C_EVENT_S2C_ReadMail):
		c.printResp(data, &bff.S2CReadMail{})
	case uint32(cpb.S2C_EVENT_S2C_GetMailRewards):
		c.printResp(data, &bff.S2CGetMailRewards{})
	case uint32(cpb.S2C_EVENT_S2C_DeleteReadMail):
		c.printResp(data, &bff.S2CDeleteReadMail{})
	case uint32(cpb.S2C_EVENT_S2C_CanPurchase):
		c.printResp(data, &bff.S2CCanPurchase{})
	case uint32(cpb.S2C_EVENT_S2C_Purchase):
		c.printResp(data, &bff.S2CPurchase{})
	case uint32(cpb.S2C_EVENT_S2C_ShopBuy):
		c.printResp(data, &bff.S2CShopBuy{})
	case uint32(cpb.S2C_EVENT_S2C_ShopProducts):
		c.printResp(data, &bff.S2CShopProducts{})
	case uint32(cpb.S2C_EVENT_S2C_FirstPayActInfo):
		c.printResp(data, &bff.S2CFirstPayActInfo{})
	case uint32(cpb.S2C_EVENT_S2C_FirstPayActDone):
		c.printResp(data, &bff.S2CFirstPayActDone{})
	case uint32(cpb.S2C_EVENT_S2C_Day7ActBuy):
		c.printResp(data, &bff.S2CDay7ActBuy{})
	case uint32(cpb.S2C_EVENT_S2C_Day7ActInfos):
		c.printResp(data, &bff.S2CDay7ActInfos{})
	case uint32(cpb.S2C_EVENT_S2C_Day7ActReward):
		c.printResp(data, &bff.S2CDay7ActReward{})
	case uint32(cpb.S2C_EVENT_S2C_MonthCardActInfo):
		c.printResp(data, &bff.S2CMonthCardActInfo{})
	case uint32(cpb.S2C_EVENT_S2C_MonthCardActDone):
		c.printResp(data, &bff.S2CMonthCardActDone{})
	case uint32(cpb.S2C_EVENT_S2C_WatchTriggerAct):
		c.printResp(data, &bff.S2CWatchTriggerAct{})
	case uint32(cpb.S2C_EVENT_S2C_NewPlayerGuide):
		c.printResp(data, &bff.S2CNewPlayerGuide{})
	case uint32(cpb.S2C_EVENT_S2C_CDKeyReward):
		c.printResp(data, &bff.S2CCDKeyReward{})
	case uint32(cpb.S2C_EVENT_S2C_GetLeaderboard):
		c.printResp(data, &bff.S2CGetLeaderboard{})
	case uint32(cpb.S2C_EVENT_S2C_BffGMCommand):
		c.printResp(data, &bff.S2CBffGMCommand{})
	case uint32(cpb.S2C_EVENT_S2C_SignActInfo):
		c.printResp(data, &bff.S2CSignActInfo{})
	case uint32(cpb.S2C_EVENT_S2C_SignBuyDay):
		c.printResp(data, &bff.S2CSignBuyDay{})
	case uint32(cpb.S2C_EVENT_S2C_SignDayDone):
		c.printResp(data, &bff.S2CSignDayDone{})
	case uint32(cpb.S2C_EVENT_S2C_GetPlayerRoomInfo):
		c.printResp(data, &bff.S2CGetPlayerRoomInfo{})
	default:
		if msg.GetMsgID() != uint32(cpb.S2C_EVENT_S2C_Heartbeat) {
			c.shell.Printf("receive msg:%d, data:%v \n", msg.GetMsgID(), msg.GetData())
		}
	}
}

func (c *Client) printResp(data []byte, pMsg proto.Message) {
	if err := proto.Unmarshal(data, pMsg); err != nil {
		c.shell.Printf("unmarshal err:%v , exit! \n", err)
	} else {
		c.shell.Printf("receive msg: %s \n", pMsg)
	}
}

func (c *Client) initShells(sh *ishell.Shell) {
	login := NewLogin(c.conn)
	sh.AddCmd(login.GetCmd())

	profile := NewProfile(c.conn)
	sh.AddCmd(profile.GetCmd())

	room := NewRoom(c.conn)
	sh.AddCmd(room.GetCmd())

	knapsack := NewKnapsack(c.conn)
	sh.AddCmd(knapsack.GetCmd())

	chat := NewChat(c.conn)
	sh.AddCmd(chat.GetCmd())

	party := NewParty(c.conn)
	sh.AddCmd(party.GetCmd())

	match := NewLeaderboard(c.conn)
	sh.AddCmd(match.GetCmd())

	buddy := NewBuddy(c.conn)
	sh.AddCmd(buddy.GetCmd())

	mail := NewMail(c.conn)
	sh.AddCmd(mail.GetCmd())

	shop := NewShop(c.conn)
	sh.AddCmd(shop.GetCmd())

	growRoad := NewGrowRoad(c.conn)
	sh.AddCmd(growRoad.GetCmd())

	leaderboard := NewLeaderboard(c.conn)
	sh.AddCmd(leaderboard.GetCmd())

	mission := NewMission(c.conn)
	sh.AddCmd(mission.GetCmd())
}

func (c *Client) Close() {
	_ = c.conn.Close()
	c.shell.Close()
}

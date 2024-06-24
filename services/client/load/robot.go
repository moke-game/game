package load

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/spf13/cobra"

	bff "github.com/moke-game/game/api/gen/bff"
	cpb "github.com/moke-game/game/api/gen/common"
	roompb "github.com/moke-game/game/api/gen/room"
	"github.com/moke-game/game/services/client/common"
)

type Robot struct {
	username string
	uid      int64
	token    string
	conn     net.Conn
	roomConn net.Conn
	shell    *cobra.Command
	pos      *roompb.Vector
	mapID    int32
	duration time.Duration
	closeSig <-chan struct{}
	speed    float64

	lastDir   int32
	moveStart *roompb.Vector
	moveDest  *roompb.Vector
}

func CreateRobot(
	cmd *cobra.Command,
	addr string,
	duration time.Duration,
	close <-chan struct{},
) (*Robot, error) {
	r := &Robot{
		shell:    cmd,
		duration: duration,
		closeSig: close,
	}
	if err := r.Init(addr); err != nil {
		return nil, err
	}
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	r.username = fmt.Sprintf("load_%s_%v", hostname, time.Now().Unix())
	return r, nil
}

func (r *Robot) Init(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	r.conn = conn
	go r.watchResponse(r.conn)
	return nil
}

func (r *Robot) Run() error {
	if err := r.login(); err != nil {
		return err
	} else if err := r.newProfile(); err != nil {
		return err
	} else if err := r.getRoomInfo(); err != nil {
		return err
	} else if err := r.watchKnapsack(); err != nil {
		return err
	} else if err := r.watchMail(); err != nil {
		return err
	}
	r.Update()
	return nil
}
func (r *Robot) isHall() bool {
	return r.mapID == 1001
}

func (r *Robot) Update() {
	ticker := time.NewTicker(10 * time.Second)
	timeout := time.NewTicker(1 * time.Minute)
	//moving := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-r.closeSig:
			return
		case <-ticker.C:
			if err := r.heartbeat(); err != nil {
				r.shell.PrintErrln(err)
			}
			if err := r.randomMove(); err != nil {
				r.shell.PrintErrln(err)
			}
		case <-timeout.C:
			if r.isHall() {
				if err := r.matchmaking(); err != nil {
					r.shell.PrintErrln(err)
				}
			} else {
				if err := r.getRoomInfo(); err != nil {
					r.shell.PrintErrln(err)
				}
			}
			//case <-moving.C:
			//if err := r.moving(); err != nil {
			//	return
			//}
		}
	}
}

func (r *Robot) OnStop(pos *roompb.Vector) {
	r.pos = pos
	r.lastDir = math.MinInt32
	r.moveDest = nil
	r.moveStart = nil
}

func (r *Robot) Destroy() {
	if r == nil {
		return
	}
	if r.shell == nil {
		return
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
	if r.roomConn != nil {
		_ = r.roomConn.Close()
	}
}

func (r *Robot) login() error {
	req := &bff.C2SAuth{
		Token:  r.username,
		Openid: r.username,
	}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_Auth, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) newProfile() error {
	heroIds := []int32{10001, 10002, 10003, 10004, 10005, 10008, 10010, 10012, 10014, 10015}
	hid := random.RandInt(0, len(heroIds))
	req := &bff.C2SNewPlayer{
		Name:   r.username,
		HeroId: heroIds[hid],
	}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_NewPlayer, req)
	if err != nil {
		return err
	}
	return nil
}
func (r *Robot) chooseHero() error {
	heros := []int32{10001, 10002, 10003, 10004, 10005, 10008, 10010, 10012, 10014, 10015}
	req := &bff.C2SChooseHero{
		HeroId: heros[rand.Intn(len(heros))],
	}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_ChooseHero, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) watchKnapsack() error {
	req := &bff.C2SWatchingKnapsack{}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_WatchingKnapsack, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) watchMail() error {
	req := &bff.C2SWatchMail{}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_WatchMail, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) unlockFeats() error {
	cmd := fmt.Sprintf("feat %d", r.uid)
	req := &bff.C2SBffGMCommand{
		Command: cmd,
	}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_BffGMCommand, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) getMissions() error {
	req := &bff.C2SGetMission{}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_GetMission, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) unlockHero() error {
	cmd := fmt.Sprintf("hero %d", r.uid)
	req := &bff.C2SBffGMCommand{
		Command: cmd,
	}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_BffGMCommand, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) getRoomInfo() error {
	req := &bff.C2SGetRoomInfo{}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_GetRoomInfo, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) enterRoom(msg *bff.S2CGetRoomInfo) error {
	if r.roomConn != nil {
		_ = r.roomConn.Close()
	}
	if conn, err := net.Dial("tcp", msg.RoomHost); err != nil {
		return err
	} else {
		r.roomConn = conn
	}
	go r.watchResponse(r.roomConn)
	req := &roompb.C2SEnterRoom{
		Token:     r.token,
		RoomId:    msg.RoomId,
		RoomToken: msg.RoomToken,
	}
	err := common.SendMsg(r.roomConn, cpb.C2S_EVENT_C2S_EnterRoom, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) enterScene() error {
	req := &roompb.C2SEnterScene{}
	err := common.SendMsg(r.roomConn, cpb.C2S_EVENT_C2S_EnterScene, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) randomMove() error {
	if r.roomConn == nil {
		return nil
	}
	return common.SendMsg(r.roomConn, cpb.C2S_EVENT_C2S_RandomMove, &roompb.C2SRandomMove{
		Radius: 10,
	})
}

func (r *Robot) heartbeat() error {
	req := &bff.C2SHeartbeat{}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_Heartbeat, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) matchmaking() error {
	playIDs := []int32{1, 2, 3, 5}
	req := &bff.C2SMatchingSingleStart{
		PlayMap: playIDs[random.RandInt(0, len(playIDs))],
	}
	err := common.SendMsg(r.conn, cpb.C2S_EVENT_C2S_MatchingSingleStart, req)
	if err != nil {
		return err
	}
	return nil
}

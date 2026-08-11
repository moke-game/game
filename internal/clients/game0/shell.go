package game0

import (
	"net"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/tools"

	authclient "github.com/moke-game/platform/services/auth/client"
)

func RunGrpc(url string) {
	sh := ishell.New()
	slogger.Info(sh, "interactive game connect to "+url)
	slogger.Info(sh, "flow: auth token <id> → game token <access> → game hi")

	if conn, err := tools.DialInsecure(url); err != nil {
		slogger.Die(sh, err)
	} else {
		gameGrpc := NewDemoGrpcCli(conn)
		sh.AddCmd(gameGrpc.GetCmd())

		if authCmd, err := authclient.CreateAuthClient(url); err != nil {
			slogger.Warn(sh, err)
		} else {
			sh.AddCmd(authCmd)
		}

		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				conn.Close()
				slogger.Done(c, "interrupted, press again to exit")
			}
		})
	}
	sh.Run()
}

func RunTcp(url string) {
	sh := ishell.New()
	slogger.Info(sh, "interactive game tcp connect to "+url)
	if conn, err := net.Dial("tcp", url); err != nil {
		slogger.Die(sh, err)
	} else {
		gameTcp := NewTcpCli(conn)
		sh.AddCmd(gameTcp.GetCmd())

		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				conn.Close()
				slogger.Done(c, "interrupted, press again to exit")
			}
		})
	}
	sh.Run()
}

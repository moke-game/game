package game0

import (
	"net"

	"github.com/abiosoft/ishell"

	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/tools"
)

func RunGrpc(url string) {
	sh := ishell.New()
	slogger.Info(sh, "interactive game connect to "+url)

	if conn, err := tools.DialInsecure(url); err != nil {
		slogger.Die(sh, err)
	} else {
		gameGrpc := NewDemoGrpcCli(conn)
		sh.AddCmd(gameGrpc.GetCmd())

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

package game0

import (
	"net"
	"os"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/tools"

	authclient "github.com/moke-game/platform/services/auth/client"
)

// RunGrpcOptions configures the interactive gRPC shell.
type RunGrpcOptions struct {
	GameURL string
	AuthURL string // if empty, falls back to GameURL (aggregate) or AUTH_URL env
}

func RunGrpc(url string) {
	RunGrpcWithOptions(RunGrpcOptions{GameURL: url})
}

func RunGrpcWithOptions(opts RunGrpcOptions) {
	sh := ishell.New()
	gameURL := opts.GameURL
	authURL := opts.AuthURL
	if authURL == "" {
		authURL = os.Getenv("AUTH_URL")
	}
	if authURL == "" {
		authURL = gameURL
	}

	slogger.Info(sh, "interactive game connect to "+gameURL)
	slogger.Info(sh, "auth client connect to "+authURL)
	slogger.Info(sh, "flow: auth token → copy access → game token → game hi")

	if conn, err := tools.DialInsecure(gameURL); err != nil {
		slogger.Die(sh, err)
	} else {
		gameGrpc := NewDemoGrpcCli(conn)
		sh.AddCmd(gameGrpc.GetCmd())

		if authCmd, err := authclient.CreateAuthClient(authURL); err != nil {
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

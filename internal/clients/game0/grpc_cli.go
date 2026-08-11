package game0

import (
	"context"
	"fmt"
	"sync"

	pb "github.com/moke-game/game/api/gen/game0/api"

	"github.com/abiosoft/ishell"
	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/gstones/moke-kit/logging/slogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type DemoGrpc struct {
	client pb.Game0ServiceClient
	cmd    *ishell.Cmd

	mu    sync.RWMutex
	token string
}

func NewDemoGrpcCli(conn *grpc.ClientConn) *DemoGrpc {
	cmd := &ishell.Cmd{
		Name:    "game",
		Help:    "game interactive",
		Aliases: []string{"D"},
	}
	p := &DemoGrpc{
		client: pb.NewGame0ServiceClient(conn),
		cmd:    cmd,
	}
	p.initSubShells()
	return p
}

func (p *DemoGrpc) GetCmd() *ishell.Cmd {
	return p.cmd
}

// SetAccessToken stores the bearer token used by hi/watch.
func (p *DemoGrpc) SetAccessToken(token string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.token = token
}

func (p *DemoGrpc) accessToken() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.token
}

func (p *DemoGrpc) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "token",
		Help:    "set access token for subsequent RPCs",
		Aliases: []string{"t"},
		Func:    p.setToken,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "hi",
		Help:    "say hi",
		Aliases: []string{"hi"},
		Func:    p.sayHi,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "watch",
		Help:    "watch topic",
		Aliases: []string{"w"},
		Func:    p.watch,
	})
}

func (p *DemoGrpc) setToken(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	tok := slogger.ReadLine(c, "access token: ")
	p.SetAccessToken(tok)
	slogger.Infof(c, "token set (%d chars)", len(tok))
}

func (p *DemoGrpc) authCtx() context.Context {
	tok := p.accessToken()
	if tok == "" {
		tok = "test"
	}
	md := metadata.Pairs("authorization", fmt.Sprintf("%s %s", "bearer", tok))
	return mm.MD(md).ToOutgoing(context.Background())
}

func (p *DemoGrpc) sayHi(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	msg := "hello"
	in := slogger.ReadLine(c, "message(default:hello): ")
	if in != "" {
		msg = in
	}
	topic := "game"
	t := slogger.ReadLine(c, "topic(default:game): ")
	if t != "" {
		topic = t
	}
	if response, err := p.client.Hi(p.authCtx(), &pb.HiRequest{
		Uid:     "10000",
		Message: msg,
		Topic:   topic,
	}); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: %s", response.Message)
	}
}

func (p *DemoGrpc) watch(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	topic := "game"
	t := slogger.ReadLine(c, "topic(default:game): ")
	if t != "" {
		topic = t
	}

	if stream, err := p.client.Watch(p.authCtx(), &pb.WatchRequest{
		Topic: topic,
	}); err != nil {
		slogger.Warn(c, err)
	} else {
		for {
			if response, err := stream.Recv(); err != nil {
				slogger.Warn(c, err)
				break
			} else {
				slogger.Infof(c, "Response: %s \r\n", response.Message)
			}
		}
	}
}

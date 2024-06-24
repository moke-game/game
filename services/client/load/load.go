package load

import (
	"time"

	"github.com/gstones/zinx/zconf"
	"github.com/spf13/cobra"
	"go.uber.org/atomic"
)

type Loader struct {
	addr  string
	num   *atomic.Int32
	close chan struct{}
	cmd   *cobra.Command
}

func CreateLoader(addr string) *Loader {
	zconf.GlobalObject.MaxPacketSize = 8192
	return &Loader{
		addr:  addr,
		num:   atomic.NewInt32(0),
		close: make(chan struct{}),
	}
}

func (l *Loader) Shell(cmd *cobra.Command, num int32, duration int64) {
	l.cmd = cmd
	l.start(int(num), duration)
}

func (l *Loader) start(num int, duration int64) {
	hours := time.Duration(duration) * time.Hour
	for i := 1; i < num+1; i++ {
		l.cmd.Printf("progress:%d/%d \r\n", i, num)
		time.Sleep(1 * time.Second)
		go l.runRobot(l.addr, hours)
	}
	if duration > 0 {
		time.Sleep(hours)
		close(l.close)
	} else {
		<-l.close
	}
}

func (l *Loader) runRobot(addr string, duration time.Duration) {
	l.num.Inc()
	robot, err := CreateRobot(l.cmd, addr, duration, l.close)
	defer func() {
		l.cmd.PrintErrf("robot destroy,current:%d \n", l.num.Dec())
		robot.Destroy()
	}()
	if err != nil {
		l.cmd.PrintErrln(err)
		return
	}
	if err := robot.Run(); err != nil {
		l.cmd.PrintErrln(err)
		return
	}

}

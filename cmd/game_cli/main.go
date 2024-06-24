package main

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/moke-game/game/services/bff/pkg/module"
	"github.com/moke-game/game/services/client/load"
)

var options struct {
	game     string
	num      int32
	duration int64
}

/*
模拟客户端
*/
func main() {
	rootCmd := &cobra.Command{
		Use:   "game",
		Short: "Run a go-game CLI",
	}
	ctx := context.Background()
	rootCmd.PersistentFlags().StringVar(
		&options.game,
		"game",
		"127.0.0.1:8888",
		"game service (<host>:<port>)",
	)
	rootCmd.PersistentFlags().Int32Var(
		&options.num,
		"num",
		2000,
		"number of load robots",
	)
	rootCmd.PersistentFlags().Int64Var(
		&options.duration,
		"duration",
		-1,
		"duration for the load test",
	)

	{
		shell := &cobra.Command{
			Use:   "client",
			Short: "Run interactive game service client",
			Run: func(cmd *cobra.Command, args []string) {
				if cs, err := module.NewClientShell(options.game); err != nil {
					cmd.Println(err)
				} else {
					cs.Run()
				}
			},
		}

		loader := &cobra.Command{
			Use:   "load",
			Short: "Run a load test client",
			Run: func(cmd *cobra.Command, args []string) {
				load.CreateLoader(options.game).Shell(cmd, options.num, options.duration)
			},
		}
		rootCmd.AddCommand(shell, loader)
	}
	_ = rootCmd.ExecuteContext(ctx)
}

package main

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/moke-game/game/internal/clients/game0"
)

var options struct {
	host    string
	tcpHost string
}

const (
	defaultHost    = "localhost:8081"
	defaultTcpHost = "localhost:8888"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "game",
		Short:   "cli",
		Aliases: []string{"cli"},
	}
	rootCmd.PersistentFlags().StringVar(
		&options.host,
		"host",
		defaultHost,
		"grpc http service (<host>:<port>)",
	)

	rootCmd.PersistentFlags().StringVar(
		&options.tcpHost,
		"tcp_host",
		defaultTcpHost,
		"tcp service (<host>:<port>)",
	)

	sGrpc := &cobra.Command{
		Use:   "grpc",
		Short: "Run an interactive grpc client",
		Run: func(cmd *cobra.Command, args []string) {
			game0.RunGrpc(options.host)
		},
	}
	sTcp := &cobra.Command{
		Use:   "tcp",
		Short: "Run an interactive tcp client",
		Run: func(cmd *cobra.Command, args []string) {
			game0.RunTcp(options.tcpHost)
		},
	}
	rootCmd.AddCommand(sGrpc, sTcp)
	_ = rootCmd.ExecuteContext(context.Background())
}

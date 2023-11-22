package cmd_monitor

import (
	"bcd-util/cmd_monitor/client"
	"bcd-util/cmd_monitor/server"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "monitor",
		Short: "信息采集监控",
	}
	cmd.AddCommand(client.Cmd())
	cmd.AddCommand(server.Cmd())
	return &cmd
}

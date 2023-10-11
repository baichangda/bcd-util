package cmd_monitor

import (
	"bcd-util/cmd_monitor/client"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "monitor",
		Short: "信息采集监控",
	}
	cmd.AddCommand(client.Cmd())

	return &cmd
}

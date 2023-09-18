package monitor

import (
	"bcd-util/monitor/client"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "monitor",
		Short: "monitor",
	}
	cmd.AddCommand(client.Cmd())

	return &cmd
}

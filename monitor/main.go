package monitor

import (
	"github.com/spf13/cobra"
	"gmmc-tool/monitor/client"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "monitor",
		Short: "monitor",
	}
	cmd.AddCommand(client.Cmd())

	return &cmd
}

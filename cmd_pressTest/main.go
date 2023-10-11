package cmd_pressTest

import (
	"bcd-util/cmd_pressTest/gb32960"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pressTest",
		Short: "压力测试",
	}
	cmd.AddCommand(gb32960.Cmd())
	return &cmd
}

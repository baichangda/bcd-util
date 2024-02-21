package cmd_pressTest

import (
	"bcd-util/cmd_pressTest/gb32960"
	"bcd-util/cmd_pressTest/immotors_bin"
	"bcd-util/cmd_pressTest/immotors_json"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pressTest",
		Short: "压力测试",
	}
	cmd.AddCommand(gb32960.Cmd())
	cmd.AddCommand(immotors_bin.Cmd())
	cmd.AddCommand(immotors_json.Cmd())
	return &cmd
}

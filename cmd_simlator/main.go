package cmd_simlator

import (
	"bcd-util/cmd_simlator/gb32960"
	"bcd-util/cmd_simlator/immotors_bin"
	"bcd-util/cmd_simlator/immotors_json"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "simlator",
		Short: "模拟器",
	}

	cmd.AddCommand(gb32960.Cmd())
	cmd.AddCommand(immotors_bin.Cmd())
	cmd.AddCommand(immotors_json.Cmd())
	return &cmd
}

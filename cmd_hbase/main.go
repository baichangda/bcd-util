package cmd_hbase

import (
	"bcd-util/cmd_hbase/export"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "hbase",
		Short: "hbase工具",
	}
	cmd.AddCommand(export.Cmd())
	return &cmd
}

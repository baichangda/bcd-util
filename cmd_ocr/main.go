package cmd_ocr

import (
	"bcd-util/cmd_ocr/table"
	"bcd-util/cmd_ocr/web"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "ocr",
		Short: "信息采集监控",
	}
	cmd.AddCommand(web.Cmd())
	cmd.AddCommand(table.Cmd())
	return &cmd
}

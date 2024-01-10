package cmd_kafka

import (
	"bcd-util/cmd_kafka/producer"
	"bcd-util/cmd_kafka/web"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "kafka",
		Short: "kafka工具",
	}
	cmd.AddCommand(producer.Cmd())
	cmd.AddCommand(web.Cmd())
	return &cmd
}

package cmd_kafka

import (
	"bcd-util/cmd_kafka/producer"
	"bcd-util/cmd_kafka/prop"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "kafka",
		Short: "kafka工具",
	}
	cmd.PersistentFlags().StringSliceVarP(&prop.Addrs, "addrs", "a", []string{}, "kafka地址")
	_ = cmd.MarkPersistentFlagRequired("addrs")

	cmd.AddCommand(producer.Cmd())
	return &cmd
}

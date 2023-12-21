package producer

import (
	"bcd-util/cmd_kafka/producer/send_file"
	"bcd-util/cmd_kafka/prop"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "producer",
		Short: "生产者工具",
	}
	cmd.PersistentFlags().StringVarP(&prop.Topic, "topic", "t", "", "kafka topic")
	_ = cmd.MarkPersistentFlagRequired("topic")

	cmd.AddCommand(send_file.Cmd())
	return &cmd
}

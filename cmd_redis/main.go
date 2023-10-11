package cmd_redis

import (
	"bcd-util/cmd_redis/cluster"
	"bcd-util/cmd_redis/prop"
	"bcd-util/cmd_redis/single"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "redis",
		Short: "redis工具",
	}
	cmd.PersistentFlags().StringVarP(&prop.Password, "password", "p", "", "密码")
	_ = cmd.MarkPersistentFlagRequired("password")

	cmd.AddCommand(single.Cmd())
	cmd.AddCommand(cluster.Cmd())

	return &cmd
}

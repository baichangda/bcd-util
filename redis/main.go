package redis

import (
	"bcd-util/redis/cluster"
	"bcd-util/redis/prop"
	"bcd-util/redis/single"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "redis",
		Short: "redis",
	}
	cmd.PersistentFlags().StringVarP(&prop.Password, "password", "p", "", "密码")
	_ = cmd.MarkPersistentFlagRequired("password")

	cmd.AddCommand(single.Cmd())
	cmd.AddCommand(cluster.Cmd())

	return &cmd
}

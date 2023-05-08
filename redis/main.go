package redis

import (
	"github.com/spf13/cobra"
	"gmmc-tool/redis/cluster"
	"gmmc-tool/redis/prop"
	"gmmc-tool/redis/single"
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

package cluster

import (
	"github.com/spf13/cobra"
	"gmmc-tool/redis/cluster/dump"
	"gmmc-tool/redis/cluster/flush"
	"gmmc-tool/redis/cluster/restore"
	"gmmc-tool/redis/prop"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "cluster",
		Short: "集群模式",
	}
	cmd.PersistentFlags().StringSliceVarP(&prop.Addrs, "addr", "a", []string{}, "集群地址")
	_ = cmd.MarkPersistentFlagRequired("addr")

	cmd.AddCommand(dump.Cmd())
	cmd.AddCommand(restore.Cmd())
	cmd.AddCommand(flush.Cmd())

	return &cmd
}

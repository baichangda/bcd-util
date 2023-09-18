package cluster

import (
	"bcd-util/redis/cluster/dump"
	"bcd-util/redis/cluster/flush"
	"bcd-util/redis/cluster/restore"
	"bcd-util/redis/prop"
	"github.com/spf13/cobra"
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

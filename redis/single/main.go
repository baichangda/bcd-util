package single

import (
	"bcd-util/redis/prop"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "single",
		Short: "单机模式",
	}
	cmd.PersistentFlags().StringVarP(&prop.Addr, "addr", "a", "", "单机地址")
	_ = cmd.MarkPersistentFlagRequired("addr")

	return &cmd
}

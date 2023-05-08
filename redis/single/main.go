package single

import (
	"github.com/spf13/cobra"
	"gmmc-tool/redis/prop"
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

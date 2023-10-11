package flush

import (
	"bcd-util/cmd_redis/prop"
	"bcd-util/util"
	"bufio"
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "flush",
		Short: "清空数据库",
		Run: func(cmd *cobra.Command, args []string) {
			client := redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:        prop.Addrs,
				Password:     prop.Password,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			})

			util.Log.Infof("will flush cmd_redis[%s],enter [yes] to be continue...", prop.Addrs)

			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				text := scanner.Text()
				if text != "yes" {
					return
				}
			}

			ctx := context.Background()
			err := client.ForEachMaster(ctx, func(_ctx context.Context, _client *redis.Client) error {
				_, err := _client.FlushDB(_ctx).Result()
				if err != nil {
					return errors.WithStack(err)
				}
				return nil
			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
		},
	}
	return &cmd
}

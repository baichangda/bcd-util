package dump

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/hex"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gmmc-tool/redis/prop"
	"gmmc-tool/util"
	"os"
	"strings"
	"sync"
	"time"
)

var resGzFile string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "dump",
		Short: "备份redis成文件、使用dump命令、不会存储ttl信息",
		Run: func(cmd *cobra.Command, args []string) {
			index := strings.LastIndex(resGzFile, "/")
			var resDir string
			if index == 0 || index == -1 {
				resDir = ""
			} else {
				resDir = resGzFile[0:index]
			}

			if resDir != "" {
				err := os.MkdirAll(resDir, 0666)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
			}

			create, err := os.Create(resGzFile)
			defer create.Close()

			gzWriter := gzip.NewWriter(create)
			defer gzWriter.Close()

			writer := bufio.NewWriter(gzWriter)

			client := redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:        prop.Addrs,
				Password:     prop.Password,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			})
			ctx := context.Background()

			var nodeKeys [][]string
			err = client.ForEachMaster(ctx, func(_ctx context.Context, _client *redis.Client) error {
				var keys []string
				ks, err := _client.Keys(_ctx, "*").Result()
				if err != nil {
					return errors.WithStack(err)
				}
				keys = append(keys, ks...)
				nodeKeys = append(nodeKeys, keys)
				return nil
			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			ch := make(chan string, 10000)
			waitGroup := sync.WaitGroup{}
			for i, keys := range nodeKeys {
				util.Log.Infof("find node[%d] keys[%d]", i, len(keys))
				_keys := keys
				go func() {
					waitGroup.Add(1)
					for _, key := range _keys {
						result, _err := client.Dump(ctx, key).Bytes()
						if _err != nil {
							util.Log.Errorf("key[%s],%+v", key, _err)
							continue
						}
						ch <- key + " " + hex.EncodeToString(result) + "\n"
					}
					waitGroup.Done()
				}()
			}

			go func() {
				time.Sleep(1 * time.Second)
				waitGroup.Wait()
				close(ch)
			}()

			for e := range ch {
				_, err := writer.WriteString(e)
				if err != nil {
					util.Log.Errorf("%+v", err)
				}
			}
		},
	}
	cmd.Flags().StringVarP(&resGzFile, "resGzFile", "r", "", "结果gz文件路径")
	_ = cmd.MarkFlagRequired("resGzFile")

	return &cmd
}

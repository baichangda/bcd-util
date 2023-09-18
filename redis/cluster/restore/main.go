package restore

import (
	"bcd-util/redis/prop"
	"bcd-util/util"
	"bufio"
	"compress/gzip"
	"context"
	"encoding/hex"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
	"time"
)

var resGzFile string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "restore",
		Short: "恢复gz文件到redis中",
		Run: func(cmd *cobra.Command, args []string) {
			open, err := os.Open(resGzFile)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			defer open.Close()

			reader, err := gzip.NewReader(open)
			defer reader.Close()

			bufReader := bufio.NewReader(reader)

			client := redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:        prop.Addrs,
				Password:     prop.Password,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			})
			ctx := context.Background()

			count := 0

			lastLine := false
			for {
				line, err := bufReader.ReadString('\n')
				if err == nil {
					line = line[0 : len(line)-1]
				} else {
					if err == io.EOF {
						lastLine = true
					} else {
						util.Log.Errorf("%+v", err)
						return
					}
				}
				index := strings.Index(line, " ")
				key := line[0:index]
				valHex := line[index+1:]
				decodeString, err := hex.DecodeString(valHex)
				if err != nil {
					util.Log.Errorf("key[%s],%+v", key, err)
					if lastLine {
						break
					} else {
						continue
					}
				}
				restore := client.Restore(ctx, key, 0, string(decodeString))
				if restore.Err() != nil {
					util.Log.Errorf("key[%s],%+v", key, restore.Err())
					if lastLine {
						break
					} else {
						continue
					}
				}

				count++
				if lastLine {
					break
				}
			}
			util.Log.Infof("restore keys[%d]", count)
		},
	}
	cmd.Flags().StringVarP(&resGzFile, "resGzFile", "r", "", "结果gz文件路径")
	_ = cmd.MarkFlagRequired("resGzFile")

	return &cmd
}

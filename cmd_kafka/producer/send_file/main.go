package send_file

import (
	"bcd-util/cmd_kafka/prop"
	"bcd-util/util"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
	"time"
)

var filePath string
var split string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "send_file",
		Short: "发送文件中的数据到kafka",
		Run: func(cmd *cobra.Command, args []string) {
			stat, err := os.Stat(filePath)
			if err != nil {
				util.Log.Errorf("文件[%s]不存在", filePath)
				return
			}
			var files []string
			if stat.IsDir() {
				dir, err := os.ReadDir(filePath)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				for _, e := range dir {
					if !e.IsDir() && strings.HasSuffix(e.Name(), ".txt") {
						files = append(files, path.Join(filePath, e.Name()))
					}
				}
			} else {
				files = []string{filePath}
			}

			messages := make([]kafka.Message, len(files))
			for i, f := range files {
				file, err := os.Open(f)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				all, err := util.ReadSplitAll_reader(file, split[0])
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				util.Log.Infof("read file[%s] len[%d]", f, len(all))
				for _, e := range all {
					messages[i] = kafka.Message{
						Value: e,
					}
				}
			}
			//连接kafka
			kafkaWriter := &kafka.Writer{
				Addr:         kafka.TCP(prop.Addrs...),
				Topic:        prop.Topic,
				Balancer:     &kafka.LeastBytes{},
				BatchTimeout: 100 * time.Millisecond,
				BatchSize:    1000,
				//Async:                  true,
				AllowAutoTopicCreation: true,
			}
			err = kafkaWriter.WriteMessages(context.Background(), messages...)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			util.Log.Infof("send to kafka[%s] topic[%s] succeed", strings.Join(prop.Addrs, ","), prop.Topic)
		},
	}
	cmd.Flags().StringVarP(&filePath, "filePath", "f", "data.txt", "要发送数据文件路径(可以是文件夹、如果是文件夹则发送文件夹下面所有.txt文件、其中子级目录忽略)")
	cmd.Flags().StringVarP(&filePath, "split", "s", "\n", "针对每个发送的数据文件、按照分隔符分割、发送多条记录)")
	return &cmd
}

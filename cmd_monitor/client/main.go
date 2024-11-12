package client

import (
	"bcd-util/cmd_monitor/server"
	"bcd-util/support_system"
	"bcd-util/util"
	"context"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

var redisAddrs []string
var redisPassword string
var redisTopic string
var redisListName string
var serverId string
var serverName string
var serverType int

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "server",
		Short: `
通过redis通知接收到监控服务的采集监控信息请求、采集本机的监控信息、然后放在redis list中、供监控服务获取存储
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(serverName) == 0 {
				serverName = serverId
			}
			var client redis.UniversalClient
			if len(redisAddrs) == 1 {
				client = redis.NewClient(&redis.Options{
					Addr:         redisAddrs[0],
					Password:     redisPassword,
					ReadTimeout:  10 * time.Second,
					WriteTimeout: 10 * time.Second,
				})
			} else {
				client = redis.NewClusterClient(&redis.ClusterOptions{
					Addrs:        redisAddrs,
					Password:     redisPassword,
					ReadTimeout:  10 * time.Second,
					WriteTimeout: 10 * time.Second,
				})
			}
			subscribe := client.Subscribe(context.Background(), redisTopic)
			util.Log.Infof("start listen redis addrs[%s] topic[%s] serverId[%s] serverName[%s]", strings.Join(redisAddrs, ","), redisTopic, serverId, serverName)
			for message := range subscribe.Channel() {
				util.Log.Infof("receive batch[%s]", message.Payload)
				batch, err := strconv.ParseInt(message.Payload, 10, 64)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}

				systemData, err := support_system.Collect()
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				marshal, err := json.Marshal(systemData)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}

				responseData := server.ResponseData{
					ServerId:   serverId,
					Batch:      batch,
					Data:       string(marshal),
					ServerName: serverName,
					ServerType: serverType,
				}

				bytes, err := json.Marshal(responseData)

				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				lPush := client.LPush(context.Background(), redisListName, bytes)
				if lPush.Err() != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				util.Log.Infof("response batch[%s]", message.Payload)
			}
		},
	}
	cmd.Flags().StringSliceVarP(&redisAddrs, "redisAddrs", "a", []string{"127.0.0.1:3306"}, "redis地址(如果只有1个元素、视为单机、否则为集群)")
	cmd.Flags().StringVarP(&redisPassword, "redisPassword", "p", "bcd", "redis密码")
	cmd.Flags().StringVarP(&redisTopic, "redisTopic", "t", "topic_monitor", "redis下发采集指令集的topic")
	cmd.Flags().StringVarP(&redisListName, "redisListName", "l", "list_monitor", "redis结果集合名称")
	cmd.Flags().StringVarP(&serverId, "serverId", "i", "", "服务id")
	cmd.Flags().StringVarP(&serverName, "serverName", "n", "", "服务名称(默认值为serverId)")
	cmd.Flags().IntVarP(&serverType, "serverType", "y", 0, "服务类型(0:服务器)")

	_ = cmd.MarkFlagRequired("redisAddrs")
	_ = cmd.MarkFlagRequired("redisPassword")
	_ = cmd.MarkFlagRequired("serverId")
	return &cmd
}

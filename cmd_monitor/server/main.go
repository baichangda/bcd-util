package server

import (
	"bcd-util/util"
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"strings"
	"time"
)

var redisAddrs []string
var redisPassword string
var redisTopic string
var redisListName string
var mysqlUrl string
var period int
var check bool

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "server",
		Short: `
本服务通过redis广播定时广播采集信息指令、redis list名称到其他服务器、等待5s后从redis指定的list中获取所有的服务器上报的指标信息、最后存储到mysql中
服务器执行接收到采集指令后、采集本机的指标、上传到redis list
需要预先创建服务静态信息表、服务监控信息采集表
`,
		Run: func(cmd *cobra.Command, args []string) {
			db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   "t_",
					SingularTable: true,
				},
				//Logger: logger.New(util.LogWriter, logger.Config{LogLevel: logger.Info}),
			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			db.Table("t_server_data").Assign()

			err = db.AutoMigrate(&MonitorData{})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			if check {
				err = db.AutoMigrate(&ServerData{})
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
			}
			var client redis.Cmdable
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

			c := cron.New(cron.WithSeconds())
			cronExpr := "0 0/" + strconv.Itoa(period) + " * * * *"
			util.Log.Infof("start cron[%s]", cronExpr)
			_, err = c.AddFunc(cronExpr, func() {
				collectTime := time.Now()

				//往前取最近一个整点时间
				collectTime.Add(time.Duration(collectTime.Second()-(collectTime.Second()%period)) * time.Second)
				//发送采集指令
				collectTimeStr := collectTime.Format("20060102150405")
				batch, err := strconv.ParseInt(collectTimeStr, 10, 64)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				util.Log.Infof("start batch[%d]", batch)
				client.Publish(context.Background(), redisTopic, collectTimeStr)
				//等待10s
				time.Sleep(10 * time.Second)
				//获取结果、并清空结果集合
				lRange := client.LRange(context.Background(), redisListName, 0, -1)
				result, err := lRange.Result()
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				client.Del(context.Background(), redisListName)

				//解析结果
				responseDataMap := make(map[string]ResponseData)
				for _, e := range result {
					data := ResponseData{}
					err := json.Unmarshal([]byte(e), &data)
					if err != nil {
						util.Log.Errorf("response data error:\n%s", e)
						continue
					}
					if data.Batch != batch {
						util.Log.Warnf("response serverId[%s] batch[%d]!=except[%d]\n%s", data.ServerId, data.Batch, batch)
						continue
					}
					_, ok := responseDataMap[data.ServerId]
					if ok {
						util.Log.Warnf("response serverId[%s] data repeat", data.ServerId)
					}
					responseDataMap[data.ServerId] = data
				}

				if check {
					//查找所有服务器
					var serverDatas []ServerData
					db.Find(&serverDatas)
					//监控信息
					var ids1 []string
					var ids2 []string
					monitorDatas := make([]MonitorData, len(serverDatas))
					for i, serverData := range serverDatas {
						responseData, ok := responseDataMap[serverData.Id]
						monitorData := MonitorData{
							ServerId:   serverData.Id,
							ServerType: serverData.ServerType,
							ServerName: serverData.ServerName,
							Batch:      batch,
						}
						if ok {
							ids1 = append(ids1, serverData.Id)
							monitorData.ServerStatus = 0
							if len(responseData.Data) == 0 {
								monitorData.Data = sql.NullString{Valid: false}
							} else {
								monitorData.Data = sql.NullString{
									String: responseData.Data,
									Valid:  true,
								}
							}
						} else {
							ids2 = append(ids2, serverData.Id)
							monitorData.ServerStatus = 1
							monitorData.Data = sql.NullString{Valid: false}
						}
						monitorDatas[i] = monitorData
					}
					db.Create(monitorDatas)
					util.Log.Infof("finish batch[%d] alive server[%s] dead server[%s]",
						batch,
						strings.Join(ids1, ","),
						strings.Join(ids2, ","))
				} else {
					var monitorDatas []MonitorData
					var ids1 []string
					for _, v := range responseDataMap {
						ids1 = append(ids1, v.ServerId)
						monitorData := MonitorData{
							ServerId:     v.ServerId,
							ServerType:   v.ServerType,
							ServerName:   v.ServerName,
							ServerStatus: 0,
							Batch:        batch,
						}
						if len(v.Data) == 0 {
							monitorData.Data = sql.NullString{Valid: false}
						} else {
							monitorData.Data = sql.NullString{
								String: v.Data,
								Valid:  true,
							}
						}
						monitorDatas = append(monitorDatas, monitorData)
					}
					db.Create(monitorDatas)
					util.Log.Infof("finish batch[%d] alive server[%s]",
						batch,
						strings.Join(ids1, ","))
				}
			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			c.Run()
		},
	}
	cmd.Flags().IntVarP(&period, "period", "p", 1, "定时任务采集周期(分钟)")
	cmd.Flags().StringVarP(&mysqlUrl, "mysqlUrl", "u", "root:incar@2023@tcp(10.0.11.50:39005)/rvm3?multiStatements=true&charset=utf8mb4&parseTime=True&loc=Local", "mysql url连接")
	cmd.Flags().StringSliceVarP(&redisAddrs, "redisAddrs", "a", []string{"127.0.0.1:3306"}, "redis地址(如果只有1个元素、视为单机、否则为集群)")
	cmd.Flags().StringVarP(&redisPassword, "redisPassword", "w", "bcd", "redis密码")
	cmd.Flags().StringVarP(&redisTopic, "redisTopic", "t", "topic_monitor", "redis下发采集指令集的topic")
	cmd.Flags().StringVarP(&redisListName, "redisListName", "l", "list_monitor", "redis结果集合名称")
	cmd.Flags().BoolVarP(&check, "check", "c", false, "是否验证服务器信息(如果是、则需要t_server_data表)")

	_ = cmd.MarkFlagRequired("mysqlUrl")
	_ = cmd.MarkFlagRequired("redisAddrs")
	_ = cmd.MarkFlagRequired("redisPassword")
	return &cmd
}

type ServerData struct {
	Id         string         `gorm:"primaryKey;size:50;comment:服务id"`
	ServerName string         `gorm:"size:50;not null;comment:服务名称"`
	ServerType int            `gorm:"not null;comment:服务类型(0:服务器)"`
	Remark     sql.NullString `gorm:"size:200;comment:服务描述"`
	CreateTime time.Time      `gorm:"autoCreateTime"`
}

type MonitorData struct {
	Id           uint           `gorm:"primaryKey"`
	ServerId     string         `gorm:"not null;comment:服务id;size:50"`
	ServerType   int            `gorm:"not null;comment:服务类型"`
	ServerName   string         `gorm:"not null;comment:服务名称;size:50"`
	ServerStatus int            `gorm:"not null;comment:服务状态(0:正常;1:故障)"`
	Batch        int64          `gorm:"not null;comment:批次"`
	Data         sql.NullString `gorm:"size:1000;comment:服务监控数据(json格式)"`
	CreateTime   time.Time      `gorm:"autoCreateTime"`
}

type ResponseData struct {
	//服务id
	ServerId string `json:"serverId"`
	//采集批次
	Batch int64 `json:"batch"`
	//上报数据
	Data string `json:"data"`

	//以下属性如果开启了服务器校验、则不需要
	//服务名称
	ServerName string `json:"serverName"`
	//服务类型
	ServerType int `json:"serverType"`
}

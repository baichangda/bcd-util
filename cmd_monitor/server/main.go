package server

import (
	"bcd-util/support_mysql"
	"bcd-util/support_system"
	"bcd-util/util"
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var redisAddrs []string
var redisPassword string
var redisTopic string
var redisListName string
var mysqlUrl string
var period int

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "server",
		Short: `
本服务通过redis广播定时广播采集信息指令、redis list名称到其他服务器、等待5s后从redis指定的list中获取所有的服务器上报的指标信息、最后存储到mysql中
服务器执行接收到采集指令后、采集本机的指标、上传到redis list
需要预先创建服务静态信息表、服务监控信息采集表
`,
		Run: func(cmd *cobra.Command, args []string) {
			//初始化mysql
			open, err := sql.Open("mysql", mysqlUrl)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			_, err = open.Exec(CreateTableSql_serverData)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			_, err = open.Exec(CreateTableSql_monitorData)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			//加载全量服务信息
			serverDatas, err := ListServerData(open)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
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
				lRange := client.LRange(context.Background(), redisListName, 1, -1)
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
					responseDataMap[data.Id] = data
				}

				//监控信息
				var ids1 []string
				var ids2 []string
				monitorDatas := make([]MonitorData, len(result))
				for i, serverData := range serverDatas {
					responseData, ok := responseDataMap[serverData.Id]
					if ok {
						if responseData.Batch != batch {

						}
						ids1 = append(ids1, serverData.Id)
						marshal, err := json.Marshal(responseData)
						if err != nil {
							util.Log.Errorf("%+v", err)
							continue
						}
						monitorDatas[i] = MonitorData{
							ServerId:     serverData.Id,
							ServerType:   serverData.Flag,
							ServerName:   serverData.Name,
							ServerStatus: 0,
							Batch:        batch,
							Data:         string(marshal),
						}
					} else {
						ids2 = append(ids2, serverData.Id)
						monitorDatas[i] = MonitorData{
							ServerId:     serverData.Id,
							ServerType:   serverData.Flag,
							ServerName:   serverData.Name,
							ServerStatus: 1,
							Batch:        batch,
							Data:         "",
						}
					}
				}

				util.Log.Infof("finish batch[%d] alive server[%v] dead server[%v]", batch, ids1, ids2)
			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			c.Run()
		},
	}
	cmd.Flags().IntVarP(&period, "period", "p", 1, "定时任务采集周期(分钟)")
	cmd.Flags().StringVarP(&mysqlUrl, "mysqlUrl", "u", "root:incar@2023@tcp(10.0.11.50:39005)/rvm3?multiStatements=true&charset=utf8", "mysql url连接")
	cmd.Flags().StringSliceVarP(&redisAddrs, "redisAddrs", "a", []string{"127.0.0.1:3306"}, "redis地址(如果只有1个元素、视为单机、否则为集群)")
	cmd.Flags().StringVarP(&redisPassword, "redisPassword", "w", "bcd", "redis密码")
	cmd.Flags().StringVarP(&redisTopic, "redisTopic", "t", "topic_monitor", "redis下发采集指令集的topic")
	cmd.Flags().StringVarP(&redisListName, "redisListName", "l", "list_monitor", "redis结果集合名称")

	_ = cmd.MarkFlagRequired("mysqlUrl")
	_ = cmd.MarkFlagRequired("redisAddrs")
	_ = cmd.MarkFlagRequired("redisPassword")
	return &cmd
}

type MonitorData struct {
	ServerId   string
	ServerType int
	ServerName string
	//服务状态0:正常,1:故障
	ServerStatus int
	Batch        int64
	Data         string
}

func (data *MonitorData) Insert(db *sql.DB) error {
	err := support_mysql.Insert(db,
		"insert into t_monitor_data(server_id,server_type,server_name,server_status,batch,data)", []any{
			data.ServerId, data.ServerType, data.ServerName, data.ServerStatus, data.Batch, data.Data,
		})
	if err != nil {
		return errors.WithStack(err)
	} else {
		return nil
	}
}

type ResponseData struct {
	//服务id
	Id string `json:"id"`
	//系统信息
	System support_system.SystemData `json:"system"`
	//附加信息
	Ext map[string]any `json:"ext"`
	//采集时间
	Batch int64 `json:"batch"`
}

type ServerData struct {
	//服务id
	Id string
	//服务名称
	Name string
	//服务类型
	Flag int
	//服务描述
	Remark string
}

func (data *ServerData) Insert(db *sql.DB) error {
	err := support_mysql.Insert(db,
		"insert into t_server_data(id,name,flag,remark)", []any{
			data.Id, data.Name, data.Flag, data.Remark,
		})
	if err != nil {
		return errors.WithStack(err)
	} else {
		return nil
	}
}

func ListServerData(db *sql.DB) ([]ServerData, error) {
	query, err := db.Query("select id,name,flag,remark from t_server_data")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var datas []ServerData
	for query.Next() {
		data := ServerData{}
		err := query.Scan(&data.Id, &data.Name, &data.Flag, &data.Remark)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		datas = append(datas, data)
	}
	return datas, err
}

const CreateTableSql_serverData = `
create table if not exists t_server_data(
    id varchar(20) primary key not null comment '服务id',
    name varchar(50) not null comment '服务名称',
    flag int not null comment '服务类型',
    remark varchar(100) null comment '服务描述'
)
`

const CreateTableSql_monitorData = `
create table if not exists t_monitor_data(
    id bigint primary key not null auto_increment comment '主键',
    server_id varchar(20) comment '服务id',
    server_name varchar(50) comment '服务名称',
    server_type int comment '服务类型',
    server_status int comment '服务状态(0:正常;1:故障)',
    batch int comment '日期批次',
    data text comment '监控数据'
)
`

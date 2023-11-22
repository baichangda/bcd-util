package client

import (
	"bcd-util/support_system"
	"bcd-util/util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var mysqlUrl string
var period int

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "client",
		Short: `
定时采集操作系统硬件信息并存入到mysql中
定时任务时间从每分钟0s开始、以传入的参数p为间隔执行
每次执行时候会取当前时间、往前取最近一个应该执行定时任务的时间、这样可以有效解决各个服务器时间误差不大时候统一采集时间、计算方法为 采集时间=当前时间秒-(当前时间秒%p)
`,
		Run: func(cmd *cobra.Command, args []string) {
			open, err := sql.Open("mysql", mysqlUrl)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			_, err = open.Exec(support_system.CreateTableSql)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			c := cron.New(cron.WithSeconds())
			_, err = c.AddFunc("0/"+strconv.Itoa(period)+" * * * * *", func() {
				collectTime := time.Now()
				//往前取最近一个整点时间
				collectTime.Add(time.Duration(collectTime.Second()-(collectTime.Second()%period)) * time.Second)
				systemData, err := support_system.Collect()
				if err != nil {
					util.Log.Errorf("%+v", err)
					c.Stop()
					return
				}
				err = systemData.Insert(open)
				if err != nil {
					util.Log.Errorf("%+v", err)
					c.Stop()
					return
				}

				util.Log.Infof("collect system data collectTime[%s]", collectTime.Format("20060102150405"))
			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			c.Run()
		},
	}
	cmd.Flags().IntVarP(&period, "period", "p", 10, "定时任务执行间隔、从每分钟0秒开始、以间隔执行、如果不能被60整除、则每分钟0秒会执行一次")
	cmd.Flags().StringVarP(&mysqlUrl, "mysqlUrl", "u", "root:incar@2023@tcp(10.0.11.50:39005)/rvm3?multiStatements=true&charset=utf8", "mysql url连接")

	return &cmd
}

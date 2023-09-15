package client

import (
	"database/sql"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
	"gmmc-tool/util"
	"time"
)

type SystemData struct {
}

var mysqlUrl string
var period uint

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "client",
		Short: "client",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := sql.Open("mysql", mysqlUrl)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			c := cron.New()
			_, err = c.AddFunc("0/10 * * * * *", func() {
				_, err := cpu.Percent(1000, false)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}

			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			c.Run()
		},
	}
	cmd.Flags().UintVarP(&period, "period", "p", 10, "定时任务执行间隔、从每分钟0秒开始、以间隔执行、必须要被60整除")
	cmd.Flags().StringVarP(&mysqlUrl, "mysqlUrl", "u", "root:incar@2023@tcp(10.0.11.50:39005)/rvm2?multiStatements=true&charset=utf8", "mysql url连接")
	_ = cmd.MarkFlagRequired("mysqlUrl")

	return &cmd
}

func Collect() {
	percent, err := cpu.Percent(1000*time.Millisecond, false)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	util.Log.Infof("%.2f", percent[0])

	memory, err := mem.VirtualMemory()
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	util.Log.Infof("%d,%d", memory.Available, memory.Total)

	partitions, err := disk.Partitions(true)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	for _, e := range partitions {
		usage, err := disk.Usage(e.Mountpoint)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		util.Log.Infof("%s,%d,%d", e.String(), usage.Free, usage.Total)

		counters, err := disk.IOCounters(e.Mountpoint)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		for k, v := range counters {
			util.Log.Infof("%s,%d,%d", k, v.ReadBytes, v.WriteBytes)
		}
	}

	counters, err := net.IOCounters(false)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	for _, e := range counters {
		util.Log.Infof("%s,%d,%d", e.String(), e.BytesRecv, e.BytesSent)
	}
}

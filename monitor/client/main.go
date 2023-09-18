package client

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
	"gmmc-tool/util"
	"time"
)

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

type SystemData struct {
	//cpu物理核心
	PhysicalProcessorNum int `json:"physicalProcessorNum"`
	//cpu逻辑核心
	LogicalProcessorNum int `json:"logicalProcessorNum"`
	//cpu使用百分比
	CpuUsePercent float64 `json:"cpuUsePercent"`
	//内存使用百分比
	MemoryUsePercent float64 `json:"memoryUsePercent"`
	//最大内存(GB)
	MemoryMax float64 `json:"memoryMax"`
	//已使用内存(GB)
	MemoryUse float64 `json:"memoryUse"`
	//磁盘最大容量(GB)
	DiskMax float64 `json:"diskMax"`
	//磁盘使用容量(GB)
	DiskUse float64 `json:"diskUse"`
	//磁盘使用百分比
	DiskUsePercent float64 `json:"diskUsePercent"`
	//磁盘读取速度(KB/s)
	DiskReadSpeed float64 `json:"diskReadSpeed"`
	//磁盘写入速度(KB/s)
	DiskWriteSpeed float64 `json:"diskWriteSpeed"`
	//网络流入速度(KB/s)
	NetRecvSpeed float64 `json:"netRecvSpeed"`
	//网络流出速度(KB/s)
	NetSentSpeed float64 `json:"netSentSpeed"`
}

const gb float64 = 1024 * 1024 * 1024

const mb float64 = 1024 * 1024

const kb float64 = 1024

func Collect() (*SystemData, error) {
	systemData := SystemData{}
	//cpu
	physicalProcessorNum, err := cpu.Counts(false)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	logicalProcessorNum, err := cpu.Counts(true)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	systemData.PhysicalProcessorNum = physicalProcessorNum
	systemData.LogicalProcessorNum = logicalProcessorNum

	//内存
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	systemData.MemoryUse = util.Format(float64(memory.Total-memory.Available)/gb, 2)
	systemData.MemoryMax = util.Format(float64(memory.Total)/gb, 2)
	systemData.MemoryUsePercent = util.Format(float64(memory.Total-memory.Available)*100/float64(memory.Total), 2)

	//磁盘
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var disk_free uint64 = 0
	var disk_total uint64 = 0
	var prev_disk_io_read uint64 = 0
	var prev_disk_io_write uint64 = 0
	for _, e := range partitions {
		usage, err := disk.Usage(e.Mountpoint)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		disk_free += usage.Free
		disk_total += usage.Total

		counters, err := disk.IOCounters(e.Mountpoint)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, v := range counters {
			prev_disk_io_read += v.ReadBytes
			prev_disk_io_write += v.WriteBytes
		}
	}

	systemData.DiskMax = util.Format(float64(disk_total)/gb, 2)
	systemData.DiskUse = util.Format(float64(disk_total-disk_free)/gb, 2)
	systemData.DiskUsePercent = util.Format(float64(disk_total-disk_free)*100/float64(disk_total), 2)

	//网络
	var prev_net_recv uint64 = 0
	var prev_net_send uint64 = 0
	prev_net_counters, err := net.IOCounters(false)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	prev_net_recv = prev_net_counters[0].BytesRecv
	prev_net_send = prev_net_counters[0].BytesSent

	//cpu
	cpuPercent, err := cpu.Percent(1000*time.Millisecond, false)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	systemData.CpuUsePercent = util.Format(cpuPercent[0], 2)

	//磁盘
	var cur_disk_io_read uint64 = 0
	var cur_disk_io_write uint64 = 0
	for _, e := range partitions {
		counters, err := disk.IOCounters(e.Mountpoint)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, v := range counters {
			cur_disk_io_read += v.ReadBytes
			cur_disk_io_write += v.WriteBytes
		}
	}

	systemData.DiskReadSpeed = util.Format(float64(cur_disk_io_read-prev_disk_io_read)/kb, 2)
	systemData.DiskWriteSpeed = util.Format(float64(cur_disk_io_write-prev_disk_io_write)/kb, 2)

	//网络
	var cur_net_recv uint64 = 0
	var cur_net_send uint64 = 0
	cur_net_counters, err := net.IOCounters(false)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cur_net_recv = cur_net_counters[0].BytesRecv
	cur_net_send = cur_net_counters[0].BytesSent
	systemData.NetRecvSpeed = util.Format(float64(cur_net_recv-prev_net_recv)/kb, 2)
	systemData.NetSentSpeed = util.Format(float64(cur_net_send-prev_net_send)/kb, 2)

	return &systemData, nil
}

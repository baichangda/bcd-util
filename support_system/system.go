package support_system

import (
	"bcd-util/support_sql"
	"bcd-util/util"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

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
	//采集实际
	CollectTime time.Time `json:"collectTime"`
}

func (systemData *SystemData) Insert(db *sql.DB) error {
	err := support_sql.Insert(db,
		"insert into t_monitor_system_data(collect_time,physical_processor_num,logical_processor_num,cpu_use_percent,memory_use_percent,memory_max,memory_use,disk_max,disk_use,disk_use_percent,disk_read_speed,disk_write_speed,net_recv_speed,net_sent_speed)", []any{
			systemData.CollectTime,
			systemData.PhysicalProcessorNum, systemData.LogicalProcessorNum, systemData.CpuUsePercent,
			systemData.MemoryUsePercent, systemData.MemoryMax, systemData.MemoryUse, systemData.DiskMax, systemData.DiskUse,
			systemData.DiskUsePercent, systemData.DiskReadSpeed, systemData.DiskWriteSpeed, systemData.NetRecvSpeed, systemData.NetSentSpeed,
		})
	if err != nil {
		return errors.WithStack(err)
	} else {
		return nil
	}
}

const CreateTableSql = `
create table if not exists t_monitor_system_data(
    id bigint primary key not null auto_increment comment '主键',
    physical_processor_num int not null comment 'cpu物理核心',
    logical_processor_num int not null comment 'cpu逻辑核心',
    cpu_use_percent float not null comment 'cpu使用百分比',
    memory_use_percent float not null comment '内存使用百分比',
    memory_max float not null comment '最大内存(GB)',
    memory_use float not null comment '已使用内存(GB)',
    disk_max float not null comment '磁盘最大容量(GB)',
    disk_use float not null comment '磁盘使用容量(GB)',
    disk_use_percent float not null comment '磁盘使用百分比',
    disk_read_speed float not null comment '磁盘读取速度(KB/s)',
    disk_write_speed float not null comment '磁盘写入速度(KB/s)',
    net_recv_speed float not null comment '网络流入速度(KB/s)',
    net_sent_speed float not null comment '网络流出速度(KB/s)',
    collect_time timestamp not null comment '采集时间'
)
`

const gb float64 = 1024 * 1024 * 1024

const mb float64 = 1024 * 1024

const kb float64 = 1024

func Collect() (*SystemData, error) {
	systemData := SystemData{}
	systemData.CollectTime = time.Now()

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

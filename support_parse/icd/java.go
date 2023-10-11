package icd

import (
	"bcd-util/support_parse/parse"
	"reflect"
	"time"
	"unsafe"
)

type Dev_func_list struct {
	F_target_detection    uint8 `json:"target_detection"`
	F_lane_detection      uint8 `json:"lane_detection"`
	F_event_detection     uint8 `json:"event_detection"`
	F_period_detection    uint8 `json:"period_detection"`
	F_area_detection      uint8 `json:"area_detection"`
	F_trigger_detection   uint8 `json:"trigger_detection"`
	F_queue_detection     uint8 `json:"queue_detection"`
	F_status_monitor      uint8 `json:"status_monitor"`
	F_environment_monitor uint8 `json:"environment_monitor"`
}

func To_Dev_func_list(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Dev_func_list {
	_instance := Dev_func_list{}
	_start_index := _byteBuf.ReaderIndex()
	F_target_detection_v := _byteBuf.Read_uint8()
	_instance.F_target_detection = F_target_detection_v

	F_lane_detection_v := _byteBuf.Read_uint8()
	_instance.F_lane_detection = F_lane_detection_v

	F_event_detection_v := _byteBuf.Read_uint8()
	_instance.F_event_detection = F_event_detection_v

	F_period_detection_v := _byteBuf.Read_uint8()
	_instance.F_period_detection = F_period_detection_v

	F_area_detection_v := _byteBuf.Read_uint8()
	_instance.F_area_detection = F_area_detection_v

	F_trigger_detection_v := _byteBuf.Read_uint8()
	_instance.F_trigger_detection = F_trigger_detection_v

	F_queue_detection_v := _byteBuf.Read_uint8()
	_instance.F_queue_detection = F_queue_detection_v

	F_status_monitor_v := _byteBuf.Read_uint8()
	_instance.F_status_monitor = F_status_monitor_v

	F_environment_monitor_v := _byteBuf.Read_uint8()
	_instance.F_environment_monitor = F_environment_monitor_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Dev_func_list) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_target_detection)
	_byteBuf.Write_uint8(_instance.F_lane_detection)
	_byteBuf.Write_uint8(_instance.F_event_detection)
	_byteBuf.Write_uint8(_instance.F_period_detection)
	_byteBuf.Write_uint8(_instance.F_area_detection)
	_byteBuf.Write_uint8(_instance.F_trigger_detection)
	_byteBuf.Write_uint8(_instance.F_queue_detection)
	_byteBuf.Write_uint8(_instance.F_status_monitor)
	_byteBuf.Write_uint8(_instance.F_environment_monitor)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Dev_hw_list struct {
	F_computing_power string `json:"computing_power"`
	F_mem_cap         uint16 `json:"mem_cap"`
	F_in_storage_cap  uint16 `json:"in_storage_cap"`
	F_ext_storage_cap uint16 `json:"ext_storage_cap"`
	F_bandwidth       uint16 `json:"bandwidth"`
}

func To_Dev_hw_list(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Dev_hw_list {
	_instance := Dev_hw_list{}
	_start_index := _byteBuf.ReaderIndex()
	F_computing_power_len := 16
	F_computing_power_v := _byteBuf.Read_slice_uint8(F_computing_power_len)
	F_computing_power_count := 0
	for i := F_computing_power_len - 1; i >= 0; i-- {
		if F_computing_power_v[i] == 0 {
			F_computing_power_count++
		} else {
			break
		}
	}
	_instance.F_computing_power = string(F_computing_power_v[:(F_computing_power_len - F_computing_power_count)])

	F_mem_cap_v := _byteBuf.Read_uint16()
	_instance.F_mem_cap = F_mem_cap_v / 10

	F_in_storage_cap_v := _byteBuf.Read_uint16()
	_instance.F_in_storage_cap = F_in_storage_cap_v

	F_ext_storage_cap_v := _byteBuf.Read_uint16()
	_instance.F_ext_storage_cap = F_ext_storage_cap_v

	F_bandwidth_v := _byteBuf.Read_uint16()
	_instance.F_bandwidth = F_bandwidth_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Dev_hw_list) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	F_computing_power_len := 16
	F_computing_power_v := []byte(_instance.F_computing_power)
	_byteBuf.Write_slice_uint8(F_computing_power_v)
	_byteBuf.Write_zero(F_computing_power_len - len(F_computing_power_v))
	_byteBuf.Write_uint16(_instance.F_mem_cap * 10)
	_byteBuf.Write_uint16(_instance.F_in_storage_cap)
	_byteBuf.Write_uint16(_instance.F_ext_storage_cap)
	_byteBuf.Write_uint16(_instance.F_bandwidth)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Dev_sw_list struct {
	F_data_transfer int8 `json:"data_transfer"`
}

func To_Dev_sw_list(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Dev_sw_list {
	_instance := Dev_sw_list{}
	_start_index := _byteBuf.ReaderIndex()
	F_data_transfer_v := _byteBuf.Read_int8()
	_instance.F_data_transfer = F_data_transfer_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Dev_sw_list) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_int8(_instance.F_data_transfer)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Disk_infos struct {
	F_disk_id string `json:"disk_id"`
	F_size    int64  `json:"size"`
	F_usage   uint8  `json:"usage"`
}

func To_Disk_infos(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Disk_infos {
	_instance := Disk_infos{}
	F_disk_id_len := 32
	F_disk_id_v := _byteBuf.Read_slice_uint8(F_disk_id_len)
	F_disk_id_count := 0
	for i := F_disk_id_len - 1; i >= 0; i-- {
		if F_disk_id_v[i] == 0 {
			F_disk_id_count++
		} else {
			break
		}
	}
	_instance.F_disk_id = string(F_disk_id_v[:(F_disk_id_len - F_disk_id_count)])

	F_size_v := _byteBuf.Read_int64()
	_instance.F_size = F_size_v

	F_usage_v := _byteBuf.Read_uint8()
	_instance.F_usage = F_usage_v

	return &_instance
}

func (__instance *Disk_infos) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_disk_id_len := 32
	F_disk_id_v := []byte(_instance.F_disk_id)
	_byteBuf.Write_slice_uint8(F_disk_id_v)
	_byteBuf.Write_zero(F_disk_id_len - len(F_disk_id_v))
	_byteBuf.Write_int64(_instance.F_size)
	_byteBuf.Write_uint8(_instance.F_usage)
}

type Event_info struct {
	F_event_id           uint16          `json:"event_id"`
	F_event_type         uint16          `json:"event_type"`
	F_event_lon          float64         `json:"event_lon"`
	F_event_lat          float64         `json:"event_lat"`
	F_event_alt          uint32          `json:"event_alt"`
	F_event_road_id      uint32          `json:"event_road_id"`
	F_src_count          uint16          `json:"src_count"`
	F_target_count       uint16          `json:"target_count"`
	F_src_array          []uint32        `json:"src_array"`
	F_event_target_array []*Event_target `json:"event_target_array"`
}

func To_Event_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Event_info {
	_instance := Event_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_event_id_v := _byteBuf.Read_uint16()
	_instance.F_event_id = F_event_id_v

	F_event_type_v := _byteBuf.Read_uint16()
	_instance.F_event_type = F_event_type_v

	F_event_lon_v := _byteBuf.Read_int64()
	_instance.F_event_lon = float64(F_event_lon_v) / 10000000

	F_event_lat_v := _byteBuf.Read_int64()
	_instance.F_event_lat = float64(F_event_lat_v) / 10000000

	F_event_alt_v := _byteBuf.Read_uint32()
	_instance.F_event_alt = F_event_alt_v

	F_event_road_id_v := _byteBuf.Read_uint32()
	_instance.F_event_road_id = F_event_road_id_v

	F_src_count_v := _byteBuf.Read_uint16()
	_instance.F_src_count = F_src_count_v

	F_target_count_v := _byteBuf.Read_uint16()
	_instance.F_target_count = F_target_count_v

	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_src_array_len := (int)(F_src_count_v)
	F_src_array_arr := make([]uint32, F_src_array_len, F_src_array_len)
	for i := 0; i < F_src_array_len; i++ {
		e := _byteBuf.Read_uint32()
		F_src_array_arr[i] = e
	}
	_instance.F_src_array = F_src_array_arr
	F_event_target_array_len := (int)(F_target_count_v)
	F_event_target_array_arr := make([]*Event_target, F_event_target_array_len, F_event_target_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < F_event_target_array_len; i++ {
		F_event_target_array_arr[i] = To_Event_target(_byteBuf, _parseContext)
	}
	_instance.F_event_target_array = F_event_target_array_arr
	return &_instance
}

func (__instance *Event_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.F_event_id)
	_byteBuf.Write_uint16(_instance.F_event_type)
	_byteBuf.Write_int64(int64(parse.Round(_instance.F_event_lon * 10000000)))
	_byteBuf.Write_int64(int64(parse.Round(_instance.F_event_lat * 10000000)))
	_byteBuf.Write_uint32(_instance.F_event_alt)
	_byteBuf.Write_uint32(_instance.F_event_road_id)
	_byteBuf.Write_uint16(_instance.F_src_count)
	_byteBuf.Write_uint16(_instance.F_target_count)
	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_src_array_arr := _instance.F_src_array
	for i := 0; i < len(F_src_array_arr); i++ {
		_byteBuf.Write_uint32(F_src_array_arr[i])
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	F_event_target_array_arr := _instance.F_event_target_array
	for i := 0; i < len(F_event_target_array_arr); i++ {
		F_event_target_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Event_target struct {
	F_track_id    uint32  `json:"track_id"`
	F_lon         float64 `json:"lon"`
	F_lat         float64 `json:"lat"`
	F_alt         uint32  `json:"alt"`
	F_targetClass uint8   `json:"targetClass"`
	F_extras      any     `json:"extras"`
}

func To_Event_target(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Event_target {
	_instance := Event_target{}
	_start_index := _byteBuf.ReaderIndex()
	F_track_id_v := _byteBuf.Read_uint32()
	_instance.F_track_id = F_track_id_v

	F_lon_v := _byteBuf.Read_uint32()
	_instance.F_lon = float64(F_lon_v) / 10000000

	F_lat_v := _byteBuf.Read_uint32()
	_instance.F_lat = float64(F_lat_v) / 10000000

	F_alt_v := _byteBuf.Read_uint32()
	_instance.F_alt = F_alt_v

	F_targetClass_v := _byteBuf.Read_uint8()
	_instance.F_targetClass = F_targetClass_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_extras = To_F_extras(_byteBuf, _parseContext)
	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Event_target) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_track_id)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lon * 10000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lat * 10000000)))
	_byteBuf.Write_uint32(_instance.F_alt)
	_byteBuf.Write_uint8(_instance.F_targetClass)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_extras(_byteBuf, _instance.F_extras, _parseContext)
	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Lane_info_area struct {
	F_lane_id        uint8  `json:"lane_id"`
	F_car_count      uint16 `json:"car_count"`
	F_occupancy      uint8  `json:"occupancy"`
	F_ave_car_speed  uint16 `json:"ave_car_speed"`
	F_car_distribute uint32 `json:"car_distribute"`
	F_head_car_pos   uint16 `json:"head_car_pos"`
	F_head_car_speed uint16 `json:"head_car_speed"`
	F_tail_car_pos   uint16 `json:"tail_car_pos"`
	F_tail_car_speed uint16 `json:"tail_car_speed"`
}

func To_Lane_info_area(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Lane_info_area {
	_instance := Lane_info_area{}
	_start_index := _byteBuf.ReaderIndex()
	F_lane_id_v := _byteBuf.Read_uint8()
	_instance.F_lane_id = F_lane_id_v

	F_car_count_v := _byteBuf.Read_uint16()
	_instance.F_car_count = F_car_count_v

	F_occupancy_v := _byteBuf.Read_uint8()
	_instance.F_occupancy = F_occupancy_v

	F_ave_car_speed_v := _byteBuf.Read_uint16()
	_instance.F_ave_car_speed = F_ave_car_speed_v

	F_car_distribute_v := _byteBuf.Read_uint32()
	_instance.F_car_distribute = F_car_distribute_v

	F_head_car_pos_v := _byteBuf.Read_uint16()
	_instance.F_head_car_pos = F_head_car_pos_v

	F_head_car_speed_v := _byteBuf.Read_uint16()
	_instance.F_head_car_speed = F_head_car_speed_v

	F_tail_car_pos_v := _byteBuf.Read_uint16()
	_instance.F_tail_car_pos = F_tail_car_pos_v

	F_tail_car_speed_v := _byteBuf.Read_uint16()
	_instance.F_tail_car_speed = F_tail_car_speed_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Lane_info_area) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_lane_id)
	_byteBuf.Write_uint16(_instance.F_car_count)
	_byteBuf.Write_uint8(_instance.F_occupancy)
	_byteBuf.Write_uint16(_instance.F_ave_car_speed)
	_byteBuf.Write_uint32(_instance.F_car_distribute)
	_byteBuf.Write_uint16(_instance.F_head_car_pos)
	_byteBuf.Write_uint16(_instance.F_head_car_speed)
	_byteBuf.Write_uint16(_instance.F_tail_car_pos)
	_byteBuf.Write_uint16(_instance.F_tail_car_speed)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Lane_info_cycle struct {
	F_lane_id          uint8  `json:"lane_id"`
	F_total_car        uint16 `json:"total_car"`
	F_car_a_count      uint16 `json:"car_a_count"`
	F_car_b_count      uint16 `json:"car_b_count"`
	F_car_c_count      uint16 `json:"car_c_count"`
	F_occupancy        uint8  `json:"occupancy"`
	F_ave_speed        uint16 `json:"ave_speed"`
	F_ave_car_len      uint8  `json:"ave_car_len"`
	F_ave_car_head_dis uint8  `json:"ave_car_head_dis"`
	F_ave_car_body_dis uint8  `json:"ave_car_body_dis"`
}

func To_Lane_info_cycle(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Lane_info_cycle {
	_instance := Lane_info_cycle{}
	_start_index := _byteBuf.ReaderIndex()
	F_lane_id_v := _byteBuf.Read_uint8()
	_instance.F_lane_id = F_lane_id_v

	F_total_car_v := _byteBuf.Read_uint16()
	_instance.F_total_car = F_total_car_v

	F_car_a_count_v := _byteBuf.Read_uint16()
	_instance.F_car_a_count = F_car_a_count_v

	F_car_b_count_v := _byteBuf.Read_uint16()
	_instance.F_car_b_count = F_car_b_count_v

	F_car_c_count_v := _byteBuf.Read_uint16()
	_instance.F_car_c_count = F_car_c_count_v

	F_occupancy_v := _byteBuf.Read_uint8()
	_instance.F_occupancy = F_occupancy_v

	F_ave_speed_v := _byteBuf.Read_uint16()
	_instance.F_ave_speed = F_ave_speed_v

	F_ave_car_len_v := _byteBuf.Read_uint8()
	_instance.F_ave_car_len = F_ave_car_len_v

	F_ave_car_head_dis_v := _byteBuf.Read_uint8()
	_instance.F_ave_car_head_dis = F_ave_car_head_dis_v

	F_ave_car_body_dis_v := _byteBuf.Read_uint8()
	_instance.F_ave_car_body_dis = F_ave_car_body_dis_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Lane_info_cycle) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_lane_id)
	_byteBuf.Write_uint16(_instance.F_total_car)
	_byteBuf.Write_uint16(_instance.F_car_a_count)
	_byteBuf.Write_uint16(_instance.F_car_b_count)
	_byteBuf.Write_uint16(_instance.F_car_c_count)
	_byteBuf.Write_uint8(_instance.F_occupancy)
	_byteBuf.Write_uint16(_instance.F_ave_speed)
	_byteBuf.Write_uint8(_instance.F_ave_car_len)
	_byteBuf.Write_uint8(_instance.F_ave_car_head_dis)
	_byteBuf.Write_uint8(_instance.F_ave_car_body_dis)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Lane_info_queue struct {
	F_lane_id      uint8  `json:"lane_id"`
	F_len          uint16 `json:"len"`
	F_head_car_pos uint16 `json:"head_car_pos"`
	F_tail_car_pos uint16 `json:"tail_car_pos"`
	F_car_count    uint16 `json:"car_count"`
}

func To_Lane_info_queue(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Lane_info_queue {
	_instance := Lane_info_queue{}
	_start_index := _byteBuf.ReaderIndex()
	F_lane_id_v := _byteBuf.Read_uint8()
	_instance.F_lane_id = F_lane_id_v

	F_len_v := _byteBuf.Read_uint16()
	_instance.F_len = F_len_v

	F_head_car_pos_v := _byteBuf.Read_uint16()
	_instance.F_head_car_pos = F_head_car_pos_v

	F_tail_car_pos_v := _byteBuf.Read_uint16()
	_instance.F_tail_car_pos = F_tail_car_pos_v

	F_car_count_v := _byteBuf.Read_uint16()
	_instance.F_car_count = F_car_count_v

	F_reserved_len := 16
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Lane_info_queue) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_lane_id)
	_byteBuf.Write_uint16(_instance.F_len)
	_byteBuf.Write_uint16(_instance.F_head_car_pos)
	_byteBuf.Write_uint16(_instance.F_tail_car_pos)
	_byteBuf.Write_uint16(_instance.F_car_count)
	F_reserved_len := 16
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Lane_info_trigger struct {
	F_track_id uint32 `json:"track_id"`
	F_lane_id  uint8  `json:"lane_id"`
	F_lane_dis uint16 `json:"lane_dis"`
	F_speed    uint16 `json:"speed"`
	F_status   uint8  `json:"status"`
}

func To_Lane_info_trigger(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Lane_info_trigger {
	_instance := Lane_info_trigger{}
	_start_index := _byteBuf.ReaderIndex()
	F_track_id_v := _byteBuf.Read_uint32()
	_instance.F_track_id = F_track_id_v

	F_lane_id_v := _byteBuf.Read_uint8()
	_instance.F_lane_id = F_lane_id_v

	F_lane_dis_v := _byteBuf.Read_uint16()
	_instance.F_lane_dis = F_lane_dis_v

	F_speed_v := _byteBuf.Read_uint16()
	_instance.F_speed = F_speed_v

	F_status_v := _byteBuf.Read_uint8()
	_instance.F_status = F_status_v

	F_reserved_len := 16
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Lane_info_trigger) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_track_id)
	_byteBuf.Write_uint8(_instance.F_lane_id)
	_byteBuf.Write_uint16(_instance.F_lane_dis)
	_byteBuf.Write_uint16(_instance.F_speed)
	_byteBuf.Write_uint8(_instance.F_status)
	F_reserved_len := 16
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Msg struct {
	F_msg_header *Msg_header `json:"msg_header"`
	F_msg_body   any         `json:"msg_body"`
	F_msg_tailer *Msg_tailer `json:"msg_tailer"`
}

func To_Msg(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg {
	_instance := Msg{}
	_instance.F_msg_header = To_Msg_header(_byteBuf, nil)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_msg_body = To_F_msg_body(_byteBuf, _parseContext)
	_instance.F_msg_tailer = To_Msg_tailer(_byteBuf, nil)
	return &_instance
}

func (__instance *Msg) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_instance.F_msg_header.Write(_byteBuf, nil)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_msg_body(_byteBuf, _instance.F_msg_body, _parseContext)
	_instance.F_msg_tailer.Write(_byteBuf, nil)
}

type Msg_body_area_statistics_info struct {
	F_period          float32           `json:"period"`
	F_area_dis_near   uint16            `json:"area_dis_near"`
	F_area_dis_far    uint16            `json:"area_dis_far"`
	F_src_count       uint16            `json:"src_count"`
	F_lane_count      uint8             `json:"lane_count"`
	F_src_array       []uint32          `json:"src_array"`
	F_lane_info_array []*Lane_info_area `json:"lane_info_array"`
}

func To_Msg_body_area_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_area_statistics_info {
	_instance := Msg_body_area_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_period_v := _byteBuf.Read_uint16()
	_instance.F_period = float32(F_period_v) / 10

	F_area_dis_near_v := _byteBuf.Read_uint16()
	_instance.F_area_dis_near = F_area_dis_near_v

	F_area_dis_far_v := _byteBuf.Read_uint16()
	_instance.F_area_dis_far = F_area_dis_far_v

	F_src_count_v := _byteBuf.Read_uint16()
	_instance.F_src_count = F_src_count_v

	F_lane_count_v := _byteBuf.Read_uint8()
	_instance.F_lane_count = F_lane_count_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_src_array_len := (int)(F_src_count_v)
	F_src_array_arr := make([]uint32, F_src_array_len, F_src_array_len)
	for i := 0; i < F_src_array_len; i++ {
		e := _byteBuf.Read_uint32()
		F_src_array_arr[i] = e
	}
	_instance.F_src_array = F_src_array_arr
	F_lane_info_array_len := (int)(F_lane_count_v)
	F_lane_info_array_arr := make([]*Lane_info_area, F_lane_info_array_len, F_lane_info_array_len)
	for i := 0; i < F_lane_info_array_len; i++ {
		F_lane_info_array_arr[i] = To_Lane_info_area(_byteBuf, nil)
	}
	_instance.F_lane_info_array = F_lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_area_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_period * 10)))
	_byteBuf.Write_uint16(_instance.F_area_dis_near)
	_byteBuf.Write_uint16(_instance.F_area_dis_far)
	_byteBuf.Write_uint16(_instance.F_src_count)
	_byteBuf.Write_uint8(_instance.F_lane_count)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_src_array_arr := _instance.F_src_array
	for i := 0; i < len(F_src_array_arr); i++ {
		_byteBuf.Write_uint32(F_src_array_arr[i])
	}
	F_lane_info_array_arr := _instance.F_lane_info_array
	for i := 0; i < len(F_lane_info_array_arr); i++ {
		F_lane_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Msg_body_cycle_statistics_info struct {
	F_period          uint16             `json:"period"`
	F_len_A           uint8              `json:"len_A"`
	F_len_B           uint8              `json:"len_B"`
	F_len_C           uint8              `json:"len_C"`
	F_src_count       uint16             `json:"src_count"`
	F_lane_count      uint8              `json:"lane_count"`
	F_src_array       []uint32           `json:"src_array"`
	F_lane_info_array []*Lane_info_cycle `json:"lane_info_array"`
}

func To_Msg_body_cycle_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_cycle_statistics_info {
	_instance := Msg_body_cycle_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_period_v := _byteBuf.Read_uint16()
	_instance.F_period = F_period_v

	F_len_A_v := _byteBuf.Read_uint8()
	_instance.F_len_A = F_len_A_v

	F_len_B_v := _byteBuf.Read_uint8()
	_instance.F_len_B = F_len_B_v

	F_len_C_v := _byteBuf.Read_uint8()
	_instance.F_len_C = F_len_C_v

	F_src_count_v := _byteBuf.Read_uint16()
	_instance.F_src_count = F_src_count_v

	F_lane_count_v := _byteBuf.Read_uint8()
	_instance.F_lane_count = F_lane_count_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_src_array_len := (int)(F_src_count_v)
	F_src_array_arr := make([]uint32, F_src_array_len, F_src_array_len)
	for i := 0; i < F_src_array_len; i++ {
		e := _byteBuf.Read_uint32()
		F_src_array_arr[i] = e
	}
	_instance.F_src_array = F_src_array_arr
	F_lane_info_array_len := (int)(F_lane_count_v)
	F_lane_info_array_arr := make([]*Lane_info_cycle, F_lane_info_array_len, F_lane_info_array_len)
	for i := 0; i < F_lane_info_array_len; i++ {
		F_lane_info_array_arr[i] = To_Lane_info_cycle(_byteBuf, nil)
	}
	_instance.F_lane_info_array = F_lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_cycle_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.F_period)
	_byteBuf.Write_uint8(_instance.F_len_A)
	_byteBuf.Write_uint8(_instance.F_len_B)
	_byteBuf.Write_uint8(_instance.F_len_C)
	_byteBuf.Write_uint16(_instance.F_src_count)
	_byteBuf.Write_uint8(_instance.F_lane_count)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_src_array_arr := _instance.F_src_array
	for i := 0; i < len(F_src_array_arr); i++ {
		_byteBuf.Write_uint32(F_src_array_arr[i])
	}
	F_lane_info_array_arr := _instance.F_lane_info_array
	for i := 0; i < len(F_lane_info_array_arr); i++ {
		F_lane_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Msg_body_device_status_info struct {
	F_dev_sn        uint32         `json:"dev_sn"`
	F_dev_status    uint8          `json:"dev_status"`
	F_dev_hw_list   *Dev_hw_list   `json:"dev_hw_list"`
	F_dev_sw_list   *Dev_sw_list   `json:"dev_sw_list"`
	F_dev_func_list *Dev_func_list `json:"dev_func_list"`
}

func To_Msg_body_device_status_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_device_status_info {
	_instance := Msg_body_device_status_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_dev_sn_v := _byteBuf.Read_uint32()
	_instance.F_dev_sn = F_dev_sn_v

	F_dev_status_v := _byteBuf.Read_uint8()
	_instance.F_dev_status = F_dev_status_v

	_instance.F_dev_hw_list = To_Dev_hw_list(_byteBuf, nil)
	_instance.F_dev_sw_list = To_Dev_sw_list(_byteBuf, nil)
	_instance.F_dev_func_list = To_Dev_func_list(_byteBuf, nil)
	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Msg_body_device_status_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_dev_sn)
	_byteBuf.Write_uint8(_instance.F_dev_status)
	_instance.F_dev_hw_list.Write(_byteBuf, nil)
	_instance.F_dev_sw_list.Write(_byteBuf, nil)
	_instance.F_dev_func_list.Write(_byteBuf, nil)
	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Msg_body_event_info struct {
	F_event_count      uint16        `json:"event_count"`
	F_event_info_array []*Event_info `json:"event_info_array"`
}

func To_Msg_body_event_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_event_info {
	_instance := Msg_body_event_info{}
	F_event_count_v := _byteBuf.Read_uint16()
	_instance.F_event_count = F_event_count_v

	F_event_info_array_len := (int)(F_event_count_v)
	F_event_info_array_arr := make([]*Event_info, F_event_info_array_len, F_event_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < F_event_info_array_len; i++ {
		F_event_info_array_arr[i] = To_Event_info(_byteBuf, _parseContext)
	}
	_instance.F_event_info_array = F_event_info_array_arr
	return &_instance
}

func (__instance *Msg_body_event_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_event_count)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	F_event_info_array_arr := _instance.F_event_info_array
	for i := 0; i < len(F_event_info_array_arr); i++ {
		F_event_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Msg_body_lane_detect_info struct {
	F_frame_id        uint32       `json:"frame_id"`
	F_src_count       uint16       `json:"src_count"`
	F_road_count      uint32       `json:"road_count"`
	F_src_array       []uint32     `json:"src_array"`
	F_road_info_array []*Road_info `json:"road_info_array"`
}

func To_Msg_body_lane_detect_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_lane_detect_info {
	_instance := Msg_body_lane_detect_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_frame_id_v := _byteBuf.Read_uint32()
	_instance.F_frame_id = F_frame_id_v

	F_src_count_v := _byteBuf.Read_uint16()
	_instance.F_src_count = F_src_count_v

	F_road_count_v := _byteBuf.Read_uint32()
	_instance.F_road_count = F_road_count_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_src_array_len := (int)(F_src_count_v)
	F_src_array_arr := make([]uint32, F_src_array_len, F_src_array_len)
	for i := 0; i < F_src_array_len; i++ {
		e := _byteBuf.Read_uint32()
		F_src_array_arr[i] = e
	}
	_instance.F_src_array = F_src_array_arr
	F_road_info_array_len := (int)(F_road_count_v)
	F_road_info_array_arr := make([]*Road_info, F_road_info_array_len, F_road_info_array_len)
	for i := 0; i < F_road_info_array_len; i++ {
		F_road_info_array_arr[i] = To_Road_info(_byteBuf, nil)
	}
	_instance.F_road_info_array = F_road_info_array_arr
	return &_instance
}

func (__instance *Msg_body_lane_detect_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_frame_id)
	_byteBuf.Write_uint16(_instance.F_src_count)
	_byteBuf.Write_uint32(_instance.F_road_count)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_src_array_arr := _instance.F_src_array
	for i := 0; i < len(F_src_array_arr); i++ {
		_byteBuf.Write_uint32(F_src_array_arr[i])
	}
	F_road_info_array_arr := _instance.F_road_info_array
	for i := 0; i < len(F_road_info_array_arr); i++ {
		F_road_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Msg_body_queue_statistics_info struct {
	F_src_count       uint16             `json:"src_count"`
	F_lane_count      uint8              `json:"lane_count"`
	F_src_array       []uint32           `json:"src_array"`
	F_lane_info_array []*Lane_info_queue `json:"lane_info_array"`
}

func To_Msg_body_queue_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_queue_statistics_info {
	_instance := Msg_body_queue_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_src_count_v := _byteBuf.Read_uint16()
	_instance.F_src_count = F_src_count_v

	F_lane_count_v := _byteBuf.Read_uint8()
	_instance.F_lane_count = F_lane_count_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_src_array_len := (int)(F_src_count_v)
	F_src_array_arr := make([]uint32, F_src_array_len, F_src_array_len)
	for i := 0; i < F_src_array_len; i++ {
		e := _byteBuf.Read_uint32()
		F_src_array_arr[i] = e
	}
	_instance.F_src_array = F_src_array_arr
	F_lane_info_array_len := (int)(F_lane_count_v)
	F_lane_info_array_arr := make([]*Lane_info_queue, F_lane_info_array_len, F_lane_info_array_len)
	for i := 0; i < F_lane_info_array_len; i++ {
		F_lane_info_array_arr[i] = To_Lane_info_queue(_byteBuf, nil)
	}
	_instance.F_lane_info_array = F_lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_queue_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.F_src_count)
	_byteBuf.Write_uint8(_instance.F_lane_count)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_src_array_arr := _instance.F_src_array
	for i := 0; i < len(F_src_array_arr); i++ {
		_byteBuf.Write_uint32(F_src_array_arr[i])
	}
	F_lane_info_array_arr := _instance.F_lane_info_array
	for i := 0; i < len(F_lane_info_array_arr); i++ {
		F_lane_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Msg_body_road_info struct {
	F_road_count      uint16        `json:"road_count"`
	F_road_info_array []*Road2_info `json:"road_info_array"`
}

func To_Msg_body_road_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_road_info {
	_instance := Msg_body_road_info{}
	F_road_count_v := _byteBuf.Read_uint16()
	_instance.F_road_count = F_road_count_v

	F_road_info_array_len := (int)(F_road_count_v)
	F_road_info_array_arr := make([]*Road2_info, F_road_info_array_len, F_road_info_array_len)
	for i := 0; i < F_road_info_array_len; i++ {
		F_road_info_array_arr[i] = To_Road2_info(_byteBuf, nil)
	}
	_instance.F_road_info_array = F_road_info_array_arr
	return &_instance
}

func (__instance *Msg_body_road_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_road_count)
	F_road_info_array_arr := _instance.F_road_info_array
	for i := 0; i < len(F_road_info_array_arr); i++ {
		F_road_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Msg_body_sensor_status_info struct {
	F_sensor_count      uint16         `json:"sensor_count"`
	F_sensor_info_array []*Sensor_info `json:"sensor_info_array"`
}

func To_Msg_body_sensor_status_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_sensor_status_info {
	_instance := Msg_body_sensor_status_info{}
	F_sensor_count_v := _byteBuf.Read_uint16()
	_instance.F_sensor_count = F_sensor_count_v

	F_sensor_info_array_len := (int)(F_sensor_count_v)
	F_sensor_info_array_arr := make([]*Sensor_info, F_sensor_info_array_len, F_sensor_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < F_sensor_info_array_len; i++ {
		F_sensor_info_array_arr[i] = To_Sensor_info(_byteBuf, _parseContext)
	}
	_instance.F_sensor_info_array = F_sensor_info_array_arr
	return &_instance
}

func (__instance *Msg_body_sensor_status_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_sensor_count)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	F_sensor_info_array_arr := _instance.F_sensor_info_array
	for i := 0; i < len(F_sensor_info_array_arr); i++ {
		F_sensor_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Msg_body_system_runtime_info struct {
	F_cpu_num    uint8              `json:"cpu_num"`
	F_cpu_usage  parse.JsonUint8Arr `json:"cpu_usage"`
	F_mem_size   uint32             `json:"mem_size"`
	F_mem_usage  uint8              `json:"mem_usage"`
	F_gpu_num    uint8              `json:"gpu_num"`
	F_gpu_usage  parse.JsonUint8Arr `json:"gpu_usage"`
	F_disk_num   uint8              `json:"disk_num"`
	F_disk_infos []*Disk_infos      `json:"disk_infos"`
	F_net_num    uint8              `json:"net_num"`
	F_net_infos  []*Net_infos       `json:"net_infos"`
	F_temp_num   uint8              `json:"temp_num"`
	F_temp_val   parse.JsonUint8Arr `json:"temp_val"`
	F_fans_num   uint8              `json:"fans_num"`
	F_fans_speed parse.JsonUint8Arr `json:"fans_speed"`
}

func To_Msg_body_system_runtime_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_system_runtime_info {
	_instance := Msg_body_system_runtime_info{}
	F_cpu_num_v := _byteBuf.Read_uint8()
	_instance.F_cpu_num = F_cpu_num_v

	F_cpu_usage_len := (int)(F_cpu_num_v)
	F_cpu_usage_arr := _byteBuf.Read_slice_uint8(F_cpu_usage_len)
	_instance.F_cpu_usage = F_cpu_usage_arr
	F_mem_size_v := _byteBuf.Read_uint32()
	_instance.F_mem_size = F_mem_size_v

	F_mem_usage_v := _byteBuf.Read_uint8()
	_instance.F_mem_usage = F_mem_usage_v

	F_gpu_num_v := _byteBuf.Read_uint8()
	_instance.F_gpu_num = F_gpu_num_v

	F_gpu_usage_len := (int)(F_gpu_num_v)
	F_gpu_usage_arr := _byteBuf.Read_slice_uint8(F_gpu_usage_len)
	_instance.F_gpu_usage = F_gpu_usage_arr
	F_disk_num_v := _byteBuf.Read_uint8()
	_instance.F_disk_num = F_disk_num_v

	F_disk_infos_len := (int)(F_disk_num_v)
	F_disk_infos_arr := make([]*Disk_infos, F_disk_infos_len, F_disk_infos_len)
	for i := 0; i < F_disk_infos_len; i++ {
		F_disk_infos_arr[i] = To_Disk_infos(_byteBuf, nil)
	}
	_instance.F_disk_infos = F_disk_infos_arr
	F_net_num_v := _byteBuf.Read_uint8()
	_instance.F_net_num = F_net_num_v

	F_net_infos_len := (int)(F_net_num_v)
	F_net_infos_arr := make([]*Net_infos, F_net_infos_len, F_net_infos_len)
	for i := 0; i < F_net_infos_len; i++ {
		F_net_infos_arr[i] = To_Net_infos(_byteBuf, nil)
	}
	_instance.F_net_infos = F_net_infos_arr
	F_temp_num_v := _byteBuf.Read_uint8()
	_instance.F_temp_num = F_temp_num_v

	F_temp_val_len := (int)(F_temp_num_v)
	F_temp_val_arr := _byteBuf.Read_slice_uint8(F_temp_val_len)
	_instance.F_temp_val = F_temp_val_arr
	F_fans_num_v := _byteBuf.Read_uint8()
	_instance.F_fans_num = F_fans_num_v

	F_fans_speed_len := (int)(F_fans_num_v)
	F_fans_speed_arr := _byteBuf.Read_slice_uint8(F_fans_speed_len)
	_instance.F_fans_speed = F_fans_speed_arr
	return &_instance
}

func (__instance *Msg_body_system_runtime_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_cpu_num)
	F_cpu_usage_arr := _instance.F_cpu_usage
	_byteBuf.Write_slice_uint8(F_cpu_usage_arr)
	_byteBuf.Write_uint32(_instance.F_mem_size)
	_byteBuf.Write_uint8(_instance.F_mem_usage)
	_byteBuf.Write_uint8(_instance.F_gpu_num)
	F_gpu_usage_arr := _instance.F_gpu_usage
	_byteBuf.Write_slice_uint8(F_gpu_usage_arr)
	_byteBuf.Write_uint8(_instance.F_disk_num)
	F_disk_infos_arr := _instance.F_disk_infos
	for i := 0; i < len(F_disk_infos_arr); i++ {
		F_disk_infos_arr[i].Write(_byteBuf, nil)
	}
	_byteBuf.Write_uint8(_instance.F_net_num)
	F_net_infos_arr := _instance.F_net_infos
	for i := 0; i < len(F_net_infos_arr); i++ {
		F_net_infos_arr[i].Write(_byteBuf, nil)
	}
	_byteBuf.Write_uint8(_instance.F_temp_num)
	F_temp_val_arr := _instance.F_temp_val
	_byteBuf.Write_slice_uint8(F_temp_val_arr)
	_byteBuf.Write_uint8(_instance.F_fans_num)
	F_fans_speed_arr := _instance.F_fans_speed
	_byteBuf.Write_slice_uint8(F_fans_speed_arr)
}

type Msg_body_target_detect_info struct {
	F_frame_id          uint32         `json:"frame_id"`
	F_src_count         uint16         `json:"src_count"`
	F_target_count      uint16         `json:"target_count"`
	F_src_array         []uint32       `json:"src_array"`
	F_target_info_array []*Target_info `json:"target_info_array"`
}

func To_Msg_body_target_detect_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_target_detect_info {
	_instance := Msg_body_target_detect_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_frame_id_v := _byteBuf.Read_uint32()
	_instance.F_frame_id = F_frame_id_v

	F_src_count_v := _byteBuf.Read_uint16()
	_instance.F_src_count = F_src_count_v

	F_target_count_v := _byteBuf.Read_uint16()
	_instance.F_target_count = F_target_count_v

	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_src_array_len := (int)(F_src_count_v)
	F_src_array_arr := make([]uint32, F_src_array_len, F_src_array_len)
	for i := 0; i < F_src_array_len; i++ {
		e := _byteBuf.Read_uint32()
		F_src_array_arr[i] = e
	}
	_instance.F_src_array = F_src_array_arr
	F_target_info_array_len := (int)(F_target_count_v)
	F_target_info_array_arr := make([]*Target_info, F_target_info_array_len, F_target_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < F_target_info_array_len; i++ {
		F_target_info_array_arr[i] = To_Target_info(_byteBuf, _parseContext)
	}
	_instance.F_target_info_array = F_target_info_array_arr
	return &_instance
}

func (__instance *Msg_body_target_detect_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_frame_id)
	_byteBuf.Write_uint16(_instance.F_src_count)
	_byteBuf.Write_uint16(_instance.F_target_count)
	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_src_array_arr := _instance.F_src_array
	for i := 0; i < len(F_src_array_arr); i++ {
		_byteBuf.Write_uint32(F_src_array_arr[i])
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	F_target_info_array_arr := _instance.F_target_info_array
	for i := 0; i < len(F_target_info_array_arr); i++ {
		F_target_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Msg_body_trigger_statistics_info struct {
	F_src_count       uint16               `json:"src_count"`
	F_lane_count      uint8                `json:"lane_count"`
	F_src_array       []uint32             `json:"src_array"`
	F_lane_info_array []*Lane_info_trigger `json:"lane_info_array"`
}

func To_Msg_body_trigger_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_trigger_statistics_info {
	_instance := Msg_body_trigger_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_src_count_v := _byteBuf.Read_uint16()
	_instance.F_src_count = F_src_count_v

	F_lane_count_v := _byteBuf.Read_uint8()
	_instance.F_lane_count = F_lane_count_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_src_array_len := (int)(F_src_count_v)
	F_src_array_arr := make([]uint32, F_src_array_len, F_src_array_len)
	for i := 0; i < F_src_array_len; i++ {
		e := _byteBuf.Read_uint32()
		F_src_array_arr[i] = e
	}
	_instance.F_src_array = F_src_array_arr
	F_lane_info_array_len := (int)(F_lane_count_v)
	F_lane_info_array_arr := make([]*Lane_info_trigger, F_lane_info_array_len, F_lane_info_array_len)
	for i := 0; i < F_lane_info_array_len; i++ {
		F_lane_info_array_arr[i] = To_Lane_info_trigger(_byteBuf, nil)
	}
	_instance.F_lane_info_array = F_lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_trigger_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.F_src_count)
	_byteBuf.Write_uint8(_instance.F_lane_count)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_src_array_arr := _instance.F_src_array
	for i := 0; i < len(F_src_array_arr); i++ {
		_byteBuf.Write_uint32(F_src_array_arr[i])
	}
	F_lane_info_array_arr := _instance.F_lane_info_array
	for i := 0; i < len(F_lane_info_array_arr); i++ {
		F_lane_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Msg_header struct {
	F_header      [4]uint8  `json:"header"`
	F_header_len  uint16    `json:"header_len"`
	F_frame_len   uint32    `json:"frame_len"`
	F_frame_type  uint16    `json:"frame_type"`
	F_version     uint32    `json:"version"`
	F_device_sn   uint32    `json:"device_sn"`
	F_count       uint32    `json:"count"`
	F_timestamp   time.Time `json:"timestamp"`
	F_fps         float32   `json:"fps"`
	F_dev_lon     float64   `json:"dev_lon"`
	F_dev_lat     float64   `json:"dev_lat"`
	F_dev_alt     uint32    `json:"dev_alt"`
	F_dev_azimuth float32   `json:"dev_azimuth"`
}

func To_Msg_header(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_header {
	_instance := Msg_header{}
	_start_index := _byteBuf.ReaderIndex()
	F_header_arr := [4]uint8(_byteBuf.Read_slice_uint8(4))
	_instance.F_header = F_header_arr
	F_header_len_v := _byteBuf.Read_uint16()
	_instance.F_header_len = F_header_len_v

	F_frame_len_v := _byteBuf.Read_uint32()
	_instance.F_frame_len = F_frame_len_v

	F_frame_type_v := _byteBuf.Read_uint16()
	_instance.F_frame_type = F_frame_type_v

	F_version_v := _byteBuf.Read_uint32()
	_instance.F_version = F_version_v

	F_device_sn_v := _byteBuf.Read_uint32()
	_instance.F_device_sn = F_device_sn_v

	F_count_v := _byteBuf.Read_uint32()
	_instance.F_count = F_count_v

	F_timestamp_v := _byteBuf.Read_float64()
	_instance.F_timestamp = time.UnixMilli(int64(F_timestamp_v * 1000))
	F_fps_v := _byteBuf.Read_uint16()
	_instance.F_fps = float32(F_fps_v) / 10

	F_dev_lon_v := _byteBuf.Read_uint32()
	_instance.F_dev_lon = float64(F_dev_lon_v) / 10000000

	F_dev_lat_v := _byteBuf.Read_uint32()
	_instance.F_dev_lat = float64(F_dev_lat_v) / 10000000

	F_dev_alt_v := _byteBuf.Read_uint32()
	_instance.F_dev_alt = F_dev_alt_v

	F_dev_azimuth_v := _byteBuf.Read_uint16()
	_instance.F_dev_azimuth = float32(F_dev_azimuth_v) / 100

	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Msg_header) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	F_header_arr := _instance.F_header
	_byteBuf.Write_slice_uint8(F_header_arr[:])
	_byteBuf.Write_uint16(_instance.F_header_len)
	_byteBuf.Write_uint32(_instance.F_frame_len)
	_byteBuf.Write_uint16(_instance.F_frame_type)
	_byteBuf.Write_uint32(_instance.F_version)
	_byteBuf.Write_uint32(_instance.F_device_sn)
	_byteBuf.Write_uint32(_instance.F_count)
	F_timestamp_v := _instance.F_timestamp
	_byteBuf.Write_float64(float64(F_timestamp_v.UnixMilli()) / 1000)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_fps * 10)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_dev_lon * 10000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_dev_lat * 10000000)))
	_byteBuf.Write_uint32(_instance.F_dev_alt)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_dev_azimuth * 100)))
	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Msg_tailer struct {
	F_check_sum uint32   `json:"check_sum"`
	F_tail      [4]uint8 `json:"tail"`
}

func To_Msg_tailer(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_tailer {
	return (*Msg_tailer)(unsafe.Pointer(unsafe.SliceData(_byteBuf.Read_slice_uint8(8))))
}
func (__instance *Msg_tailer) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_slice_uint8(*(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(__instance)),
		Len:  8,
		Cap:  8,
	})))
}

type Net_infos struct {
	F_net_name     string   `json:"net_name"`
	F_ipv4_addr    [4]uint8 `json:"ipv4_addr"`
	F_ipv4_mask    [4]uint8 `json:"ipv4_mask"`
	F_ipv4_gateway [4]uint8 `json:"ipv4_gateway"`
	F_send_rate    int64    `json:"send_rate"`
	F_recv_rate    int64    `json:"recv_rate"`
}

func To_Net_infos(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Net_infos {
	_instance := Net_infos{}
	F_net_name_len := 16
	F_net_name_v := _byteBuf.Read_slice_uint8(F_net_name_len)
	F_net_name_count := 0
	for i := F_net_name_len - 1; i >= 0; i-- {
		if F_net_name_v[i] == 0 {
			F_net_name_count++
		} else {
			break
		}
	}
	_instance.F_net_name = string(F_net_name_v[:(F_net_name_len - F_net_name_count)])

	F_ipv4_addr_arr := [4]uint8(_byteBuf.Read_slice_uint8(4))
	_instance.F_ipv4_addr = F_ipv4_addr_arr
	F_ipv4_mask_arr := [4]uint8(_byteBuf.Read_slice_uint8(4))
	_instance.F_ipv4_mask = F_ipv4_mask_arr
	F_ipv4_gateway_arr := [4]uint8(_byteBuf.Read_slice_uint8(4))
	_instance.F_ipv4_gateway = F_ipv4_gateway_arr
	F_send_rate_v := _byteBuf.Read_int64()
	_instance.F_send_rate = F_send_rate_v

	F_recv_rate_v := _byteBuf.Read_int64()
	_instance.F_recv_rate = F_recv_rate_v

	return &_instance
}

func (__instance *Net_infos) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_net_name_len := 16
	F_net_name_v := []byte(_instance.F_net_name)
	_byteBuf.Write_slice_uint8(F_net_name_v)
	_byteBuf.Write_zero(F_net_name_len - len(F_net_name_v))
	F_ipv4_addr_arr := _instance.F_ipv4_addr
	_byteBuf.Write_slice_uint8(F_ipv4_addr_arr[:])
	F_ipv4_mask_arr := _instance.F_ipv4_mask
	_byteBuf.Write_slice_uint8(F_ipv4_mask_arr[:])
	F_ipv4_gateway_arr := _instance.F_ipv4_gateway
	_byteBuf.Write_slice_uint8(F_ipv4_gateway_arr[:])
	_byteBuf.Write_int64(_instance.F_send_rate)
	_byteBuf.Write_int64(_instance.F_recv_rate)
}

type Road2_info struct {
	F_road_id         uint32             `json:"road_id"`
	F_road_type       uint8              `json:"road_type"`
	F_road_lon        float64            `json:"road_lon"`
	F_road_lat        float64            `json:"road_lat"`
	F_road_alt        uint32             `json:"road_alt"`
	F_lane_count      uint32             `json:"lane_count"`
	F_lane_info_array []*Road2_info_lane `json:"lane_info_array"`
}

func To_Road2_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Road2_info {
	_instance := Road2_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_road_id_v := _byteBuf.Read_uint32()
	_instance.F_road_id = F_road_id_v

	F_road_type_v := _byteBuf.Read_uint8()
	_instance.F_road_type = F_road_type_v

	F_road_lon_v := _byteBuf.Read_uint32()
	_instance.F_road_lon = float64(F_road_lon_v) / 10000000

	F_road_lat_v := _byteBuf.Read_uint32()
	_instance.F_road_lat = float64(F_road_lat_v) / 10000000

	F_road_alt_v := _byteBuf.Read_uint32()
	_instance.F_road_alt = F_road_alt_v

	F_lane_count_v := _byteBuf.Read_uint32()
	_instance.F_lane_count = F_lane_count_v

	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_lane_info_array_len := (int)(F_lane_count_v)
	F_lane_info_array_arr := make([]*Road2_info_lane, F_lane_info_array_len, F_lane_info_array_len)
	for i := 0; i < F_lane_info_array_len; i++ {
		F_lane_info_array_arr[i] = To_Road2_info_lane(_byteBuf, nil)
	}
	_instance.F_lane_info_array = F_lane_info_array_arr
	return &_instance
}

func (__instance *Road2_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_road_id)
	_byteBuf.Write_uint8(_instance.F_road_type)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_road_lon * 10000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_road_lat * 10000000)))
	_byteBuf.Write_uint32(_instance.F_road_alt)
	_byteBuf.Write_uint32(_instance.F_lane_count)
	F_reserved_len := 64
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_lane_info_array_arr := _instance.F_lane_info_array
	for i := 0; i < len(F_lane_info_array_arr); i++ {
		F_lane_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Road2_info_lane struct {
	F_lane_id           uint32                        `json:"lane_id"`
	F_lane_azimuth      float32                       `json:"lane_azimuth"`
	F_lane_canalization uint8                         `json:"lane_canalization"`
	F_area_point_count  uint16                        `json:"area_point_count"`
	F_area_point_array  []*Road2_info_lane_area_point `json:"area_point_array"`
}

func To_Road2_info_lane(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Road2_info_lane {
	_instance := Road2_info_lane{}
	_start_index := _byteBuf.ReaderIndex()
	F_lane_id_v := _byteBuf.Read_uint32()
	_instance.F_lane_id = F_lane_id_v

	F_lane_azimuth_v := _byteBuf.Read_uint16()
	_instance.F_lane_azimuth = float32(F_lane_azimuth_v) / 100

	F_lane_canalization_v := _byteBuf.Read_uint8()
	_instance.F_lane_canalization = F_lane_canalization_v

	F_area_point_count_v := _byteBuf.Read_uint16()
	_instance.F_area_point_count = F_area_point_count_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	F_area_point_array_len := (int)(F_area_point_count_v)
	F_area_point_array_arr := make([]*Road2_info_lane_area_point, F_area_point_array_len, F_area_point_array_len)
	for i := 0; i < F_area_point_array_len; i++ {
		F_area_point_array_arr[i] = To_Road2_info_lane_area_point(_byteBuf, nil)
	}
	_instance.F_area_point_array = F_area_point_array_arr
	return &_instance
}

func (__instance *Road2_info_lane) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_lane_id)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_lane_azimuth * 100)))
	_byteBuf.Write_uint8(_instance.F_lane_canalization)
	_byteBuf.Write_uint16(_instance.F_area_point_count)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

	F_area_point_array_arr := _instance.F_area_point_array
	for i := 0; i < len(F_area_point_array_arr); i++ {
		F_area_point_array_arr[i].Write(_byteBuf, nil)
	}
}

type Road2_info_lane_area_point struct {
	F_area_point_id  uint16  `json:"area_point_id"`
	F_area_point_lon float64 `json:"area_point_lon"`
	F_area_point_lat float64 `json:"area_point_lat"`
	F_area_point_alt uint32  `json:"area_point_alt"`
}

func To_Road2_info_lane_area_point(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Road2_info_lane_area_point {
	_instance := Road2_info_lane_area_point{}
	F_area_point_id_v := _byteBuf.Read_uint16()
	_instance.F_area_point_id = F_area_point_id_v

	F_area_point_lon_v := _byteBuf.Read_uint32()
	_instance.F_area_point_lon = float64(F_area_point_lon_v) / 10000000

	F_area_point_lat_v := _byteBuf.Read_uint32()
	_instance.F_area_point_lat = float64(F_area_point_lat_v) / 10000000

	F_area_point_alt_v := _byteBuf.Read_uint32()
	_instance.F_area_point_alt = F_area_point_alt_v

	return &_instance
}

func (__instance *Road2_info_lane_area_point) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_area_point_id)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_area_point_lon * 10000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_area_point_lat * 10000000)))
	_byteBuf.Write_uint32(_instance.F_area_point_alt)
}

type Road_info struct {
	F_road_id         uint32            `json:"road_id"`
	F_lane_count      uint32            `json:"lane_count"`
	F_lane_info_array []*Road_info_lane `json:"lane_info_array"`
}

func To_Road_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Road_info {
	_instance := Road_info{}
	F_road_id_v := _byteBuf.Read_uint32()
	_instance.F_road_id = F_road_id_v

	F_lane_count_v := _byteBuf.Read_uint32()
	_instance.F_lane_count = F_lane_count_v

	F_lane_info_array_len := (int)(F_lane_count_v)
	F_lane_info_array_arr := make([]*Road_info_lane, F_lane_info_array_len, F_lane_info_array_len)
	for i := 0; i < F_lane_info_array_len; i++ {
		F_lane_info_array_arr[i] = To_Road_info_lane(_byteBuf, nil)
	}
	_instance.F_lane_info_array = F_lane_info_array_arr
	return &_instance
}

func (__instance *Road_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint32(_instance.F_road_id)
	_byteBuf.Write_uint32(_instance.F_lane_count)
	F_lane_info_array_arr := _instance.F_lane_info_array
	for i := 0; i < len(F_lane_info_array_arr); i++ {
		F_lane_info_array_arr[i].Write(_byteBuf, nil)
	}
}

type Road_info_lane struct {
	F_lane_id           uint32                   `json:"lane_id"`
	F_target_count      uint16                   `json:"target_count"`
	F_lane_target_array []*Road_info_lane_target `json:"lane_target_array"`
}

func To_Road_info_lane(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Road_info_lane {
	_instance := Road_info_lane{}
	F_lane_id_v := _byteBuf.Read_uint32()
	_instance.F_lane_id = F_lane_id_v

	F_target_count_v := _byteBuf.Read_uint16()
	_instance.F_target_count = F_target_count_v

	F_lane_target_array_len := (int)(F_target_count_v)
	F_lane_target_array_arr := make([]*Road_info_lane_target, F_lane_target_array_len, F_lane_target_array_len)
	for i := 0; i < F_lane_target_array_len; i++ {
		F_lane_target_array_arr[i] = To_Road_info_lane_target(_byteBuf, nil)
	}
	_instance.F_lane_target_array = F_lane_target_array_arr
	return &_instance
}

func (__instance *Road_info_lane) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint32(_instance.F_lane_id)
	_byteBuf.Write_uint16(_instance.F_target_count)
	F_lane_target_array_arr := _instance.F_lane_target_array
	for i := 0; i < len(F_lane_target_array_arr); i++ {
		F_lane_target_array_arr[i].Write(_byteBuf, nil)
	}
}

type Road_info_lane_target struct {
	F_track_id uint32 `json:"track_id"`
	F_lane_dis uint32 `json:"lane_dis"`
	F_lane_v   uint32 `json:"lane_v"`
}

func To_Road_info_lane_target(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Road_info_lane_target {
	_instance := Road_info_lane_target{}
	_start_index := _byteBuf.ReaderIndex()
	F_track_id_v := _byteBuf.Read_uint32()
	_instance.F_track_id = F_track_id_v

	F_lane_dis_v := _byteBuf.Read_uint32()
	_instance.F_lane_dis = F_lane_dis_v

	F_lane_v_v := _byteBuf.Read_uint32()
	_instance.F_lane_v = F_lane_v_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Road_info_lane_target) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_track_id)
	_byteBuf.Write_uint32(_instance.F_lane_dis)
	_byteBuf.Write_uint32(_instance.F_lane_v)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Sensor_body_camera struct {
	F_pixel           uint16  `json:"pixel"`
	F_focal           uint16  `json:"focal"`
	F_hori_view_angle float32 `json:"hori_view_angle"`
	F_vert_view_angle float32 `json:"vert_view_angle"`
}

func To_Sensor_body_camera(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Sensor_body_camera {
	_instance := Sensor_body_camera{}
	_start_index := _byteBuf.ReaderIndex()
	F_pixel_v := _byteBuf.Read_uint16()
	_instance.F_pixel = F_pixel_v

	F_focal_v := _byteBuf.Read_uint16()
	_instance.F_focal = F_focal_v

	F_hori_view_angle_v := _byteBuf.Read_uint16()
	_instance.F_hori_view_angle = float32(F_hori_view_angle_v) / 100

	F_vert_view_angle_v := _byteBuf.Read_uint16()
	_instance.F_vert_view_angle = float32(F_vert_view_angle_v) / 100

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Sensor_body_camera) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.F_pixel)
	_byteBuf.Write_uint16(_instance.F_focal)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_hori_view_angle * 100)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_vert_view_angle * 100)))
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Sensor_body_lidar struct {
	F_distance        uint32  `json:"distance"`
	F_line_count      uint16  `json:"line_count"`
	F_hori_view_angle float32 `json:"hori_view_angle"`
	F_vert_view_angle float32 `json:"vert_view_angle"`
}

func To_Sensor_body_lidar(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Sensor_body_lidar {
	_instance := Sensor_body_lidar{}
	_start_index := _byteBuf.ReaderIndex()
	F_distance_v := _byteBuf.Read_uint32()
	_instance.F_distance = F_distance_v

	F_line_count_v := _byteBuf.Read_uint16()
	_instance.F_line_count = F_line_count_v

	F_hori_view_angle_v := _byteBuf.Read_uint16()
	_instance.F_hori_view_angle = float32(F_hori_view_angle_v) / 100

	F_vert_view_angle_v := _byteBuf.Read_uint16()
	_instance.F_vert_view_angle = float32(F_vert_view_angle_v) / 100

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Sensor_body_lidar) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_distance)
	_byteBuf.Write_uint16(_instance.F_line_count)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_hori_view_angle * 100)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_vert_view_angle * 100)))
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Sensor_body_millimeter_wave_radar struct {
	F_distance int32 `json:"distance"`
}

func To_Sensor_body_millimeter_wave_radar(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Sensor_body_millimeter_wave_radar {
	_instance := Sensor_body_millimeter_wave_radar{}
	_start_index := _byteBuf.ReaderIndex()
	F_distance_v := _byteBuf.Read_int32()
	_instance.F_distance = F_distance_v

	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Sensor_body_millimeter_wave_radar) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_int32(_instance.F_distance)
	F_reserved_len := 32
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Sensor_info struct {
	F_sensor_id      uint32  `json:"sensor_id"`
	F_sensor_sn      uint32  `json:"sensor_sn"`
	F_sensor_type    uint8   `json:"sensor_type"`
	F_data_type      uint8   `json:"data_type"`
	F_sensor_lon     float64 `json:"sensor_lon"`
	F_sensor_lat     float64 `json:"sensor_lat"`
	F_sensor_alt     uint32  `json:"sensor_alt"`
	F_sensor_azimuth float32 `json:"sensor_azimuth"`
	F_sensor_pitch   float32 `json:"sensor_pitch"`
	F_sensor_roll    float32 `json:"sensor_roll"`
	F_sensor_status  uint8   `json:"sensor_status"`
	F_sensor_body    any     `json:"sensor_body"`
}

func To_Sensor_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Sensor_info {
	_instance := Sensor_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_sensor_id_v := _byteBuf.Read_uint32()
	_instance.F_sensor_id = F_sensor_id_v

	F_sensor_sn_v := _byteBuf.Read_uint32()
	_instance.F_sensor_sn = F_sensor_sn_v

	F_sensor_type_v := _byteBuf.Read_uint8()
	_instance.F_sensor_type = F_sensor_type_v

	F_data_type_v := _byteBuf.Read_uint8()
	_instance.F_data_type = F_data_type_v

	F_sensor_lon_v := _byteBuf.Read_uint32()
	_instance.F_sensor_lon = float64(F_sensor_lon_v) / 10000000

	F_sensor_lat_v := _byteBuf.Read_uint32()
	_instance.F_sensor_lat = float64(F_sensor_lat_v) / 10000000

	F_sensor_alt_v := _byteBuf.Read_uint32()
	_instance.F_sensor_alt = F_sensor_alt_v

	F_sensor_azimuth_v := _byteBuf.Read_uint16()
	_instance.F_sensor_azimuth = float32(F_sensor_azimuth_v) / 100

	F_sensor_pitch_v := _byteBuf.Read_uint16()
	_instance.F_sensor_pitch = float32(F_sensor_pitch_v) / 100

	F_sensor_roll_v := _byteBuf.Read_uint16()
	_instance.F_sensor_roll = float32(F_sensor_roll_v) / 100

	F_sensor_status_v := _byteBuf.Read_uint8()
	_instance.F_sensor_status = F_sensor_status_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_sensor_body = To_F_sensor_body(_byteBuf, _parseContext)
	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Sensor_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_sensor_id)
	_byteBuf.Write_uint32(_instance.F_sensor_sn)
	_byteBuf.Write_uint8(_instance.F_sensor_type)
	_byteBuf.Write_uint8(_instance.F_data_type)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_sensor_lon * 10000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_sensor_lat * 10000000)))
	_byteBuf.Write_uint32(_instance.F_sensor_alt)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_sensor_azimuth * 100)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_sensor_pitch * 100)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_sensor_roll * 100)))
	_byteBuf.Write_uint8(_instance.F_sensor_status)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_sensor_body(_byteBuf, _instance.F_sensor_body, _parseContext)
	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Target_info struct {
	F_target_num       uint16  `json:"target_num"`
	F_track_id         uint32  `json:"track_id"`
	F_lane_id          uint32  `json:"lane_id"`
	F_lon              float64 `json:"lon"`
	F_lat              float64 `json:"lat"`
	F_alt              uint32  `json:"alt"`
	F_dev_x            int32   `json:"dev_x"`
	F_dev_y            int32   `json:"dev_y"`
	F_dev_z            int32   `json:"dev_z"`
	F_azimuth_angle    float32 `json:"azimuth_angle"`
	F_dev_vx           int32   `json:"dev_vx"`
	F_dev_vy           int32   `json:"dev_vy"`
	F_dev_vz           int32   `json:"dev_vz"`
	F_len              uint16  `json:"len"`
	F_width            uint16  `json:"width"`
	F_height           uint16  `json:"height"`
	F_clazz            uint8   `json:"clazz"`
	F_class_cfd        uint8   `json:"class_cfd"`
	F_img_x            uint16  `json:"img_x"`
	F_img_y            uint16  `json:"img_y"`
	F_img_len          uint16  `json:"img_len"`
	F_img_width        uint16  `json:"img_width"`
	F_img_height       uint16  `json:"img_height"`
	F_img_direc_len    int16   `json:"img_direc_len"`
	F_img_direc_width  int16   `json:"img_direc_width"`
	F_img_direc_height int16   `json:"img_direc_height"`
	F_if_extras        int8    `json:"if_extras"`
	F_extras           any     `json:"extras"`
}

func To_Target_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Target_info {
	_instance := Target_info{}
	_start_index := _byteBuf.ReaderIndex()
	F_target_num_v := _byteBuf.Read_uint16()
	_instance.F_target_num = F_target_num_v

	F_track_id_v := _byteBuf.Read_uint32()
	_instance.F_track_id = F_track_id_v

	F_lane_id_v := _byteBuf.Read_uint32()
	_instance.F_lane_id = F_lane_id_v

	F_lon_v := _byteBuf.Read_uint32()
	_instance.F_lon = float64(F_lon_v) / 10000000

	F_lat_v := _byteBuf.Read_uint32()
	_instance.F_lat = float64(F_lat_v) / 10000000

	F_alt_v := _byteBuf.Read_uint32()
	_instance.F_alt = F_alt_v

	F_dev_x_v := _byteBuf.Read_int32()
	_instance.F_dev_x = F_dev_x_v

	F_dev_y_v := _byteBuf.Read_int32()
	_instance.F_dev_y = F_dev_y_v

	F_dev_z_v := _byteBuf.Read_int32()
	_instance.F_dev_z = F_dev_z_v

	F_azimuth_angle_v := _byteBuf.Read_uint16()
	_instance.F_azimuth_angle = float32(F_azimuth_angle_v) / 100

	F_dev_vx_v := _byteBuf.Read_int32()
	_instance.F_dev_vx = F_dev_vx_v

	F_dev_vy_v := _byteBuf.Read_int32()
	_instance.F_dev_vy = F_dev_vy_v

	F_dev_vz_v := _byteBuf.Read_int32()
	_instance.F_dev_vz = F_dev_vz_v

	F_len_v := _byteBuf.Read_uint16()
	_instance.F_len = F_len_v

	F_width_v := _byteBuf.Read_uint16()
	_instance.F_width = F_width_v

	F_height_v := _byteBuf.Read_uint16()
	_instance.F_height = F_height_v

	F_clazz_v := _byteBuf.Read_uint8()
	_instance.F_clazz = F_clazz_v

	F_class_cfd_v := _byteBuf.Read_uint8()
	_instance.F_class_cfd = F_class_cfd_v

	F_img_x_v := _byteBuf.Read_uint16()
	_instance.F_img_x = F_img_x_v

	F_img_y_v := _byteBuf.Read_uint16()
	_instance.F_img_y = F_img_y_v

	F_img_len_v := _byteBuf.Read_uint16()
	_instance.F_img_len = F_img_len_v

	F_img_width_v := _byteBuf.Read_uint16()
	_instance.F_img_width = F_img_width_v

	F_img_height_v := _byteBuf.Read_uint16()
	_instance.F_img_height = F_img_height_v

	F_img_direc_len_v := _byteBuf.Read_int16()
	_instance.F_img_direc_len = F_img_direc_len_v / 100

	F_img_direc_width_v := _byteBuf.Read_int16()
	_instance.F_img_direc_width = F_img_direc_width_v / 100

	F_img_direc_height_v := _byteBuf.Read_int16()
	_instance.F_img_direc_height = F_img_direc_height_v / 100

	F_if_extras_v := _byteBuf.Read_int8()
	_instance.F_if_extras = F_if_extras_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_extras = To_F_extras(_byteBuf, _parseContext)
	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.ReaderIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Skip(F_reserved_skipLen)
	}
	return &_instance
}

func (__instance *Target_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.F_target_num)
	_byteBuf.Write_uint32(_instance.F_track_id)
	_byteBuf.Write_uint32(_instance.F_lane_id)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lon * 10000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lat * 10000000)))
	_byteBuf.Write_uint32(_instance.F_alt)
	_byteBuf.Write_int32(_instance.F_dev_x)
	_byteBuf.Write_int32(_instance.F_dev_y)
	_byteBuf.Write_int32(_instance.F_dev_z)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_azimuth_angle * 100)))
	_byteBuf.Write_int32(_instance.F_dev_vx)
	_byteBuf.Write_int32(_instance.F_dev_vy)
	_byteBuf.Write_int32(_instance.F_dev_vz)
	_byteBuf.Write_uint16(_instance.F_len)
	_byteBuf.Write_uint16(_instance.F_width)
	_byteBuf.Write_uint16(_instance.F_height)
	_byteBuf.Write_uint8(_instance.F_clazz)
	_byteBuf.Write_uint8(_instance.F_class_cfd)
	_byteBuf.Write_uint16(_instance.F_img_x)
	_byteBuf.Write_uint16(_instance.F_img_y)
	_byteBuf.Write_uint16(_instance.F_img_len)
	_byteBuf.Write_uint16(_instance.F_img_width)
	_byteBuf.Write_uint16(_instance.F_img_height)
	_byteBuf.Write_int16(_instance.F_img_direc_len * 100)
	_byteBuf.Write_int16(_instance.F_img_direc_width * 100)
	_byteBuf.Write_int16(_instance.F_img_direc_height * 100)
	_byteBuf.Write_int8(_instance.F_if_extras)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_extras(_byteBuf, _instance.F_extras, _parseContext)
	F_reserved_len := 128
	F_reserved_skipLen := F_reserved_len + _start_index - _byteBuf.WriterIndex()
	if F_reserved_skipLen > 0 {
		_byteBuf.Write_zero(F_reserved_skipLen)
	}

}

type Target_info_extras_barrier struct {
	F_type uint16 `json:"type"`
}

func To_Target_info_extras_barrier(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Target_info_extras_barrier {
	return (*Target_info_extras_barrier)(unsafe.Pointer(unsafe.SliceData(_byteBuf.Read_slice_uint8(2))))
}
func (__instance *Target_info_extras_barrier) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_slice_uint8(*(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(__instance)),
		Len:  2,
		Cap:  2,
	})))
}

type Target_info_extras_car struct {
	F_type      uint32 `json:"type"`
	F_lic_plate string `json:"lic_plate"`
	F_color     uint16 `json:"color"`
	F_status    int8   `json:"status"`
}

func To_Target_info_extras_car(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Target_info_extras_car {
	_instance := Target_info_extras_car{}
	F_type_v := _byteBuf.Read_uint32()
	_instance.F_type = F_type_v

	F_lic_plate_len := 16
	F_lic_plate_v := _byteBuf.Read_slice_uint8(F_lic_plate_len)
	F_lic_plate_count := 0
	for i := F_lic_plate_len - 1; i >= 0; i-- {
		if F_lic_plate_v[i] == 0 {
			F_lic_plate_count++
		} else {
			break
		}
	}
	_instance.F_lic_plate = string(F_lic_plate_v[:(F_lic_plate_len - F_lic_plate_count)])

	F_color_v := _byteBuf.Read_uint16()
	_instance.F_color = F_color_v

	F_status_v := _byteBuf.Read_int8()
	_instance.F_status = F_status_v

	return &_instance
}

func (__instance *Target_info_extras_car) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint32(_instance.F_type)
	F_lic_plate_len := 16
	F_lic_plate_v := []byte(_instance.F_lic_plate)
	_byteBuf.Write_slice_uint8(F_lic_plate_v)
	_byteBuf.Write_zero(F_lic_plate_len - len(F_lic_plate_v))
	_byteBuf.Write_uint16(_instance.F_color)
	_byteBuf.Write_int8(_instance.F_status)
}

type Target_info_extras_person struct {
	F_type     uint8 `json:"type"`
	F_behavior uint8 `json:"behavior"`
}

func To_Target_info_extras_person(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Target_info_extras_person {
	return (*Target_info_extras_person)(unsafe.Pointer(unsafe.SliceData(_byteBuf.Read_slice_uint8(2))))
}
func (__instance *Target_info_extras_person) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_slice_uint8(*(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(__instance)),
		Len:  2,
		Cap:  2,
	})))
}

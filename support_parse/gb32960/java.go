package gb32960

import (
	"bcd-util/support_parse/parse"
	"time"
)

var _Location0 = time.FixedZone("_Location0", 28800)

type MotorData struct {
	F_no                    uint8   `json:"no"`
	F_status                uint8   `json:"status"`
	F_controllerTemperature int16   `json:"controllerTemperature"`
	F_rotateSpeed           int32   `json:"rotateSpeed"`
	F_rotateRectangle       float32 `json:"rotateRectangle"`
	F_temperature           int16   `json:"temperature"`
	F_inputVoltage          float32 `json:"inputVoltage"`
	F_current               float32 `json:"current"`
}

func To_MotorData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *MotorData {
	_instance := MotorData{}
	F_no_v := _byteBuf.Read_uint8()
	_instance.F_no = F_no_v

	F_status_v := _byteBuf.Read_uint8()
	_instance.F_status = F_status_v

	F_controllerTemperature_v := _byteBuf.Read_uint8()
	_instance.F_controllerTemperature = int16(F_controllerTemperature_v) - 40

	F_rotateSpeed_v := _byteBuf.Read_uint16()
	_instance.F_rotateSpeed = int32(F_rotateSpeed_v) - 20000

	F_rotateRectangle_v := _byteBuf.Read_uint16()
	_instance.F_rotateRectangle = float32(F_rotateRectangle_v)/10 - 2000

	F_temperature_v := _byteBuf.Read_uint8()
	_instance.F_temperature = int16(F_temperature_v) - 40

	F_inputVoltage_v := _byteBuf.Read_uint16()
	_instance.F_inputVoltage = float32(F_inputVoltage_v) / 10

	F_current_v := _byteBuf.Read_uint16()
	_instance.F_current = float32(F_current_v)/10 - 1000

	return &_instance
}

func (__instance *MotorData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_no)
	_byteBuf.Write_uint8(_instance.F_status)
	_byteBuf.Write_uint8(uint8((_instance.F_controllerTemperature + 40)))
	_byteBuf.Write_uint16(uint16((_instance.F_rotateSpeed + 20000)))
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_rotateRectangle + 2000) * 10)))
	_byteBuf.Write_uint8(uint8((_instance.F_temperature + 40)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_inputVoltage * 10)))
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_current + 1000) * 10)))
}

type Packet struct {
	F_header        [2]int8 `json:"header"`
	F_flag          uint8   `json:"flag"`
	F_replyFlag     uint8   `json:"replyFlag"`
	F_vin           string  `json:"vin"`
	F_encodeWay     uint8   `json:"encodeWay"`
	F_contentLength uint16  `json:"contentLength"`
	F_data          any     `json:"data"`
	F_code          int8    `json:"code"`
}

func To_Packet(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Packet {
	_instance := Packet{}
	F_header_arr := [2]int8(_byteBuf.Read_slice_int8(2))
	_instance.F_header = F_header_arr
	F_flag_v := _byteBuf.Read_uint8()
	_instance.F_flag = F_flag_v

	F_replyFlag_v := _byteBuf.Read_uint8()
	_instance.F_replyFlag = F_replyFlag_v

	F_vin_len := 17
	F_vin_v := _byteBuf.Read_slice_uint8(F_vin_len)
	F_vin_count := 0
	for i := F_vin_len - 1; i >= 0; i-- {
		if F_vin_v[i] == 0 {
			F_vin_count++
		} else {
			break
		}
	}
	_instance.F_vin = string(F_vin_v[:(F_vin_len - F_vin_count)])

	F_encodeWay_v := _byteBuf.Read_uint8()
	_instance.F_encodeWay = F_encodeWay_v

	F_contentLength_v := _byteBuf.Read_uint16()
	_instance.F_contentLength = F_contentLength_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_data = To_F_data(_byteBuf, _parseContext)
	F_code_v := _byteBuf.Read_int8()
	_instance.F_code = F_code_v

	return &_instance
}

func (__instance *Packet) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_header_arr := _instance.F_header
	_byteBuf.Write_slice_int8(F_header_arr[:])
	_byteBuf.Write_uint8(_instance.F_flag)
	_byteBuf.Write_uint8(_instance.F_replyFlag)
	F_vin_len := 17
	F_vin_v := []byte(_instance.F_vin)
	_byteBuf.Write_slice_uint8(F_vin_v)
	_byteBuf.Write_zero(F_vin_len - len(F_vin_v))
	_byteBuf.Write_uint8(_instance.F_encodeWay)
	_byteBuf.Write_uint16(_instance.F_contentLength)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_data(_byteBuf, _instance.F_data, _parseContext)
	_byteBuf.Write_int8(_instance.F_code)
}

type PlatformLoginData struct {
	F_collectTime time.Time `json:"collectTime"`
	F_sn          uint16    `json:"sn"`
	F_username    string    `json:"username"`
	F_password    string    `json:"password"`
	F_encode      uint8     `json:"encode"`
}

func To_PlatformLoginData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *PlatformLoginData {
	_instance := PlatformLoginData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _Location0)
	F_sn_v := _byteBuf.Read_uint16()
	_instance.F_sn = F_sn_v

	F_username_len := 12
	F_username_v := _byteBuf.Read_slice_uint8(F_username_len)
	F_username_count := 0
	for i := F_username_len - 1; i >= 0; i-- {
		if F_username_v[i] == 0 {
			F_username_count++
		} else {
			break
		}
	}
	_instance.F_username = string(F_username_v[:(F_username_len - F_username_count)])

	F_password_len := 20
	F_password_v := _byteBuf.Read_slice_uint8(F_password_len)
	F_password_count := 0
	for i := F_password_len - 1; i >= 0; i-- {
		if F_password_v[i] == 0 {
			F_password_count++
		} else {
			break
		}
	}
	_instance.F_password = string(F_password_v[:(F_password_len - F_password_count)])

	F_encode_v := _byteBuf.Read_uint8()
	_instance.F_encode = F_encode_v

	return &_instance
}

func (__instance *PlatformLoginData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.F_sn)
	F_username_len := 12
	F_username_v := []byte(_instance.F_username)
	_byteBuf.Write_slice_uint8(F_username_v)
	_byteBuf.Write_zero(F_username_len - len(F_username_v))
	F_password_len := 20
	F_password_v := []byte(_instance.F_password)
	_byteBuf.Write_slice_uint8(F_password_v)
	_byteBuf.Write_zero(F_password_len - len(F_password_v))
	_byteBuf.Write_uint8(_instance.F_encode)
}

type PlatformLogoutData struct {
	F_collectTime time.Time `json:"collectTime"`
	F_sn          uint16    `json:"sn"`
}

func To_PlatformLogoutData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *PlatformLogoutData {
	_instance := PlatformLogoutData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _Location0)
	F_sn_v := _byteBuf.Read_uint16()
	_instance.F_sn = F_sn_v

	return &_instance
}

func (__instance *PlatformLogoutData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.F_sn)
}

type StorageTemperatureData struct {
	F_no           uint8   `json:"no"`
	F_num          uint16  `json:"num"`
	F_temperatures []int16 `json:"temperatures"`
}

func To_StorageTemperatureData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *StorageTemperatureData {
	_instance := StorageTemperatureData{}
	F_no_v := _byteBuf.Read_uint8()
	_instance.F_no = F_no_v

	F_num_v := _byteBuf.Read_uint16()
	_instance.F_num = F_num_v

	F_temperatures_len := (int)(F_num_v)
	F_temperatures_arr := make([]int16, F_temperatures_len, F_temperatures_len)
	for i := 0; i < F_temperatures_len; i++ {
		e := _byteBuf.Read_uint8()
		F_temperatures_arr[i] = int16(e) - 40
	}
	_instance.F_temperatures = F_temperatures_arr
	return &_instance
}

func (__instance *StorageTemperatureData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_no)
	_byteBuf.Write_uint16(_instance.F_num)
	F_temperatures_arr := _instance.F_temperatures
	for i := 0; i < len(F_temperatures_arr); i++ {
		_byteBuf.Write_uint8(uint8((F_temperatures_arr[i] + 40)))
	}
}

type StorageVoltageData struct {
	F_no            uint8     `json:"no"`
	F_voltage       float32   `json:"voltage"`
	F_current       float32   `json:"current"`
	F_total         uint16    `json:"total"`
	F_frameNo       uint16    `json:"frameNo"`
	F_frameTotal    uint8     `json:"frameTotal"`
	F_singleVoltage []float32 `json:"singleVoltage"`
}

func To_StorageVoltageData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *StorageVoltageData {
	_instance := StorageVoltageData{}
	F_no_v := _byteBuf.Read_uint8()
	_instance.F_no = F_no_v

	F_voltage_v := _byteBuf.Read_uint16()
	_instance.F_voltage = float32(F_voltage_v) / 10

	F_current_v := _byteBuf.Read_uint16()
	_instance.F_current = float32(F_current_v)/10 - 1000

	F_total_v := _byteBuf.Read_uint16()
	_instance.F_total = F_total_v

	F_frameNo_v := _byteBuf.Read_uint16()
	_instance.F_frameNo = F_frameNo_v

	F_frameTotal_v := _byteBuf.Read_uint8()
	_instance.F_frameTotal = F_frameTotal_v

	F_singleVoltage_len := (int)(F_frameTotal_v)
	F_singleVoltage_arr := make([]float32, F_singleVoltage_len, F_singleVoltage_len)
	for i := 0; i < F_singleVoltage_len; i++ {
		e := _byteBuf.Read_uint16()
		F_singleVoltage_arr[i] = float32(e) / 1000
	}
	_instance.F_singleVoltage = F_singleVoltage_arr
	return &_instance
}

func (__instance *StorageVoltageData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_no)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_voltage * 10)))
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_current + 1000) * 10)))
	_byteBuf.Write_uint16(_instance.F_total)
	_byteBuf.Write_uint16(_instance.F_frameNo)
	_byteBuf.Write_uint8(_instance.F_frameTotal)
	F_singleVoltage_arr := _instance.F_singleVoltage
	for i := 0; i < len(F_singleVoltage_arr); i++ {
		_byteBuf.Write_uint16(uint16(parse.Round(F_singleVoltage_arr[i] * 1000)))
	}
}

type VehicleAlarmData struct {
	F_maxAlarmLevel  uint8    `json:"maxAlarmLevel"`
	F_alarmFlag      int32    `json:"alarmFlag"`
	F_chargeBadNum   uint8    `json:"chargeBadNum"`
	F_chargeBadCodes []uint32 `json:"chargeBadCodes"`
	F_driverBadNum   uint8    `json:"driverBadNum"`
	F_driverBadCodes []uint32 `json:"driverBadCodes"`
	F_engineBadNum   uint8    `json:"engineBadNum"`
	F_engineBadCodes []uint32 `json:"engineBadCodes"`
	F_otherBadNum    uint8    `json:"otherBadNum"`
	F_otherBadCodes  []uint32 `json:"otherBadCodes"`
}

func To_VehicleAlarmData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleAlarmData {
	_instance := VehicleAlarmData{}
	F_maxAlarmLevel_v := _byteBuf.Read_uint8()
	_instance.F_maxAlarmLevel = F_maxAlarmLevel_v

	F_alarmFlag_v := _byteBuf.Read_int32()
	_instance.F_alarmFlag = F_alarmFlag_v

	F_chargeBadNum_v := _byteBuf.Read_uint8()
	_instance.F_chargeBadNum = F_chargeBadNum_v

	F_chargeBadCodes_len := (int)(F_chargeBadNum_v)
	F_chargeBadCodes_arr := make([]uint32, F_chargeBadCodes_len, F_chargeBadCodes_len)
	for i := 0; i < F_chargeBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_chargeBadCodes_arr[i] = e
	}
	_instance.F_chargeBadCodes = F_chargeBadCodes_arr
	F_driverBadNum_v := _byteBuf.Read_uint8()
	_instance.F_driverBadNum = F_driverBadNum_v

	F_driverBadCodes_len := (int)(F_driverBadNum_v)
	F_driverBadCodes_arr := make([]uint32, F_driverBadCodes_len, F_driverBadCodes_len)
	for i := 0; i < F_driverBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_driverBadCodes_arr[i] = e
	}
	_instance.F_driverBadCodes = F_driverBadCodes_arr
	F_engineBadNum_v := _byteBuf.Read_uint8()
	_instance.F_engineBadNum = F_engineBadNum_v

	F_engineBadCodes_len := (int)(F_engineBadNum_v)
	F_engineBadCodes_arr := make([]uint32, F_engineBadCodes_len, F_engineBadCodes_len)
	for i := 0; i < F_engineBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_engineBadCodes_arr[i] = e
	}
	_instance.F_engineBadCodes = F_engineBadCodes_arr
	F_otherBadNum_v := _byteBuf.Read_uint8()
	_instance.F_otherBadNum = F_otherBadNum_v

	F_otherBadCodes_len := (int)(F_otherBadNum_v)
	F_otherBadCodes_arr := make([]uint32, F_otherBadCodes_len, F_otherBadCodes_len)
	for i := 0; i < F_otherBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_otherBadCodes_arr[i] = e
	}
	_instance.F_otherBadCodes = F_otherBadCodes_arr
	return &_instance
}

func (__instance *VehicleAlarmData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_maxAlarmLevel)
	_byteBuf.Write_int32(_instance.F_alarmFlag)
	_byteBuf.Write_uint8(_instance.F_chargeBadNum)
	F_chargeBadCodes_arr := _instance.F_chargeBadCodes
	for i := 0; i < len(F_chargeBadCodes_arr); i++ {
		_byteBuf.Write_uint32(F_chargeBadCodes_arr[i])
	}
	_byteBuf.Write_uint8(_instance.F_driverBadNum)
	F_driverBadCodes_arr := _instance.F_driverBadCodes
	for i := 0; i < len(F_driverBadCodes_arr); i++ {
		_byteBuf.Write_uint32(F_driverBadCodes_arr[i])
	}
	_byteBuf.Write_uint8(_instance.F_engineBadNum)
	F_engineBadCodes_arr := _instance.F_engineBadCodes
	for i := 0; i < len(F_engineBadCodes_arr); i++ {
		_byteBuf.Write_uint32(F_engineBadCodes_arr[i])
	}
	_byteBuf.Write_uint8(_instance.F_otherBadNum)
	F_otherBadCodes_arr := _instance.F_otherBadCodes
	for i := 0; i < len(F_otherBadCodes_arr); i++ {
		_byteBuf.Write_uint32(F_otherBadCodes_arr[i])
	}
}

type VehicleBaseData struct {
	F_vehicleStatus uint8   `json:"vehicleStatus"`
	F_chargeStatus  uint8   `json:"chargeStatus"`
	F_runMode       uint8   `json:"runMode"`
	F_vehicleSpeed  float32 `json:"vehicleSpeed"`
	F_totalMileage  float64 `json:"totalMileage"`
	F_totalVoltage  float32 `json:"totalVoltage"`
	F_totalCurrent  float32 `json:"totalCurrent"`
	F_soc           uint8   `json:"soc"`
	F_dcStatus      uint8   `json:"dcStatus"`
	F_gearPosition  uint8   `json:"gearPosition"`
	F_resistance    uint16  `json:"resistance"`
	F_pedalVal      uint8   `json:"pedalVal"`
	F_pedalStatus   uint8   `json:"pedalStatus"`
}

func To_VehicleBaseData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleBaseData {
	_instance := VehicleBaseData{}
	F_vehicleStatus_v := _byteBuf.Read_uint8()
	_instance.F_vehicleStatus = F_vehicleStatus_v

	F_chargeStatus_v := _byteBuf.Read_uint8()
	_instance.F_chargeStatus = F_chargeStatus_v

	F_runMode_v := _byteBuf.Read_uint8()
	_instance.F_runMode = F_runMode_v

	F_vehicleSpeed_v := _byteBuf.Read_uint16()
	_instance.F_vehicleSpeed = float32(F_vehicleSpeed_v) / 10

	F_totalMileage_v := _byteBuf.Read_uint32()
	_instance.F_totalMileage = float64(F_totalMileage_v) / 10

	F_totalVoltage_v := _byteBuf.Read_uint16()
	_instance.F_totalVoltage = float32(F_totalVoltage_v) / 10

	F_totalCurrent_v := _byteBuf.Read_uint16()
	_instance.F_totalCurrent = float32(F_totalCurrent_v)/10 - 1000

	F_soc_v := _byteBuf.Read_uint8()
	_instance.F_soc = F_soc_v

	F_dcStatus_v := _byteBuf.Read_uint8()
	_instance.F_dcStatus = F_dcStatus_v

	F_gearPosition_v := _byteBuf.Read_uint8()
	_instance.F_gearPosition = F_gearPosition_v

	F_resistance_v := _byteBuf.Read_uint16()
	_instance.F_resistance = F_resistance_v

	F_pedalVal_v := _byteBuf.Read_uint8()
	_instance.F_pedalVal = F_pedalVal_v

	F_pedalStatus_v := _byteBuf.Read_uint8()
	_instance.F_pedalStatus = F_pedalStatus_v

	return &_instance
}

func (__instance *VehicleBaseData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_vehicleStatus)
	_byteBuf.Write_uint8(_instance.F_chargeStatus)
	_byteBuf.Write_uint8(_instance.F_runMode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_vehicleSpeed * 10)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_totalMileage * 10)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_totalVoltage * 10)))
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_totalCurrent + 1000) * 10)))
	_byteBuf.Write_uint8(_instance.F_soc)
	_byteBuf.Write_uint8(_instance.F_dcStatus)
	_byteBuf.Write_uint8(_instance.F_gearPosition)
	_byteBuf.Write_uint16(_instance.F_resistance)
	_byteBuf.Write_uint8(_instance.F_pedalVal)
	_byteBuf.Write_uint8(_instance.F_pedalStatus)
}

type VehicleEngineData struct {
	F_status uint8   `json:"status"`
	F_speed  uint16  `json:"speed"`
	F_rate   float32 `json:"rate"`
}

func To_VehicleEngineData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleEngineData {
	_instance := VehicleEngineData{}
	F_status_v := _byteBuf.Read_uint8()
	_instance.F_status = F_status_v

	F_speed_v := _byteBuf.Read_uint16()
	_instance.F_speed = F_speed_v

	F_rate_v := _byteBuf.Read_uint16()
	_instance.F_rate = float32(F_rate_v) / 100

	return &_instance
}

func (__instance *VehicleEngineData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_status)
	_byteBuf.Write_uint16(_instance.F_speed)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_rate * 100)))
}

type VehicleFuelBatteryData struct {
	F_voltage              float32 `json:"voltage"`
	F_current              float32 `json:"current"`
	F_consumptionRate      float32 `json:"consumptionRate"`
	F_num                  uint32  `json:"num"`
	F_temperatures         []int16 `json:"temperatures"`
	F_maxTemperature       float32 `json:"maxTemperature"`
	F_maxTemperatureCode   uint8   `json:"maxTemperatureCode"`
	F_maxConcentration     int32   `json:"maxConcentration"`
	F_maxConcentrationCode uint8   `json:"maxConcentrationCode"`
	F_maxPressure          float32 `json:"maxPressure"`
	F_maxPressureCode      uint8   `json:"maxPressureCode"`
	F_dcStatus             uint8   `json:"dcStatus"`
}

func To_VehicleFuelBatteryData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleFuelBatteryData {
	_instance := VehicleFuelBatteryData{}
	F_voltage_v := _byteBuf.Read_uint16()
	_instance.F_voltage = float32(F_voltage_v) / 10

	F_current_v := _byteBuf.Read_uint16()
	_instance.F_current = float32(F_current_v) / 10

	F_consumptionRate_v := _byteBuf.Read_uint16()
	_instance.F_consumptionRate = float32(F_consumptionRate_v) / 100

	F_num_v := _byteBuf.Read_uint32()
	_instance.F_num = F_num_v

	F_temperatures_len := (int)(F_num_v)
	F_temperatures_arr := make([]int16, F_temperatures_len, F_temperatures_len)
	for i := 0; i < F_temperatures_len; i++ {
		e := _byteBuf.Read_uint8()
		F_temperatures_arr[i] = int16(e) - 40
	}
	_instance.F_temperatures = F_temperatures_arr
	F_maxTemperature_v := _byteBuf.Read_uint16()
	_instance.F_maxTemperature = float32(F_maxTemperature_v)/10 - 40

	F_maxTemperatureCode_v := _byteBuf.Read_uint8()
	_instance.F_maxTemperatureCode = F_maxTemperatureCode_v

	F_maxConcentration_v := _byteBuf.Read_uint16()
	_instance.F_maxConcentration = int32(F_maxConcentration_v) - 10000

	F_maxConcentrationCode_v := _byteBuf.Read_uint8()
	_instance.F_maxConcentrationCode = F_maxConcentrationCode_v

	F_maxPressure_v := _byteBuf.Read_uint16()
	_instance.F_maxPressure = float32(F_maxPressure_v) / 10

	F_maxPressureCode_v := _byteBuf.Read_uint8()
	_instance.F_maxPressureCode = F_maxPressureCode_v

	F_dcStatus_v := _byteBuf.Read_uint8()
	_instance.F_dcStatus = F_dcStatus_v

	return &_instance
}

func (__instance *VehicleFuelBatteryData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_voltage * 10)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_current * 10)))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_consumptionRate * 100)))
	_byteBuf.Write_uint32(_instance.F_num)
	F_temperatures_arr := _instance.F_temperatures
	for i := 0; i < len(F_temperatures_arr); i++ {
		_byteBuf.Write_uint8(uint8((F_temperatures_arr[i] + 40)))
	}
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_maxTemperature + 40) * 10)))
	_byteBuf.Write_uint8(_instance.F_maxTemperatureCode)
	_byteBuf.Write_uint16(uint16((_instance.F_maxConcentration + 10000)))
	_byteBuf.Write_uint8(_instance.F_maxConcentrationCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_maxPressure * 10)))
	_byteBuf.Write_uint8(_instance.F_maxPressureCode)
	_byteBuf.Write_uint8(_instance.F_dcStatus)
}

type VehicleLimitValueData struct {
	F_maxVoltageSystemNo     uint8   `json:"maxVoltageSystemNo"`
	F_maxVoltageCode         uint8   `json:"maxVoltageCode"`
	F_maxVoltage             float32 `json:"maxVoltage"`
	F_minVoltageSystemNo     uint8   `json:"minVoltageSystemNo"`
	F_minVoltageCode         uint8   `json:"minVoltageCode"`
	F_minVoltage             float32 `json:"minVoltage"`
	F_maxTemperatureSystemNo uint8   `json:"maxTemperatureSystemNo"`
	F_maxTemperatureNo       uint8   `json:"maxTemperatureNo"`
	F_maxTemperature         int16   `json:"maxTemperature"`
	F_minTemperatureSystemNo uint8   `json:"minTemperatureSystemNo"`
	F_minTemperatureNo       uint8   `json:"minTemperatureNo"`
	F_minTemperature         int16   `json:"minTemperature"`
}

func To_VehicleLimitValueData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleLimitValueData {
	_instance := VehicleLimitValueData{}
	F_maxVoltageSystemNo_v := _byteBuf.Read_uint8()
	_instance.F_maxVoltageSystemNo = F_maxVoltageSystemNo_v

	F_maxVoltageCode_v := _byteBuf.Read_uint8()
	_instance.F_maxVoltageCode = F_maxVoltageCode_v

	F_maxVoltage_v := _byteBuf.Read_uint16()
	_instance.F_maxVoltage = float32(F_maxVoltage_v) / 1000

	F_minVoltageSystemNo_v := _byteBuf.Read_uint8()
	_instance.F_minVoltageSystemNo = F_minVoltageSystemNo_v

	F_minVoltageCode_v := _byteBuf.Read_uint8()
	_instance.F_minVoltageCode = F_minVoltageCode_v

	F_minVoltage_v := _byteBuf.Read_uint16()
	_instance.F_minVoltage = float32(F_minVoltage_v) / 1000

	F_maxTemperatureSystemNo_v := _byteBuf.Read_uint8()
	_instance.F_maxTemperatureSystemNo = F_maxTemperatureSystemNo_v

	F_maxTemperatureNo_v := _byteBuf.Read_uint8()
	_instance.F_maxTemperatureNo = F_maxTemperatureNo_v

	F_maxTemperature_v := _byteBuf.Read_uint8()
	_instance.F_maxTemperature = int16(F_maxTemperature_v) - 40

	F_minTemperatureSystemNo_v := _byteBuf.Read_uint8()
	_instance.F_minTemperatureSystemNo = F_minTemperatureSystemNo_v

	F_minTemperatureNo_v := _byteBuf.Read_uint8()
	_instance.F_minTemperatureNo = F_minTemperatureNo_v

	F_minTemperature_v := _byteBuf.Read_uint8()
	_instance.F_minTemperature = int16(F_minTemperature_v) - 40

	return &_instance
}

func (__instance *VehicleLimitValueData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_maxVoltageSystemNo)
	_byteBuf.Write_uint8(_instance.F_maxVoltageCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_maxVoltage * 1000)))
	_byteBuf.Write_uint8(_instance.F_minVoltageSystemNo)
	_byteBuf.Write_uint8(_instance.F_minVoltageCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_minVoltage * 1000)))
	_byteBuf.Write_uint8(_instance.F_maxTemperatureSystemNo)
	_byteBuf.Write_uint8(_instance.F_maxTemperatureNo)
	_byteBuf.Write_uint8(uint8((_instance.F_maxTemperature + 40)))
	_byteBuf.Write_uint8(_instance.F_minTemperatureSystemNo)
	_byteBuf.Write_uint8(_instance.F_minTemperatureNo)
	_byteBuf.Write_uint8(uint8((_instance.F_minTemperature + 40)))
}

type VehicleLoginData struct {
	F_collectTime   time.Time `json:"collectTime"`
	F_sn            uint16    `json:"sn"`
	F_iccid         string    `json:"iccid"`
	F_subSystemNum  uint8     `json:"subSystemNum"`
	F_systemCodeLen uint8     `json:"systemCodeLen"`
	F_systemCode    string    `json:"systemCode"`
}

func To_VehicleLoginData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleLoginData {
	_instance := VehicleLoginData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _Location0)
	F_sn_v := _byteBuf.Read_uint16()
	_instance.F_sn = F_sn_v

	F_iccid_len := 20
	F_iccid_v := _byteBuf.Read_slice_uint8(F_iccid_len)
	F_iccid_count := 0
	for i := F_iccid_len - 1; i >= 0; i-- {
		if F_iccid_v[i] == 0 {
			F_iccid_count++
		} else {
			break
		}
	}
	_instance.F_iccid = string(F_iccid_v[:(F_iccid_len - F_iccid_count)])

	F_subSystemNum_v := _byteBuf.Read_uint8()
	_instance.F_subSystemNum = F_subSystemNum_v

	F_systemCodeLen_v := _byteBuf.Read_uint8()
	_instance.F_systemCodeLen = F_systemCodeLen_v

	F_systemCode_len := (int)(_instance.F_subSystemNum) * (int)(_instance.F_systemCodeLen)
	F_systemCode_v := _byteBuf.Read_slice_uint8(F_systemCode_len)
	F_systemCode_count := 0
	for i := F_systemCode_len - 1; i >= 0; i-- {
		if F_systemCode_v[i] == 0 {
			F_systemCode_count++
		} else {
			break
		}
	}
	_instance.F_systemCode = string(F_systemCode_v[:(F_systemCode_len - F_systemCode_count)])

	return &_instance
}

func (__instance *VehicleLoginData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.F_sn)
	F_iccid_len := 20
	F_iccid_v := []byte(_instance.F_iccid)
	_byteBuf.Write_slice_uint8(F_iccid_v)
	_byteBuf.Write_zero(F_iccid_len - len(F_iccid_v))
	_byteBuf.Write_uint8(_instance.F_subSystemNum)
	_byteBuf.Write_uint8(_instance.F_systemCodeLen)
	F_systemCode_len := (int)(_instance.F_subSystemNum) * (int)(_instance.F_systemCodeLen)
	F_systemCode_v := []byte(_instance.F_systemCode)
	_byteBuf.Write_slice_uint8(F_systemCode_v)
	_byteBuf.Write_zero(F_systemCode_len - len(F_systemCode_v))
}

type VehicleLogoutData struct {
	F_collectTime time.Time `json:"collectTime"`
	F_sn          uint16    `json:"sn"`
}

func To_VehicleLogoutData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleLogoutData {
	_instance := VehicleLogoutData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _Location0)
	F_sn_v := _byteBuf.Read_uint16()
	_instance.F_sn = F_sn_v

	return &_instance
}

func (__instance *VehicleLogoutData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.F_sn)
}

type VehicleMotorData struct {
	F_num     uint8        `json:"num"`
	F_content []*MotorData `json:"content"`
}

func To_VehicleMotorData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleMotorData {
	_instance := VehicleMotorData{}
	F_num_v := _byteBuf.Read_uint8()
	_instance.F_num = F_num_v

	F_content_len := (int)(F_num_v)
	F_content_arr := make([]*MotorData, F_content_len, F_content_len)
	for i := 0; i < F_content_len; i++ {
		F_content_arr[i] = To_MotorData(_byteBuf, nil)
	}
	_instance.F_content = F_content_arr
	return &_instance
}

func (__instance *VehicleMotorData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_num)
	F_content_arr := _instance.F_content
	for i := 0; i < len(F_content_arr); i++ {
		F_content_arr[i].Write(_byteBuf, nil)
	}
}

type VehiclePositionData struct {
	F_status int8    `json:"status"`
	F_lng    float64 `json:"lng"`
	F_lat    float64 `json:"lat"`
}

func To_VehiclePositionData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehiclePositionData {
	_instance := VehiclePositionData{}
	F_status_v := _byteBuf.Read_int8()
	_instance.F_status = F_status_v

	F_lng_v := _byteBuf.Read_uint32()
	_instance.F_lng = float64(F_lng_v) / 1000000

	F_lat_v := _byteBuf.Read_uint32()
	_instance.F_lat = float64(F_lat_v) / 1000000

	return &_instance
}

func (__instance *VehiclePositionData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_int8(_instance.F_status)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lng * 1000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lat * 1000000)))
}

type VehicleRunData struct {
	F_collectTime       time.Time          `json:"collectTime"`
	F_vehicleCommonData *VehicleCommonData `json:"vehicleCommonData"`
}

func To_VehicleRunData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleRunData {
	_instance := VehicleRunData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _Location0)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_vehicleCommonData = To_F_vehicleCommonData(_byteBuf, _parseContext)
	return &_instance
}

func (__instance *VehicleRunData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_vehicleCommonData(_byteBuf, _instance.F_vehicleCommonData, _parseContext)
}

type VehicleStorageTemperatureData struct {
	F_num     uint8                     `json:"num"`
	F_content []*StorageTemperatureData `json:"content"`
}

func To_VehicleStorageTemperatureData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleStorageTemperatureData {
	_instance := VehicleStorageTemperatureData{}
	F_num_v := _byteBuf.Read_uint8()
	_instance.F_num = F_num_v

	F_content_len := (int)(F_num_v)
	F_content_arr := make([]*StorageTemperatureData, F_content_len, F_content_len)
	for i := 0; i < F_content_len; i++ {
		F_content_arr[i] = To_StorageTemperatureData(_byteBuf, nil)
	}
	_instance.F_content = F_content_arr
	return &_instance
}

func (__instance *VehicleStorageTemperatureData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_num)
	F_content_arr := _instance.F_content
	for i := 0; i < len(F_content_arr); i++ {
		F_content_arr[i].Write(_byteBuf, nil)
	}
}

type VehicleStorageVoltageData struct {
	F_num     uint8                 `json:"num"`
	F_content []*StorageVoltageData `json:"content"`
}

func To_VehicleStorageVoltageData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleStorageVoltageData {
	_instance := VehicleStorageVoltageData{}
	F_num_v := _byteBuf.Read_uint8()
	_instance.F_num = F_num_v

	F_content_len := (int)(F_num_v)
	F_content_arr := make([]*StorageVoltageData, F_content_len, F_content_len)
	for i := 0; i < F_content_len; i++ {
		F_content_arr[i] = To_StorageVoltageData(_byteBuf, nil)
	}
	_instance.F_content = F_content_arr
	return &_instance
}

func (__instance *VehicleStorageVoltageData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_num)
	F_content_arr := _instance.F_content
	for i := 0; i < len(F_content_arr); i++ {
		F_content_arr[i].Write(_byteBuf, nil)
	}
}

type VehicleSupplementData struct {
	F_collectTime       time.Time `json:"collectTime"`
	F_vehicleCommonData any       `json:"vehicleCommonData"`
}

func To_VehicleSupplementData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleSupplementData {
	_instance := VehicleSupplementData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _Location0)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_vehicleCommonData = To_F_vehicleCommonData(_byteBuf, _parseContext)
	return &_instance
}

func (__instance *VehicleSupplementData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_vehicleCommonData(_byteBuf, _instance.F_vehicleCommonData, _parseContext)
}

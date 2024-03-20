package gb32960

import (
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"time"
)

var _location_china = time.FixedZone("_location_china", 28800)

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

func To_MotorData(_byteBuf *parse.ByteBuf) MotorData {
	_instance := MotorData{}
	_instance.F_no = _byteBuf.Read_uint8()
	_instance.F_status = _byteBuf.Read_uint8()
	_instance.F_controllerTemperature = int16(_byteBuf.Read_uint8()) - 40
	_instance.F_rotateSpeed = int32(_byteBuf.Read_uint16()) - 20000
	_instance.F_rotateRectangle = float32(_byteBuf.Read_uint16())/10 - 2000
	_instance.F_temperature = int16(_byteBuf.Read_uint8()) - 40
	_instance.F_inputVoltage = float32(_byteBuf.Read_uint16()) / 10
	_instance.F_current = float32(_byteBuf.Read_uint16())/10 - 1000

	return _instance
}

func (_instance MotorData) Write(_byteBuf *parse.ByteBuf) {
	_byteBuf.Write_uint8(_instance.F_no)
	_byteBuf.Write_uint8(_instance.F_status)
	_byteBuf.Write_uint8(uint8(_instance.F_controllerTemperature + 40))
	_byteBuf.Write_uint16(uint16(_instance.F_rotateSpeed + 20000))
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_rotateRectangle + 2000) * 10)))
	_byteBuf.Write_uint8(uint8(_instance.F_temperature + 40))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_inputVoltage * 10)))
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_current + 1000) * 10)))
}

type Packet struct {
	F_header        [2]uint8 `json:"header"`
	F_flag          uint8    `json:"flag"`
	F_replyFlag     uint8    `json:"replyFlag"`
	F_vin           string   `json:"vin"`
	F_encodeWay     uint8    `json:"encodeWay"`
	F_contentLength uint16   `json:"contentLength"`
	F_data          any      `json:"data"`
	F_code          uint8    `json:"code"`
}

func To_Packet(_byteBuf *parse.ByteBuf) *Packet {
	_instance := Packet{}
	_instance.F_header = [2]uint8(_byteBuf.Read_slice_uint8(2))
	_instance.F_flag = _byteBuf.Read_uint8()
	_instance.F_replyFlag = _byteBuf.Read_uint8()
	_instance.F_vin = _byteBuf.Read_string_utf8(17)
	_instance.F_encodeWay = _byteBuf.Read_uint8()
	_instance.F_contentLength = _byteBuf.Read_uint16()
	_instance.F_data = To_F_data(_byteBuf, _instance)
	_instance.F_code = _byteBuf.Read_uint8()
	return &_instance
}

func (__instance *Packet) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_slice_uint8(_instance.F_header[:])
	_byteBuf.Write_uint8(_instance.F_flag)
	_byteBuf.Write_uint8(_instance.F_replyFlag)
	_byteBuf.Write_string_utf8(_instance.F_vin)
	_byteBuf.Write_zero(17 - len(_instance.F_vin))
	_byteBuf.Write_uint8(_instance.F_encodeWay)
	_byteBuf.Write_uint16(_instance.F_contentLength)
	Write_F_data(_byteBuf, _instance)
	_byteBuf.Write_uint8(_instance.F_code)
}

type PlatformLoginData struct {
	F_collectTime time.Time `json:"collectTime"`
	F_sn          uint16    `json:"sn"`
	F_username    string    `json:"username"`
	F_password    string    `json:"password"`
	F_encode      uint8     `json:"encode"`
}

func To_PlatformLoginData(_byteBuf *parse.ByteBuf) *PlatformLoginData {
	_instance := PlatformLoginData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _location_china)
	_instance.F_sn = _byteBuf.Read_uint16()

	F_username_v := _byteBuf.Read_slice_uint8(12)
	F_username_count := 0
	for i := 11; i >= 0; i-- {
		if F_username_v[i] == 0 {
			F_username_count++
		} else {
			break
		}
	}
	_instance.F_username = string(F_username_v[:(12 - F_username_count)])

	F_password_v := _byteBuf.Read_slice_uint8(20)
	F_password_count := 0
	for i := 19; i >= 0; i-- {
		if F_password_v[i] == 0 {
			F_password_count++
		} else {
			break
		}
	}
	_instance.F_password = string(F_password_v[:(20 - F_password_count)])
	_instance.F_encode = _byteBuf.Read_uint8()

	return &_instance
}

func (__instance *PlatformLoginData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.F_sn)
	_byteBuf.Write_string_utf8(_instance.F_username)
	_byteBuf.Write_zero(12 - len(_instance.F_username))
	_byteBuf.Write_string_utf8(_instance.F_password)
	_byteBuf.Write_zero(20 - len(_instance.F_password))
	_byteBuf.Write_uint8(_instance.F_encode)
}

type PlatformLogoutData struct {
	F_collectTime time.Time `json:"collectTime"`
	F_sn          uint16    `json:"sn"`
}

func To_PlatformLogoutData(_byteBuf *parse.ByteBuf) *PlatformLogoutData {
	_instance := PlatformLogoutData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _location_china)
	_instance.F_sn = _byteBuf.Read_uint16()

	return &_instance
}

func (__instance *PlatformLogoutData) Write(_byteBuf *parse.ByteBuf) {
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

func To_StorageTemperatureData(_byteBuf *parse.ByteBuf) *StorageTemperatureData {
	_instance := StorageTemperatureData{}
	_instance.F_no = _byteBuf.Read_uint8()

	F_num_v := _byteBuf.Read_uint16()
	_instance.F_num = F_num_v

	F_temperatures_len := (int)(F_num_v)
	F_temperatures_arr := make([]int16, F_temperatures_len)
	for i := 0; i < F_temperatures_len; i++ {
		e := _byteBuf.Read_uint8()
		F_temperatures_arr[i] = int16(e) - 40
	}
	_instance.F_temperatures = F_temperatures_arr
	return &_instance
}

func (__instance *StorageTemperatureData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_no)
	_byteBuf.Write_uint16(_instance.F_num)
	F_temperatures_arr := _instance.F_temperatures
	for i := 0; i < len(F_temperatures_arr); i++ {
		_byteBuf.Write_uint8(uint8(F_temperatures_arr[i] + 40))
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

func To_StorageVoltageData(_byteBuf *parse.ByteBuf) StorageVoltageData {
	_instance := StorageVoltageData{}
	_instance.F_no = _byteBuf.Read_uint8()
	_instance.F_voltage = float32(_byteBuf.Read_uint16()) / 10
	_instance.F_current = float32(_byteBuf.Read_uint16())/10 - 1000
	_instance.F_total = _byteBuf.Read_uint16()
	_instance.F_frameNo = _byteBuf.Read_uint16()
	F_frameTotal_v := _byteBuf.Read_uint8()
	_instance.F_frameTotal = F_frameTotal_v
	F_singleVoltage_len := (int)(F_frameTotal_v)
	F_singleVoltage_arr := make([]float32, F_singleVoltage_len)
	for i := 0; i < F_singleVoltage_len; i++ {
		F_singleVoltage_arr[i] = float32(_byteBuf.Read_uint16()) / 1000
	}
	_instance.F_singleVoltage = F_singleVoltage_arr
	return _instance
}

func (_instance StorageVoltageData) Write(_byteBuf *parse.ByteBuf) {
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

func To_VehicleAlarmData(_byteBuf *parse.ByteBuf) *VehicleAlarmData {
	_instance := VehicleAlarmData{}
	_instance.F_maxAlarmLevel = _byteBuf.Read_uint8()
	_instance.F_alarmFlag = _byteBuf.Read_int32()
	F_chargeBadNum_v := _byteBuf.Read_uint8()
	_instance.F_chargeBadNum = F_chargeBadNum_v
	F_chargeBadCodes_len := (int)(F_chargeBadNum_v)
	F_chargeBadCodes_arr := make([]uint32, F_chargeBadCodes_len)
	for i := 0; i < F_chargeBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_chargeBadCodes_arr[i] = e
	}
	_instance.F_chargeBadCodes = F_chargeBadCodes_arr

	F_driverBadNum_v := _byteBuf.Read_uint8()
	_instance.F_driverBadNum = F_driverBadNum_v
	F_driverBadCodes_len := (int)(F_driverBadNum_v)
	F_driverBadCodes_arr := make([]uint32, F_driverBadCodes_len)
	for i := 0; i < F_driverBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_driverBadCodes_arr[i] = e
	}
	_instance.F_driverBadCodes = F_driverBadCodes_arr

	F_engineBadNum_v := _byteBuf.Read_uint8()
	_instance.F_engineBadNum = F_engineBadNum_v
	F_engineBadCodes_len := (int)(F_engineBadNum_v)
	F_engineBadCodes_arr := make([]uint32, F_engineBadCodes_len)
	for i := 0; i < F_engineBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_engineBadCodes_arr[i] = e
	}
	_instance.F_engineBadCodes = F_engineBadCodes_arr

	F_otherBadNum_v := _byteBuf.Read_uint8()
	_instance.F_otherBadNum = F_otherBadNum_v
	F_otherBadCodes_len := (int)(F_otherBadNum_v)
	F_otherBadCodes_arr := make([]uint32, F_otherBadCodes_len)
	for i := 0; i < F_otherBadCodes_len; i++ {
		e := _byteBuf.Read_uint32()
		F_otherBadCodes_arr[i] = e
	}
	_instance.F_otherBadCodes = F_otherBadCodes_arr
	return &_instance
}

func (__instance *VehicleAlarmData) Write(_byteBuf *parse.ByteBuf) {
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

func To_VehicleBaseData(_byteBuf *parse.ByteBuf) *VehicleBaseData {
	_instance := VehicleBaseData{}
	_instance.F_vehicleStatus = _byteBuf.Read_uint8()
	_instance.F_chargeStatus = _byteBuf.Read_uint8()
	_instance.F_runMode = _byteBuf.Read_uint8()
	_instance.F_vehicleSpeed = float32(_byteBuf.Read_uint16()) / 10
	_instance.F_totalMileage = float64(_byteBuf.Read_uint32()) / 10
	_instance.F_totalVoltage = float32(_byteBuf.Read_uint16()) / 10
	_instance.F_totalCurrent = float32(_byteBuf.Read_uint16())/10 - 1000
	_instance.F_soc = _byteBuf.Read_uint8()
	_instance.F_dcStatus = _byteBuf.Read_uint8()
	_instance.F_gearPosition = _byteBuf.Read_uint8()
	_instance.F_resistance = _byteBuf.Read_uint16()
	_instance.F_pedalVal = _byteBuf.Read_uint8()
	_instance.F_pedalStatus = _byteBuf.Read_uint8()

	return &_instance
}

func (__instance *VehicleBaseData) Write(_byteBuf *parse.ByteBuf) {
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

func To_VehicleEngineData(_byteBuf *parse.ByteBuf) *VehicleEngineData {
	_instance := VehicleEngineData{}
	_instance.F_status = _byteBuf.Read_uint8()
	_instance.F_speed = _byteBuf.Read_uint16()
	_instance.F_rate = float32(_byteBuf.Read_uint16()) / 100

	return &_instance
}

func (__instance *VehicleEngineData) Write(_byteBuf *parse.ByteBuf) {
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

func To_VehicleFuelBatteryData(_byteBuf *parse.ByteBuf) *VehicleFuelBatteryData {
	_instance := VehicleFuelBatteryData{}
	_instance.F_voltage = float32(_byteBuf.Read_uint16()) / 10
	_instance.F_current = float32(_byteBuf.Read_uint16()) / 10
	_instance.F_consumptionRate = float32(_byteBuf.Read_uint16()) / 100

	F_num_v := _byteBuf.Read_uint32()
	_instance.F_num = F_num_v

	F_temperatures_len := (int)(F_num_v)
	F_temperatures_arr := make([]int16, F_temperatures_len)
	for i := 0; i < F_temperatures_len; i++ {
		e := _byteBuf.Read_uint8()
		F_temperatures_arr[i] = int16(e) - 40
	}
	_instance.F_temperatures = F_temperatures_arr

	_instance.F_maxTemperature = float32(_byteBuf.Read_uint16())/10 - 40

	_instance.F_maxTemperatureCode = _byteBuf.Read_uint8()

	_instance.F_maxConcentration = int32(_byteBuf.Read_uint16()) - 10000

	_instance.F_maxConcentrationCode = _byteBuf.Read_uint8()

	_instance.F_maxPressure = float32(_byteBuf.Read_uint16()) / 10

	_instance.F_maxPressureCode = _byteBuf.Read_uint8()

	_instance.F_dcStatus = _byteBuf.Read_uint8()

	return &_instance
}

func (__instance *VehicleFuelBatteryData) Write(_byteBuf *parse.ByteBuf) {
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
	_byteBuf.Write_uint16(uint16(_instance.F_maxConcentration + 10000))
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

func To_VehicleLimitValueData(_byteBuf *parse.ByteBuf) *VehicleLimitValueData {
	_instance := VehicleLimitValueData{}
	_instance.F_maxVoltageSystemNo = _byteBuf.Read_uint8()
	_instance.F_maxVoltageCode = _byteBuf.Read_uint8()
	_instance.F_maxVoltage = float32(_byteBuf.Read_uint16()) / 1000
	_instance.F_minVoltageSystemNo = _byteBuf.Read_uint8()
	_instance.F_minVoltageCode = _byteBuf.Read_uint8()
	_instance.F_minVoltage = float32(_byteBuf.Read_uint16()) / 1000
	_instance.F_maxTemperatureSystemNo = _byteBuf.Read_uint8()
	_instance.F_maxTemperatureNo = _byteBuf.Read_uint8()
	_instance.F_maxTemperature = int16(_byteBuf.Read_uint8()) - 40
	_instance.F_minTemperatureSystemNo = _byteBuf.Read_uint8()
	_instance.F_minTemperatureNo = _byteBuf.Read_uint8()
	_instance.F_minTemperature = int16(_byteBuf.Read_uint8()) - 40

	return &_instance
}

func (__instance *VehicleLimitValueData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_maxVoltageSystemNo)
	_byteBuf.Write_uint8(_instance.F_maxVoltageCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_maxVoltage * 1000)))
	_byteBuf.Write_uint8(_instance.F_minVoltageSystemNo)
	_byteBuf.Write_uint8(_instance.F_minVoltageCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_minVoltage * 1000)))
	_byteBuf.Write_uint8(_instance.F_maxTemperatureSystemNo)
	_byteBuf.Write_uint8(_instance.F_maxTemperatureNo)
	_byteBuf.Write_uint8(uint8(_instance.F_maxTemperature + 40))
	_byteBuf.Write_uint8(_instance.F_minTemperatureSystemNo)
	_byteBuf.Write_uint8(_instance.F_minTemperatureNo)
	_byteBuf.Write_uint8(uint8(_instance.F_minTemperature + 40))
}

type VehicleLoginData struct {
	F_collectTime   time.Time `json:"collectTime"`
	F_sn            uint16    `json:"sn"`
	F_iccid         string    `json:"iccid"`
	F_subSystemNum  uint8     `json:"subSystemNum"`
	F_systemCodeLen uint8     `json:"systemCodeLen"`
	F_systemCode    string    `json:"systemCode"`
}

func To_VehicleLoginData(_byteBuf *parse.ByteBuf) *VehicleLoginData {
	_instance := VehicleLoginData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _location_china)
	_instance.F_sn = _byteBuf.Read_uint16()

	F_iccid_v := _byteBuf.Read_slice_uint8(20)
	F_iccid_count := 0
	for i := 19; i >= 0; i-- {
		if F_iccid_v[i] == 0 {
			F_iccid_count++
		} else {
			break
		}
	}
	_instance.F_iccid = string(F_iccid_v[:(20 - F_iccid_count)])

	_instance.F_subSystemNum = _byteBuf.Read_uint8()
	_instance.F_systemCodeLen = _byteBuf.Read_uint8()

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

func (__instance *VehicleLoginData) Write(_byteBuf *parse.ByteBuf) {
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

func To_VehicleLogoutData(_byteBuf *parse.ByteBuf) *VehicleLogoutData {
	_instance := VehicleLogoutData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _location_china)
	_instance.F_sn = _byteBuf.Read_uint16()

	return &_instance
}

func (__instance *VehicleLogoutData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	F_collectTime_v := _instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.F_sn)
}

type VehicleMotorData struct {
	F_num     uint8       `json:"num"`
	F_content []MotorData `json:"content"`
}

func To_VehicleMotorData(_byteBuf *parse.ByteBuf) *VehicleMotorData {
	_instance := VehicleMotorData{}
	F_num_v := _byteBuf.Read_uint8()
	_instance.F_num = F_num_v

	F_content_len := (int)(F_num_v)
	F_content_arr := make([]MotorData, F_content_len, F_content_len)
	for i := 0; i < F_content_len; i++ {
		F_content_arr[i] = To_MotorData(_byteBuf)
	}
	_instance.F_content = F_content_arr
	return &_instance
}

func (__instance *VehicleMotorData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_num)
	F_content_arr := _instance.F_content
	for i := 0; i < len(F_content_arr); i++ {
		F_content_arr[i].Write(_byteBuf)
	}
}

type VehiclePositionData struct {
	F_status uint8   `json:"status"`
	F_lng    float64 `json:"lng"`
	F_lat    float64 `json:"lat"`
}

func To_VehiclePositionData(_byteBuf *parse.ByteBuf) *VehiclePositionData {
	_instance := VehiclePositionData{}
	_instance.F_status = _byteBuf.Read_uint8()
	_instance.F_lng = float64(_byteBuf.Read_uint32()) / 1000000
	_instance.F_lat = float64(_byteBuf.Read_uint32()) / 1000000
	return &_instance
}

func (__instance *VehiclePositionData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_status)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lng * 1000000)))
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.F_lat * 1000000)))
}

type VehicleStorageTemperatureData struct {
	F_num     uint8                     `json:"num"`
	F_content []*StorageTemperatureData `json:"content"`
}

func To_VehicleStorageTemperatureData(_byteBuf *parse.ByteBuf) *VehicleStorageTemperatureData {
	_instance := VehicleStorageTemperatureData{}
	F_num_v := _byteBuf.Read_uint8()
	_instance.F_num = F_num_v

	F_content_len := (int)(F_num_v)
	F_content_arr := make([]*StorageTemperatureData, F_content_len)
	for i := 0; i < F_content_len; i++ {
		F_content_arr[i] = To_StorageTemperatureData(_byteBuf)
	}
	_instance.F_content = F_content_arr
	return &_instance
}

func (__instance *VehicleStorageTemperatureData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_num)
	F_content_arr := _instance.F_content
	for i := 0; i < len(F_content_arr); i++ {
		F_content_arr[i].Write(_byteBuf)
	}
}

type VehicleStorageVoltageData struct {
	F_num     uint8                `json:"num"`
	F_content []StorageVoltageData `json:"content"`
}

func To_VehicleStorageVoltageData(_byteBuf *parse.ByteBuf) *VehicleStorageVoltageData {
	_instance := VehicleStorageVoltageData{}
	F_num_v := _byteBuf.Read_uint8()
	_instance.F_num = F_num_v

	F_content_len := (int)(F_num_v)
	F_content_arr := make([]StorageVoltageData, F_content_len)
	for i := 0; i < F_content_len; i++ {
		F_content_arr[i] = To_StorageVoltageData(_byteBuf)
	}
	_instance.F_content = F_content_arr
	return &_instance
}

func (__instance *VehicleStorageVoltageData) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.F_num)
	F_content_arr := _instance.F_content
	for i := 0; i < len(F_content_arr); i++ {
		F_content_arr[i].Write(_byteBuf)
	}
}

type VehicleRunData struct {
	F_collectTime                   time.Time                      `json:"collectTime"`
	F_vehicleBaseData               *VehicleBaseData               `json:"vehicleBaseData,omitempty"`
	F_vehicleMotorData              *VehicleMotorData              `json:"vehicleMotorData,omitempty"`
	F_vehicleFuelBatteryData        *VehicleFuelBatteryData        `json:"vehicleFuelBatteryData,omitempty"`
	F_vehicleEngineData             *VehicleEngineData             `json:"vehicleEngineData,omitempty"`
	F_vehiclePositionData           *VehiclePositionData           `json:"vehiclePositionData,omitempty"`
	F_vehicleLimitValueData         *VehicleLimitValueData         `json:"vehicleLimitValueData,omitempty"`
	F_vehicleAlarmData              *VehicleAlarmData              `json:"vehicleAlarmData,omitempty"`
	F_vehicleStorageVoltageData     *VehicleStorageVoltageData     `json:"vehicleStorageVoltageData,omitempty"`
	F_vehicleStorageTemperatureData *VehicleStorageTemperatureData `json:"vehicleStorageTemperatureData,omitempty"`
}

func To_VehicleRunData(_byteBuf *parse.ByteBuf, contentLength int) *VehicleRunData {
	_instance := VehicleRunData{}
	F_collectTime_bytes := _byteBuf.Read_slice_uint8(6)
	_instance.F_collectTime = time.Date(2000+int(F_collectTime_bytes[0]), time.Month(int(F_collectTime_bytes[1])), int(F_collectTime_bytes[2]), int(F_collectTime_bytes[3]), int(F_collectTime_bytes[4]), int(F_collectTime_bytes[5]), 0, _location_china)
	beginLeave := _byteBuf.ReadableBytes()
	contentLength = contentLength - 6
	for _byteBuf.Readable() {
		curLeave := _byteBuf.ReadableBytes()
		if beginLeave-curLeave >= contentLength {
			break
		}
		flag := _byteBuf.Read_uint8()
		switch flag {
		case 1:
			_instance.F_vehicleBaseData = To_VehicleBaseData(_byteBuf)
		case 2:
			_instance.F_vehicleMotorData = To_VehicleMotorData(_byteBuf)
		case 3:
			_instance.F_vehicleFuelBatteryData = To_VehicleFuelBatteryData(_byteBuf)
		case 4:
			_instance.F_vehicleEngineData = To_VehicleEngineData(_byteBuf)
		case 5:
			_instance.F_vehiclePositionData = To_VehiclePositionData(_byteBuf)
		case 6:
			_instance.F_vehicleLimitValueData = To_VehicleLimitValueData(_byteBuf)
		case 7:
			_instance.F_vehicleAlarmData = To_VehicleAlarmData(_byteBuf)
		case 8:
			_instance.F_vehicleStorageVoltageData = To_VehicleStorageVoltageData(_byteBuf)
		case 9:
			_instance.F_vehicleStorageTemperatureData = To_VehicleStorageTemperatureData(_byteBuf)
		default:
			util.Log.Warnf("Parse VehicleRunData Interrupted,Unknown Flag[%d]", flag)
			return &_instance
		}
	}
	return &_instance
}

func (__instance *VehicleRunData) Write(_byteBuf *parse.ByteBuf) {
	F_collectTime_v := __instance.F_collectTime
	_byteBuf.Write_slice_uint8([]byte{byte(F_collectTime_v.Year() - 2000), byte(F_collectTime_v.Month()), byte(F_collectTime_v.Day()), byte(F_collectTime_v.Hour()), byte(F_collectTime_v.Minute()), byte(F_collectTime_v.Second())})
	if __instance.F_vehicleBaseData != nil {
		_byteBuf.Write_uint8(1)
		__instance.F_vehicleBaseData.Write(_byteBuf)
	}
	if __instance.F_vehicleMotorData != nil {
		_byteBuf.Write_uint8(2)
		__instance.F_vehicleMotorData.Write(_byteBuf)
	}
	if __instance.F_vehicleFuelBatteryData != nil {
		_byteBuf.Write_uint8(3)
		__instance.F_vehicleFuelBatteryData.Write(_byteBuf)
	}
	if __instance.F_vehicleEngineData != nil {
		_byteBuf.Write_uint8(4)
		__instance.F_vehicleEngineData.Write(_byteBuf)
	}
	if __instance.F_vehiclePositionData != nil {
		_byteBuf.Write_uint8(5)
		__instance.F_vehiclePositionData.Write(_byteBuf)
	}
	if __instance.F_vehicleLimitValueData != nil {
		_byteBuf.Write_uint8(6)
		__instance.F_vehicleLimitValueData.Write(_byteBuf)
	}
	if __instance.F_vehicleAlarmData != nil {
		_byteBuf.Write_uint8(7)
		__instance.F_vehicleAlarmData.Write(_byteBuf)
	}
	if __instance.F_vehicleStorageVoltageData != nil {
		_byteBuf.Write_uint8(8)
		__instance.F_vehicleStorageVoltageData.Write(_byteBuf)
	}
	if __instance.F_vehicleStorageTemperatureData != nil {
		_byteBuf.Write_uint8(9)
		__instance.F_vehicleStorageTemperatureData.Write(_byteBuf)
	}
}

func To_F_data(_byteBuf *parse.ByteBuf, packet Packet) any {
	//if packet.F_replyFlag == 0xfe {
	switch packet.F_flag {
	case 1:
		return To_VehicleLoginData(_byteBuf)
	case 2, 3:
		return To_VehicleRunData(_byteBuf, int(packet.F_contentLength))
	case 4:
		return To_VehicleLogoutData(_byteBuf)
	case 5:
		return To_PlatformLoginData(_byteBuf)
	case 6:
		return To_PlatformLogoutData(_byteBuf)
	default:
		util.Log.Warnf("Parse PacketData Interrupted,Unknown Flag[%d]", packet.F_flag)
		return nil
	}
	//} else {
	//	return ResponseData{content: _byteBuf.Read_slice_uint8(int(packet.F_contentLength))}
	//}
}

func Write_F_data(_byteBuf *parse.ByteBuf, packet Packet) {
	f_data := packet.F_data
	//if packet.F_replyFlag == 0xfe {
	switch packet.F_flag {
	case 1:
		f_data.(*VehicleLoginData).Write(_byteBuf)
	case 2, 3:
		f_data.(*VehicleRunData).Write(_byteBuf)
	case 4:
		f_data.(*VehicleLogoutData).Write(_byteBuf)
	case 5:
		f_data.(*PlatformLoginData).Write(_byteBuf)
	case 6:
		f_data.(*PlatformLogoutData).Write(_byteBuf)
	default:
		util.Log.Warnf("Parse PacketData Interrupted,Unknown Flag[%d]", packet.F_flag)
	}
	//} else {
	//	_byteBuf.Write_slice_uint8(f_data.(ResponseData).content)
	//}
}

type ResponseData struct {
	content []byte
}

type Packet_runData struct {
	F_header        [2]uint8        `json:"header"`
	F_flag          uint8           `json:"flag"`
	F_replyFlag     uint8           `json:"replyFlag"`
	F_vin           string          `json:"vin"`
	F_encodeWay     uint8           `json:"encodeWay"`
	F_contentLength uint16          `json:"contentLength"`
	F_data          *VehicleRunData `json:"data"`
	F_code          uint8           `json:"code"`
}

func (e *Packet_runData) ToBytes() []byte {
	p := Packet{}
	p.F_header = e.F_header
	p.F_flag = e.F_flag
	p.F_replyFlag = e.F_replyFlag
	p.F_vin = e.F_vin
	p.F_encodeWay = e.F_encodeWay
	p.F_contentLength = e.F_contentLength
	p.F_data = e.F_data
	p.F_code = e.F_code
	buf := parse.ToByteBuf_empty()
	p.Write(buf)
	return buf.ToBytes()
}

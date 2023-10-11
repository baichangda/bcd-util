package gb32960

import (
	"bcd-util/support_parse/parse"
	"bcd-util/util"
)

type VehicleCommonData struct {
	VehicleBaseData               *VehicleBaseData
	VehicleMotorData              *VehicleMotorData
	VehicleFuelBatteryData        *VehicleFuelBatteryData
	VehicleEngineData             *VehicleEngineData
	VehiclePositionData           *VehiclePositionData
	VehicleLimitValueData         *VehicleLimitValueData
	VehicleAlarmData              *VehicleAlarmData
	VehicleStorageVoltageData     *VehicleStorageVoltageData
	VehicleStorageTemperatureData *VehicleStorageTemperatureData
}

func To_F_vehicleCommonData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *VehicleCommonData {
	_instance := VehicleCommonData{}
	packet := _parentParseContext.ParentContext.Instance.(*Packet)
	contentLength := int(packet.F_contentLength) - 6
	beginLeave := _byteBuf.ReadableBytes()
	for _byteBuf.Readable() {
		curLeave := _byteBuf.ReadableBytes()
		if beginLeave-curLeave >= contentLength {
			break
		}
		flag := _byteBuf.Read_uint8()
		switch flag {
		case 1:
			_instance.VehicleBaseData = To_VehicleBaseData(_byteBuf, _parentParseContext)
		case 2:
			_instance.VehicleMotorData = To_VehicleMotorData(_byteBuf, _parentParseContext)
		case 3:
			_instance.VehicleFuelBatteryData = To_VehicleFuelBatteryData(_byteBuf, _parentParseContext)
		case 4:
			_instance.VehicleEngineData = To_VehicleEngineData(_byteBuf, _parentParseContext)
		case 5:
			_instance.VehiclePositionData = To_VehiclePositionData(_byteBuf, _parentParseContext)
		case 6:
			_instance.VehicleLimitValueData = To_VehicleLimitValueData(_byteBuf, _parentParseContext)
		case 7:
			_instance.VehicleAlarmData = To_VehicleAlarmData(_byteBuf, _parentParseContext)
		case 8:
			_instance.VehicleStorageVoltageData = To_VehicleStorageVoltageData(_byteBuf, _parentParseContext)
		case 9:
			_instance.VehicleStorageTemperatureData = To_VehicleStorageTemperatureData(_byteBuf, _parentParseContext)
		default:
			util.Log.Warnf("Parse VehicleCommonData Interrupted,Unknown Flag[%d]", flag)
			return &_instance
		}
	}
	return &_instance
}

func Write_F_vehicleCommonData(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	_instance := __instance.(*VehicleCommonData)

	if _instance.VehicleBaseData != nil {
		_byteBuf.Write_uint8(1)
		_instance.VehicleBaseData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehicleMotorData != nil {
		_byteBuf.Write_uint8(2)
		_instance.VehicleMotorData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehicleFuelBatteryData != nil {
		_byteBuf.Write_uint8(3)
		_instance.VehicleFuelBatteryData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehicleEngineData != nil {
		_byteBuf.Write_uint8(4)
		_instance.VehicleEngineData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehiclePositionData != nil {
		_byteBuf.Write_uint8(5)
		_instance.VehiclePositionData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehicleLimitValueData != nil {
		_byteBuf.Write_uint8(6)
		_instance.VehicleLimitValueData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehicleAlarmData != nil {
		_byteBuf.Write_uint8(7)
		_instance.VehicleAlarmData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehicleStorageVoltageData != nil {
		_byteBuf.Write_uint8(8)
		_instance.VehicleStorageVoltageData.Write(_byteBuf, _parentParseContext)
	}
	if _instance.VehicleStorageTemperatureData != nil {
		_byteBuf.Write_uint8(9)
		_instance.VehicleStorageTemperatureData.Write(_byteBuf, _parentParseContext)
	}
}

func To_F_data(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	packet := _parentParseContext.Instance.(*Packet)
	if packet.F_replyFlag == 0xfe {
		switch packet.F_flag {
		case 1:
			return To_VehicleLoginData(_byteBuf, _parentParseContext)
		case 2:
			return To_VehicleRunData(_byteBuf, _parentParseContext)
		case 3:
			return To_VehicleSupplementData(_byteBuf, _parentParseContext)
		case 4:
			return To_VehicleLogoutData(_byteBuf, _parentParseContext)
		case 5:
			return To_PlatformLoginData(_byteBuf, _parentParseContext)
		case 6:
			return To_PlatformLogoutData(_byteBuf, _parentParseContext)
		default:
			util.Log.Warnf("Parse PacketData Interrupted,Unknown Flag[%d]", packet.F_flag)
			return nil
		}
	} else {
		return ResponseData{content: _byteBuf.Read_slice_uint8(int(packet.F_contentLength))}
	}
}

func Write_F_data(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	packet := _parentParseContext.Instance.(*Packet)
	if packet.F_replyFlag == 0xfe {
		__instance.(parse.Writeable).Write(_byteBuf, _parentParseContext)
	} else {
		_byteBuf.Write_slice_uint8(__instance.(ResponseData).content)
	}

}

type ResponseData struct {
	content []byte
}

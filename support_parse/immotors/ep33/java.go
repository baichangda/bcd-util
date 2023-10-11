package ep33

import (
	"bcd-util/support_parse/parse"
	"reflect"
	"unsafe"
)

type Evt_0001 struct {
	F_evtId      uint16 `json:"evtId"`
	F_TBOXSysTim int64  `json:"TBOXSysTim"`
}

func To_Evt_0001(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0001 {
	_instance := Evt_0001{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_TBOXSysTim_v := _bitBuf.Read(48, true, true)
	_bitBuf.Finish()
	_instance.F_TBOXSysTim = int64(F_TBOXSysTim_v)

	return &_instance
}

func (__instance *Evt_0001) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_TBOXSysTim), 48, true, true)
	_bitBuf.Finish()
}

type Evt_0004 struct {
	F_evtId     uint16  `json:"evtId"`
	F_GnssAlt   float64 `json:"GnssAlt"`
	F_Longitude float64 `json:"Longitude"`
	F_GPSSts    uint8   `json:"GPSSts"`
}

func To_Evt_0004(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0004 {
	_instance := Evt_0004{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_GnssAlt_v := _bitBuf.Read(16, true, true)
	_instance.F_GnssAlt = float64(F_GnssAlt_v)*0.1 - 500

	F_Longitude_v := _bitBuf.Read(29, true, true)
	_instance.F_Longitude = float64(F_Longitude_v) * 0.000001

	F_GPSSts_v := _bitBuf.Read(2, true, true)
	_bitBuf.Finish()
	_instance.F_GPSSts = uint8(F_GPSSts_v)

	return &_instance
}

func (__instance *Evt_0004) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round((_instance.F_GnssAlt+500)/0.1)), 16, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_Longitude/0.000001)), 29, true, true)
	_bitBuf.Write(int64(_instance.F_GPSSts), 2, true, true)
	_bitBuf.Finish()
}

type Evt_0005 struct {
	F_evtId         uint16  `json:"evtId"`
	F_Latitude      float64 `json:"Latitude"`
	F_VehTyp        uint8   `json:"VehTyp"`
	F_GNSSDirection float64 `json:"GNSSDirection"`
}

func To_Evt_0005(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0005 {
	_instance := Evt_0005{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_Latitude_v := _bitBuf.Read(28, true, true)
	_instance.F_Latitude = float64(F_Latitude_v) * 0.000001

	F_VehTyp_v := _bitBuf.Read(4, true, true)
	_instance.F_VehTyp = uint8(F_VehTyp_v)

	F_GNSSDirection_v := _bitBuf.Read(16, true, true)
	_bitBuf.Finish()
	_instance.F_GNSSDirection = float64(F_GNSSDirection_v) * 0.01

	return &_instance
}

func (__instance *Evt_0005) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_Latitude/0.000001)), 28, true, true)
	_bitBuf.Write(int64(_instance.F_VehTyp), 4, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_GNSSDirection/0.01)), 16, true, true)
	_bitBuf.Finish()
}

type Evt_0006 struct {
	F_evtId uint16  `json:"evtId"`
	F_HDop  float64 `json:"HDop"`
	F_VDop  float64 `json:"VDop"`
}

func To_Evt_0006(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0006 {
	_instance := Evt_0006{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_HDop_v := _bitBuf.Read(24, true, true)
	_instance.F_HDop = float64(F_HDop_v) * 0.1

	F_VDop_v := _bitBuf.Read(24, true, true)
	_bitBuf.Finish()
	_instance.F_VDop = float64(F_VDop_v) * 0.1

	return &_instance
}

func (__instance *Evt_0006) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_HDop/0.1)), 24, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_VDop/0.1)), 24, true, true)
	_bitBuf.Finish()
}

type Evt_0007 struct {
	F_evtId uint16  `json:"evtId"`
	F_AcceX float64 `json:"AcceX"`
	F_AcceY float64 `json:"AcceY"`
	F_AcceZ float64 `json:"AcceZ"`
}

func To_Evt_0007(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0007 {
	_instance := Evt_0007{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_AcceX_v := _bitBuf.Read(14, true, false)
	_bitBuf.Finish()
	_instance.F_AcceX = float64(F_AcceX_v) * 0.0009765625

	F_AcceY_v := _bitBuf.Read(14, true, false)
	_bitBuf.Finish()
	_instance.F_AcceY = float64(F_AcceY_v) * 0.0009765625

	F_AcceZ_v := _bitBuf.Read(14, true, false)
	_bitBuf.Finish()
	_instance.F_AcceZ = float64(F_AcceZ_v) * 0.0009765625

	return &_instance
}

func (__instance *Evt_0007) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_AcceX/0.0009765625)), 14, true, false)
	_bitBuf.Finish()
	_bitBuf.Write(int64(parse.Round(_instance.F_AcceY/0.0009765625)), 14, true, false)
	_bitBuf.Finish()
	_bitBuf.Write(int64(parse.Round(_instance.F_AcceZ/0.0009765625)), 14, true, false)
	_bitBuf.Finish()
}

type Evt_0008 struct {
	F_evtId   uint16 `json:"evtId"`
	F_cellMCC uint16 `json:"cellMCC"`
	F_cellMNC uint16 `json:"cellMNC"`
}

func To_Evt_0008(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0008 {
	_instance := Evt_0008{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_cellMCC_v := _bitBuf.Read(16, true, true)
	_instance.F_cellMCC = uint16(F_cellMCC_v)

	F_cellMNC_v := _bitBuf.Read(16, true, true)
	_bitBuf.Finish()
	_instance.F_cellMNC = uint16(F_cellMNC_v)

	F_skip_len := 2
	_byteBuf.Skip(F_skip_len)
	return &_instance
}

func (__instance *Evt_0008) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_cellMCC), 16, true, true)
	_bitBuf.Write(int64(_instance.F_cellMNC), 16, true, true)
	_bitBuf.Finish()
	F_skip_len := 2
	_byteBuf.Write_zero(F_skip_len)

}

type Evt_0009 struct {
	F_evtId   uint16 `json:"evtId"`
	F_cellLAC uint16 `json:"cellLAC"`
	F_CellID  uint32 `json:"CellID"`
}

func To_Evt_0009(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0009 {
	_instance := Evt_0009{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_cellLAC_v := _bitBuf.Read(16, true, true)
	_instance.F_cellLAC = uint16(F_cellLAC_v)

	F_CellID_v := _bitBuf.Read(32, true, true)
	_bitBuf.Finish()
	_instance.F_CellID = uint32(F_CellID_v)

	return &_instance
}

func (__instance *Evt_0009) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_cellLAC), 16, true, true)
	_bitBuf.Write(int64(_instance.F_CellID), 32, true, true)
	_bitBuf.Finish()
}

type Evt_000A struct {
	F_evtId              uint16 `json:"evtId"`
	F_cellSignalStrength int8   `json:"cellSignalStrength"`
	F_cellRAT            uint8  `json:"cellRAT"`
	F_cellRATadd         uint8  `json:"cellRATadd"`
	F_cellChanID         uint16 `json:"cellChanID"`
	F_GNSSSATS           uint8  `json:"GNSSSATS"`
}

func To_Evt_000A(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_000A {
	_instance := Evt_000A{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_cellSignalStrength_v := _bitBuf.Read(8, true, true)
	_instance.F_cellSignalStrength = int8(F_cellSignalStrength_v)

	F_cellRAT_v := _bitBuf.Read(3, true, true)
	_instance.F_cellRAT = uint8(F_cellRAT_v)

	F_cellRATadd_v := _bitBuf.Read(3, true, true)
	_instance.F_cellRATadd = uint8(F_cellRATadd_v)

	F_cellChanID_v := _bitBuf.Read(9, true, true)
	_instance.F_cellChanID = uint16(F_cellChanID_v)

	F_GNSSSATS_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.F_GNSSSATS = uint8(F_GNSSSATS_v)

	F_skip_len := 2
	_byteBuf.Skip(F_skip_len)
	return &_instance
}

func (__instance *Evt_000A) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_cellSignalStrength), 8, true, true)
	_bitBuf.Write(int64(_instance.F_cellRAT), 3, true, true)
	_bitBuf.Write(int64(_instance.F_cellRATadd), 3, true, true)
	_bitBuf.Write(int64(_instance.F_cellChanID), 9, true, true)
	_bitBuf.Write(int64(_instance.F_GNSSSATS), 8, true, true)
	_bitBuf.Finish()
	F_skip_len := 2
	_byteBuf.Write_zero(F_skip_len)

}

type Evt_0800 struct {
	F_evtId         uint16  `json:"evtId"`
	F_SysPwrMd      uint8   `json:"SysPwrMd"`
	F_SysPwrMdV     uint8   `json:"SysPwrMdV"`
	F_SysVolV       uint8   `json:"SysVolV"`
	F_TrShftLvrPos  uint8   `json:"TrShftLvrPos"`
	F_SysVol        float32 `json:"SysVol"`
	F_TrShftLvrPosV uint8   `json:"TrShftLvrPosV"`
}

func To_Evt_0800(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0800 {
	_instance := Evt_0800{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_SysPwrMd_v := _bitBuf.Read(2, true, true)
	_instance.F_SysPwrMd = uint8(F_SysPwrMd_v)

	F_SysPwrMdV_v := _bitBuf.Read(1, true, true)
	_instance.F_SysPwrMdV = uint8(F_SysPwrMdV_v)

	F_SysVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_SysVolV = uint8(F_SysVolV_v)

	F_TrShftLvrPos_v := _bitBuf.Read(4, true, true)
	_instance.F_TrShftLvrPos = uint8(F_TrShftLvrPos_v)

	F_SysVol_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.F_SysVol = float32(F_SysVol_v)*0.1 + 3

	F_skip_len := 3
	_byteBuf.Skip(F_skip_len)
	F_TrShftLvrPosV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.F_TrShftLvrPosV = uint8(F_TrShftLvrPosV_v)

	return &_instance
}

func (__instance *Evt_0800) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_SysPwrMd), 2, true, true)
	_bitBuf.Write(int64(_instance.F_SysPwrMdV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_SysVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_TrShftLvrPos), 4, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_SysVol-3)/0.1)), 8, true, true)
	_bitBuf.Finish()
	F_skip_len := 3
	_byteBuf.Write_zero(F_skip_len)

	_bitBuf.Write(int64(_instance.F_TrShftLvrPosV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_0801 struct {
	F_evtId     uint16  `json:"evtId"`
	F_BrkPdlPos float32 `json:"BrkPdlPos"`
}

func To_Evt_0801(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0801 {
	_instance := Evt_0801{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_skip_len := 5
	_byteBuf.Skip(F_skip_len)
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_BrkPdlPos_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.F_BrkPdlPos = float32(F_BrkPdlPos_v) * 0.392157

	return &_instance
}

func (__instance *Evt_0801) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	F_skip_len := 5
	_byteBuf.Write_zero(F_skip_len)

	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_BrkPdlPos/0.392157)), 8, true, true)
	_bitBuf.Finish()
}

type Evt_0802 struct {
	F_evtId          uint16  `json:"evtId"`
	F_VehSpdAvgDrvn  float64 `json:"VehSpdAvgDrvn"`
	F_VehSpdAvgDrvnV uint8   `json:"VehSpdAvgDrvnV"`
}

func To_Evt_0802(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0802 {
	_instance := Evt_0802{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_VehSpdAvgDrvn_v := _bitBuf.Read(15, true, true)
	_instance.F_VehSpdAvgDrvn = float64(F_VehSpdAvgDrvn_v) * 0.015625

	F_VehSpdAvgDrvnV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.F_VehSpdAvgDrvnV = uint8(F_VehSpdAvgDrvnV_v)

	F_skip_len := 4
	_byteBuf.Skip(F_skip_len)
	return &_instance
}

func (__instance *Evt_0802) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_VehSpdAvgDrvn/0.015625)), 15, true, true)
	_bitBuf.Write(int64(_instance.F_VehSpdAvgDrvnV), 1, true, true)
	_bitBuf.Finish()
	F_skip_len := 4
	_byteBuf.Write_zero(F_skip_len)

}

type Evt_0803 struct {
	F_evtId      uint16 `json:"evtId"`
	F_VehOdo     uint32 `json:"VehOdo"`
	F_VehOdoV    uint8  `json:"VehOdoV"`
	F_BrkPdlPosV uint8  `json:"BrkPdlPosV"`
}

func To_Evt_0803(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_0803 {
	_instance := Evt_0803{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_VehOdo_v := _bitBuf.Read(24, true, true)
	_instance.F_VehOdo = uint32(F_VehOdo_v)

	F_VehOdoV_v := _bitBuf.Read(1, true, true)
	_instance.F_VehOdoV = uint8(F_VehOdoV_v)

	F_BrkPdlPosV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.F_BrkPdlPosV = uint8(F_BrkPdlPosV_v)

	F_skip_len := 2
	_byteBuf.Skip(F_skip_len)
	return &_instance
}

func (__instance *Evt_0803) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_VehOdo), 24, true, true)
	_bitBuf.Write(int64(_instance.F_VehOdoV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BrkPdlPosV), 1, true, true)
	_bitBuf.Finish()
	F_skip_len := 2
	_byteBuf.Write_zero(F_skip_len)

}

type Evt_2_6_unknown struct {
	F_evtId uint16  `json:"evtId"`
	F_data  [6]int8 `json:"data"`
}

func To_Evt_2_6_unknown(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_2_6_unknown {
	return (*Evt_2_6_unknown)(unsafe.Pointer(unsafe.SliceData(_byteBuf.Read_slice_uint8(8))))
}
func (__instance *Evt_2_6_unknown) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_slice_uint8(*(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(__instance)),
		Len:  8,
		Cap:  8,
	})))
}

type Evt_4_x_unknown struct {
	F_evtId  uint16 `json:"evtId"`
	F_evtLen uint16 `json:"evtLen"`
	F_data   []int8 `json:"data"`
}

func To_Evt_4_x_unknown(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_4_x_unknown {
	_instance := Evt_4_x_unknown{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	F_data_len := (int)(F_evtLen_v)
	F_data_arr := _byteBuf.Read_slice_int8(F_data_len)
	_instance.F_data = F_data_arr
	return &_instance
}

func (__instance *Evt_4_x_unknown) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	F_data_arr := _instance.F_data
	_byteBuf.Write_slice_int8(F_data_arr)
}

type Evt_D006 struct {
	F_evtId                  uint16  `json:"evtId"`
	F_evtLen                 uint16  `json:"evtLen"`
	F_EPTRdy                 uint8   `json:"EPTRdy"`
	F_BMSBscSta              uint8   `json:"BMSBscSta"`
	F_BMSPackCrnt            float32 `json:"BMSPackCrnt"`
	F_BMSPackCrntV           uint8   `json:"BMSPackCrntV"`
	F_BMSPackSOC             float32 `json:"BMSPackSOC"`
	F_BMSPackSOCV            uint8   `json:"BMSPackSOCV"`
	F_BMSPackSOCDsp          float32 `json:"BMSPackSOCDsp"`
	F_BMSPackSOCDspV         uint8   `json:"BMSPackSOCDspV"`
	F_ElecVehSysMd           uint8   `json:"ElecVehSysMd"`
	F_BMSPackVol             float32 `json:"BMSPackVol"`
	F_BMSPackVolV            uint8   `json:"BMSPackVolV"`
	F_HVDCDCSta              uint8   `json:"HVDCDCSta"`
	F_EPTTrInptShaftToq      float32 `json:"EPTTrInptShaftToq"`
	F_EPTTrInptShaftToqV     uint8   `json:"EPTTrInptShaftToqV"`
	F_EPTTrOtptShaftToq      int16   `json:"EPTTrOtptShaftToq"`
	F_EPTTrOtptShaftToqV     uint8   `json:"EPTTrOtptShaftToqV"`
	F_EPTBrkPdlDscrtInptSts  uint8   `json:"EPTBrkPdlDscrtInptSts"`
	F_EPTBrkPdlDscrtInptStsV uint8   `json:"EPTBrkPdlDscrtInptStsV"`
	F_BrkSysBrkLghtsReqd     uint8   `json:"BrkSysBrkLghtsReqd"`
	F_EPBSysBrkLghtsReqd     uint8   `json:"EPBSysBrkLghtsReqd"`
	F_EPBSysBrkLghtsReqdA    uint8   `json:"EPBSysBrkLghtsReqdA"`
	F_BMSPtIsltnRstc         float32 `json:"BMSPtIsltnRstc"`
	F_EPTAccelActuPos        float32 `json:"EPTAccelActuPos"`
	F_EPTAccelActuPosV       uint8   `json:"EPTAccelActuPosV"`
	F_TMInvtrCrntV           uint8   `json:"TMInvtrCrntV"`
	F_TMInvtrCrnt            int16   `json:"TMInvtrCrnt"`
	F_ISGInvtrCrntV          uint8   `json:"ISGInvtrCrntV"`
	F_ISGInvtrCrnt           int16   `json:"ISGInvtrCrnt"`
	F_SAMInvtrCrnt           int16   `json:"SAMInvtrCrnt"`
	F_SAMInvtrCrntV          uint8   `json:"SAMInvtrCrntV"`
	F_TMSta                  uint8   `json:"TMSta"`
	F_ISGSta                 uint8   `json:"ISGSta"`
	F_SAMSta                 uint8   `json:"SAMSta"`
	F_TMInvtrTem             int16   `json:"TMInvtrTem"`
	F_ISGInvtrTem            int16   `json:"ISGInvtrTem"`
	F_SAMInvtrTem            int16   `json:"SAMInvtrTem"`
	F_TMSpd                  int32   `json:"TMSpd"`
	F_TMSpdV                 int8    `json:"TMSpdV"`
	F_ISGSpd                 int32   `json:"ISGSpd"`
	F_ISGSpdV                int8    `json:"ISGSpdV"`
	F_SAMSpdV                int8    `json:"SAMSpdV"`
	F_SAMSpd                 int32   `json:"SAMSpd"`
	F_TMActuToq              float32 `json:"TMActuToq"`
	F_TMActuToqV             int8    `json:"TMActuToqV"`
	F_ISGActuToq             float32 `json:"ISGActuToq"`
	F_ISGActuToqV            int8    `json:"ISGActuToqV"`
	F_SAMActuToqV            int8    `json:"SAMActuToqV"`
	F_SAMActuToq             float32 `json:"SAMActuToq"`
	F_TMSttrTem              int16   `json:"TMSttrTem"`
	F_ISGSttrTem             int16   `json:"ISGSttrTem"`
	F_SAMSttrTem             int16   `json:"SAMSttrTem"`
	F_HVDCDCHVSideVol        uint16  `json:"HVDCDCHVSideVol"`
	F_HVDCDCHVSideVolV       uint8   `json:"HVDCDCHVSideVolV"`
	F_AvgFuelCsump           float32 `json:"AvgFuelCsump"`
	F_TMInvtrVolV            uint8   `json:"TMInvtrVolV"`
	F_TMInvtrVol             uint16  `json:"TMInvtrVol"`
	F_ISGInvtrVolV           uint8   `json:"ISGInvtrVolV"`
	F_ISGInvtrVol            uint16  `json:"ISGInvtrVol"`
	F_SAMInvtrVolV           uint8   `json:"SAMInvtrVolV"`
	F_SAMInvtrVol            uint16  `json:"SAMInvtrVol"`
	F_BMSCellMaxTemIndx      uint8   `json:"BMSCellMaxTemIndx"`
	F_BMSCellMaxTem          float32 `json:"BMSCellMaxTem"`
	F_BMSCellMaxTemV         uint8   `json:"BMSCellMaxTemV"`
	F_BMSCellMinTemIndx      uint8   `json:"BMSCellMinTemIndx"`
	F_BMSCellMinTem          float32 `json:"BMSCellMinTem"`
	F_BMSCellMinTemV         uint8   `json:"BMSCellMinTemV"`
	F_BMSCellMaxVolIndx      uint8   `json:"BMSCellMaxVolIndx"`
	F_BMSCellMaxVol          float32 `json:"BMSCellMaxVol"`
	F_BMSCellMaxVolV         uint8   `json:"BMSCellMaxVolV"`
	F_BMSCellMinVolIndx      uint8   `json:"BMSCellMinVolIndx"`
	F_BMSCellMinVol          float32 `json:"BMSCellMinVol"`
	F_BMSCellMinVolV         uint8   `json:"BMSCellMinVolV"`
	F_BMSPtIsltnRstcV        uint8   `json:"BMSPtIsltnRstcV"`
	F_HVDCDCTem              int16   `json:"HVDCDCTem"`
	F_BrkFludLvlLow          uint8   `json:"BrkFludLvlLow"`
	F_BrkSysRedBrkTlltReq    uint8   `json:"BrkSysRedBrkTlltReq"`
	F_ABSF                   uint8   `json:"ABSF"`
	F_VSESts                 uint8   `json:"VSESts"`
	F_IbstrWrnngIO           uint8   `json:"IbstrWrnngIO"`
	F_BMSHVILClsd            uint8   `json:"BMSHVILClsd"`
	F_EPTTrOtptShaftTotToq   float32 `json:"EPTTrOtptShaftTotToq"`
	F_EPTTrOtptShaftTotToqV  uint8   `json:"EPTTrOtptShaftTotToqV"`
	F_BrkFludLvlLowV         uint8   `json:"BrkFludLvlLowV"`
	F_EnSpd                  float32 `json:"EnSpd"`
	F_EnSpdSts               uint8   `json:"EnSpdSts"`
	F_FuelCsump              uint16  `json:"FuelCsump"`
}

func To_Evt_D006(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D006 {
	_instance := Evt_D006{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_EPTRdy_v := _bitBuf.Read(1, true, true)
	_instance.F_EPTRdy = uint8(F_EPTRdy_v)

	F_BMSBscSta_v := _bitBuf.Read(5, true, true)
	_instance.F_BMSBscSta = uint8(F_BMSBscSta_v)

	F_BMSPackCrnt_v := _bitBuf.Read(16, true, true)
	_instance.F_BMSPackCrnt = float32(F_BMSPackCrnt_v)*0.05 - 1000

	F_BMSPackCrntV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPackCrntV = uint8(F_BMSPackCrntV_v)

	F_BMSPackSOC_v := _bitBuf.Read(10, true, true)
	_instance.F_BMSPackSOC = float32(F_BMSPackSOC_v) * 0.1

	F_BMSPackSOCV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPackSOCV = uint8(F_BMSPackSOCV_v)

	F_BMSPackSOCDsp_v := _bitBuf.Read(10, true, true)
	_instance.F_BMSPackSOCDsp = float32(F_BMSPackSOCDsp_v) * 0.1

	F_BMSPackSOCDspV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPackSOCDspV = uint8(F_BMSPackSOCDspV_v)

	F_ElecVehSysMd_v := _bitBuf.Read(4, true, true)
	_instance.F_ElecVehSysMd = uint8(F_ElecVehSysMd_v)

	F_BMSPackVol_v := _bitBuf.Read(12, true, true)
	_instance.F_BMSPackVol = float32(F_BMSPackVol_v) * 0.25

	F_BMSPackVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPackVolV = uint8(F_BMSPackVolV_v)

	F_HVDCDCSta_v := _bitBuf.Read(3, true, true)
	_instance.F_HVDCDCSta = uint8(F_HVDCDCSta_v)

	F_EPTTrInptShaftToq_v := _bitBuf.Read(12, true, true)
	_instance.F_EPTTrInptShaftToq = float32(F_EPTTrInptShaftToq_v)*0.5 - 848

	F_EPTTrInptShaftToqV_v := _bitBuf.Read(1, true, true)
	_instance.F_EPTTrInptShaftToqV = uint8(F_EPTTrInptShaftToqV_v)

	F_EPTTrOtptShaftToq_v := _bitBuf.Read(12, true, true)
	_instance.F_EPTTrOtptShaftToq = int16(F_EPTTrOtptShaftToq_v)*2 - 3392

	F_EPTTrOtptShaftToqV_v := _bitBuf.Read(1, true, true)
	_instance.F_EPTTrOtptShaftToqV = uint8(F_EPTTrOtptShaftToqV_v)

	F_EPTBrkPdlDscrtInptSts_v := _bitBuf.Read(1, true, true)
	_instance.F_EPTBrkPdlDscrtInptSts = uint8(F_EPTBrkPdlDscrtInptSts_v)

	F_EPTBrkPdlDscrtInptStsV_v := _bitBuf.Read(1, true, true)
	_instance.F_EPTBrkPdlDscrtInptStsV = uint8(F_EPTBrkPdlDscrtInptStsV_v)

	F_BrkSysBrkLghtsReqd_v := _bitBuf.Read(1, true, true)
	_instance.F_BrkSysBrkLghtsReqd = uint8(F_BrkSysBrkLghtsReqd_v)

	F_EPBSysBrkLghtsReqd_v := _bitBuf.Read(1, true, true)
	_instance.F_EPBSysBrkLghtsReqd = uint8(F_EPBSysBrkLghtsReqd_v)

	F_EPBSysBrkLghtsReqdA_v := _bitBuf.Read(1, true, true)
	_instance.F_EPBSysBrkLghtsReqdA = uint8(F_EPBSysBrkLghtsReqdA_v)

	F_BMSPtIsltnRstc_v := _bitBuf.Read(14, true, true)
	_instance.F_BMSPtIsltnRstc = float32(F_BMSPtIsltnRstc_v) * 0.5

	F_EPTAccelActuPos_v := _bitBuf.Read(8, true, true)
	_instance.F_EPTAccelActuPos = float32(F_EPTAccelActuPos_v) * 0.392157

	F_EPTAccelActuPosV_v := _bitBuf.Read(1, true, true)
	_instance.F_EPTAccelActuPosV = uint8(F_EPTAccelActuPosV_v)

	F_TMInvtrCrntV_v := _bitBuf.Read(1, true, true)
	_instance.F_TMInvtrCrntV = uint8(F_TMInvtrCrntV_v)

	F_TMInvtrCrnt_v := _bitBuf.Read(11, true, true)
	_instance.F_TMInvtrCrnt = int16(F_TMInvtrCrnt_v) - 1024

	F_ISGInvtrCrntV_v := _bitBuf.Read(1, true, true)
	_instance.F_ISGInvtrCrntV = uint8(F_ISGInvtrCrntV_v)

	F_ISGInvtrCrnt_v := _bitBuf.Read(11, true, true)
	_instance.F_ISGInvtrCrnt = int16(F_ISGInvtrCrnt_v) - 1024

	F_SAMInvtrCrnt_v := _bitBuf.Read(11, true, true)
	_instance.F_SAMInvtrCrnt = int16(F_SAMInvtrCrnt_v) - 1024

	F_SAMInvtrCrntV_v := _bitBuf.Read(1, true, true)
	_instance.F_SAMInvtrCrntV = uint8(F_SAMInvtrCrntV_v)

	F_TMSta_v := _bitBuf.Read(4, true, true)
	_instance.F_TMSta = uint8(F_TMSta_v)

	F_ISGSta_v := _bitBuf.Read(4, true, true)
	_instance.F_ISGSta = uint8(F_ISGSta_v)

	F_SAMSta_v := _bitBuf.Read(4, true, true)
	_instance.F_SAMSta = uint8(F_SAMSta_v)

	F_TMInvtrTem_v := _bitBuf.Read(8, true, true)
	_instance.F_TMInvtrTem = int16(F_TMInvtrTem_v) - 40

	F_ISGInvtrTem_v := _bitBuf.Read(8, true, true)
	_instance.F_ISGInvtrTem = int16(F_ISGInvtrTem_v) - 40

	F_SAMInvtrTem_v := _bitBuf.Read(8, true, true)
	_instance.F_SAMInvtrTem = int16(F_SAMInvtrTem_v) - 40

	F_TMSpd_v := _bitBuf.Read(16, true, true)
	_instance.F_TMSpd = int32(F_TMSpd_v) - 32768

	F_TMSpdV_v := _bitBuf.Read(1, true, true)
	_instance.F_TMSpdV = int8(F_TMSpdV_v)

	F_ISGSpd_v := _bitBuf.Read(16, true, true)
	_instance.F_ISGSpd = int32(F_ISGSpd_v) - 32768

	F_ISGSpdV_v := _bitBuf.Read(1, true, true)
	_instance.F_ISGSpdV = int8(F_ISGSpdV_v)

	F_SAMSpdV_v := _bitBuf.Read(1, true, true)
	_instance.F_SAMSpdV = int8(F_SAMSpdV_v)

	F_SAMSpd_v := _bitBuf.Read(16, true, true)
	_instance.F_SAMSpd = int32(F_SAMSpd_v) - 32768

	F_TMActuToq_v := _bitBuf.Read(11, true, true)
	_instance.F_TMActuToq = float32(F_TMActuToq_v)*0.5 - 512

	F_TMActuToqV_v := _bitBuf.Read(1, true, true)
	_instance.F_TMActuToqV = int8(F_TMActuToqV_v)

	F_ISGActuToq_v := _bitBuf.Read(11, true, true)
	_instance.F_ISGActuToq = float32(F_ISGActuToq_v)*0.5 - 512

	F_ISGActuToqV_v := _bitBuf.Read(1, true, true)
	_instance.F_ISGActuToqV = int8(F_ISGActuToqV_v)

	F_SAMActuToqV_v := _bitBuf.Read(1, true, true)
	_instance.F_SAMActuToqV = int8(F_SAMActuToqV_v)

	F_SAMActuToq_v := _bitBuf.Read(11, true, true)
	_instance.F_SAMActuToq = float32(F_SAMActuToq_v)*0.5 - 512

	F_TMSttrTem_v := _bitBuf.Read(8, true, true)
	_instance.F_TMSttrTem = int16(F_TMSttrTem_v) - 40

	F_ISGSttrTem_v := _bitBuf.Read(8, true, true)
	_instance.F_ISGSttrTem = int16(F_ISGSttrTem_v) - 40

	F_SAMSttrTem_v := _bitBuf.Read(8, true, true)
	_instance.F_SAMSttrTem = int16(F_SAMSttrTem_v) - 40

	F_HVDCDCHVSideVol_v := _bitBuf.Read(10, true, true)
	_instance.F_HVDCDCHVSideVol = uint16(F_HVDCDCHVSideVol_v)

	F_HVDCDCHVSideVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_HVDCDCHVSideVolV = uint8(F_HVDCDCHVSideVolV_v)

	F_AvgFuelCsump_v := _bitBuf.Read(8, true, true)
	_instance.F_AvgFuelCsump = float32(F_AvgFuelCsump_v) * 0.1

	F_TMInvtrVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_TMInvtrVolV = uint8(F_TMInvtrVolV_v)

	F_TMInvtrVol_v := _bitBuf.Read(10, true, true)
	_instance.F_TMInvtrVol = uint16(F_TMInvtrVol_v)

	F_ISGInvtrVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_ISGInvtrVolV = uint8(F_ISGInvtrVolV_v)

	F_ISGInvtrVol_v := _bitBuf.Read(10, true, true)
	_instance.F_ISGInvtrVol = uint16(F_ISGInvtrVol_v)

	F_SAMInvtrVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_SAMInvtrVolV = uint8(F_SAMInvtrVolV_v)

	F_SAMInvtrVol_v := _bitBuf.Read(10, true, true)
	_instance.F_SAMInvtrVol = uint16(F_SAMInvtrVol_v)

	F_BMSCellMaxTemIndx_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSCellMaxTemIndx = uint8(F_BMSCellMaxTemIndx_v)

	F_BMSCellMaxTem_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSCellMaxTem = float32(F_BMSCellMaxTem_v)*0.5 - 40

	F_BMSCellMaxTemV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSCellMaxTemV = uint8(F_BMSCellMaxTemV_v)

	F_BMSCellMinTemIndx_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSCellMinTemIndx = uint8(F_BMSCellMinTemIndx_v)

	F_BMSCellMinTem_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSCellMinTem = float32(F_BMSCellMinTem_v)*0.5 - 40

	F_BMSCellMinTemV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSCellMinTemV = uint8(F_BMSCellMinTemV_v)

	F_BMSCellMaxVolIndx_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSCellMaxVolIndx = uint8(F_BMSCellMaxVolIndx_v)

	F_BMSCellMaxVol_v := _bitBuf.Read(13, true, true)
	_instance.F_BMSCellMaxVol = float32(F_BMSCellMaxVol_v) * 0.001

	F_BMSCellMaxVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSCellMaxVolV = uint8(F_BMSCellMaxVolV_v)

	F_BMSCellMinVolIndx_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSCellMinVolIndx = uint8(F_BMSCellMinVolIndx_v)

	F_BMSCellMinVol_v := _bitBuf.Read(13, true, true)
	_instance.F_BMSCellMinVol = float32(F_BMSCellMinVol_v) * 0.001

	F_BMSCellMinVolV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSCellMinVolV = uint8(F_BMSCellMinVolV_v)

	F_BMSPtIsltnRstcV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPtIsltnRstcV = uint8(F_BMSPtIsltnRstcV_v)

	F_HVDCDCTem_v := _bitBuf.Read(8, true, true)
	_instance.F_HVDCDCTem = int16(F_HVDCDCTem_v) - 40

	F_BrkFludLvlLow_v := _bitBuf.Read(1, true, true)
	_instance.F_BrkFludLvlLow = uint8(F_BrkFludLvlLow_v)

	F_BrkSysRedBrkTlltReq_v := _bitBuf.Read(1, true, true)
	_instance.F_BrkSysRedBrkTlltReq = uint8(F_BrkSysRedBrkTlltReq_v)

	F_ABSF_v := _bitBuf.Read(1, true, true)
	_instance.F_ABSF = uint8(F_ABSF_v)

	F_VSESts_v := _bitBuf.Read(3, true, true)
	_instance.F_VSESts = uint8(F_VSESts_v)

	F_IbstrWrnngIO_v := _bitBuf.Read(1, true, true)
	_instance.F_IbstrWrnngIO = uint8(F_IbstrWrnngIO_v)

	F_BMSHVILClsd_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSHVILClsd = uint8(F_BMSHVILClsd_v)

	F_EPTTrOtptShaftTotToq_v := _bitBuf.Read(12, true, true)
	_instance.F_EPTTrOtptShaftTotToq = float32(F_EPTTrOtptShaftTotToq_v)*0.5 - 848

	F_EPTTrOtptShaftTotToqV_v := _bitBuf.Read(1, true, true)
	_instance.F_EPTTrOtptShaftTotToqV = uint8(F_EPTTrOtptShaftTotToqV_v)

	F_BrkFludLvlLowV_v := _bitBuf.Read(1, true, true)
	_instance.F_BrkFludLvlLowV = uint8(F_BrkFludLvlLowV_v)

	F_EnSpd_v := _bitBuf.Read(16, true, true)
	_instance.F_EnSpd = float32(F_EnSpd_v) * 0.25

	F_EnSpdSts_v := _bitBuf.Read(2, true, true)
	_instance.F_EnSpdSts = uint8(F_EnSpdSts_v)

	F_FuelCsump_v := _bitBuf.Read(12, true, true)
	_bitBuf.Finish()
	_instance.F_FuelCsump = uint16(F_FuelCsump_v) * 16

	return &_instance
}

func (__instance *Evt_D006) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_EPTRdy), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBscSta), 5, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_BMSPackCrnt+1000)/0.05)), 16, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackCrntV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSPackSOC/0.1)), 10, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackSOCV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSPackSOCDsp/0.1)), 10, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackSOCDspV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_ElecVehSysMd), 4, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSPackVol/0.25)), 12, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_HVDCDCSta), 3, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_EPTTrInptShaftToq+848)/0.5)), 12, true, true)
	_bitBuf.Write(int64(_instance.F_EPTTrInptShaftToqV), 1, true, true)
	_bitBuf.Write(int64((_instance.F_EPTTrOtptShaftToq+3392)/2), 12, true, true)
	_bitBuf.Write(int64(_instance.F_EPTTrOtptShaftToqV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_EPTBrkPdlDscrtInptSts), 1, true, true)
	_bitBuf.Write(int64(_instance.F_EPTBrkPdlDscrtInptStsV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BrkSysBrkLghtsReqd), 1, true, true)
	_bitBuf.Write(int64(_instance.F_EPBSysBrkLghtsReqd), 1, true, true)
	_bitBuf.Write(int64(_instance.F_EPBSysBrkLghtsReqdA), 1, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSPtIsltnRstc/0.5)), 14, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_EPTAccelActuPos/0.392157)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_EPTAccelActuPosV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_TMInvtrCrntV), 1, true, true)
	_bitBuf.Write(int64((_instance.F_TMInvtrCrnt + 1024)), 11, true, true)
	_bitBuf.Write(int64(_instance.F_ISGInvtrCrntV), 1, true, true)
	_bitBuf.Write(int64((_instance.F_ISGInvtrCrnt + 1024)), 11, true, true)
	_bitBuf.Write(int64((_instance.F_SAMInvtrCrnt + 1024)), 11, true, true)
	_bitBuf.Write(int64(_instance.F_SAMInvtrCrntV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_TMSta), 4, true, true)
	_bitBuf.Write(int64(_instance.F_ISGSta), 4, true, true)
	_bitBuf.Write(int64(_instance.F_SAMSta), 4, true, true)
	_bitBuf.Write(int64((_instance.F_TMInvtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.F_ISGInvtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.F_SAMInvtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.F_TMSpd + 32768)), 16, true, true)
	_bitBuf.Write(int64(_instance.F_TMSpdV), 1, true, true)
	_bitBuf.Write(int64((_instance.F_ISGSpd + 32768)), 16, true, true)
	_bitBuf.Write(int64(_instance.F_ISGSpdV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_SAMSpdV), 1, true, true)
	_bitBuf.Write(int64((_instance.F_SAMSpd + 32768)), 16, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_TMActuToq+512)/0.5)), 11, true, true)
	_bitBuf.Write(int64(_instance.F_TMActuToqV), 1, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_ISGActuToq+512)/0.5)), 11, true, true)
	_bitBuf.Write(int64(_instance.F_ISGActuToqV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_SAMActuToqV), 1, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_SAMActuToq+512)/0.5)), 11, true, true)
	_bitBuf.Write(int64((_instance.F_TMSttrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.F_ISGSttrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.F_SAMSttrTem + 40)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_HVDCDCHVSideVol), 10, true, true)
	_bitBuf.Write(int64(_instance.F_HVDCDCHVSideVolV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_AvgFuelCsump/0.1)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_TMInvtrVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_TMInvtrVol), 10, true, true)
	_bitBuf.Write(int64(_instance.F_ISGInvtrVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_ISGInvtrVol), 10, true, true)
	_bitBuf.Write(int64(_instance.F_SAMInvtrVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_SAMInvtrVol), 10, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMaxTemIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_BMSCellMaxTem+40)/0.5)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMaxTemV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMinTemIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_BMSCellMinTem+40)/0.5)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMinTemV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMaxVolIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSCellMaxVol/0.001)), 13, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMaxVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMinVolIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSCellMinVol/0.001)), 13, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellMinVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPtIsltnRstcV), 1, true, true)
	_bitBuf.Write(int64((_instance.F_HVDCDCTem + 40)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BrkFludLvlLow), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BrkSysRedBrkTlltReq), 1, true, true)
	_bitBuf.Write(int64(_instance.F_ABSF), 1, true, true)
	_bitBuf.Write(int64(_instance.F_VSESts), 3, true, true)
	_bitBuf.Write(int64(_instance.F_IbstrWrnngIO), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSHVILClsd), 1, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_EPTTrOtptShaftTotToq+848)/0.5)), 12, true, true)
	_bitBuf.Write(int64(_instance.F_EPTTrOtptShaftTotToqV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BrkFludLvlLowV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_EnSpd/0.25)), 16, true, true)
	_bitBuf.Write(int64(_instance.F_EnSpdSts), 2, true, true)
	_bitBuf.Write(int64(_instance.F_FuelCsump/16), 12, true, true)
	_bitBuf.Finish()
}

type Evt_D008 struct {
	F_evtId             uint16 `json:"evtId"`
	F_evtLen            uint16 `json:"evtLen"`
	F_DTCInfomationBMS  int64  `json:"DTCInfomationBMS"`
	F_DTCInfomationECM  int64  `json:"DTCInfomationECM"`
	F_DTCInfomationEPB  int64  `json:"DTCInfomationEPB"`
	F_DTCInfomationPLCM int64  `json:"DTCInfomationPLCM"`
	F_DTCInfomationTCM  int64  `json:"DTCInfomationTCM"`
	F_DTCInfomationTPMS int64  `json:"DTCInfomationTPMS"`
	F_DTCInfomationTC   int64  `json:"DTCInfomationTC"`
	F_DTCInfomationISC  int64  `json:"DTCInfomationISC"`
	F_DTCInfomationSAC  int64  `json:"DTCInfomationSAC"`
	F_DTCInfomationIMCU int64  `json:"DTCInfomationIMCU"`
}

func To_Evt_D008(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D008 {
	_instance := Evt_D008{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_DTCInfomationBMS_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationBMS = int64(F_DTCInfomationBMS_v)

	F_DTCInfomationECM_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationECM = int64(F_DTCInfomationECM_v)

	F_DTCInfomationEPB_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationEPB = int64(F_DTCInfomationEPB_v)

	F_DTCInfomationPLCM_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationPLCM = int64(F_DTCInfomationPLCM_v)

	F_DTCInfomationTCM_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationTCM = int64(F_DTCInfomationTCM_v)

	F_DTCInfomationTPMS_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationTPMS = int64(F_DTCInfomationTPMS_v)

	F_DTCInfomationTC_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationTC = int64(F_DTCInfomationTC_v)

	F_DTCInfomationISC_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationISC = int64(F_DTCInfomationISC_v)

	F_DTCInfomationSAC_v := _bitBuf.Read(56, true, true)
	_instance.F_DTCInfomationSAC = int64(F_DTCInfomationSAC_v)

	F_DTCInfomationIMCU_v := _bitBuf.Read(56, true, true)
	_bitBuf.Finish()
	_instance.F_DTCInfomationIMCU = int64(F_DTCInfomationIMCU_v)

	return &_instance
}

func (__instance *Evt_D008) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_DTCInfomationBMS), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationECM), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationEPB), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationPLCM), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationTCM), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationTPMS), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationTC), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationISC), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationSAC), 56, true, true)
	_bitBuf.Write(int64(_instance.F_DTCInfomationIMCU), 56, true, true)
	_bitBuf.Finish()
}

type Evt_D009 struct {
	F_evtId                  uint16 `json:"evtId"`
	F_evtLen                 uint16 `json:"evtLen"`
	F_BMSCMUFlt              uint8  `json:"BMSCMUFlt"`
	F_BMSCellVoltFlt         uint8  `json:"BMSCellVoltFlt"`
	F_BMSPackTemFlt          uint8  `json:"BMSPackTemFlt"`
	F_BMSPackVoltFlt         uint8  `json:"BMSPackVoltFlt"`
	F_BMSWrnngInfo           uint8  `json:"BMSWrnngInfo"`
	F_BMSWrnngInfoPV         uint8  `json:"BMSWrnngInfoPV"`
	F_BMSWrnngInfoRC         uint8  `json:"BMSWrnngInfoRC"`
	F_BMSPreThrmFltInd       uint8  `json:"BMSPreThrmFltInd"`
	F_BMSKeepSysAwkScene     uint8  `json:"BMSKeepSysAwkScene"`
	F_BMSTemOverDifAlrm      uint8  `json:"BMSTemOverDifAlrm"`
	F_BMSOverTemAlrm         uint8  `json:"BMSOverTemAlrm"`
	F_BMSOverPackVolAlrm     uint8  `json:"BMSOverPackVolAlrm"`
	F_BMSUnderPackVolAlrm    uint8  `json:"BMSUnderPackVolAlrm"`
	F_BMSHVILAlrm            uint8  `json:"BMSHVILAlrm"`
	F_BMSOverCellVolAlrm     uint8  `json:"BMSOverCellVolAlrm"`
	F_BMSUnderCellVolAlrm    uint8  `json:"BMSUnderCellVolAlrm"`
	F_BMSLowSOCAlrm          uint8  `json:"BMSLowSOCAlrm"`
	F_BMSJumpngSOCAlrm       uint8  `json:"BMSJumpngSOCAlrm"`
	F_BMSHiSOCAlrm           uint8  `json:"BMSHiSOCAlrm"`
	F_BMSPackVolMsmchAlrm    uint8  `json:"BMSPackVolMsmchAlrm"`
	F_BMSPoorCellCnstncyAlrm uint8  `json:"BMSPoorCellCnstncyAlrm"`
	F_BMSCellOverChrgdAlrm   uint8  `json:"BMSCellOverChrgdAlrm"`
	F_BMSLowPtIsltnRstcAlrm  uint8  `json:"BMSLowPtIsltnRstcAlrm"`
	F_TMRtrTem               int16  `json:"TMRtrTem"`
	F_TMStrOvTempAlrm        uint8  `json:"TMStrOvTempAlrm"`
	F_TMInvtrOvTempAlrm      uint8  `json:"TMInvtrOvTempAlrm"`
	F_ISCStrOvTempAlrm       uint8  `json:"ISCStrOvTempAlrm"`
	F_ISCInvtrOvTempAlrm     uint8  `json:"ISCInvtrOvTempAlrm"`
	F_SAMStrOvTempAlrm       uint8  `json:"SAMStrOvTempAlrm"`
	F_SAMInvtrOvTempAlrm     uint8  `json:"SAMInvtrOvTempAlrm"`
	F_EPTHVDCDCMdReq         uint8  `json:"EPTHVDCDCMdReq"`
	F_VCUSecyWrnngInfo       uint8  `json:"VCUSecyWrnngInfo"`
	F_VCUSecyWrnngInfoPV     uint8  `json:"VCUSecyWrnngInfoPV"`
	F_VCUSecyWrnngInfoRC     uint8  `json:"VCUSecyWrnngInfoRC"`
	F_VCUSecyWrnngInfoCRC    uint8  `json:"VCUSecyWrnngInfoCRC"`
	F_BMSOnbdChrgSpRsn       uint8  `json:"BMSOnbdChrgSpRsn"`
}

func To_Evt_D009(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D009 {
	_instance := Evt_D009{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_BMSCMUFlt_v := _bitBuf.Read(2, true, true)
	_instance.F_BMSCMUFlt = uint8(F_BMSCMUFlt_v)

	F_BMSCellVoltFlt_v := _bitBuf.Read(2, true, true)
	_instance.F_BMSCellVoltFlt = uint8(F_BMSCellVoltFlt_v)

	F_BMSPackTemFlt_v := _bitBuf.Read(2, true, true)
	_instance.F_BMSPackTemFlt = uint8(F_BMSPackTemFlt_v)

	F_BMSPackVoltFlt_v := _bitBuf.Read(2, true, true)
	_instance.F_BMSPackVoltFlt = uint8(F_BMSPackVoltFlt_v)

	F_BMSWrnngInfo_v := _bitBuf.Read(6, true, true)
	_instance.F_BMSWrnngInfo = uint8(F_BMSWrnngInfo_v)

	F_BMSWrnngInfoPV_v := _bitBuf.Read(6, true, true)
	_instance.F_BMSWrnngInfoPV = uint8(F_BMSWrnngInfoPV_v)

	F_BMSWrnngInfoRC_v := _bitBuf.Read(4, true, true)
	_instance.F_BMSWrnngInfoRC = uint8(F_BMSWrnngInfoRC_v)

	F_BMSPreThrmFltInd_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPreThrmFltInd = uint8(F_BMSPreThrmFltInd_v)

	_bitBuf.Skip(5)
	F_BMSKeepSysAwkScene_v := _bitBuf.Read(4, true, true)
	_instance.F_BMSKeepSysAwkScene = uint8(F_BMSKeepSysAwkScene_v)

	F_BMSTemOverDifAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSTemOverDifAlrm = uint8(F_BMSTemOverDifAlrm_v)

	F_BMSOverTemAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSOverTemAlrm = uint8(F_BMSOverTemAlrm_v)

	F_BMSOverPackVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSOverPackVolAlrm = uint8(F_BMSOverPackVolAlrm_v)

	F_BMSUnderPackVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSUnderPackVolAlrm = uint8(F_BMSUnderPackVolAlrm_v)

	F_BMSHVILAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSHVILAlrm = uint8(F_BMSHVILAlrm_v)

	F_BMSOverCellVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSOverCellVolAlrm = uint8(F_BMSOverCellVolAlrm_v)

	F_BMSUnderCellVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSUnderCellVolAlrm = uint8(F_BMSUnderCellVolAlrm_v)

	F_BMSLowSOCAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSLowSOCAlrm = uint8(F_BMSLowSOCAlrm_v)

	F_BMSJumpngSOCAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSJumpngSOCAlrm = uint8(F_BMSJumpngSOCAlrm_v)

	F_BMSHiSOCAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSHiSOCAlrm = uint8(F_BMSHiSOCAlrm_v)

	F_BMSPackVolMsmchAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSPackVolMsmchAlrm = uint8(F_BMSPackVolMsmchAlrm_v)

	F_BMSPoorCellCnstncyAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSPoorCellCnstncyAlrm = uint8(F_BMSPoorCellCnstncyAlrm_v)

	F_BMSCellOverChrgdAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSCellOverChrgdAlrm = uint8(F_BMSCellOverChrgdAlrm_v)

	F_BMSLowPtIsltnRstcAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSLowPtIsltnRstcAlrm = uint8(F_BMSLowPtIsltnRstcAlrm_v)

	F_TMRtrTem_v := _bitBuf.Read(8, true, true)
	_instance.F_TMRtrTem = int16(F_TMRtrTem_v) - 40

	F_TMStrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_TMStrOvTempAlrm = uint8(F_TMStrOvTempAlrm_v)

	F_TMInvtrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_TMInvtrOvTempAlrm = uint8(F_TMInvtrOvTempAlrm_v)

	F_ISCStrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_ISCStrOvTempAlrm = uint8(F_ISCStrOvTempAlrm_v)

	F_ISCInvtrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_ISCInvtrOvTempAlrm = uint8(F_ISCInvtrOvTempAlrm_v)

	F_SAMStrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_SAMStrOvTempAlrm = uint8(F_SAMStrOvTempAlrm_v)

	F_SAMInvtrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.F_SAMInvtrOvTempAlrm = uint8(F_SAMInvtrOvTempAlrm_v)

	F_EPTHVDCDCMdReq_v := _bitBuf.Read(3, true, true)
	_instance.F_EPTHVDCDCMdReq = uint8(F_EPTHVDCDCMdReq_v)

	F_VCUSecyWrnngInfo_v := _bitBuf.Read(6, true, true)
	_instance.F_VCUSecyWrnngInfo = uint8(F_VCUSecyWrnngInfo_v)

	F_VCUSecyWrnngInfoPV_v := _bitBuf.Read(6, true, true)
	_instance.F_VCUSecyWrnngInfoPV = uint8(F_VCUSecyWrnngInfoPV_v)

	F_VCUSecyWrnngInfoRC_v := _bitBuf.Read(4, true, true)
	_instance.F_VCUSecyWrnngInfoRC = uint8(F_VCUSecyWrnngInfoRC_v)

	F_VCUSecyWrnngInfoCRC_v := _bitBuf.Read(8, true, true)
	_instance.F_VCUSecyWrnngInfoCRC = uint8(F_VCUSecyWrnngInfoCRC_v)

	F_BMSOnbdChrgSpRsn_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.F_BMSOnbdChrgSpRsn = uint8(F_BMSOnbdChrgSpRsn_v)

	return &_instance
}

func (__instance *Evt_D009) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_BMSCMUFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellVoltFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackTemFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackVoltFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.F_BMSWrnngInfo), 6, true, true)
	_bitBuf.Write(int64(_instance.F_BMSWrnngInfoPV), 6, true, true)
	_bitBuf.Write(int64(_instance.F_BMSWrnngInfoRC), 4, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPreThrmFltInd), 1, true, true)
	_bitBuf.Skip(5)
	_bitBuf.Write(int64(_instance.F_BMSKeepSysAwkScene), 4, true, true)
	_bitBuf.Write(int64(_instance.F_BMSTemOverDifAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSOverTemAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSOverPackVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSUnderPackVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSHVILAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSOverCellVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSUnderCellVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSLowSOCAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSJumpngSOCAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSHiSOCAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackVolMsmchAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPoorCellCnstncyAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellOverChrgdAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSLowPtIsltnRstcAlrm), 3, true, true)
	_bitBuf.Write(int64((_instance.F_TMRtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_TMStrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_TMInvtrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_ISCStrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_ISCInvtrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_SAMStrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_SAMInvtrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.F_EPTHVDCDCMdReq), 3, true, true)
	_bitBuf.Write(int64(_instance.F_VCUSecyWrnngInfo), 6, true, true)
	_bitBuf.Write(int64(_instance.F_VCUSecyWrnngInfoPV), 6, true, true)
	_bitBuf.Write(int64(_instance.F_VCUSecyWrnngInfoRC), 4, true, true)
	_bitBuf.Write(int64(_instance.F_VCUSecyWrnngInfoCRC), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BMSOnbdChrgSpRsn), 8, true, true)
	_bitBuf.Finish()
}

type Evt_D00A struct {
	F_evtId     uint16 `json:"evtId"`
	F_evtLen    uint16 `json:"evtLen"`
	F_VIN       string `json:"VIN"`
	F_IAMSN     string `json:"IAMSN"`
	F_EsimIccid string `json:"EsimIccid"`
	F_EsimID    string `json:"EsimID"`
}

func To_Evt_D00A(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00A {
	_instance := Evt_D00A{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	F_VIN_len := 17
	F_VIN_v := _byteBuf.Read_slice_uint8(F_VIN_len)
	_instance.F_VIN = string(F_VIN_v)

	F_IAMSN_len := 16
	F_IAMSN_v := _byteBuf.Read_slice_uint8(F_IAMSN_len)
	_instance.F_IAMSN = string(F_IAMSN_v)

	F_EsimIccid_len := 20
	F_EsimIccid_v := _byteBuf.Read_slice_uint8(F_EsimIccid_len)
	_instance.F_EsimIccid = string(F_EsimIccid_v)

	F_EsimID_len := 32
	F_EsimID_v := _byteBuf.Read_slice_uint8(F_EsimID_len)
	_instance.F_EsimID = string(F_EsimID_v)

	return &_instance
}

func (__instance *Evt_D00A) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_string_utf8(_instance.F_VIN)
	_byteBuf.Write_string_utf8(_instance.F_IAMSN)
	_byteBuf.Write_string_utf8(_instance.F_EsimIccid)
	_byteBuf.Write_string_utf8(_instance.F_EsimID)
}

type Evt_D00B_BMSCellVol struct {
	F_BMSCellVol  float32 `json:"BMSCellVol"`
	F_BMSCellVolV uint8   `json:"BMSCellVolV"`
}

func To_Evt_D00B_BMSCellVol(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00B_BMSCellVol {
	_instance := Evt_D00B_BMSCellVol{}
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_BMSCellVol_v := _bitBuf.Read(13, true, true)
	_instance.F_BMSCellVol = float32(F_BMSCellVol_v) * 0.001

	F_BMSCellVolV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.F_BMSCellVolV = uint8(F_BMSCellVolV_v)

	return &_instance
}

func (__instance *Evt_D00B_BMSCellVol) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSCellVol/0.001)), 13, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellVolV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_D00B struct {
	F_evtId            uint16                 `json:"evtId"`
	F_evtLen           uint16                 `json:"evtLen"`
	F_BMSCellVolSumNum uint8                  `json:"BMSCellVolSumNum"`
	F_BMSCellVols      []*Evt_D00B_BMSCellVol `json:"BMSCellVols"`
}

func To_Evt_D00B(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00B {
	_instance := Evt_D00B{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	F_BMSCellVolSumNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSCellVolSumNum = F_BMSCellVolSumNum_v

	F_BMSCellVols_len := (int)(F_BMSCellVolSumNum_v)
	F_BMSCellVols_arr := make([]*Evt_D00B_BMSCellVol, F_BMSCellVols_len, F_BMSCellVols_len)
	for i := 0; i < F_BMSCellVols_len; i++ {
		F_BMSCellVols_arr[i] = To_Evt_D00B_BMSCellVol(_byteBuf, nil)
	}
	_instance.F_BMSCellVols = F_BMSCellVols_arr
	return &_instance
}

func (__instance *Evt_D00B) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_uint8(_instance.F_BMSCellVolSumNum)
	F_BMSCellVols_arr := _instance.F_BMSCellVols
	for i := 0; i < len(F_BMSCellVols_arr); i++ {
		F_BMSCellVols_arr[i].Write(_byteBuf, nil)
	}
}

type Evt_D00C_BMSCellTem struct {
	F_BMSCellTem  float32 `json:"BMSCellTem"`
	F_BMSCellTemV uint8   `json:"BMSCellTemV"`
}

func To_Evt_D00C_BMSCellTem(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00C_BMSCellTem {
	_instance := Evt_D00C_BMSCellTem{}
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_BMSCellTem_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSCellTem = float32(F_BMSCellTem_v) - 40

	F_BMSCellTemV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.F_BMSCellTemV = uint8(F_BMSCellTemV_v)

	return &_instance
}

func (__instance *Evt_D00C_BMSCellTem) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round((_instance.F_BMSCellTem + 40))), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BMSCellTemV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_D00C struct {
	F_evtId            uint16                 `json:"evtId"`
	F_evtLen           uint16                 `json:"evtLen"`
	F_BMSCellTemSumNum uint8                  `json:"BMSCellTemSumNum"`
	F_BMSCellTems      []*Evt_D00C_BMSCellTem `json:"BMSCellTems"`
}

func To_Evt_D00C(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00C {
	_instance := Evt_D00C{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	F_BMSCellTemSumNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSCellTemSumNum = F_BMSCellTemSumNum_v

	F_BMSCellTems_len := (int)(F_BMSCellTemSumNum_v)
	F_BMSCellTems_arr := make([]*Evt_D00C_BMSCellTem, F_BMSCellTems_len, F_BMSCellTems_len)
	for i := 0; i < F_BMSCellTems_len; i++ {
		F_BMSCellTems_arr[i] = To_Evt_D00C_BMSCellTem(_byteBuf, nil)
	}
	_instance.F_BMSCellTems = F_BMSCellTems_arr
	return &_instance
}

func (__instance *Evt_D00C) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_uint8(_instance.F_BMSCellTemSumNum)
	F_BMSCellTems_arr := _instance.F_BMSCellTems
	for i := 0; i < len(F_BMSCellTems_arr); i++ {
		F_BMSCellTems_arr[i].Write(_byteBuf, nil)
	}
}

type Evt_D00D_BMSBusbarTem struct {
	F_BMSBusbarTem  float32 `json:"BMSBusbarTem"`
	F_BMSBusbarTemV uint8   `json:"BMSBusbarTemV"`
}

func To_Evt_D00D_BMSBusbarTem(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00D_BMSBusbarTem {
	_instance := Evt_D00D_BMSBusbarTem{}
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_BMSBusbarTem_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSBusbarTem = float32(F_BMSBusbarTem_v) - 40

	F_BMSBusbarTemV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.F_BMSBusbarTemV = uint8(F_BMSBusbarTemV_v)

	return &_instance
}

func (__instance *Evt_D00D_BMSBusbarTem) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round((_instance.F_BMSBusbarTem + 40))), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBusbarTemV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_D00D struct {
	F_evtId              uint16                   `json:"evtId"`
	F_evtLen             uint16                   `json:"evtLen"`
	F_BMSBusbarTemSumNum uint8                    `json:"BMSBusbarTemSumNum"`
	F_BMSBusbarTems      []*Evt_D00D_BMSBusbarTem `json:"BMSBusbarTems"`
}

func To_Evt_D00D(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00D {
	_instance := Evt_D00D{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	F_BMSBusbarTemSumNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSBusbarTemSumNum = F_BMSBusbarTemSumNum_v

	F_BMSBusbarTems_len := (int)(F_BMSBusbarTemSumNum_v)
	F_BMSBusbarTems_arr := make([]*Evt_D00D_BMSBusbarTem, F_BMSBusbarTems_len, F_BMSBusbarTems_len)
	for i := 0; i < F_BMSBusbarTems_len; i++ {
		F_BMSBusbarTems_arr[i] = To_Evt_D00D_BMSBusbarTem(_byteBuf, nil)
	}
	_instance.F_BMSBusbarTems = F_BMSBusbarTems_arr
	return &_instance
}

func (__instance *Evt_D00D) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_uint8(_instance.F_BMSBusbarTemSumNum)
	F_BMSBusbarTems_arr := _instance.F_BMSBusbarTems
	for i := 0; i < len(F_BMSBusbarTems_arr); i++ {
		F_BMSBusbarTems_arr[i].Write(_byteBuf, nil)
	}
}

type Evt_D00E struct {
	F_evtId            uint16             `json:"evtId"`
	F_evtLen           uint16             `json:"evtLen"`
	F_BMSRptBatCodeNum uint8              `json:"BMSRptBatCodeNum"`
	F_BMSRptBatCodeAsc parse.JsonUint8Arr `json:"BMSRptBatCodeAsc"`
}

func To_Evt_D00E(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00E {
	_instance := Evt_D00E{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	F_BMSRptBatCodeNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSRptBatCodeNum = F_BMSRptBatCodeNum_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_BMSRptBatCodeAsc_len := (int)(F_BMSRptBatCodeNum_v)
	F_BMSRptBatCodeAsc_arr := make([]uint8, F_BMSRptBatCodeAsc_len, F_BMSRptBatCodeAsc_len)
	for i := 0; i < F_BMSRptBatCodeAsc_len; i++ {
		e := _bitBuf.Read(8, true, true)
		F_BMSRptBatCodeAsc_arr[i] = uint8(e)
	}
	_instance.F_BMSRptBatCodeAsc = F_BMSRptBatCodeAsc_arr
	return &_instance
}

func (__instance *Evt_D00E) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_uint8(_instance.F_BMSRptBatCodeNum)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	F_BMSRptBatCodeAsc_arr := _instance.F_BMSRptBatCodeAsc
	for i := 0; i < len(F_BMSRptBatCodeAsc_arr); i++ {
		_bitBuf.Write(int64(F_BMSRptBatCodeAsc_arr[i]), 8, true, true)
	}
}

type Evt_D00F struct {
	F_evtId                uint16  `json:"evtId"`
	F_evtLen               uint16  `json:"evtLen"`
	F_BMSWrnngInfoCRC      uint8   `json:"BMSWrnngInfoCRC"`
	F_BMSBusbarTempMax     float32 `json:"BMSBusbarTempMax"`
	F_BMSPreThrmFltIndBkup uint8   `json:"BMSPreThrmFltIndBkup"`
	F_BMSWrnngInfoRCBkup   uint8   `json:"BMSWrnngInfoRCBkup"`
	F_BMSBatPrsFlt         uint8   `json:"BMSBatPrsFlt"`
	F_BMSWrnngInfoBkup     uint8   `json:"BMSWrnngInfoBkup"`
	F_BMSBatPrsAlrm        uint8   `json:"BMSBatPrsAlrm"`
	F_BMSBatPrsAlrmV       uint8   `json:"BMSBatPrsAlrmV"`
	F_BMSBatPrsSnsrV       uint8   `json:"BMSBatPrsSnsrV"`
	F_BMSBatPrsSnsrValBkup float32 `json:"BMSBatPrsSnsrValBkup"`
	F_BMSBatPrsSnsrVBkup   uint8   `json:"BMSBatPrsSnsrVBkup"`
	F_BMSBatPrsSnsrVal     float32 `json:"BMSBatPrsSnsrVal"`
	F_BMSClntPumpPWMReq    float32 `json:"BMSClntPumpPWMReq"`
	F_BMSPumpPwrOnReq      uint8   `json:"BMSPumpPwrOnReq"`
	F_BMSBatPrsAlrmVBkup   uint8   `json:"BMSBatPrsAlrmVBkup"`
	F_BMSBatPrsAlrmBkup    uint8   `json:"BMSBatPrsAlrmBkup"`
	F_BMSWrnngInfoCRCBkup  uint8   `json:"BMSWrnngInfoCRCBkup"`
	F_VCUBatPrsAlrm        uint8   `json:"VCUBatPrsAlrm"`
	F_OtsdAirTemCrVal      float32 `json:"OtsdAirTemCrVal"`
	F_VCUBatPrsAlrmV       uint8   `json:"VCUBatPrsAlrmV"`
	F_OtsdAirTemCrValV     uint8   `json:"OtsdAirTemCrValV"`
}

func To_Evt_D00F(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D00F {
	_instance := Evt_D00F{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_BMSWrnngInfoCRC_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSWrnngInfoCRC = uint8(F_BMSWrnngInfoCRC_v)

	F_BMSBusbarTempMax_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSBusbarTempMax = float32(F_BMSBusbarTempMax_v)*0.5 - 40

	F_BMSPreThrmFltIndBkup_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPreThrmFltIndBkup = uint8(F_BMSPreThrmFltIndBkup_v)

	F_BMSWrnngInfoRCBkup_v := _bitBuf.Read(4, true, true)
	_instance.F_BMSWrnngInfoRCBkup = uint8(F_BMSWrnngInfoRCBkup_v)

	F_BMSBatPrsFlt_v := _bitBuf.Read(3, true, true)
	_instance.F_BMSBatPrsFlt = uint8(F_BMSBatPrsFlt_v)

	F_BMSWrnngInfoBkup_v := _bitBuf.Read(6, true, true)
	_instance.F_BMSWrnngInfoBkup = uint8(F_BMSWrnngInfoBkup_v)

	F_BMSBatPrsAlrm_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSBatPrsAlrm = uint8(F_BMSBatPrsAlrm_v)

	F_BMSBatPrsAlrmV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSBatPrsAlrmV = uint8(F_BMSBatPrsAlrmV_v)

	F_BMSBatPrsSnsrV_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSBatPrsSnsrV = uint8(F_BMSBatPrsSnsrV_v)

	F_BMSBatPrsSnsrValBkup_v := _bitBuf.Read(15, true, true)
	_instance.F_BMSBatPrsSnsrValBkup = float32(F_BMSBatPrsSnsrValBkup_v) * 0.05

	F_BMSBatPrsSnsrVBkup_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSBatPrsSnsrVBkup = uint8(F_BMSBatPrsSnsrVBkup_v)

	F_BMSBatPrsSnsrVal_v := _bitBuf.Read(15, true, true)
	_instance.F_BMSBatPrsSnsrVal = float32(F_BMSBatPrsSnsrVal_v) * 0.05

	F_BMSClntPumpPWMReq_v := _bitBuf.Read(8, true, true)
	_instance.F_BMSClntPumpPWMReq = float32(F_BMSClntPumpPWMReq_v) * 0.4

	F_BMSPumpPwrOnReq_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSPumpPwrOnReq = uint8(F_BMSPumpPwrOnReq_v)

	F_BMSBatPrsAlrmVBkup_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSBatPrsAlrmVBkup = uint8(F_BMSBatPrsAlrmVBkup_v)

	F_BMSBatPrsAlrmBkup_v := _bitBuf.Read(1, true, true)
	_instance.F_BMSBatPrsAlrmBkup = uint8(F_BMSBatPrsAlrmBkup_v)

	F_BMSWrnngInfoCRCBkup_v := _bitBuf.Read(4, true, true)
	_instance.F_BMSWrnngInfoCRCBkup = uint8(F_BMSWrnngInfoCRCBkup_v)

	F_VCUBatPrsAlrm_v := _bitBuf.Read(1, true, true)
	_instance.F_VCUBatPrsAlrm = uint8(F_VCUBatPrsAlrm_v)

	F_OtsdAirTemCrVal_v := _bitBuf.Read(8, true, true)
	_instance.F_OtsdAirTemCrVal = float32(F_OtsdAirTemCrVal_v)*0.5 - 40

	F_VCUBatPrsAlrmV_v := _bitBuf.Read(1, true, true)
	_instance.F_VCUBatPrsAlrmV = uint8(F_VCUBatPrsAlrmV_v)

	F_OtsdAirTemCrValV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.F_OtsdAirTemCrValV = uint8(F_OtsdAirTemCrValV_v)

	return &_instance
}

func (__instance *Evt_D00F) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_BMSWrnngInfoCRC), 8, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_BMSBusbarTempMax+40)/0.5)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPreThrmFltIndBkup), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSWrnngInfoRCBkup), 4, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBatPrsFlt), 3, true, true)
	_bitBuf.Write(int64(_instance.F_BMSWrnngInfoBkup), 6, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBatPrsAlrm), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBatPrsAlrmV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBatPrsSnsrV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSBatPrsSnsrValBkup/0.05)), 15, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBatPrsSnsrVBkup), 1, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSBatPrsSnsrVal/0.05)), 15, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_BMSClntPumpPWMReq/0.4)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPumpPwrOnReq), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBatPrsAlrmVBkup), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSBatPrsAlrmBkup), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BMSWrnngInfoCRCBkup), 4, true, true)
	_bitBuf.Write(int64(_instance.F_VCUBatPrsAlrm), 1, true, true)
	_bitBuf.Write(int64(parse.Round((_instance.F_OtsdAirTemCrVal+40)/0.5)), 8, true, true)
	_bitBuf.Write(int64(_instance.F_VCUBatPrsAlrmV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_OtsdAirTemCrValV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_D01D struct {
	F_evtId     uint16 `json:"evtId"`
	F_evtLen    uint16 `json:"evtLen"`
	F_cellLAC5G uint32 `json:"cellLAC5G"`
	F_CellID5G  uint64 `json:"CellID5G"`
}

func To_Evt_D01D(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_D01D {
	return (*Evt_D01D)(unsafe.Pointer(unsafe.SliceData(_byteBuf.Read_slice_uint8(16))))
}
func (__instance *Evt_D01D) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_slice_uint8(*(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(__instance)),
		Len:  16,
		Cap:  16,
	})))
}

type Evt_FFFF struct {
	F_evtId  uint16 `json:"evtId"`
	F_EvtCRC uint64 `json:"EvtCRC"`
}

func To_Evt_FFFF(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Evt_FFFF {
	_instance := Evt_FFFF{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_EvtCRC_v := _bitBuf.Read(48, true, true)
	_bitBuf.Finish()
	_instance.F_EvtCRC = uint64(F_EvtCRC_v)

	return &_instance
}

func (__instance *Evt_FFFF) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_EvtCRC), 48, true, true)
	_bitBuf.Finish()
}

type Packet struct {
	F_evts any `json:"evts"`
}

func To_Packet(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Packet {
	_instance := Packet{}
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.F_evts = To_F_evts(_byteBuf, _parseContext)
	return &_instance
}

func (__instance *Packet) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Write_F_evts(_byteBuf, _instance.F_evts, _parseContext)
}

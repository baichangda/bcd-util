package immotors

import (
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"encoding/hex"
	"unsafe"
)

type Evt_0001 struct {
	F_evtId      uint16 `json:"evtId"`
	F_TBOXSysTim int64  `json:"TBOXSysTim"`
}

func To_Evt_0001(_byteBuf *parse.ByteBuf) *Evt_0001 {
	_instance := Evt_0001{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_TBOXSysTim = _bitBuf.Read(48, true, true)
	_bitBuf.Finish()

	return &_instance
}

func (__instance *Evt_0001) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(_instance.F_TBOXSysTim, 48, true, true)
	_bitBuf.Finish()
}

type Evt_0003 struct {
	F_evtId        uint16 `json:"evtId"`
	F_RelwakeupTim int64  `json:"RelwakeupTim"`
}

func To_Evt_0003(_byteBuf *parse.ByteBuf) *Evt_0003 {
	_instance := Evt_0003{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_RelwakeupTim = _bitBuf.Read(48, true, true)
	_bitBuf.Finish()

	return &_instance
}

func (__instance *Evt_0003) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(_instance.F_RelwakeupTim, 48, true, true)
	_bitBuf.Finish()
}

type Evt_0004 struct {
	F_evtId     uint16  `json:"evtId"`
	F_GnssAlt   float32 `json:"GnssAlt"`
	F_Longitude float64 `json:"Longitude"`
	F_GPSSts    uint8   `json:"GPSSts"`
}

func To_Evt_0004(_byteBuf *parse.ByteBuf) *Evt_0004 {
	_instance := Evt_0004{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_GnssAlt = float32(_byteBuf.Read_uint16())*0.1 - 500
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_Longitude = float64(_bitBuf.Read(29, true, false)) * 0.000001
	_instance.F_GPSSts = uint8(_bitBuf.Read(2, true, true))
	_bitBuf.Finish()

	return &_instance
}

func (__instance *Evt_0004) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.F_GnssAlt + 500) / 0.1)))
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_Longitude/0.000001)), 29, true, false)
	_bitBuf.Write(int64(_instance.F_GPSSts), 2, true, true)
	_bitBuf.Finish()
}

type Evt_0005 struct {
	F_evtId         uint16  `json:"evtId"`
	F_Latitude      float64 `json:"Latitude"`
	F_VehTyp        uint8   `json:"VehTyp"`
	F_GNSSDirection float32 `json:"GNSSDirection"`
}

func To_Evt_0005(_byteBuf *parse.ByteBuf) *Evt_0005 {
	_instance := Evt_0005{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_Latitude = float64(_bitBuf.Read(28, true, false)) * 0.000001
	_bitBuf.Skip(2)
	_instance.F_VehTyp = uint8(_bitBuf.Read(2, true, true))
	_bitBuf.Finish()
	_instance.F_GNSSDirection = float32(_byteBuf.Read_uint16()) * 0.01

	return &_instance
}

func (__instance *Evt_0005) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_Latitude/0.000001)), 28, true, false)
	_bitBuf.Skip(2)
	_bitBuf.Write(int64(_instance.F_VehTyp), 2, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.F_GNSSDirection / 0.01)))
}

type Evt_0006 struct {
	F_evtId uint16  `json:"evtId"`
	F_HDop  float32 `json:"HDop"`
	F_VDop  float32 `json:"VDop"`
}

func To_Evt_0006(_byteBuf *parse.ByteBuf) *Evt_0006 {
	_instance := Evt_0006{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_HDop = float32(_bitBuf.Read(24, true, true)) * 0.1
	_instance.F_VDop = float32(_bitBuf.Read(24, true, true)) * 0.1
	_bitBuf.Finish()

	return &_instance
}

func (__instance *Evt_0006) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_HDop/0.1)), 24, true, true)
	_bitBuf.Write(int64(parse.Round(_instance.F_VDop/0.1)), 24, true, true)
	_bitBuf.Finish()
}

type Evt_0007 struct {
	F_evtId uint16  `json:"evtId"`
	F_AcceX float32 `json:"AcceX"`
	F_AcceY float32 `json:"AcceY"`
	F_AcceZ float32 `json:"AcceZ"`
}

func To_Evt_0007(_byteBuf *parse.ByteBuf) *Evt_0007 {
	_instance := Evt_0007{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_AcceX = float32(_bitBuf.Read(14, true, false)) * 0.0009765625
	_bitBuf.Finish()
	_instance.F_AcceY = float32(_bitBuf.Read(14, true, false)) * 0.0009765625
	_bitBuf.Finish()
	_instance.F_AcceZ = float32(_bitBuf.Read(14, true, false)) * 0.0009765625
	_bitBuf.Finish()
	return &_instance
}

func (__instance *Evt_0007) Write(_byteBuf *parse.ByteBuf) {
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
	F_evtId       uint16 `json:"evtId"`
	F_cellMCC     uint16 `json:"cellMCC"`
	F_cellMNC     uint16 `json:"cellMNC"`
	F_millisecond uint16 `json:"millisecond"`
	F_spistatus   uint8  `json:"spistatus"`
}

func To_Evt_0008(_byteBuf *parse.ByteBuf) *Evt_0008 {
	_instance := Evt_0008{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_cellMCC = _byteBuf.Read_uint16()
	_instance.F_cellMNC = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_millisecond = uint16(_bitBuf.Read(10, true, true))
	_instance.F_spistatus = uint8(_bitBuf.Read(1, true, true))
	_bitBuf.Finish()
	return &_instance
}

func (__instance *Evt_0008) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_cellMCC)
	_byteBuf.Write_uint16(_instance.F_cellMNC)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_millisecond), 10, true, true)
	_bitBuf.Write(int64(_instance.F_spistatus), 1, true, true)
	_bitBuf.Finish()
}

type Evt_0009 struct {
	F_evtId   uint16 `json:"evtId"`
	F_cellLAC uint16 `json:"cellLAC"`
	F_CellID  uint32 `json:"CellID"`
}

func To_Evt_0009(_byteBuf *parse.ByteBuf) *Evt_0009 {
	_instance := Evt_0009{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_cellLAC = _byteBuf.Read_uint16()
	_instance.F_CellID = _byteBuf.Read_uint32()
	return &_instance
}

func (__instance *Evt_0009) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_cellLAC)
	_byteBuf.Write_uint32(_instance.F_CellID)
}

type Evt_000A struct {
	F_evtId              uint16 `json:"evtId"`
	F_cellSignalStrength int8   `json:"cellSignalStrength"`
	F_cellRAT            uint8  `json:"cellRAT"`
	F_cellRATadd         uint8  `json:"cellRATadd"`
	F_cellChanID         uint16 `json:"cellChanID"`
	F_GNSSSATS           uint8  `json:"GNSSSATS"`
}

func To_Evt_000A(_byteBuf *parse.ByteBuf) *Evt_000A {
	_instance := Evt_000A{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_cellSignalStrength = _byteBuf.Read_int8()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_cellRAT = uint8(_bitBuf.Read(3, true, true))
	_instance.F_cellRATadd = uint8(_bitBuf.Read(3, true, true))
	_instance.F_cellChanID = uint16(_bitBuf.Read(9, true, true))
	_instance.F_GNSSSATS = uint8(_bitBuf.Read(8, true, true))
	_bitBuf.Finish()

	_byteBuf.Skip(2)
	return &_instance
}

func (__instance *Evt_000A) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_int8(_instance.F_cellSignalStrength)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_cellRAT), 3, true, true)
	_bitBuf.Write(int64(_instance.F_cellRATadd), 3, true, true)
	_bitBuf.Write(int64(_instance.F_cellChanID), 9, true, true)
	_bitBuf.Write(int64(_instance.F_GNSSSATS), 8, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_zero(2)

}

type Evt_000B struct {
	F_evtId               uint16 `json:"evtId"`
	F_ModemStates         uint8  `json:"ModemStates"`
	F_iNetworkSts         uint8  `json:"iNetworkSts"`
	F_iNetworkSts_ErrCode uint16 `json:"iNetworkSts_ErrCode"`
}

func To_Evt_000B(_byteBuf *parse.ByteBuf) *Evt_000B {
	_instance := Evt_000B{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_ModemStates = _byteBuf.Read_uint8()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_iNetworkSts = uint8(_bitBuf.Read(1, true, true))
	_instance.F_iNetworkSts_ErrCode = uint16(_bitBuf.Read(16, true, true))
	_bitBuf.Finish()
	_byteBuf.Skip(2)
	return &_instance
}

func (__instance *Evt_000B) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint8(_instance.F_ModemStates)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_iNetworkSts), 1, true, true)
	_bitBuf.Write(int64(_instance.F_iNetworkSts_ErrCode), 16, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_zero(2)
}

type Evt_000C struct {
	F_evtId         uint16 `json:"evtId"`
	F_PotclVer      uint8  `json:"PotclVer"`
	F_PotclSecyVer  uint8  `json:"PotclSecyVer"`
	F_CalendarYear  uint16 `json:"CalendarYear"`
	F_CalendarDay   uint8  `json:"CalendarDay"`
	F_CalendarMonth uint8  `json:"CalendarMonth"`
}

func To_Evt_000C(_byteBuf *parse.ByteBuf) *Evt_000C {
	_instance := Evt_000C{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_PotclVer = uint8(_bitBuf.Read(4, true, true))
	_instance.F_PotclSecyVer = uint8(_bitBuf.Read(4, true, true))
	_bitBuf.Finish()
	_instance.F_CalendarYear = uint16(_byteBuf.Read_uint8()) + 2000
	_instance.F_CalendarDay = uint8(_bitBuf.Read(5, true, true))
	_instance.F_CalendarMonth = uint8(_bitBuf.Read(4, true, true))
	_bitBuf.Finish()
	_byteBuf.Skip(2)
	return &_instance
}

func (__instance *Evt_000C) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_PotclVer), 4, true, true)
	_bitBuf.Write(int64(_instance.F_PotclSecyVer), 4, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_uint8(uint8(_instance.F_CalendarYear - 2000))
	_bitBuf.Write(int64(_instance.F_CalendarDay), 5, true, true)
	_bitBuf.Write(int64(_instance.F_CalendarMonth), 4, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_zero(2)
}

type Evt_000D struct {
	F_evtId         uint16 `json:"evtId"`
	F_CellFrequency uint32 `json:"CellFrequency"`
}

func To_Evt_000D(_byteBuf *parse.ByteBuf) *Evt_000D {
	_instance := Evt_000D{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_CellFrequency = _byteBuf.Read_uint32()
	_byteBuf.Skip(2)
	return &_instance
}

func (__instance *Evt_000D) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint32(_instance.F_CellFrequency)
	_byteBuf.Write_zero(2)
}

type Evt_000E struct {
	F_evtId            uint16  `json:"evtId"`
	F_BMSChrgSts       uint8   `json:"BMSChrgSts"`
	F_BMSPackSOCBkup   float32 `json:"BMSPackSOCBkup"`
	F_BMSPackSOCVBkup  uint8   `json:"BMSPackSOCVBkup"`
	F_BMSOfbdChrgSpRsn uint8   `json:"BMSOfbdChrgSpRsn"`
	F_BMSWrlsChrgSpRsn uint8   `json:"BMSWrlsChrgSpRsn"`
}

func To_Evt_000E(_byteBuf *parse.ByteBuf) *Evt_000E {
	_instance := Evt_000E{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_BMSChrgSts = uint8(_bitBuf.Read(5, true, true))
	_instance.F_BMSPackSOCBkup = float32(_bitBuf.Read(10, true, true)) * 0.1
	_instance.F_BMSPackSOCVBkup = uint8(_bitBuf.Read(1, true, true))
	_bitBuf.Finish()
	_instance.F_BMSOfbdChrgSpRsn = _byteBuf.Read_uint8()
	_instance.F_BMSWrlsChrgSpRsn = _byteBuf.Read_uint8()
	_byteBuf.Skip(2)
	return &_instance
}

func (__instance *Evt_000E) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_BMSChrgSts), 5, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackSOCBkup*10), 10, true, true)
	_bitBuf.Write(int64(_instance.F_BMSPackSOCVBkup), 1, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_uint8(_instance.F_BMSOfbdChrgSpRsn)
	_byteBuf.Write_uint8(_instance.F_BMSWrlsChrgSpRsn)
	_byteBuf.Write_zero(2)
}

type Evt_000F struct {
	F_evtId            uint16  `json:"evtId"`
	F_TMActuToqHiPre   float32 `json:"TMActuToqHiPre"`
	F_TMInvtrCrntHiPre float32 `json:"TMInvtrCrntHiPre"`
}

func To_Evt_000F(_byteBuf *parse.ByteBuf) *Evt_000F {
	_instance := Evt_000F{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_TMActuToqHiPre = float32(_byteBuf.Read_uint16())*0.1 - 2000
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_TMInvtrCrntHiPre = float32(_bitBuf.Read(15, true, true))*0.1 - 1000
	_bitBuf.Finish()
	_byteBuf.Skip(2)
	return &_instance
}

func (__instance *Evt_000F) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(uint16(_instance.F_TMActuToqHiPre*10 + 20000))
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_TMInvtrCrntHiPre*10+10000), 15, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_zero(2)
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

func To_Evt_0800(_byteBuf *parse.ByteBuf) *Evt_0800 {
	_instance := Evt_0800{}
	_instance.F_evtId = _byteBuf.Read_uint16()

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_SysPwrMd = uint8(_bitBuf.Read(2, true, true))
	_instance.F_SysPwrMdV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_SysVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_TrShftLvrPos = uint8(_bitBuf.Read(4, true, true))
	_bitBuf.Finish()
	_instance.F_SysVol = float32(_byteBuf.Read_uint8())*0.1 + 3
	_byteBuf.Skip(3)
	_instance.F_TrShftLvrPosV = uint8(_bitBuf.Read(1, true, true))
	_bitBuf.Finish()

	return &_instance
}

func (__instance *Evt_0800) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_SysPwrMd), 2, true, true)
	_bitBuf.Write(int64(_instance.F_SysPwrMdV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_SysVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_TrShftLvrPos), 4, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_uint8(uint8(parse.Round((_instance.F_SysVol - 3) / 0.1)))
	_byteBuf.Write_zero(3)
	_bitBuf.Write(int64(_instance.F_TrShftLvrPosV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_0801 struct {
	F_evtId     uint16  `json:"evtId"`
	F_BrkPdlPos float32 `json:"BrkPdlPos"`
}

func To_Evt_0801(_byteBuf *parse.ByteBuf) *Evt_0801 {
	_instance := Evt_0801{}
	_instance.F_evtId = _byteBuf.Read_uint16()

	_byteBuf.Skip(5)
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_BrkPdlPos = float32(_bitBuf.Read(8, true, true)) * 0.392157
	_bitBuf.Finish()

	return &_instance
}

func (__instance *Evt_0801) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_zero(5)

	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_BrkPdlPos/0.392157)), 8, true, true)
	_bitBuf.Finish()
}

type Evt_0802 struct {
	F_evtId          uint16  `json:"evtId"`
	F_VehSpdAvgDrvn  float64 `json:"VehSpdAvgDrvn"`
	F_VehSpdAvgDrvnV uint8   `json:"VehSpdAvgDrvnV"`
}

func To_Evt_0802(_byteBuf *parse.ByteBuf) *Evt_0802 {
	_instance := Evt_0802{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_VehSpdAvgDrvn = float64(_bitBuf.Read(15, true, true)) * 0.015625
	_instance.F_VehSpdAvgDrvnV = uint8(_bitBuf.Read(1, true, true))
	_bitBuf.Finish()
	_byteBuf.Skip(4)
	return &_instance
}

func (__instance *Evt_0802) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(parse.Round(_instance.F_VehSpdAvgDrvn/0.015625)), 15, true, true)
	_bitBuf.Write(int64(_instance.F_VehSpdAvgDrvnV), 1, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_zero(4)

}

type Evt_0803 struct {
	F_evtId      uint16 `json:"evtId"`
	F_VehOdo     uint32 `json:"VehOdo"`
	F_VehOdoV    uint8  `json:"VehOdoV"`
	F_BrkPdlPosV uint8  `json:"BrkPdlPosV"`
}

func To_Evt_0803(_byteBuf *parse.ByteBuf) *Evt_0803 {
	_instance := Evt_0803{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_VehOdo = uint32(_bitBuf.Read(24, true, true))
	_instance.F_VehOdoV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BrkPdlPosV = uint8(_bitBuf.Read(1, true, true))
	_bitBuf.Finish()
	_byteBuf.Skip(2)
	return &_instance
}

func (__instance *Evt_0803) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_VehOdo), 24, true, true)
	_bitBuf.Write(int64(_instance.F_VehOdoV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_BrkPdlPosV), 1, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_zero(2)

}

type Evt_2_6_unknown struct {
	F_evtId uint16  `json:"evtId"`
	F_data  [6]byte `json:"data"`
}

func To_Evt_2_6_unknown(_byteBuf *parse.ByteBuf) *Evt_2_6_unknown {
	_instance := Evt_2_6_unknown{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	arr := _byteBuf.Read_slice_uint8(6)
	_instance.F_data = *(*[6]byte)(unsafe.Pointer(&arr[0]))
	return &_instance
}
func (__instance *Evt_2_6_unknown) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_slice_uint8(_instance.F_data[:])
}

type Evt_4_x_unknown struct {
	F_evtId  uint16             `json:"evtId"`
	F_evtLen uint16             `json:"evtLen"`
	F_data   parse.JsonUint8Arr `json:"data"`
}

func To_Evt_4_x_unknown(_byteBuf *parse.ByteBuf) *Evt_4_x_unknown {
	_instance := Evt_4_x_unknown{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	F_evtLen_v := _byteBuf.Read_uint16()
	_instance.F_evtLen = F_evtLen_v
	F_data_arr := _byteBuf.Read_slice_uint8((int)(F_evtLen_v))
	_instance.F_data = F_data_arr
	return &_instance
}

func (__instance *Evt_4_x_unknown) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	F_data_arr := _instance.F_data
	_byteBuf.Write_slice_uint8(F_data_arr)
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

func To_Evt_D006(_byteBuf *parse.ByteBuf) *Evt_D006 {
	_instance := Evt_D006{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_EPTRdy = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSBscSta = uint8(_bitBuf.Read(5, true, true))
	_instance.F_BMSPackCrnt = float32(_bitBuf.Read(16, true, true))*0.05 - 1000
	_instance.F_BMSPackCrntV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSPackSOC = float32(_bitBuf.Read(10, true, true)) * 0.1
	_instance.F_BMSPackSOCV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSPackSOCDsp = float32(_bitBuf.Read(10, true, true)) * 0.1
	_instance.F_BMSPackSOCDspV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_ElecVehSysMd = uint8(_bitBuf.Read(4, true, true))
	_instance.F_BMSPackVol = float32(_bitBuf.Read(12, true, true)) * 0.25
	_instance.F_BMSPackVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_HVDCDCSta = uint8(_bitBuf.Read(3, true, true))
	_instance.F_EPTTrInptShaftToq = float32(_bitBuf.Read(12, true, true))*0.5 - 848
	_instance.F_EPTTrInptShaftToqV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EPTTrOtptShaftToq = int16(_bitBuf.Read(12, true, true))*2 - 3392
	_instance.F_EPTTrOtptShaftToqV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EPTBrkPdlDscrtInptSts = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EPTBrkPdlDscrtInptStsV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BrkSysBrkLghtsReqd = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EPBSysBrkLghtsReqd = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EPBSysBrkLghtsReqdA = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSPtIsltnRstc = float32(_bitBuf.Read(14, true, true)) * 0.5
	_instance.F_EPTAccelActuPos = float32(_bitBuf.Read(8, true, true)) * 0.392157
	_instance.F_EPTAccelActuPosV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_TMInvtrCrntV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_TMInvtrCrnt = int16(_bitBuf.Read(11, true, true)) - 1024
	_instance.F_ISGInvtrCrntV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_ISGInvtrCrnt = int16(_bitBuf.Read(11, true, true)) - 1024
	_instance.F_SAMInvtrCrnt = int16(_bitBuf.Read(11, true, true)) - 1024
	_instance.F_SAMInvtrCrntV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_TMSta = uint8(_bitBuf.Read(4, true, true))
	_instance.F_ISGSta = uint8(_bitBuf.Read(4, true, true))
	_instance.F_SAMSta = uint8(_bitBuf.Read(4, true, true))
	_instance.F_TMInvtrTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_ISGInvtrTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_SAMInvtrTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_TMSpd = int32(_bitBuf.Read(16, true, true)) - 32768
	_instance.F_TMSpdV = int8(_bitBuf.Read(1, true, true))
	_instance.F_ISGSpd = int32(_bitBuf.Read(16, true, true)) - 32768
	_instance.F_ISGSpdV = int8(_bitBuf.Read(1, true, true))
	_instance.F_SAMSpdV = int8(_bitBuf.Read(1, true, true))
	_instance.F_SAMSpd = int32(_bitBuf.Read(16, true, true)) - 32768
	_instance.F_TMActuToq = float32(_bitBuf.Read(11, true, true))*0.5 - 512
	_instance.F_TMActuToqV = int8(_bitBuf.Read(1, true, true))
	_instance.F_ISGActuToq = float32(_bitBuf.Read(11, true, true))*0.5 - 512
	_instance.F_ISGActuToqV = int8(_bitBuf.Read(1, true, true))
	_instance.F_SAMActuToqV = int8(_bitBuf.Read(1, true, true))
	_instance.F_SAMActuToq = float32(_bitBuf.Read(11, true, true))*0.5 - 512
	_instance.F_TMSttrTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_ISGSttrTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_SAMSttrTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_HVDCDCHVSideVol = uint16(_bitBuf.Read(10, true, true))
	_instance.F_HVDCDCHVSideVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_AvgFuelCsump = float32(_bitBuf.Read(8, true, true)) * 0.1
	_instance.F_TMInvtrVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_TMInvtrVol = uint16(_bitBuf.Read(10, true, true))
	_instance.F_ISGInvtrVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_ISGInvtrVol = uint16(_bitBuf.Read(10, true, true))
	_instance.F_SAMInvtrVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_SAMInvtrVol = uint16(_bitBuf.Read(10, true, true))
	_instance.F_BMSCellMaxTemIndx = uint8(_bitBuf.Read(8, true, true))
	_instance.F_BMSCellMaxTem = float32(_bitBuf.Read(8, true, true))*0.5 - 40
	_instance.F_BMSCellMaxTemV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSCellMinTemIndx = uint8(_bitBuf.Read(8, true, true))
	_instance.F_BMSCellMinTem = float32(_bitBuf.Read(8, true, true))*0.5 - 40
	_instance.F_BMSCellMinTemV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSCellMaxVolIndx = uint8(_bitBuf.Read(8, true, true))
	_instance.F_BMSCellMaxVol = float32(_bitBuf.Read(13, true, true)) * 0.001
	_instance.F_BMSCellMaxVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSCellMinVolIndx = uint8(_bitBuf.Read(8, true, true))
	_instance.F_BMSCellMinVol = float32(_bitBuf.Read(13, true, true)) * 0.001
	_instance.F_BMSCellMinVolV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSPtIsltnRstcV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_HVDCDCTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_BrkFludLvlLow = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BrkSysRedBrkTlltReq = uint8(_bitBuf.Read(1, true, true))
	_instance.F_ABSF = uint8(_bitBuf.Read(1, true, true))
	_instance.F_VSESts = uint8(_bitBuf.Read(3, true, true))
	_instance.F_IbstrWrnngIO = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSHVILClsd = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EPTTrOtptShaftTotToq = float32(_bitBuf.Read(12, true, true))*0.5 - 848
	_instance.F_EPTTrOtptShaftTotToqV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BrkFludLvlLowV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EnSpd = float32(_bitBuf.Read(16, true, true)) * 0.25
	_instance.F_EnSpdSts = uint8(_bitBuf.Read(2, true, true))
	_instance.F_FuelCsump = uint16(_bitBuf.Read(12, true, true)) * 16
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}

	return &_instance
}

func (__instance *Evt_D006) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
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
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
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

func To_Evt_D008(_byteBuf *parse.ByteBuf) *Evt_D008 {
	_instance := Evt_D008{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_instance.F_DTCInfomationBMS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationECM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationEPB = _byteBuf.Read_int56()
	_instance.F_DTCInfomationPLCM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationTCM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationTPMS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationTC = _byteBuf.Read_int56()
	_instance.F_DTCInfomationISC = _byteBuf.Read_int56()
	_instance.F_DTCInfomationSAC = _byteBuf.Read_int56()
	_instance.F_DTCInfomationIMCU = _byteBuf.Read_int56()
	skip := int(_instance.F_evtLen) - 70
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D008) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_int56(_instance.F_DTCInfomationBMS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationECM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationEPB)
	_byteBuf.Write_int56(_instance.F_DTCInfomationPLCM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationTCM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationTPMS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationTC)
	_byteBuf.Write_int56(_instance.F_DTCInfomationISC)
	_byteBuf.Write_int56(_instance.F_DTCInfomationSAC)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIMCU)
	skip := int(_instance.F_evtLen) - 70
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
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

func To_Evt_D009(_byteBuf *parse.ByteBuf) *Evt_D009 {
	_instance := Evt_D009{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_BMSCMUFlt = uint8(_bitBuf.Read(2, true, true))
	_instance.F_BMSCellVoltFlt = uint8(_bitBuf.Read(2, true, true))
	_instance.F_BMSPackTemFlt = uint8(_bitBuf.Read(2, true, true))
	_instance.F_BMSPackVoltFlt = uint8(_bitBuf.Read(2, true, true))
	_instance.F_BMSWrnngInfo = uint8(_bitBuf.Read(6, true, true))
	_instance.F_BMSWrnngInfoPV = uint8(_bitBuf.Read(6, true, true))
	_instance.F_BMSWrnngInfoRC = uint8(_bitBuf.Read(4, true, true))
	_instance.F_BMSPreThrmFltInd = uint8(_bitBuf.Read(1, true, true))
	_bitBuf.Skip(5)
	_instance.F_BMSKeepSysAwkScene = uint8(_bitBuf.Read(4, true, true))
	_instance.F_BMSTemOverDifAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSOverTemAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSOverPackVolAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSUnderPackVolAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSHVILAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSOverCellVolAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSUnderCellVolAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSLowSOCAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSJumpngSOCAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSHiSOCAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSPackVolMsmchAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSPoorCellCnstncyAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSCellOverChrgdAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSLowPtIsltnRstcAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_TMRtrTem = int16(_bitBuf.Read(8, true, true)) - 40
	_instance.F_TMStrOvTempAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_TMInvtrOvTempAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_ISCStrOvTempAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_ISCInvtrOvTempAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_SAMStrOvTempAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_SAMInvtrOvTempAlrm = uint8(_bitBuf.Read(3, true, true))
	_instance.F_EPTHVDCDCMdReq = uint8(_bitBuf.Read(3, true, true))
	_instance.F_VCUSecyWrnngInfo = uint8(_bitBuf.Read(6, true, true))
	_instance.F_VCUSecyWrnngInfoPV = uint8(_bitBuf.Read(6, true, true))
	_instance.F_VCUSecyWrnngInfoRC = uint8(_bitBuf.Read(4, true, true))
	_instance.F_VCUSecyWrnngInfoCRC = uint8(_bitBuf.Read(8, true, true))
	_instance.F_BMSOnbdChrgSpRsn = uint8(_bitBuf.Read(8, true, true))
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D009) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
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
	_bitBuf.Write(int64(_instance.F_TMRtrTem+40), 8, true, true)
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
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D00A struct {
	F_evtId     uint16 `json:"evtId"`
	F_evtLen    uint16 `json:"evtLen"`
	F_VIN       string `json:"VIN"`
	F_IAMSN     string `json:"IAMSN"`
	F_EsimIccid string `json:"EsimIccid"`
	F_EsimID    string `json:"EsimID"`
}

func To_Evt_D00A(_byteBuf *parse.ByteBuf) *Evt_D00A {
	_instance := Evt_D00A{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_instance.F_VIN = _byteBuf.Read_string_utf8(17)
	_instance.F_IAMSN = _byteBuf.Read_string_utf8(16)
	_instance.F_EsimIccid = _byteBuf.Read_string_utf8(20)
	_instance.F_EsimID = _byteBuf.Read_string_utf8(32)
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D00A) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_string_utf8(_instance.F_VIN)
	_byteBuf.Write_string_utf8(_instance.F_IAMSN)
	_byteBuf.Write_string_utf8(_instance.F_EsimIccid)
	_byteBuf.Write_string_utf8(_instance.F_EsimID)
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D00B_BMSCellVol struct {
	F_BMSCellVol  float32 `json:"BMSCellVol"`
	F_BMSCellVolV uint8   `json:"BMSCellVolV"`
}

type Evt_D00B struct {
	F_evtId            uint16             `json:"evtId"`
	F_evtLen           uint16             `json:"evtLen"`
	F_BMSCellVolSumNum uint8              `json:"BMSCellVolSumNum"`
	F_BMSCellVol       []float32          `json:"BMSCellVol"`
	F_BMSCellVolV      parse.JsonUint8Arr `json:"BMSCellVolV"`
}

func To_Evt_D00B(_byteBuf *parse.ByteBuf) *Evt_D00B {
	_instance := Evt_D00B{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	F_BMSCellVolSumNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSCellVolSumNum = F_BMSCellVolSumNum_v

	F_BMSCellVols_len := (int)(F_BMSCellVolSumNum_v)
	F_BMSCellVol := make([]float32, F_BMSCellVols_len, F_BMSCellVols_len)
	F_BMSCellVolV := make([]uint8, F_BMSCellVols_len, F_BMSCellVols_len)
	for i := 0; i < F_BMSCellVols_len; i++ {
		temp := _byteBuf.Read_uint16()
		F_BMSCellVol[i] = float32(temp>>3) * 0.001
		F_BMSCellVolV[i] = uint8((temp >> 2) & 0x01)
	}
	_instance.F_BMSCellVol = F_BMSCellVol
	_instance.F_BMSCellVolV = F_BMSCellVolV
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D00B) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_BMSCellVolSumNum)
	for i := 0; i < int(_instance.F_BMSCellVolSumNum); i++ {
		n1 := uint16(_instance.F_BMSCellVol[i] * 1000)
		n2 := uint16(_instance.F_BMSCellVolV[i])
		_byteBuf.Write_uint16((n1 << 3) | (n2 << 2))
	}
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D00C struct {
	F_evtId            uint16             `json:"evtId"`
	F_evtLen           uint16             `json:"evtLen"`
	F_BMSCellTemSumNum uint8              `json:"BMSCellTemSumNum"`
	F_BMSCellTem       []int16            `json:"BMSCellTem"`
	F_BMSCellTemV      parse.JsonUint8Arr `json:"BMSCellTemV"`
}

func To_Evt_D00C(_byteBuf *parse.ByteBuf) *Evt_D00C {
	_instance := Evt_D00C{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	F_BMSCellTemSumNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSCellTemSumNum = F_BMSCellTemSumNum_v

	F_BMSCellTems_len := (int)(F_BMSCellTemSumNum_v)
	F_BMSCellTem := make([]int16, F_BMSCellTems_len, F_BMSCellTems_len)
	F_BMSCellTemV := make([]uint8, F_BMSCellTems_len, F_BMSCellTems_len)
	for i := 0; i < F_BMSCellTems_len; i++ {
		temp := _byteBuf.Read_int16()
		F_BMSCellTem[i] = (temp >> 8) - 40
		F_BMSCellTemV[i] = uint8((temp >> 7) & 1)
	}
	_instance.F_BMSCellTem = F_BMSCellTem
	_instance.F_BMSCellTemV = F_BMSCellTemV
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D00C) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_BMSCellTemSumNum)
	for i := 0; i < int(_instance.F_BMSCellTemSumNum); i++ {
		n1 := _instance.F_BMSCellTem[i] + 40
		n2 := _instance.F_BMSCellTemV[i]
		_byteBuf.Write_int16((n1 << 8) | int16(n2<<7))
	}
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D00D struct {
	F_evtId              uint16             `json:"evtId"`
	F_evtLen             uint16             `json:"evtLen"`
	F_BMSBusbarTemSumNum uint8              `json:"BMSBusbarTemSumNum"`
	F_BMSBusbarTem       []int16            `json:"BMSBusbarTem"`
	F_BMSBusbarTemV      parse.JsonUint8Arr `json:"BMSBusbarTemV"`
}

func To_Evt_D00D(_byteBuf *parse.ByteBuf) *Evt_D00D {
	_instance := Evt_D00D{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	F_BMSBusbarTemSumNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSBusbarTemSumNum = F_BMSBusbarTemSumNum_v
	F_BMSBusbarTems_len := (int)(F_BMSBusbarTemSumNum_v)
	F_BMSBusbarTem := make([]int16, F_BMSBusbarTems_len, F_BMSBusbarTems_len)
	F_BMSBusbarTemV := make([]uint8, F_BMSBusbarTems_len, F_BMSBusbarTems_len)
	for i := 0; i < F_BMSBusbarTems_len; i++ {
		temp := _byteBuf.Read_int16()
		F_BMSBusbarTem[i] = (temp >> 8) - 40
		F_BMSBusbarTemV[i] = uint8((temp >> 7) & 1)
	}
	_instance.F_BMSBusbarTem = F_BMSBusbarTem
	_instance.F_BMSBusbarTemV = F_BMSBusbarTemV
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D00D) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_BMSBusbarTemSumNum)
	for i := 0; i < int(_instance.F_BMSBusbarTemSumNum); i++ {
		n1 := _instance.F_BMSBusbarTem[i] + 40
		n2 := _instance.F_BMSBusbarTemV[i]
		_byteBuf.Write_int16((n1 << 8) | int16(n2<<7))
	}
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D00E struct {
	F_evtId            uint16 `json:"evtId"`
	F_evtLen           uint16 `json:"evtLen"`
	F_BMSRptBatCodeNum uint8  `json:"BMSRptBatCodeNum"`
	F_BMSRptBatCodeAsc string `json:"BMSRptBatCodeAsc"`
}

func To_Evt_D00E(_byteBuf *parse.ByteBuf) *Evt_D00E {
	_instance := Evt_D00E{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	F_BMSRptBatCodeNum_v := _byteBuf.Read_uint8()
	_instance.F_BMSRptBatCodeNum = F_BMSRptBatCodeNum_v
	_instance.F_BMSRptBatCodeAsc = _byteBuf.Read_string_utf8((int)(F_BMSRptBatCodeNum_v))
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D00E) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_BMSRptBatCodeNum)
	_byteBuf.Write_string_utf8(_instance.F_BMSRptBatCodeAsc)
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
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

func To_Evt_D00F(_byteBuf *parse.ByteBuf) *Evt_D00F {
	_instance := Evt_D00F{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_instance.F_BMSWrnngInfoCRC = _byteBuf.Read_uint8()
	_instance.F_BMSBusbarTempMax = float32(_byteBuf.Read_uint8())*0.5 - 40
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_BMSPreThrmFltIndBkup = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSWrnngInfoRCBkup = uint8(_bitBuf.Read(4, true, true))
	_instance.F_BMSBatPrsFlt = uint8(_bitBuf.Read(3, true, true))
	_instance.F_BMSWrnngInfoBkup = uint8(_bitBuf.Read(6, true, true))
	_instance.F_BMSBatPrsAlrm = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSBatPrsAlrmV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSBatPrsSnsrV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSBatPrsSnsrValBkup = float32(_bitBuf.Read(15, true, true)) * 0.05
	_instance.F_BMSBatPrsSnsrVBkup = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSBatPrsSnsrVal = float32(_bitBuf.Read(15, true, true)) * 0.05
	_instance.F_BMSClntPumpPWMReq = float32(_bitBuf.Read(8, true, true)) * 0.4
	_instance.F_BMSPumpPwrOnReq = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSBatPrsAlrmVBkup = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSBatPrsAlrmBkup = uint8(_bitBuf.Read(1, true, true))
	_instance.F_BMSWrnngInfoCRCBkup = uint8(_bitBuf.Read(4, true, true))
	_instance.F_VCUBatPrsAlrm = uint8(_bitBuf.Read(1, true, true))
	_instance.F_OtsdAirTemCrVal = float32(_bitBuf.Read(8, true, true))*0.5 - 40
	_instance.F_VCUBatPrsAlrmV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_OtsdAirTemCrValV = uint8(_bitBuf.Read(1, true, true))
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D00F) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_BMSWrnngInfoCRC)
	_byteBuf.Write_uint8(uint8(parse.Round((_instance.F_BMSBusbarTempMax + 40) / 0.5)))
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
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
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D010 struct {
	F_evtId              uint16 `json:"evtId"`
	F_evtLen             uint16 `json:"evtLen"`
	F_DTCInfomationIAM   int64  `json:"DTCInfomationIAM"`
	F_DTCInfomationIPD   int64  `json:"DTCInfomationIPD"`
	F_DTCInfomationIECU  int64  `json:"DTCInfomationIECU"`
	F_DTCInfomationFDR   int64  `json:"DTCInfomationFDR"`
	F_DTCInfomationLFSDA int64  `json:"DTCInfomationLFSDA"`
	F_DTCInfomationRFSDA int64  `json:"DTCInfomationRFSDA"`
	F_DTCInfomationLHRDA int64  `json:"DTCInfomationLHRDA"`
	F_DTCInfomationRHRDA int64  `json:"DTCInfomationRHRDA"`
}

func To_Evt_D010(_byteBuf *parse.ByteBuf) *Evt_D010 {
	_instance := Evt_D010{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_instance.F_DTCInfomationIAM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationIPD = _byteBuf.Read_int56()
	_instance.F_DTCInfomationIECU = _byteBuf.Read_int56()
	_instance.F_DTCInfomationFDR = _byteBuf.Read_int56()
	_instance.F_DTCInfomationLFSDA = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRFSDA = _byteBuf.Read_int56()
	_instance.F_DTCInfomationLHRDA = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRHRDA = _byteBuf.Read_int56()
	skip := int(_instance.F_evtLen) - 56
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D010) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIAM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIPD)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIECU)
	_byteBuf.Write_int56(_instance.F_DTCInfomationFDR)
	_byteBuf.Write_int56(_instance.F_DTCInfomationLFSDA)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRFSDA)
	_byteBuf.Write_int56(_instance.F_DTCInfomationLHRDA)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRHRDA)
	skip := int(_instance.F_evtLen) - 56
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D011 struct {
	F_evtId              uint16 `json:"evtId"`
	F_evtLen             uint16 `json:"evtLen"`
	F_DTCInfomationEPMCU int64  `json:"DTCInfomationEPMCU"`
	F_DTCInfomationWLC   int64  `json:"DTCInfomationWLC"`
	F_DTCInfomationSCU   int64  `json:"DTCInfomationSCU"`
	F_DTCInfomationEOPC  int64  `json:"DTCInfomationEOPC"`
	F_DTCInfomationCCU   int64  `json:"DTCInfomationCCU"`
}

func To_Evt_D011(_byteBuf *parse.ByteBuf) *Evt_D011 {
	_instance := Evt_D011{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_instance.F_DTCInfomationEPMCU = _byteBuf.Read_int56()
	_instance.F_DTCInfomationWLC = _byteBuf.Read_int56()
	_instance.F_DTCInfomationSCU = _byteBuf.Read_int56()
	_instance.F_DTCInfomationEOPC = _byteBuf.Read_int56()
	_instance.F_DTCInfomationCCU = _byteBuf.Read_int56()
	skip := int(_instance.F_evtLen) - 35
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D011) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_int56(_instance.F_DTCInfomationEPMCU)
	_byteBuf.Write_int56(_instance.F_DTCInfomationWLC)
	_byteBuf.Write_int56(_instance.F_DTCInfomationSCU)
	_byteBuf.Write_int56(_instance.F_DTCInfomationEOPC)
	_byteBuf.Write_int56(_instance.F_DTCInfomationCCU)
	skip := int(_instance.F_evtLen) - 35
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D012 struct {
	F_evtId              uint16 `json:"evtId"`
	F_evtLen             uint16 `json:"evtLen"`
	F_DTCInfomationSDM   int64  `json:"DTCInfomationSDM"`
	F_DTCInfomationIBS   int64  `json:"DTCInfomationIBS"`
	F_DTCInfomationEPS   int64  `json:"DTCInfomationEPS"`
	F_DTCInfomationEPS_S int64  `json:"DTCInfomationEPS_S"`
	F_DTCInfomationSCM   int64  `json:"DTCInfomationSCM"`
	F_DTCInfomationRBM   int64  `json:"DTCInfomationRBM"`
	F_DTCInfomationSAS   int64  `json:"DTCInfomationSAS"`
	F_DTCInfomationRWSGW int64  `json:"DTCInfomationRWSGW"`
	F_DTCInfomationRWS   int64  `json:"DTCInfomationRWS"`
}

func To_Evt_D012(_byteBuf *parse.ByteBuf) *Evt_D012 {
	_instance := Evt_D012{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_DTCInfomationSDM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationIBS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationEPS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationEPS_S = _byteBuf.Read_int56()
	_instance.F_DTCInfomationSCM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRBM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationSAS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRWSGW = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRWS = _byteBuf.Read_int56()
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - 63
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D012) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_int56(_instance.F_DTCInfomationSDM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIBS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationEPS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationEPS_S)
	_byteBuf.Write_int56(_instance.F_DTCInfomationSCM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRBM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationSAS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRWSGW)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRWS)
	skip := int(_instance.F_evtLen) - 63
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D013 struct {
	F_evtId                 uint16 `json:"evtId"`
	F_evtLen                uint16 `json:"evtLen"`
	F_DTCInfomationDCM_FL   int64  `json:"DTCInfomationDCM_FL"`
	F_DTCInfomationDCM_FR   int64  `json:"DTCInfomationDCM_FR"`
	F_DTCInfomationDCM_RL   int64  `json:"DTCInfomationDCM_RL"`
	F_DTCInfomationDCM_RR   int64  `json:"DTCInfomationDCM_RR"`
	F_DTCInfomationATC      int64  `json:"DTCInfomationATC"`
	F_DTCInfomationAMR      int64  `json:"DTCInfomationAMR"`
	F_DTCInfomationBPEPS    int64  `json:"DTCInfomationBPEPS"`
	F_DTCInfomationMSM_Drv  int64  `json:"DTCInfomationMSM_Drv"`
	F_DTCInfomationMSM_Psng int64  `json:"DTCInfomationMSM_Psng"`
	F_DTCInfomationDLP      int64  `json:"DTCInfomationDLP"`
	F_DTCInfomationBCM      int64  `json:"DTCInfomationBCM"`
}

func To_Evt_D013(_byteBuf *parse.ByteBuf) *Evt_D013 {
	_instance := Evt_D013{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_instance.F_DTCInfomationDCM_FL = _byteBuf.Read_int56()
	_instance.F_DTCInfomationDCM_FR = _byteBuf.Read_int56()
	_instance.F_DTCInfomationDCM_RL = _byteBuf.Read_int56()
	_instance.F_DTCInfomationDCM_RR = _byteBuf.Read_int56()
	_instance.F_DTCInfomationATC = _byteBuf.Read_int56()
	_instance.F_DTCInfomationAMR = _byteBuf.Read_int56()
	_instance.F_DTCInfomationBPEPS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationMSM_Drv = _byteBuf.Read_int56()
	_instance.F_DTCInfomationMSM_Psng = _byteBuf.Read_int56()
	_instance.F_DTCInfomationDLP = _byteBuf.Read_int56()
	_instance.F_DTCInfomationBCM = _byteBuf.Read_int56()
	skip := int(_instance.F_evtLen) - 77
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D013) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_int56(_instance.F_DTCInfomationDCM_FL)
	_byteBuf.Write_int56(_instance.F_DTCInfomationDCM_FR)
	_byteBuf.Write_int56(_instance.F_DTCInfomationDCM_RL)
	_byteBuf.Write_int56(_instance.F_DTCInfomationDCM_RR)
	_byteBuf.Write_int56(_instance.F_DTCInfomationATC)
	_byteBuf.Write_int56(_instance.F_DTCInfomationAMR)
	_byteBuf.Write_int56(_instance.F_DTCInfomationBPEPS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationMSM_Drv)
	_byteBuf.Write_int56(_instance.F_DTCInfomationMSM_Psng)
	_byteBuf.Write_int56(_instance.F_DTCInfomationDLP)
	_byteBuf.Write_int56(_instance.F_DTCInfomationBCM)
	skip := int(_instance.F_evtLen) - 77
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D014 struct {
	F_evtId               uint16 `json:"evtId"`
	F_evtLen              uint16 `json:"evtLen"`
	F_DTCInfomationICM    int64  `json:"DTCInfomationICM"`
	F_DTCInfomationCARLog int64  `json:"DTCInfomationCARLog"`
	F_DTCInfomationIMATE  int64  `json:"DTCInfomationIMATE"`
	F_DTCInfomationAMP    int64  `json:"DTCInfomationAMP"`
	F_DTCInfomationPGM    int64  `json:"DTCInfomationPGM"`
}

func To_Evt_D014(_byteBuf *parse.ByteBuf) *Evt_D014 {
	_instance := Evt_D014{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_DTCInfomationICM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationCARLog = _byteBuf.Read_int56()
	_instance.F_DTCInfomationIMATE = _byteBuf.Read_int56()
	_instance.F_DTCInfomationAMP = _byteBuf.Read_int56()
	_instance.F_DTCInfomationPGM = _byteBuf.Read_int56()
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - 35
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D014) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_int56(_instance.F_DTCInfomationICM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationCARLog)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIMATE)
	_byteBuf.Write_int56(_instance.F_DTCInfomationAMP)
	_byteBuf.Write_int56(_instance.F_DTCInfomationPGM)
	skip := int(_instance.F_evtLen) - 35
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D015 struct {
	F_evtId            uint16 `json:"evtId"`
	F_evtLen           uint16 `json:"evtLen"`
	F_DTCInfomationICC int64  `json:"DTCInfomationICC"`
}

func To_Evt_D015(_byteBuf *parse.ByteBuf) *Evt_D015 {
	_instance := Evt_D015{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_instance.F_DTCInfomationICC = _byteBuf.Read_int56()
	skip := int(_instance.F_evtLen) - 7
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D015) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_byteBuf.Write_int56(_instance.F_DTCInfomationICC)
	skip := int(_instance.F_evtLen) - 7
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D016 struct {
	F_evtId              uint16 `json:"evtId"`
	F_evtLen             uint16 `json:"evtLen"`
	F_DTCInfomationLHCMS int64  `json:"DTCInfomationLHCMS"`
	F_DTCInfomationRHCMS int64  `json:"DTCInfomationRHCMS"`
	F_DTCInfomationRLSM  int64  `json:"DTCInfomationRLSM"`
	F_DTCInfomationRRSM  int64  `json:"DTCInfomationRRSM"`
	F_DTCInfomationPMA   int64  `json:"DTCInfomationPMA"`
	F_DTCInfomationLVBM  int64  `json:"DTCInfomationLVBM"`
	F_DTCInfomationIMU   int64  `json:"DTCInfomationIMU"`
}

func To_Evt_D016(_byteBuf *parse.ByteBuf) *Evt_D016 {
	_instance := Evt_D016{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_DTCInfomationLHCMS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRHCMS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRLSM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRRSM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationPMA = _byteBuf.Read_int56()
	_instance.F_DTCInfomationLVBM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationIMU = _byteBuf.Read_int56()
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - 49
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D016) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_byteBuf.Write_int56(_instance.F_DTCInfomationLHCMS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRHCMS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRLSM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRRSM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationPMA)
	_byteBuf.Write_int56(_instance.F_DTCInfomationLVBM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIMU)
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - 49
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D017 struct {
	F_evtId                  uint16 `json:"evtId"`
	F_evtLen                 uint16 `json:"evtLen"`
	F_DTCInfomationIPS       int64  `json:"DTCInfomationIPS"`
	F_DTCInfomationRrDetnRdr int64  `json:"DTCInfomationRrDetnRdr"`
	F_DTCInfomationHUD       int64  `json:"DTCInfomationHUD"`
	F_DTCInfomationFLIDAR    int64  `json:"DTCInfomationFLIDAR"`
	F_DTCInfomationFVCM      int64  `json:"DTCInfomationFVCM"`
	F_DTCInfomationSPD       int64  `json:"DTCInfomationSPD"`
}

func To_Evt_D017(_byteBuf *parse.ByteBuf) *Evt_D017 {
	_instance := Evt_D017{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	_instance.F_DTCInfomationIPS = _byteBuf.Read_int56()
	_instance.F_DTCInfomationRrDetnRdr = _byteBuf.Read_int56()
	_instance.F_DTCInfomationHUD = _byteBuf.Read_int56()
	_instance.F_DTCInfomationFLIDAR = _byteBuf.Read_int56()
	_instance.F_DTCInfomationFVCM = _byteBuf.Read_int56()
	_instance.F_DTCInfomationSPD = _byteBuf.Read_int56()
	skip := int(_instance.F_evtLen) - 42
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D017) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_byteBuf.Write_int56(_instance.F_DTCInfomationIPS)
	_byteBuf.Write_int56(_instance.F_DTCInfomationRrDetnRdr)
	_byteBuf.Write_int56(_instance.F_DTCInfomationHUD)
	_byteBuf.Write_int56(_instance.F_DTCInfomationFLIDAR)
	_byteBuf.Write_int56(_instance.F_DTCInfomationFVCM)
	_byteBuf.Write_int56(_instance.F_DTCInfomationSPD)
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - 42
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D018 struct {
	F_evtId           uint16  `json:"evtId"`
	F_evtLen          uint16  `json:"evtLen"`
	F_APN1ConnSts     uint8   `json:"APN1ConnSts"`
	F_APN2ConnSts     uint8   `json:"APN2ConnSts"`
	F_MqttConnFailRsn uint8   `json:"MqttConnFailRsn"`
	F_ECallSts        uint8   `json:"ECallSts"`
	F_LocDRSts        uint8   `json:"LocDRSts"`
	F_LongitudeDR     float64 `json:"LongitudeDR"`
	F_LatitudeDR      float64 `json:"LatitudeDR"`
	F_LocGnns1Sts     uint8   `json:"LocGnns1Sts"`
	F_TBOXGPSTime     int64   `json:"TBOXGPSTime"`
	F_LocGnns2Sts     uint8   `json:"LocGnns2Sts"`
	F_LocRTKSts       uint16  `json:"LocRTKSts"`
	F_LocGnns1SatNum  uint8   `json:"LocGnns1SatNum"`
	F_LocGnns2SatNum  uint8   `json:"LocGnns2SatNum"`
}

func To_Evt_D018(_byteBuf *parse.ByteBuf) *Evt_D018 {
	_instance := Evt_D018{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_APN1ConnSts = uint8(_bitBuf.Read(1, true, true))
	_instance.F_APN2ConnSts = uint8(_bitBuf.Read(1, true, true))
	_instance.F_MqttConnFailRsn = uint8(_bitBuf.Read(2, true, true))
	_instance.F_ECallSts = uint8(_bitBuf.Read(4, true, true))
	_instance.F_LocDRSts = uint8(_bitBuf.Read(4, true, true))
	_instance.F_LongitudeDR = float64(_bitBuf.Read(29, true, false)) * 0.000001
	_instance.F_LatitudeDR = float64(_bitBuf.Read(28, true, false)) * 0.000001
	_instance.F_LocGnns1Sts = uint8(_bitBuf.Read(4, true, true))
	_instance.F_TBOXGPSTime = _bitBuf.Read(48, true, true)
	_instance.F_LocGnns2Sts = uint8(_bitBuf.Read(4, true, true))
	_instance.F_LocRTKSts = uint16(_bitBuf.Read(14, true, true))
	_instance.F_LocGnns1SatNum = uint8(_bitBuf.Read(8, true, true))
	_instance.F_LocGnns2SatNum = uint8(_bitBuf.Read(8, true, true))
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D018) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_APN1ConnSts), 1, true, true)
	_bitBuf.Write(int64(_instance.F_APN2ConnSts), 1, true, true)
	_bitBuf.Write(int64(_instance.F_MqttConnFailRsn), 2, true, true)
	_bitBuf.Write(int64(_instance.F_ECallSts), 4, true, true)
	_bitBuf.Write(int64(_instance.F_LocDRSts), 4, true, true)
	_bitBuf.Write(int64(_instance.F_LongitudeDR*1000000), 29, true, false)
	_bitBuf.Write(int64(_instance.F_LatitudeDR*1000000), 28, true, false)
	_bitBuf.Write(int64(_instance.F_LocGnns1Sts), 4, true, true)
	_bitBuf.Write(int64(_instance.F_TBOXGPSTime), 48, true, true)
	_bitBuf.Write(int64(_instance.F_LocGnns2Sts), 4, true, true)
	_bitBuf.Write(int64(_instance.F_LocRTKSts), 14, true, true)
	_bitBuf.Write(int64(_instance.F_LocGnns1SatNum), 8, true, true)
	_bitBuf.Write(int64(_instance.F_LocGnns2SatNum), 8, true, true)
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D019 struct {
	F_evtId            uint16  `json:"evtId"`
	F_evtLen           uint16  `json:"evtLen"`
	F_BCMAvlbly        uint8   `json:"BCMAvlbly"`
	F_CCUAvlbly        uint8   `json:"CCUAvlbly"`
	F_EnrgSplReqEPTRdy uint8   `json:"EnrgSplReqEPTRdy"`
	F_HVDCDCLVSideVol  float32 `json:"HVDCDCLVSideVol"`
	F_BatCrnt          float32 `json:"BatCrnt"`
	F_BatSOC           float32 `json:"BatSOC"`
	F_BatSOCSts        uint8   `json:"BatSOCSts"`
	F_BatVol           float32 `json:"BatVol"`
	F_EnrgSplReq       uint8   `json:"EnrgSplReq"`
	F_EnrgSplReqScene  uint64  `json:"EnrgSplReqScene"`
	F_VehEnrgRdyLvl    uint8   `json:"VehEnrgRdyLvl"`
	F_VehEnrgRdyLvlV   uint8   `json:"VehEnrgRdyLvlV"`
	F_HVEstbCond       uint8   `json:"HVEstbCond"`
}

func To_Evt_D019(_byteBuf *parse.ByteBuf) *Evt_D019 {
	_instance := Evt_D019{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_BCMAvlbly = uint8(_bitBuf.Read(1, true, true))
	_instance.F_CCUAvlbly = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EnrgSplReqEPTRdy = uint8(_bitBuf.Read(1, true, true))
	_instance.F_HVDCDCLVSideVol = float32(_bitBuf.Read(8, true, true)) * 0.125
	_instance.F_BatCrnt = float32(_bitBuf.Read(16, true, true))*0.03125 - 1024
	_instance.F_BatSOC = float32(_bitBuf.Read(8, true, true)) * 0.4
	_instance.F_BatSOCSts = uint8(_bitBuf.Read(2, true, true))
	_instance.F_BatVol = float32(_bitBuf.Read(14, true, true))*0.00097656 + 3
	_instance.F_EnrgSplReq = uint8(_bitBuf.Read(1, true, true))
	_instance.F_EnrgSplReqScene = uint64(_bitBuf.Read(64, true, true))
	_instance.F_VehEnrgRdyLvl = uint8(_bitBuf.Read(3, true, true))
	_instance.F_VehEnrgRdyLvlV = uint8(_bitBuf.Read(1, true, true))
	_instance.F_HVEstbCond = uint8(_bitBuf.Read(2, true, true))
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D019) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_BCMAvlbly), 1, true, true)
	_bitBuf.Write(int64(_instance.F_CCUAvlbly), 1, true, true)
	_bitBuf.Write(int64(_instance.F_EnrgSplReqEPTRdy), 1, true, true)
	_bitBuf.Write(int64(_instance.F_HVDCDCLVSideVol/0.125), 8, true, true)
	_bitBuf.Write(int64((_instance.F_BatCrnt+1024)/0.03125), 16, true, true)
	_bitBuf.Write(int64(_instance.F_BatSOC/0.4), 8, true, true)
	_bitBuf.Write(int64(_instance.F_BatSOCSts), 2, true, true)
	_bitBuf.Write(int64((_instance.F_BatVol-3)/0.00097656), 14, true, true)
	_bitBuf.Write(int64(_instance.F_EnrgSplReq), 1, true, true)
	_bitBuf.Write(int64(_instance.F_EnrgSplReqScene), 64, true, true)
	_bitBuf.Write(int64(_instance.F_VehEnrgRdyLvl), 3, true, true)
	_bitBuf.Write(int64(_instance.F_VehEnrgRdyLvlV), 1, true, true)
	_bitBuf.Write(int64(_instance.F_HVEstbCond), 2, true, true)
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D01A struct {
	F_evtId                uint16 `json:"evtId"`
	F_evtLen               uint16 `json:"evtLen"`
	F_iEcuSts              uint16 `json:"iEcuSts"`
	F_iIAMInterSts         uint8  `json:"iIAMInterSts"`
	F_iMpuIPTableRuleSts   uint8  `json:"iMpuIPTableRuleSts"`
	F_iModemIPTableRuleSts uint8  `json:"iModemIPTableRuleSts"`
	F_iARPRuleSts          uint8  `json:"iARPRuleSts"`
	F_iICC2PHYSGMIISts     uint16 `json:"iICC2PHYSGMIISts"`
	F_iMpuRGMIISts         uint16 `json:"iMpuRGMIISts"`
	F_iModemRGMIISts       uint16 `json:"iModemRGMIISts"`
	F_iSwitchSGMIISts      uint16 `json:"iSwitchSGMIISts"`
	F_iUSBConnSts          uint8  `json:"iUSBConnSts"`
	F_iIPASts              uint8  `json:"iIPASts"`
	F_iAPSts               uint8  `json:"iAPSts"`
	F_networkbackupinfo    string `json:"networkbackupinfo"`
}

func To_Evt_D01A(_byteBuf *parse.ByteBuf) *Evt_D01A {
	_instance := Evt_D01A{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_instance.F_iEcuSts = _byteBuf.Read_uint16()
	_instance.F_iIAMInterSts = _byteBuf.Read_uint8()
	_instance.F_iMpuIPTableRuleSts = _byteBuf.Read_uint8()
	_instance.F_iModemIPTableRuleSts = _byteBuf.Read_uint8()
	_instance.F_iARPRuleSts = _byteBuf.Read_uint8()
	_instance.F_iICC2PHYSGMIISts = _byteBuf.Read_uint16()
	_instance.F_iMpuRGMIISts = _byteBuf.Read_uint16()
	_instance.F_iModemRGMIISts = _byteBuf.Read_uint16()
	_instance.F_iSwitchSGMIISts = _byteBuf.Read_uint16()
	_instance.F_iUSBConnSts = _byteBuf.Read_uint8()
	_instance.F_iIPASts = _byteBuf.Read_uint8()
	_instance.F_iAPSts = _byteBuf.Read_uint8()
	_instance.F_networkbackupinfo = _byteBuf.Read_string_utf8(28)
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D01A) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.F_iEcuSts)
	_byteBuf.Write_uint8(_instance.F_iIAMInterSts)
	_byteBuf.Write_uint8(_instance.F_iMpuIPTableRuleSts)
	_byteBuf.Write_uint8(_instance.F_iModemIPTableRuleSts)
	_byteBuf.Write_uint8(_instance.F_iARPRuleSts)
	_byteBuf.Write_uint16(_instance.F_iICC2PHYSGMIISts)
	_byteBuf.Write_uint16(_instance.F_iMpuRGMIISts)
	_byteBuf.Write_uint16(_instance.F_iModemRGMIISts)
	_byteBuf.Write_uint16(_instance.F_iSwitchSGMIISts)
	_byteBuf.Write_uint8(_instance.F_iUSBConnSts)
	_byteBuf.Write_uint8(_instance.F_iIPASts)
	_byteBuf.Write_uint8(_instance.F_iAPSts)
	_byteBuf.Write_string_utf8(_instance.F_networkbackupinfo)
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D01B struct {
	F_evtId          uint16 `json:"evtId"`
	F_evtLen         uint16 `json:"evtLen"`
	F_WANStatus      uint8  `json:"WANStatus"`
	F_ChannelType1   uint8  `json:"ChannelType1"`
	F_ChannelStates1 uint8  `json:"ChannelStates1"`
	F_IPAddress1     string `json:"IPAddress1"`
	F_ChannelType2   uint8  `json:"ChannelType2"`
	F_ChannelStates2 uint8  `json:"ChannelStates2"`
	F_IPAddress2     string `json:"IPAddress2"`
	F_ChannelType3   uint8  `json:"ChannelType3"`
	F_ChannelStates3 uint8  `json:"ChannelStates3"`
	F_IPAddress3     string `json:"IPAddress3"`
	F_ChannelType4   uint8  `json:"ChannelType4"`
	F_ChannelStates4 uint8  `json:"ChannelStates4"`
}

func To_Evt_D01B(_byteBuf *parse.ByteBuf) *Evt_D01B {
	_instance := Evt_D01B{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_instance.F_WANStatus = _byteBuf.Read_uint8()
	_instance.F_ChannelType1 = _byteBuf.Read_uint8()
	_instance.F_ChannelStates1 = _byteBuf.Read_uint8()
	_instance.F_IPAddress1 = _byteBuf.Read_string_utf8(18)
	_instance.F_ChannelType2 = _byteBuf.Read_uint8()
	_instance.F_ChannelStates2 = _byteBuf.Read_uint8()
	_instance.F_IPAddress2 = _byteBuf.Read_string_utf8(18)
	_instance.F_ChannelType3 = _byteBuf.Read_uint8()
	_instance.F_ChannelStates3 = _byteBuf.Read_uint8()
	_instance.F_IPAddress3 = _byteBuf.Read_string_utf8(18)
	_instance.F_ChannelType4 = _byteBuf.Read_uint8()
	_instance.F_ChannelStates4 = _byteBuf.Read_uint8()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D01B) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_WANStatus)
	_byteBuf.Write_uint8(_instance.F_ChannelType1)
	_byteBuf.Write_uint8(_instance.F_ChannelStates1)
	_byteBuf.Write_string_utf8(_instance.F_IPAddress1)
	_byteBuf.Write_uint8(_instance.F_ChannelType2)
	_byteBuf.Write_uint8(_instance.F_ChannelStates2)
	_byteBuf.Write_string_utf8(_instance.F_IPAddress2)
	_byteBuf.Write_uint8(_instance.F_ChannelType3)
	_byteBuf.Write_uint8(_instance.F_ChannelStates3)
	_byteBuf.Write_string_utf8(_instance.F_IPAddress3)
	_byteBuf.Write_uint8(_instance.F_ChannelType4)
	_byteBuf.Write_uint8(_instance.F_ChannelStates4)
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D01C struct {
	F_evtId      uint16 `json:"evtId"`
	F_evtLen     uint16 `json:"evtLen"`
	F_IPAddress4 string `json:"IPAddress4"`
	F_CurIMSI    string `json:"CurIMSI"`
	F_NetType    uint8  `json:"NetType"`
	F_rssi       uint16 `json:"rssi"`
	F_rsrp       uint16 `json:"rsrp"`
	F_rscp       uint16 `json:"rscp"`
	F_sinr       uint16 `json:"sinr"`
	F_ecio       uint16 `json:"ecio"`
}

func To_Evt_D01C(_byteBuf *parse.ByteBuf) *Evt_D01C {
	_instance := Evt_D01C{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_instance.F_IPAddress4 = _byteBuf.Read_string_utf8(18)
	_instance.F_CurIMSI = _byteBuf.Read_string_utf8(17)
	_instance.F_NetType = _byteBuf.Read_uint8()
	_instance.F_rssi = _byteBuf.Read_uint16()
	_instance.F_rsrp = _byteBuf.Read_uint16()
	_instance.F_rscp = _byteBuf.Read_uint16()
	_instance.F_sinr = _byteBuf.Read_uint16()
	_instance.F_ecio = _byteBuf.Read_uint16()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D01C) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_string_utf8(_instance.F_IPAddress4)
	_byteBuf.Write_string_utf8(_instance.F_CurIMSI)
	_byteBuf.Write_uint8(_instance.F_NetType)
	_byteBuf.Write_uint16(_instance.F_rssi)
	_byteBuf.Write_uint16(_instance.F_rsrp)
	_byteBuf.Write_uint16(_instance.F_rscp)
	_byteBuf.Write_uint16(_instance.F_sinr)
	_byteBuf.Write_uint16(_instance.F_ecio)
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D01D struct {
	F_evtId     uint16 `json:"evtId"`
	F_evtLen    uint16 `json:"evtLen"`
	F_cellLAC5G uint32 `json:"cellLAC5G"`
	F_CellID5G  uint64 `json:"CellID5G"`
}

func To_Evt_D01D(_byteBuf *parse.ByteBuf) *Evt_D01D {
	_instance := Evt_D01D{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_instance.F_cellLAC5G = _byteBuf.Read_uint32()
	_instance.F_CellID5G = _byteBuf.Read_uint64()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}
func (__instance *Evt_D01D) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.F_cellLAC5G)
	_byteBuf.Write_uint64(_instance.F_CellID5G)
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_D01F struct {
	F_evtId              uint16 `json:"evtId"`
	F_evtLen             uint16 `json:"evtLen"`
	F_NetRecvRsn         uint8  `json:"NetRecvRsn"`
	F_NetRecvActn        uint8  `json:"NetRecvActn"`
	F_NetRecvActnTimstmp int64  `json:"NetRecvActnTimstmp"`
	F_NetRecvActnCnt     uint8  `json:"NetRecvActnCnt"`
	F_NetRecvActnRst     uint8  `json:"NetRecvActnRst"`
	F_NetRecvtime        int64  `json:"NetRecvtime"`
}

func To_Evt_D01F(_byteBuf *parse.ByteBuf) *Evt_D01F {
	_instance := Evt_D01F{}
	_instance.F_evtId = _byteBuf.Read_uint16()
	_instance.F_evtLen = _byteBuf.Read_uint16()
	index := _byteBuf.ReaderIndex()
	_instance.F_NetRecvRsn = _byteBuf.Read_uint8()
	_instance.F_NetRecvActn = _byteBuf.Read_uint8()
	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	_instance.F_NetRecvActnTimstmp = _bitBuf.Read(48, true, true)
	_bitBuf.Finish()
	_instance.F_NetRecvActnCnt = _byteBuf.Read_uint8()
	_instance.F_NetRecvActnRst = _byteBuf.Read_uint8()
	_instance.F_NetRecvtime = _bitBuf.Read(48, true, true)
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.ReaderIndex() + index
	if skip > 0 {
		_byteBuf.Skip(skip)
	}
	return &_instance
}

func (__instance *Evt_D01F) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_byteBuf.Write_uint16(_instance.F_evtLen)
	index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.F_NetRecvRsn)
	_byteBuf.Write_uint8(_instance.F_NetRecvActn)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(_instance.F_NetRecvActnTimstmp, 48, true, true)
	_bitBuf.Finish()
	_byteBuf.Write_uint8(_instance.F_NetRecvActnCnt)
	_byteBuf.Write_uint8(_instance.F_NetRecvActnRst)
	_bitBuf.Write(_instance.F_NetRecvtime, 48, true, true)
	_bitBuf.Finish()
	skip := int(_instance.F_evtLen) - _byteBuf.WriterIndex() + index
	if skip > 0 {
		_byteBuf.Write_zero(skip)
	}
}

type Evt_FFFF struct {
	F_evtId  uint16 `json:"evtId"`
	F_EvtCRC uint64 `json:"EvtCRC"`
}

func To_Evt_FFFF(_byteBuf *parse.ByteBuf) *Evt_FFFF {
	_instance := Evt_FFFF{}
	F_evtId_v := _byteBuf.Read_uint16()
	_instance.F_evtId = F_evtId_v

	_bitBuf := parse.ToBitBuf_reader(_byteBuf)
	F_EvtCRC_v := _bitBuf.Read(48, true, true)
	_bitBuf.Finish()
	_instance.F_EvtCRC = uint64(F_EvtCRC_v)

	return &_instance
}

func (__instance *Evt_FFFF) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.F_evtId)
	_bitBuf := parse.ToBitBuf_writer(_byteBuf)
	_bitBuf.Write(int64(_instance.F_EvtCRC), 48, true, true)
	_bitBuf.Finish()
}

type Packet struct {
	F_evt_0001        *Evt_0001         `json:"evt_0001,omitempty"`
	F_evt_0003        *Evt_0003         `json:"evt_0003,omitempty"`
	F_evt_0004        *Evt_0004         `json:"evt_0004,omitempty"`
	F_evt_0005        *Evt_0005         `json:"evt_0005,omitempty"`
	F_evt_0006        *Evt_0006         `json:"evt_0006,omitempty"`
	F_evt_0007        *Evt_0007         `json:"evt_0007,omitempty"`
	F_evt_0008        *Evt_0008         `json:"evt_0008,omitempty"`
	F_evt_0009        *Evt_0009         `json:"evt_0009,omitempty"`
	F_evt_000A        *Evt_000A         `json:"evt_000A,omitempty"`
	F_evt_000B        *Evt_000B         `json:"evt_000B,omitempty"`
	F_evt_000C        *Evt_000C         `json:"evt_000C,omitempty"`
	F_evt_000D        *Evt_000D         `json:"evt_000D,omitempty"`
	F_evt_000E        *Evt_000E         `json:"evt_000E,omitempty"`
	F_evt_000F        *Evt_000F         `json:"evt_000F,omitempty"`
	F_evt_0800        *Evt_0800         `json:"evt_0800,omitempty"`
	F_evt_0801        *Evt_0801         `json:"evt_0801,omitempty"`
	F_evt_0802        *Evt_0802         `json:"evt_0802,omitempty"`
	F_evt_0803        *Evt_0803         `json:"evt_0803,omitempty"`
	F_evt_D006        *Evt_D006         `json:"evt_D006,omitempty"`
	F_evt_D008        *Evt_D008         `json:"evt_D008,omitempty"`
	F_evt_D009        *Evt_D009         `json:"evt_D009,omitempty"`
	F_evt_D00A        *Evt_D00A         `json:"evt_D00A,omitempty"`
	F_evt_D00B        *Evt_D00B         `json:"evt_D00B,omitempty"`
	F_evt_D00C        *Evt_D00C         `json:"evt_D00C,omitempty"`
	F_evt_D00D        *Evt_D00D         `json:"evt_D00D,omitempty"`
	F_evt_D00E        *Evt_D00E         `json:"evt_D00E,omitempty"`
	F_evt_D00F        *Evt_D00F         `json:"evt_D00F,omitempty"`
	F_evt_D010        *Evt_D010         `json:"evt_D010,omitempty"`
	F_evt_D011        *Evt_D011         `json:"evt_D011,omitempty"`
	F_evt_D012        *Evt_D012         `json:"evt_D012,omitempty"`
	F_evt_D013        *Evt_D013         `json:"evt_D013,omitempty"`
	F_evt_D014        *Evt_D014         `json:"evt_D014,omitempty"`
	F_evt_D015        *Evt_D015         `json:"evt_D015,omitempty"`
	F_evt_D016        *Evt_D016         `json:"evt_D016,omitempty"`
	F_evt_D017        *Evt_D017         `json:"evt_D017,omitempty"`
	F_evt_D018        *Evt_D018         `json:"evt_D018,omitempty"`
	F_evt_D019        *Evt_D019         `json:"evt_D019,omitempty"`
	F_evt_D01A        *Evt_D01A         `json:"evt_D01A,omitempty"`
	F_evt_D01B        *Evt_D01B         `json:"evt_D01B,omitempty"`
	F_evt_D01C        *Evt_D01C         `json:"evt_D01C,omitempty"`
	F_evt_D01D        *Evt_D01D         `json:"evt_D01D,omitempty"`
	F_evt_D01F        *Evt_D01F         `json:"evt_D01F,omitempty"`
	F_evt_2_6_unknown []Evt_2_6_unknown `json:"evt_2_6_unknown,omitempty"`
	F_evt_4_x_unknown []Evt_4_x_unknown `json:"evt_4_x_unknown,omitempty"`
	F_evt_FFFF        *Evt_FFFF         `json:"evt_FFFF,omitempty"`
}

func To_Packet(_byteBuf *parse.ByteBuf) *Packet {
	_instance := Packet{}
A:
	for _byteBuf.Readable() {
		evtId := _byteBuf.Get_uint16()
		switch evtId {
		case 0x0001:
			if _instance.F_evt_0001 != nil {
				break A
			}
			_instance.F_evt_0001 = To_Evt_0001(_byteBuf)
		case 0x0003:
			_instance.F_evt_0003 = To_Evt_0003(_byteBuf)
		case 0x0004:
			_instance.F_evt_0004 = To_Evt_0004(_byteBuf)
		case 0x0005:
			_instance.F_evt_0005 = To_Evt_0005(_byteBuf)
		case 0x0006:
			_instance.F_evt_0006 = To_Evt_0006(_byteBuf)
		case 0x0007:
			_instance.F_evt_0007 = To_Evt_0007(_byteBuf)
		case 0x0008:
			_instance.F_evt_0008 = To_Evt_0008(_byteBuf)
		case 0x0009:
			_instance.F_evt_0009 = To_Evt_0009(_byteBuf)
		case 0x000A:
			_instance.F_evt_000A = To_Evt_000A(_byteBuf)
		case 0x000B:
			_instance.F_evt_000B = To_Evt_000B(_byteBuf)
		case 0x000C:
			_instance.F_evt_000C = To_Evt_000C(_byteBuf)
		case 0x000D:
			_instance.F_evt_000D = To_Evt_000D(_byteBuf)
		case 0x000E:
			_instance.F_evt_000E = To_Evt_000E(_byteBuf)
		case 0x000F:
			_instance.F_evt_000F = To_Evt_000F(_byteBuf)
		case 0x0800:
			_instance.F_evt_0800 = To_Evt_0800(_byteBuf)
		case 0x0801:
			_instance.F_evt_0801 = To_Evt_0801(_byteBuf)
		case 0x0802:
			_instance.F_evt_0802 = To_Evt_0802(_byteBuf)
		case 0x0803:
			_instance.F_evt_0803 = To_Evt_0803(_byteBuf)
		case 0xD006:
			_instance.F_evt_D006 = To_Evt_D006(_byteBuf)
		case 0xD008:
			_instance.F_evt_D008 = To_Evt_D008(_byteBuf)
		case 0xD009:
			_instance.F_evt_D009 = To_Evt_D009(_byteBuf)
		case 0xD00A:
			_instance.F_evt_D00A = To_Evt_D00A(_byteBuf)
		case 0xD00B:
			_instance.F_evt_D00B = To_Evt_D00B(_byteBuf)
		case 0xD00C:
			_instance.F_evt_D00C = To_Evt_D00C(_byteBuf)
		case 0xD00D:
			_instance.F_evt_D00D = To_Evt_D00D(_byteBuf)
		case 0xD00E:
			_instance.F_evt_D00E = To_Evt_D00E(_byteBuf)
		case 0xD00F:
			_instance.F_evt_D00F = To_Evt_D00F(_byteBuf)
		case 0xD010:
			_instance.F_evt_D010 = To_Evt_D010(_byteBuf)
		case 0xD011:
			_instance.F_evt_D011 = To_Evt_D011(_byteBuf)
		case 0xD012:
			_instance.F_evt_D012 = To_Evt_D012(_byteBuf)
		case 0xD013:
			_instance.F_evt_D013 = To_Evt_D013(_byteBuf)
		case 0xD014:
			_instance.F_evt_D014 = To_Evt_D014(_byteBuf)
		case 0xD015:
			_instance.F_evt_D015 = To_Evt_D015(_byteBuf)
		case 0xD016:
			_instance.F_evt_D016 = To_Evt_D016(_byteBuf)
		case 0xD017:
			_instance.F_evt_D017 = To_Evt_D017(_byteBuf)
		case 0xD018:
			_instance.F_evt_D018 = To_Evt_D018(_byteBuf)
		case 0xD019:
			_instance.F_evt_D019 = To_Evt_D019(_byteBuf)
		case 0xD01A:
			_instance.F_evt_D01A = To_Evt_D01A(_byteBuf)
		case 0xD01B:
			_instance.F_evt_D01B = To_Evt_D01B(_byteBuf)
		case 0xD01C:
			_instance.F_evt_D01C = To_Evt_D01C(_byteBuf)
		case 0xD01D:
			_instance.F_evt_D01D = To_Evt_D01D(_byteBuf)
		case 0xD01F:
			_instance.F_evt_D01F = To_Evt_D01F(_byteBuf)
		case 0xFFFF:
			_instance.F_evt_FFFF = To_Evt_FFFF(_byteBuf)
		default:
			if evtId >= 0x0001 && evtId <= 0xAFFF {
				_instance.F_evt_2_6_unknown = append(_instance.F_evt_2_6_unknown, *To_Evt_2_6_unknown(_byteBuf))
			} else if evtId >= 0xD000 && evtId <= 0xDFFF {
				_instance.F_evt_4_x_unknown = append(_instance.F_evt_4_x_unknown, *To_Evt_4_x_unknown(_byteBuf))
			} else {
				evtIdHex := hex.EncodeToString([]byte{uint8(evtId >> 8), uint8(evtId)})
				util.Log.Warnf("evtId[%s] not support", evtIdHex)
				return nil
			}

		}
	}
	return &_instance
}

func (__instance *Packet) Write(_byteBuf *parse.ByteBuf) {
	_instance := *__instance
	if _instance.F_evt_0001 != nil {
		_instance.F_evt_0001.Write(_byteBuf)
	}
	if _instance.F_evt_0003 != nil {
		_instance.F_evt_0003.Write(_byteBuf)
	}
	if _instance.F_evt_0004 != nil {
		_instance.F_evt_0004.Write(_byteBuf)
	}
	if _instance.F_evt_0005 != nil {
		_instance.F_evt_0005.Write(_byteBuf)
	}
	if _instance.F_evt_0006 != nil {
		_instance.F_evt_0006.Write(_byteBuf)
	}
	if _instance.F_evt_0007 != nil {
		_instance.F_evt_0007.Write(_byteBuf)
	}
	if _instance.F_evt_0008 != nil {
		_instance.F_evt_0008.Write(_byteBuf)
	}
	if _instance.F_evt_0009 != nil {
		_instance.F_evt_0009.Write(_byteBuf)
	}
	if _instance.F_evt_000A != nil {
		_instance.F_evt_000A.Write(_byteBuf)
	}
	if _instance.F_evt_000B != nil {
		_instance.F_evt_000B.Write(_byteBuf)
	}
	if _instance.F_evt_000C != nil {
		_instance.F_evt_000C.Write(_byteBuf)
	}
	if _instance.F_evt_000D != nil {
		_instance.F_evt_000D.Write(_byteBuf)
	}
	if _instance.F_evt_000E != nil {
		_instance.F_evt_000E.Write(_byteBuf)
	}
	if _instance.F_evt_000F != nil {
		_instance.F_evt_000F.Write(_byteBuf)
	}
	if _instance.F_evt_0800 != nil {
		_instance.F_evt_0800.Write(_byteBuf)
	}
	if _instance.F_evt_0801 != nil {
		_instance.F_evt_0801.Write(_byteBuf)
	}
	if _instance.F_evt_0802 != nil {
		_instance.F_evt_0802.Write(_byteBuf)
	}
	if _instance.F_evt_0803 != nil {
		_instance.F_evt_0803.Write(_byteBuf)
	}
	if _instance.F_evt_D006 != nil {
		_instance.F_evt_D006.Write(_byteBuf)
	}
	if _instance.F_evt_D008 != nil {
		_instance.F_evt_D008.Write(_byteBuf)
	}
	if _instance.F_evt_D009 != nil {
		_instance.F_evt_D009.Write(_byteBuf)
	}
	if _instance.F_evt_D00A != nil {
		_instance.F_evt_D00A.Write(_byteBuf)
	}
	if _instance.F_evt_D00B != nil {
		_instance.F_evt_D00B.Write(_byteBuf)
	}
	if _instance.F_evt_D00C != nil {
		_instance.F_evt_D00C.Write(_byteBuf)
	}
	if _instance.F_evt_D00D != nil {
		_instance.F_evt_D00D.Write(_byteBuf)
	}
	if _instance.F_evt_D00E != nil {
		_instance.F_evt_D00E.Write(_byteBuf)
	}
	if _instance.F_evt_D00F != nil {
		_instance.F_evt_D00F.Write(_byteBuf)
	}
	if _instance.F_evt_D010 != nil {
		_instance.F_evt_D010.Write(_byteBuf)
	}
	if _instance.F_evt_D011 != nil {
		_instance.F_evt_D011.Write(_byteBuf)
	}
	if _instance.F_evt_D012 != nil {
		_instance.F_evt_D012.Write(_byteBuf)
	}
	if _instance.F_evt_D013 != nil {
		_instance.F_evt_D013.Write(_byteBuf)
	}
	if _instance.F_evt_D014 != nil {
		_instance.F_evt_D014.Write(_byteBuf)
	}
	if _instance.F_evt_D015 != nil {
		_instance.F_evt_D015.Write(_byteBuf)
	}
	if _instance.F_evt_D016 != nil {
		_instance.F_evt_D016.Write(_byteBuf)
	}
	if _instance.F_evt_D017 != nil {
		_instance.F_evt_D017.Write(_byteBuf)
	}
	if _instance.F_evt_D018 != nil {
		_instance.F_evt_D018.Write(_byteBuf)
	}
	if _instance.F_evt_D019 != nil {
		_instance.F_evt_D019.Write(_byteBuf)
	}
	if _instance.F_evt_D01A != nil {
		_instance.F_evt_D01A.Write(_byteBuf)
	}
	if _instance.F_evt_D01B != nil {
		_instance.F_evt_D01B.Write(_byteBuf)
	}
	if _instance.F_evt_D01C != nil {
		_instance.F_evt_D01C.Write(_byteBuf)
	}
	if _instance.F_evt_D01D != nil {
		_instance.F_evt_D01D.Write(_byteBuf)
	}
	if _instance.F_evt_D01F != nil {
		_instance.F_evt_D01F.Write(_byteBuf)
	}
	if _instance.F_evt_2_6_unknown != nil {
		for i, _ := range _instance.F_evt_2_6_unknown {
			_instance.F_evt_2_6_unknown[i].Write(_byteBuf)
		}
	}
	if _instance.F_evt_4_x_unknown != nil {
		for i, _ := range _instance.F_evt_4_x_unknown {
			_instance.F_evt_4_x_unknown[i].Write(_byteBuf)
		}
	}
	if _instance.F_evt_FFFF != nil {
		_instance.F_evt_FFFF.Write(_byteBuf)
	}
}

func To_Packets(_byteBuf *parse.ByteBuf) []Packet {
	var ps []Packet
	for _byteBuf.Readable() {
		ps = append(ps, *To_Packet(_byteBuf))
	}
	return ps
}

func Write_Packets(packets []Packet, _byteBuf *parse.ByteBuf) {
	for i, _ := range packets {
		packets[i].Write(_byteBuf)
	}
}

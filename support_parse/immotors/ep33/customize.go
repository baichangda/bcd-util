package ep33

import (
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"encoding/hex"
)

func To_F_evts(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	evts := make([]any, 0)
	for _byteBuf.Readable() {
		evtId := _byteBuf.Get_uint16()
		var evt any
		switch evtId {
		case 0x0001:
			evt = To_Evt_0001(_byteBuf, _parentParseContext)
		case 0x0004:
			evt = To_Evt_0004(_byteBuf, _parentParseContext)
		case 0x0005:
			evt = To_Evt_0005(_byteBuf, _parentParseContext)
		case 0x0006:
			evt = To_Evt_0006(_byteBuf, _parentParseContext)
		case 0x0007:
			evt = To_Evt_0007(_byteBuf, _parentParseContext)
		case 0x0008:
			evt = To_Evt_0008(_byteBuf, _parentParseContext)
		case 0x0009:
			evt = To_Evt_0009(_byteBuf, _parentParseContext)
		case 0x000A:
			evt = To_Evt_000A(_byteBuf, _parentParseContext)
		case 0x0801:
			evt = To_Evt_0801(_byteBuf, _parentParseContext)
		case 0x0802:
			evt = To_Evt_0802(_byteBuf, _parentParseContext)
		case 0x0803:
			evt = To_Evt_0803(_byteBuf, _parentParseContext)
		case 0xD006:
			evt = To_Evt_D006(_byteBuf, _parentParseContext)
		case 0xD008:
			evt = To_Evt_D008(_byteBuf, _parentParseContext)
		case 0xD009:
			evt = To_Evt_D009(_byteBuf, _parentParseContext)
		case 0xD00A:
			evt = To_Evt_D00A(_byteBuf, _parentParseContext)
		case 0xD00B:
			evt = To_Evt_D00B(_byteBuf, _parentParseContext)
		case 0xD00C:
			evt = To_Evt_D00C(_byteBuf, _parentParseContext)
		case 0xD00D:
			evt = To_Evt_D00D(_byteBuf, _parentParseContext)
		case 0xD00E:
			evt = To_Evt_D00E(_byteBuf, _parentParseContext)
		case 0xD00F:
			evt = To_Evt_D00F(_byteBuf, _parentParseContext)
		case 0xD01D:
			evt = To_Evt_D01D(_byteBuf, _parentParseContext)
		case 0xFFFF:
			evt = To_Evt_FFFF(_byteBuf, _parentParseContext)
		default:
			if evtId >= 0x0001 && evtId <= 0x07FF ||
				(evtId >= 0x0800 && evtId <= 0x0FFF) ||
				(evtId >= 0x1000 && evtId <= 0x2FFF) ||
				(evtId >= 0x3000 && evtId <= 0x4FFF) ||
				(evtId >= 0x5000 && evtId <= 0x5FFF) ||
				(evtId >= 0x6000 && evtId <= 0x6FFF) ||
				(evtId >= 0x7000 && evtId <= 0x8FFF) ||
				(evtId >= 0x9000 && evtId <= 0xAFFF) {
				evt = To_Evt_2_6_unknown(_byteBuf, _parentParseContext)
			} else if evtId >= 0xD000 && evtId <= 0xDFFF {
				evt = To_Evt_4_x_unknown(_byteBuf, _parentParseContext)
			} else {
				evtIdHex := hex.EncodeToString([]byte{uint8(evtId >> 8), uint8(evtId)})
				util.Log.Warnf("evtId[%s] not support", evtIdHex)
				return nil
			}

		}
		evts = append(evts, evt)
	}
	return evts
}

func Write_F_evts(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	evts := __instance.([]any)
	for _, e := range evts {
		e.(parse.Writeable).Write(_byteBuf, _parentParseContext)
	}
}

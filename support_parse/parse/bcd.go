package parse

import "bcd-util/util"

var BCD_DUMP_TABLE [308]byte

func init() {
	DIGITS := []byte("0123456789")
	for _, b1 := range DIGITS {
		n1 := int(b1) - 48
		for _, b2 := range DIGITS {
			n2 := int(b2) - 48
			i := ((n1 << 4) | n2) << 1
			BCD_DUMP_TABLE[i] = b1
			BCD_DUMP_TABLE[i+1] = b2
		}
	}
}

func BCD_bytesToString(bs []byte) string {
	res := make([]byte, len(bs)<<1)
	for i, b := range bs {
		start := int(b) << 1
		copy(res[i<<1:], BCD_DUMP_TABLE[start:start+2])
	}
	return util.Bytes2String(res)
}

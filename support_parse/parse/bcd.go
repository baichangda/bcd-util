package parse

var BCD_DUMP_TABLE [308]byte

func init() {
	DIGITS := []byte("0123456789")
	for _, c1 := range DIGITS {
		n1 := int(c1) - 48
		for _, c2 := range DIGITS {
			n2 := int(c2) - 48
			i := ((n1 << 4) | n2) << 1
			BCD_DUMP_TABLE[i] = c1
			BCD_DUMP_TABLE[i+1] = c2
		}
	}
}

func BCD_bytesToString(bs []byte) string {
	chars := make([]byte, len(bs)<<1)
	for i, b := range bs {
		start := int(b) << 1
		copy(chars[i<<1:], BCD_DUMP_TABLE[start:start+2])
	}
	return string(chars)
}

package parse

import (
	"testing"
)

func TestByteBuf_Read_uint64(t *testing.T) {
	buf := ToByteBuf_empty()
	buf.Write_uint64(111)
	readUint16 := buf.Read_uint64()
	t.Log(readUint16)
	if readUint16 != 1 {
		t.Failed()
	}
	buf.Write_uint64_le(111)
	readUint16 = buf.Read_uint64_le()
	t.Log(readUint16)
	if readUint16 != 1 {
		t.Failed()
	}
}

func BenchmarkByteBuf_Read_uint64(b *testing.B) {
	buf := ToByteBuf_empty()
	buf.Write_uint64_le(1)
	buf.MarkReaderIndex()
	for i := 0; i < b.N; i++ {
		buf.Read_uint64_le()
		buf.ResetReaderIndex()
	}
}

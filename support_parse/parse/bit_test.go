package parse

import "testing"

func BenchmarkBitBuf_reader_Read(b *testing.B) {
	bytes := []byte{0x81, 0x72, 0x40}
	byteBuf := ToByteBuf(bytes)
	bitBuf_reader := ToBitBuf_reader(byteBuf)
	byteBuf.MarkReaderIndex()
	byteBuf.MarkWriterIndex()
	for i := 0; i < b.N; i++ {
		byteBuf.ResetReaderIndex()
		byteBuf.ResetWriterIndex()
		bitBuf_reader.Read(3, true, true)
		bitBuf_reader.Read(3, true, true)
		bitBuf_reader.Skip(3)
		bitBuf_reader.Read(9, false, false)
	}
}

func TestBitBuf_reader_Read(t *testing.T) {
	bytes := []byte{0x81, 0x72, 0x40}
	byteBuf := ToByteBuf(bytes)
	bitBuf_reader := ToBitBuf_reader(byteBuf)
	read1 := bitBuf_reader.Read(3, true, true)
	read2 := bitBuf_reader.Read(3, true, true)
	bitBuf_reader.Skip(3)
	read3 := bitBuf_reader.Read(9, false, false)

	t.Logf("%d %d %d", read1, read2, read3)
	if !(read1 == 4 && read2 == 0 && read3 == -217) {
		t.Fail()
	}
}

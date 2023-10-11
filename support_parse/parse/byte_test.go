package parse

import (
	"encoding/json"
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

func TestJsonUint8Arr_MarshalJSON(t *testing.T) {
	arr1 := JsonUint8Arr{1, 2, 3}
	marshal1, err := json.Marshal(arr1)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(marshal1))

	var arr2 JsonUint8Arr = nil
	marshal2, err := json.Marshal(arr2)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(marshal2))

	byteBuf := ToByteBuf([]uint8{1, 2, 3})
	arr3 := byteBuf.Read_slice_uint8(3)
	t.Log(string(arr3))

	marshal4, err := json.Marshal([3]uint8{1, 2, 3})
	if err != nil {
		t.Error(err)
	}
	t.Log(string(marshal4))

	marshal5, err := json.Marshal([]uint8{1, 2, 3})
	if err != nil {
		t.Error(err)
	}
	t.Log(string(marshal5))
}

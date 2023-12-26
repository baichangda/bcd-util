package parse

import (
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

type JsonUint8Arr []uint8

func (e JsonUint8Arr) MarshalJSON() ([]byte, error) {
	var result string
	if e == nil {
		result = "null"
	} else {
		sb := strings.Builder{}
		sb.WriteString("[")
		for i, v := range e {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.FormatUint(uint64(v), 10))
		}
		sb.WriteString("]")
		result = sb.String()
	}
	return []byte(result), nil
}

type ByteBuf struct {
	rIndex      int
	wIndex      int
	capacity    int
	bytes       []byte
	mark_rIndex int
	mark_wIndex int
}

func ToByteBuf_capacity(capacity int) *ByteBuf {
	return &ByteBuf{
		rIndex:      0,
		wIndex:      0,
		capacity:    capacity,
		bytes:       make([]byte, capacity),
		mark_rIndex: 0,
		mark_wIndex: 0,
	}
}

func ToByteBuf_empty() *ByteBuf {
	return &ByteBuf{
		rIndex:      0,
		wIndex:      0,
		capacity:    16,
		bytes:       make([]byte, 16),
		mark_rIndex: 0,
		mark_wIndex: 0,
	}
}

func ToByteBuf(bytes []byte) *ByteBuf {
	return &ByteBuf{
		rIndex:      0,
		wIndex:      len(bytes),
		capacity:    len(bytes),
		bytes:       bytes,
		mark_rIndex: 0,
		mark_wIndex: 0,
	}
}

func (b *ByteBuf) checkGrow(n int) {
	oldLen := b.capacity
	minGrowLen := b.wIndex + n - oldLen
	if minGrowLen > 0 {
		prefGrowLen := oldLen >> 1
		oldBytes := b.bytes
		var newBytes []byte
		if prefGrowLen >= minGrowLen {
			b.capacity = oldLen + prefGrowLen
		} else {
			b.capacity = oldLen + minGrowLen
		}
		newBytes = make([]byte, b.capacity)
		copy(newBytes, oldBytes[:b.wIndex])
		b.bytes = newBytes
	}
}

func (b *ByteBuf) ToBytes() []byte {
	return b.bytes[b.rIndex:b.wIndex]
}

func (b *ByteBuf) MarkReaderIndex() {
	b.mark_rIndex = b.rIndex
}

func (b *ByteBuf) MarkWriterIndex() {
	b.mark_wIndex = b.wIndex
}

func (b *ByteBuf) ResetReaderIndex() {
	b.rIndex = b.mark_rIndex
}

func (b *ByteBuf) ResetWriterIndex() {
	b.wIndex = b.mark_wIndex
}

func (b *ByteBuf) Readable() bool {
	return b.rIndex < b.wIndex
}

func (b *ByteBuf) ReaderIndex() int {
	return b.rIndex
}

func (b *ByteBuf) WriterIndex() int {
	return b.wIndex
}

func (b *ByteBuf) ReadableBytes() int {
	return b.wIndex - b.rIndex
}

func (b *ByteBuf) Get_uint8() uint8 {
	return b.bytes[b.rIndex]
}
func (b *ByteBuf) Get_int8() int8 {
	return int8(b.bytes[b.rIndex])
}

func (b *ByteBuf) Read_uint8() uint8 {
	e := b.Get_uint8()
	b.rIndex += 1
	return e
}
func (b *ByteBuf) Read_int8() int8 {
	e := b.Get_int8()
	b.rIndex += 1
	return e
}

func (b *ByteBuf) Get_uint16() uint16 {
	return (uint16(b.bytes[b.rIndex]) << 8) | uint16(b.bytes[b.rIndex+1])
}

func (b *ByteBuf) Get_uint16_le() uint16 {
	return *((*uint16)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+2]))))
}

func (b *ByteBuf) Read_uint16() uint16 {
	res := b.Get_uint16()
	b.rIndex += 2
	return res
}

func (b *ByteBuf) Read_uint16_le() uint16 {
	res := b.Get_uint16_le()
	b.rIndex += 2
	return res
}

func (b *ByteBuf) Get_int16() int16 {
	return (int16(b.bytes[b.rIndex]) << 8) | int16(b.bytes[b.rIndex+1])
}

func (b *ByteBuf) Get_int16_le() int16 {
	return *((*int16)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+2]))))
}

func (b *ByteBuf) Read_int16() int16 {
	res := b.Get_int16()
	b.rIndex += 2
	return res
}

func (b *ByteBuf) Read_int16_le() int16 {
	res := b.Get_int16_le()
	b.rIndex += 2
	return res
}

func (b *ByteBuf) Get_uint24() uint32 {
	return (uint32(b.bytes[b.rIndex+1]) << 16) | (uint32(b.bytes[b.rIndex+2]) << 8) | uint32(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_uint24_le() uint32 {
	return *((*uint32)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+3]))))
}

func (b *ByteBuf) Read_uint24() uint32 {
	res := b.Get_uint24()
	b.rIndex += 3
	return res
}

func (b *ByteBuf) Read_uint24_le() uint32 {
	res := b.Get_uint24_le()
	b.rIndex += 3
	return res
}

func (b *ByteBuf) Get_int24() int32 {
	return (int32(b.bytes[b.rIndex+1]) << 16) | (int32(b.bytes[b.rIndex+2]) << 8) | int32(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_int24_le() int32 {
	return *((*int32)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+3]))))
}

func (b *ByteBuf) Read_int24() int32 {
	res := b.Get_int24()
	b.rIndex += 3
	return res
}

func (b *ByteBuf) Read_int24_le() int32 {
	res := b.Get_int24_le()
	b.rIndex += 3
	return res
}

func (b *ByteBuf) Get_uint32() uint32 {
	return (uint32(b.bytes[b.rIndex]) << 24) | (uint32(b.bytes[b.rIndex+1]) << 16) | (uint32(b.bytes[b.rIndex+2]) << 8) | uint32(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_uint32_le() uint32 {
	return *((*uint32)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+4]))))
}

func (b *ByteBuf) Read_uint32() uint32 {
	res := b.Get_uint32()
	b.rIndex += 4
	return res
}

func (b *ByteBuf) Read_uint32_le() uint32 {
	res := b.Get_uint32_le()
	b.rIndex += 4
	return res
}

func (b *ByteBuf) Get_int32() int32 {
	return (int32(b.bytes[b.rIndex]) << 24) | (int32(b.bytes[b.rIndex+1]) << 16) | (int32(b.bytes[b.rIndex+2]) << 8) | int32(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_int32_le() int32 {
	return *((*int32)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+4]))))
}

func (b *ByteBuf) Read_int32() int32 {
	res := b.Get_int32()
	b.rIndex += 4
	return res
}

func (b *ByteBuf) Read_int32_le() int32 {
	res := b.Get_int32_le()
	b.rIndex += 4
	return res
}

func (b *ByteBuf) Get_uint40() uint64 {
	return (uint64(b.bytes[b.rIndex]) << 32) | (uint64(b.bytes[b.rIndex]) << 24) | (uint64(b.bytes[b.rIndex+1]) << 16) | (uint64(b.bytes[b.rIndex+2]) << 8) | uint64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_uint40_le() uint64 {
	return *((*uint64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+5]))))
}

func (b *ByteBuf) Read_uint40() uint64 {
	res := b.Get_uint40()
	b.rIndex += 5
	return res
}

func (b *ByteBuf) Read_uint40_le() uint64 {
	res := b.Get_uint40_le()
	b.rIndex += 5
	return res
}

func (b *ByteBuf) Get_int40() int64 {
	return (int64(b.bytes[b.rIndex]) << 32) | (int64(b.bytes[b.rIndex]) << 24) | (int64(b.bytes[b.rIndex+1]) << 16) | (int64(b.bytes[b.rIndex+2]) << 8) | int64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_int40_le() int64 {
	return *((*int64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+5]))))
}

func (b *ByteBuf) Read_int40() int64 {
	res := b.Get_int40()
	b.rIndex += 5
	return res
}

func (b *ByteBuf) Read_int40_le() int64 {
	res := b.Get_int40_le()
	b.rIndex += 5
	return res
}

func (b *ByteBuf) Get_uint48() uint64 {
	return (uint64(b.bytes[b.rIndex]) << 40) | (uint64(b.bytes[b.rIndex]) << 32) |
		(uint64(b.bytes[b.rIndex]) << 24) | (uint64(b.bytes[b.rIndex+1]) << 16) | (uint64(b.bytes[b.rIndex+2]) << 8) | uint64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_uint48_le() uint64 {
	return *((*uint64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+6]))))
}

func (b *ByteBuf) Read_uint48() uint64 {
	res := b.Get_uint48()
	b.rIndex += 6
	return res
}

func (b *ByteBuf) Read_uint48_le() uint64 {
	res := b.Get_uint48_le()
	b.rIndex += 6
	return res
}

func (b *ByteBuf) Get_int48() int64 {
	return (int64(b.bytes[b.rIndex]) << 40) | (int64(b.bytes[b.rIndex]) << 32) |
		(int64(b.bytes[b.rIndex]) << 24) | (int64(b.bytes[b.rIndex+1]) << 16) | (int64(b.bytes[b.rIndex+2]) << 8) | int64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_int48_le() int64 {
	return *((*int64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+6]))))
}

func (b *ByteBuf) Read_int48() int64 {
	res := b.Get_int48()
	b.rIndex += 6
	return res
}

func (b *ByteBuf) Read_int48_le() int64 {
	res := b.Get_int48_le()
	b.rIndex += 6
	return res
}

func (b *ByteBuf) Get_uint56() uint64 {
	return (uint64(b.bytes[b.rIndex]) << 48) | (uint64(b.bytes[b.rIndex]) << 40) | (uint64(b.bytes[b.rIndex]) << 32) |
		(uint64(b.bytes[b.rIndex]) << 24) | (uint64(b.bytes[b.rIndex+1]) << 16) | (uint64(b.bytes[b.rIndex+2]) << 8) | uint64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_uint56_le() uint64 {
	return *((*uint64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+7]))))
}

func (b *ByteBuf) Read_uint56() uint64 {
	res := b.Get_uint56()
	b.rIndex += 7
	return res
}

func (b *ByteBuf) Read_uint56_le() uint64 {
	res := b.Get_uint56_le()
	b.rIndex += 7
	return res
}

func (b *ByteBuf) Get_int56() int64 {
	return (int64(b.bytes[b.rIndex]) << 48) | (int64(b.bytes[b.rIndex]) << 40) | (int64(b.bytes[b.rIndex]) << 32) |
		(int64(b.bytes[b.rIndex]) << 24) | (int64(b.bytes[b.rIndex+1]) << 16) | (int64(b.bytes[b.rIndex+2]) << 8) | int64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_int56_le() int64 {
	return *((*int64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+7]))))
}

func (b *ByteBuf) Read_int56() int64 {
	res := b.Get_int56()
	b.rIndex += 7
	return res
}

func (b *ByteBuf) Read_int56_le() int64 {
	res := b.Get_int56_le()
	b.rIndex += 7
	return res
}

func (b *ByteBuf) Get_uint64() uint64 {
	return (uint64(b.bytes[b.rIndex]) << 56) | (uint64(b.bytes[b.rIndex]) << 48) | (uint64(b.bytes[b.rIndex]) << 40) | (uint64(b.bytes[b.rIndex]) << 32) |
		(uint64(b.bytes[b.rIndex]) << 24) | (uint64(b.bytes[b.rIndex+1]) << 16) | (uint64(b.bytes[b.rIndex+2]) << 8) | uint64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_uint64_le() uint64 {
	return *((*uint64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+8]))))
}

func (b *ByteBuf) Read_uint64() uint64 {
	res := b.Get_uint64()
	b.rIndex += 8
	return res
}

func (b *ByteBuf) Read_uint64_le() uint64 {
	res := b.Get_uint64_le()
	b.rIndex += 8
	return res
}

func (b *ByteBuf) Get_int64() int64 {
	return (int64(b.bytes[b.rIndex]) << 56) | (int64(b.bytes[b.rIndex]) << 48) | (int64(b.bytes[b.rIndex]) << 40) | (int64(b.bytes[b.rIndex]) << 32) |
		(int64(b.bytes[b.rIndex]) << 24) | (int64(b.bytes[b.rIndex+1]) << 16) | (int64(b.bytes[b.rIndex+2]) << 8) | int64(b.bytes[b.rIndex+3])
}

func (b *ByteBuf) Get_int64_le() int64 {
	return *((*int64)(unsafe.Pointer(unsafe.SliceData(b.bytes[b.rIndex : b.rIndex+8]))))
}

func (b *ByteBuf) Read_int64() int64 {
	res := b.Get_int64()
	b.rIndex += 8
	return res
}

func (b *ByteBuf) Read_int64_le() int64 {
	res := b.Get_int64_le()
	b.rIndex += 8
	return res
}

func (b *ByteBuf) Get_float32() float32 {
	res := b.Get_uint32()
	return math.Float32frombits(res)
}

func (b *ByteBuf) Get_float32_le() float32 {
	res := b.Get_uint32_le()
	return math.Float32frombits(res)
}

func (b *ByteBuf) Read_float32() float32 {
	res := b.Read_uint32()
	return math.Float32frombits(res)
}

func (b *ByteBuf) Read_float32_le() float32 {
	res := b.Read_uint32_le()
	return math.Float32frombits(res)
}

func (b *ByteBuf) Get_float64() float64 {
	res := b.Get_uint64()
	return math.Float64frombits(res)
}

func (b *ByteBuf) Get_float64_le() float64 {
	res := b.Get_uint64_le()
	return math.Float64frombits(res)
}

func (b *ByteBuf) Read_float64() float64 {
	res := b.Read_uint64()
	return math.Float64frombits(res)
}

func (b *ByteBuf) Read_float64_le() float64 {
	res := b.Read_uint64_le()
	return math.Float64frombits(res)
}

func (b *ByteBuf) Get_bytes(n int) []byte {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	return bytes
}

func (b *ByteBuf) Get_string_utf8(n int) string {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	return string(bytes)
}

func (b *ByteBuf) Read_slice_int8(n int) []int8 {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	b.rIndex += n
	return *(*[]int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(unsafe.SliceData(bytes))),
		Len:  n,
		Cap:  n,
	}))
}

func (b *ByteBuf) Read_slice_uint8(n int) []uint8 {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	b.rIndex += n
	return bytes
}

func (b *ByteBuf) Read_string_utf8(n int) string {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	b.rIndex += n
	return string(bytes)
}

func (b *ByteBuf) Write_uint8(v uint8) {
	b.checkGrow(1)
	b.bytes[b.wIndex] = v
	b.wIndex++
}

func (b *ByteBuf) Write_int8(v int8) {
	b.checkGrow(1)
	b.bytes[b.wIndex] = byte(v)
	b.wIndex++
}

func (b *ByteBuf) Write_uint16(v uint16) {
	b.checkGrow(2)
	b.bytes[b.wIndex] = uint8(v >> 8)
	b.bytes[b.wIndex+1] = uint8(v)
	b.wIndex += 2
}

func (b *ByteBuf) Write_uint16_le(v uint16) {
	b.checkGrow(2)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.wIndex += 2
}

func (b *ByteBuf) Write_int16(v int16) {
	b.checkGrow(2)
	b.bytes[b.wIndex] = uint8(v >> 8)
	b.bytes[b.wIndex+1] = uint8(v)
	b.wIndex += 2
}

func (b *ByteBuf) Write_int16_le(v int16) {
	b.checkGrow(2)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.wIndex += 2
}

func (b *ByteBuf) Write_uint32(v uint32) {
	b.checkGrow(4)
	b.bytes[b.wIndex] = uint8(v >> 24)
	b.bytes[b.wIndex+1] = uint8(v >> 16)
	b.bytes[b.wIndex+2] = uint8(v >> 8)
	b.bytes[b.wIndex+3] = uint8(v)
	b.wIndex += 4
}

func (b *ByteBuf) Write_uint32_le(v uint32) {
	b.checkGrow(4)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.wIndex += 4
}

func (b *ByteBuf) Write_int32(v int32) {
	b.checkGrow(4)
	b.bytes[b.wIndex] = uint8(v >> 24)
	b.bytes[b.wIndex+1] = uint8(v >> 16)
	b.bytes[b.wIndex+2] = uint8(v >> 8)
	b.bytes[b.wIndex+3] = uint8(v)
	b.wIndex += 4
}

func (b *ByteBuf) Write_int32_le(v int32) {
	b.checkGrow(4)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.wIndex += 4
}

func (b *ByteBuf) Write_uint40(v uint64) {
	b.checkGrow(5)
	b.bytes[b.wIndex] = uint8(v >> 32)
	b.bytes[b.wIndex+1] = uint8(v >> 24)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 8)
	b.bytes[b.wIndex+4] = uint8(v)
	b.wIndex += 5
}

func (b *ByteBuf) Write_uint40_le(v uint64) {
	b.checkGrow(5)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.wIndex += 5
}

func (b *ByteBuf) Write_int40(v int64) {
	b.checkGrow(5)
	b.bytes[b.wIndex] = uint8(v >> 32)
	b.bytes[b.wIndex+1] = uint8(v >> 24)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 8)
	b.bytes[b.wIndex+4] = uint8(v)
	b.wIndex += 5
}

func (b *ByteBuf) Write_int40_le(v int64) {
	b.checkGrow(5)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.wIndex += 5
}

func (b *ByteBuf) Write_uint48(v uint64) {
	b.checkGrow(6)
	b.bytes[b.wIndex] = uint8(v >> 40)
	b.bytes[b.wIndex+1] = uint8(v >> 32)
	b.bytes[b.wIndex+2] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 16)
	b.bytes[b.wIndex+5] = uint8(v >> 8)
	b.bytes[b.wIndex+6] = uint8(v)
	b.wIndex += 6
}

func (b *ByteBuf) Write_uint48_le(v uint64) {
	b.checkGrow(6)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.bytes[b.wIndex+5] = uint8(v >> 40)
	b.wIndex += 6
}

func (b *ByteBuf) Write_int48(v int64) {
	b.checkGrow(6)
	b.bytes[b.wIndex] = uint8(v >> 40)
	b.bytes[b.wIndex+1] = uint8(v >> 32)
	b.bytes[b.wIndex+2] = uint8(v >> 24)
	b.bytes[b.wIndex+3] = uint8(v >> 16)
	b.bytes[b.wIndex+4] = uint8(v >> 8)
	b.bytes[b.wIndex+5] = uint8(v)
	b.wIndex += 6
}

func (b *ByteBuf) Write_int48_le(v int64) {
	b.checkGrow(6)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.bytes[b.wIndex+5] = uint8(v >> 40)
	b.wIndex += 6
}

func (b *ByteBuf) Write_uint56(v uint64) {
	b.checkGrow(7)
	b.bytes[b.wIndex] = uint8(v >> 48)
	b.bytes[b.wIndex+1] = uint8(v >> 40)
	b.bytes[b.wIndex+2] = uint8(v >> 32)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 16)
	b.bytes[b.wIndex+5] = uint8(v >> 8)
	b.bytes[b.wIndex+6] = uint8(v)
	b.wIndex += 7
}

func (b *ByteBuf) Write_uint56_le(v uint64) {
	b.checkGrow(7)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.bytes[b.wIndex+5] = uint8(v >> 40)
	b.bytes[b.wIndex+6] = uint8(v >> 48)
	b.wIndex += 7
}

func (b *ByteBuf) Write_int56(v int64) {
	b.checkGrow(7)
	b.bytes[b.wIndex] = uint8(v >> 48)
	b.bytes[b.wIndex+1] = uint8(v >> 40)
	b.bytes[b.wIndex+2] = uint8(v >> 32)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 16)
	b.bytes[b.wIndex+5] = uint8(v >> 8)
	b.bytes[b.wIndex+6] = uint8(v)
	b.wIndex += 7
}

func (b *ByteBuf) Write_int56_le(v int64) {
	b.checkGrow(7)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.bytes[b.wIndex+5] = uint8(v >> 40)
	b.bytes[b.wIndex+6] = uint8(v >> 48)
	b.bytes[b.wIndex+7] = uint8(v >> 56)
	b.wIndex += 7
}

func (b *ByteBuf) Write_uint64(v uint64) {
	b.checkGrow(8)
	b.bytes[b.wIndex] = uint8(v >> 56)
	b.bytes[b.wIndex+1] = uint8(v >> 48)
	b.bytes[b.wIndex+2] = uint8(v >> 40)
	b.bytes[b.wIndex+3] = uint8(v >> 32)
	b.bytes[b.wIndex+4] = uint8(v >> 24)
	b.bytes[b.wIndex+5] = uint8(v >> 16)
	b.bytes[b.wIndex+6] = uint8(v >> 8)
	b.bytes[b.wIndex+7] = uint8(v)
	b.wIndex += 8
}

func (b *ByteBuf) Write_uint64_le(v uint64) {
	b.checkGrow(8)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.bytes[b.wIndex+5] = uint8(v >> 40)
	b.bytes[b.wIndex+6] = uint8(v >> 48)
	b.bytes[b.wIndex+7] = uint8(v >> 56)
	b.wIndex += 8
}

func (b *ByteBuf) Write_int64(v int64) {
	b.checkGrow(8)
	b.bytes[b.wIndex] = uint8(v >> 56)
	b.bytes[b.wIndex+1] = uint8(v >> 48)
	b.bytes[b.wIndex+2] = uint8(v >> 40)
	b.bytes[b.wIndex+3] = uint8(v >> 32)
	b.bytes[b.wIndex+4] = uint8(v >> 24)
	b.bytes[b.wIndex+5] = uint8(v >> 16)
	b.bytes[b.wIndex+6] = uint8(v >> 8)
	b.bytes[b.wIndex+7] = uint8(v)
	b.wIndex += 8
}

func (b *ByteBuf) Write_int64_le(v int64) {
	b.checkGrow(8)
	b.bytes[b.wIndex] = uint8(v)
	b.bytes[b.wIndex+1] = uint8(v >> 8)
	b.bytes[b.wIndex+2] = uint8(v >> 16)
	b.bytes[b.wIndex+3] = uint8(v >> 24)
	b.bytes[b.wIndex+4] = uint8(v >> 32)
	b.bytes[b.wIndex+5] = uint8(v >> 40)
	b.bytes[b.wIndex+6] = uint8(v >> 48)
	b.bytes[b.wIndex+7] = uint8(v >> 56)
	b.wIndex += 8
}

func (b *ByteBuf) Write_float32(v float32) {
	b.Write_uint32(math.Float32bits(v))
}
func (b *ByteBuf) Write_float32_le(v float32) {
	b.Write_uint32_le(math.Float32bits(v))
}
func (b *ByteBuf) Write_float64(v float64) {
	b.Write_uint64(math.Float64bits(v))
}
func (b *ByteBuf) Write_float64_le(v float64) {
	b.Write_uint64_le(math.Float64bits(v))
}

func (b *ByteBuf) Write_slice_uint8(slice []uint8) {
	b.checkGrow(len(slice))
	copy(b.bytes[b.wIndex:], slice)
	b.wIndex += len(slice)
}

func (b *ByteBuf) Write_slice_int8(slice []int8) {
	n := len(slice)
	bytes := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(unsafe.SliceData(slice))),
		Len:  n,
		Cap:  n,
	}))
	b.checkGrow(len(bytes))
	copy(b.bytes[b.wIndex:], bytes)
	b.wIndex += len(bytes)
}

func (b *ByteBuf) Write_zero(n int) {
	if n <= 0 {
		return
	}
	b.checkGrow(n)
	b.wIndex += n
}

func (b *ByteBuf) Write_string_utf8(v string) {
	bytes := []byte(v)
	b.Write_slice_uint8(bytes)
}

func (b *ByteBuf) Skip(n int) {
	b.rIndex += n
}

func (b *ByteBuf) Clear() {
	b.rIndex = 0
	b.wIndex = 0
}

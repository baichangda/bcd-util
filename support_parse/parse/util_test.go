package parse

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestRound(t *testing.T) {
	round1 := Round(1234.5678)
	t.Log(round1)
	if round1 != 1235 {
		t.Fail()
	}

	round2 := Round(1234.4678)
	t.Log(round2)
	if round2 != 1234 {
		t.Fail()
	}

	round3 := Round(float64(-1015))
	t.Log(round3)
	if round2 != -1015 {
		t.Fail()
	}
}

type A struct {
	u1 uint8
	u2 uint8
	u3 uint16
}

func Test_slice_to_struct(t *testing.T) {
	bytes := []byte{0x01, 0x02, 0x03, 0x04}
	a := (*A)(unsafe.Pointer(unsafe.SliceData(bytes)))
	t.Logf("%+v", a)
	bytes[0] = 0x02
	t.Logf("%+v", a)
}

func Test_struct_to_slice(t *testing.T) {
	a := A{
		u1: 1,
		u2: 2,
		u3: 1027,
	}
	b := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&a)),
		Len:  4,
		Cap:  4,
	}))
	t.Logf("%+v", b)
	(*b)[0] = 2
	t.Logf("%+v", &a)
}

func Test_slice_to_array(t *testing.T) {
	bytes := []byte{0x01, 0x02, 0x03, 0x04}
	a := (*[2]uint16)(unsafe.Pointer(unsafe.SliceData(bytes)))
	t.Logf("%+v", a)
}

func Test_slice_to_uint16(t *testing.T) {
	bytes := []byte{0x00, 0x01}
	a := (*uint16)(unsafe.Pointer(unsafe.SliceData(bytes)))
	t.Logf("%+v", *a)
}

func Test_uint16_to_slice(t *testing.T) {
	var i uint16 = 1
	a := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&i)),
		Len:  2,
		Cap:  2,
	}))
	t.Logf("%+v", a)
}

func Test_array_to_uint16(t *testing.T) {
	bytes := [2]byte{0x00, 0x01}
	a := (*uint16)(unsafe.Pointer(&bytes))
	t.Logf("%+v", *a)
}

func Test_uint16_to_array(t *testing.T) {
	var i uint16 = 1
	a := (*[2]byte)(unsafe.Pointer(&i))
	t.Logf("%+v", a)
}

func Test_array_to_slice(t *testing.T) {
	uint16s := [2]uint16{513, 1027}
	a := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&uint16s)),
		Len:  4,
		Cap:  4,
	}))
	t.Logf("%+v", a)
}

type Msg_tailer struct {
	Check_sum uint32
	Tail      [4]uint8
}

func Test_Msg_tailer(t *testing.T) {
	tailer := Msg_tailer{
		Check_sum: 0,
		Tail:      [4]uint8{},
	}
	m := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&tailer)),
		Len:  8,
		Cap:  8,
	}))
	t.Logf("%+v", m)
}

func Test2(t *testing.T) {
	var i1 uint8 = 255
	t.Log(int64(i1))
	t.Log(int8(i1))
	var i2 int8 = -1
	t.Log(uint16(i2))
	t.Log(int16(i2))
	t.Log(uint8(i2))
}

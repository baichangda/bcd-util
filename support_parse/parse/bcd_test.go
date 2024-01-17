package parse

import (
	"testing"
)

func TestBCD_bytesToString(t *testing.T) {
	//for i, r := range BCD_DUMP_TABLE {
	//	println(strconv.Itoa(i) + " " + string(r))
	//}
	if BCD_bytesToString([]byte{133, 153}) != "8599" {
		t.Fail()
	}
	if BCD_bytesToString([]byte{66, 87}) != "4257" {
		t.Fail()
	}
	if BCD_bytesToString([]byte{71, 99}) != "4763" {
		t.Fail()
	}
}

package util

import (
	"encoding/json"
	"testing"
	"time"
)

func TestUint8Arr_MarshalJSON(t *testing.T) {
	arr1 := ByteSlice{1, 2, 3}
	marshal1, err := json.Marshal(arr1)
	if err != nil {
		t.Error(err)
	}
	if string(marshal1) != "[1,2,3]" {
		t.Failed()
	}

	var arr2 ByteSlice
	err = json.Unmarshal(marshal1, &arr2)
	if err != nil {
		t.Error(err)
	}
	if arr2[0] != 1 || arr2[1] != 2 || arr2[2] != 3 {
		t.Failed()
	}
}

func TestUnixTime_MarshalJSON(t *testing.T) {
	_location_china := time.FixedZone("_location_china", 28800)
	parse, err := time.ParseInLocation("20060102150405", "20240126121212", _location_china)
	if err != nil {
		t.Error(err)
	}
	ts := TimeTs(parse)
	marshal1, err := json.Marshal(ts)
	if err != nil {
		t.Error(err)
	}
	if string(marshal1) != "1706242332000" {
		t.Failed()
	}

	var t2 TimeTs
	err = json.Unmarshal(marshal1, &t2)
	if err != nil {
		t.Error(err)
	}
	if time.Time(t2).In(_location_china).Format("20060102150405") != "20240126121212" {
		t.Failed()
	}
}

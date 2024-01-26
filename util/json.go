package util

import (
	"strconv"
	"strings"
	"time"
)

type Uint8Arr []uint8

func (e Uint8Arr) MarshalJSON() ([]byte, error) {
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

type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return ([]byte)(strconv.FormatInt(time.Time(t).UnixMilli(), 10)), nil
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	num, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	*t = UnixTime(time.UnixMilli(int64(num)))
	return nil
}

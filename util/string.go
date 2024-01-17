package util

import (
	"unsafe"
)

// String2Bytes
// 注意此方法转换的[]byte不能只读、如果修改会导致错误
func String2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

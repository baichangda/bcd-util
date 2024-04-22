package util

import (
	"strings"
	"unicode"
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

func CamelCaseToSplitChar(str string, splitChar rune) string {
	if len(str) == 0 {
		return str
	}
	result := strings.Builder{}
	result.WriteRune(unicode.ToLower(rune(str[0])))
	for _, c := range str[1:] {
		if unicode.IsUpper(c) {
			result.WriteRune(splitChar)
		}
		result.WriteRune(unicode.ToLower(c))
	}
	return result.String()
}

func SplitCharToCamelCase(str string, splitChar rune) string {
	if len(str) == 0 {
		return str
	}
	result := strings.Builder{}
	nextIsUpper := false
	result.WriteRune(unicode.ToLower(rune(str[0])))
	for _, c := range str[1:] {
		if c == splitChar {
			nextIsUpper = true
		} else {
			if nextIsUpper {
				result.WriteRune(unicode.ToUpper(c))
				nextIsUpper = false
			} else {
				result.WriteRune(unicode.ToLower(c))
			}
		}
	}
	return result.String()
}

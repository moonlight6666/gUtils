package gUtils

import (
	"bytes"
	"math/rand"
	"reflect"
	"strings"
	"unicode"
	"unsafe"
)

const letterBytes = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func HasSuffixs(path string, ignoreList []string) bool {
	for _, e := range ignoreList {
		if strings.HasSuffix(path, e) {
			return true
		}
	}
	return false
}

func HasStrings(s string, ignoreList []string) bool {
	for _, i := range ignoreList {
		if strings.Contains(s, i) {
			return true
		}
	}
	return false
}

// Unsafe read only
func UnsafeString2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Unsafe read only
func UnsafeBytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func IsChinese(str string) bool {
	for _, v := range str {
		if !unicode.Is(unicode.Han, v) {
			return false
		}
	}
	return true
}

// 驼峰转下划线
func Camel2Case(name string) string {
	buffer := &bytes.Buffer{}
	for i, r := range name {
		if r == '_' || unicode.IsDigit(r) {
			buffer.WriteRune(r)
		} else if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				if i != 0 {
					buffer.WriteRune('_')
				}
				buffer.WriteRune(unicode.ToLower(r))
			} else {
				buffer.WriteRune(r)
			}
		}
	}
	return buffer.String()
}

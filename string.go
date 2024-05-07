package gUtils

import (
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

// RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
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

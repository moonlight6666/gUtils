package gUtils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Md5(s []byte) string {
	return fmt.Sprintf("%x", md5.Sum(s)) //将[]byte转成16进制
}

func Md5String(str string) string {
	data := []byte(str)
	return Md5(data)
}

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

package gUtils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/cespare/xxhash/v2"
	"io"
	"os"
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

func SHA256(s string) string {
	o := sha256.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

func HMAC_SHA256(s string, k string) string {
	h := hmac.New(sha256.New, []byte(k))
	h.Write([]byte(s))
	digest := h.Sum(nil)
	return hex.EncodeToString(digest)
}

func Md5File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func XXhashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 流式哈希，大文件无压力
	hash := xxhash.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", hash.Sum64()), nil
}

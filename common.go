package gUtils

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CheckErrorExit(err error, msg ...string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		//fileBaseName := filepath.Base(file)
		//color.RedString("[ERROR] %s:%d %s %v\n", file, line, msg, err)
		fmt.Printf("\x1b[0;31m[ERROR] %s:%d %s %v\n\x1b[0m", file, line, msg, err)
		os.Exit(1)
	}
}

func Exit(msg ...string) {
	_, file, line, _ := runtime.Caller(1)
	//fileBaseName := filepath.Base(file)
	//color.Red("[ERROR] %s:%d %s\n", file, line, msg)
	fmt.Printf("\x1b[0;31m[ERROR] %s:%d %s\n\x1b[0m", file, line, msg)
	os.Exit(1)
}

func RecoverCatch() {
	if err:=recover();err!=nil{
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		fmt.Printf("Catch Stack =>\n %s\nReason: %v\n", string(buf[:n]), err)

	}
}
func PrintStack(err interface{}) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("Catch Stack =>\n %s\nReason: %v\n", string(buf[:n]), err)
}

func CheckError(err error, msg ... string) {
	if err != nil {
		fmt.Printf("%s %v", msg, err)
	}
}


func FormatFileSize(size int64) string {
	i := float32(size) / 1024.0
	if i >= 1024 {
		i = i / 1024.0
		if i >= 1024 {
			i = i / 1024.0
			return fmt.Sprintf("%.2fGB", i)
		}
		return fmt.Sprintf("%.2fM", i)
	}
	return fmt.Sprintf("%.2fkb", float32(size)/1024.0)
}


func Md5(s []byte ) string{
	return fmt.Sprintf("%x", md5.Sum(s))
}

//将字符串加密成 md5
func Md5String(str string) string {
	data := []byte(str)
	return Md5(data) //将[]byte转成16进制
}






func IsPortInUse(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}
package gUtils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
		logs.GetBeeLogger().Error("%s %v", msg, err)
	}
}


// 获取ip归属地
func GetIpLocation(ip string) (string, error) {
	if strings.HasPrefix(ip, "192.168") || ip == "127.0.0.1" {
		return "内网", nil
	}
	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	var result struct {
		Code int
		Data struct {
			Country string
			Region  string
			City    string
			Isp     string
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		return "未知", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "未知", err
	}
	//logs.Info("result:%v", string(body))

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "未知", err
	}
	if result.Code == 0 {
		if result.Data.Country == "中国" {
			return result.Data.Region + "." + result.Data.City + " " + result.Data.Isp, nil
		}
		return result.Data.Country + "." + result.Data.Region + "." + result.Data.City + " " + result.Data.Isp, nil
	}
	return "未知", err
}

func IsInArray(v string, array [] string) bool {
	for _, e := range array {
		if e == v {
			return true
		}
	}
	return false
}

func RemoveDuplicateArray(s [] interface{}) [] interface{} {
	maps := make(map[interface{}]interface{}, len(s))
	r := make([] interface{}, 0)
	for _, v := range s {
		if _, ok := maps[v]; ok {
			continue
		}
		maps[v] = true
		r = append(r, v)
	}
	return r
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

// 保留N位小数点
func Decimal(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}



func IsPortInUse(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}
package gUtils

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
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

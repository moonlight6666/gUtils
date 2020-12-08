package gUtils

import (
	"fmt"
	"time"
)

//获取昨日0点时间戳
func GetYesterdayZeroTimestamp() int {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return int(tm1.Unix()) - 86400
}

// 获取该日0点时间戳
func GetThatZeroTimestamp(timestamp int64) int {
	if timestamp == 0 {
		return 0
	}
	t := time.Unix(timestamp, 0)
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return int(t1.Unix())
}

//是否是今天
func IsToday(timestamp int64) bool {
	t := time.Unix(timestamp, 0)
	now := time.Now()
	return t.Year() == now.Year() &&  t.Month() == now.Month() && t.Day() == now.Day()
}

// 获取当前时间戳
func GetTimestamp() int {
	return int(time.Now().Unix())
}

//获取今日0点时间戳
func GetTodayZeroTimestamp() int {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return int(tm1.Unix())
}

func FormatTimeLength(sec int) string {
	h := sec / (60 * 60)
	m := sec % (60 * 60) / 60
	s := sec % 60
	if h > 0 && m > 0 && s > 0 {
		return fmt.Sprintf("%d时%d分%d秒", h, m, s)
	} else if m > 0 && s > 0 {
		return fmt.Sprintf("%d分%d秒", m, s)
	} else {
		return fmt.Sprintf("%d秒", sec)
	}
}

func FormatTime(sec int) string {
	h := sec / (60 * 60)
	m := sec % (60 * 60) / 60
	return fmt.Sprintf("%d:%d", h, m)
}

func FormatDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}

func FormatDateTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func FormatMath(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return fmt.Sprintf("%d-%d", t.Year(), t.Month())
}


package gUtils

import (
	"fmt"
	"time"
)

const (
	MinuteSecond = 60
	HourSecond   = MinuteSecond * 60
	DaySecond    = HourSecond * 24
	YearSecond   = DaySecond * 365
)

// 获取昨日0点时间戳
func GetYesterdayZeroTimestamp() int {
	return GetTodayZeroTimestamp() - 86400
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

// 获取本月1号时间戳
func GetThisMonthFirstDayTimestamp(timestamp int) int {
	t := time.Unix(int64(timestamp), 0)
	return int(time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local).Unix())
}

// 获取下月1号时间戳
func GetNextMonthFirstDayTimestamp(timestamp int) int {
	t := time.Unix(int64(timestamp), 0)
	if t.Month() == 12 {
		return int(time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, time.Local).Unix())
	} else {
		return int(time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.Local).Unix())
	}
}

// 是否是今天
func IsToday(timestamp int64) bool {
	t := time.Unix(timestamp, 0)
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// 是否是本月
func IsThisMonth(timestamp int64) bool {
	t := time.Unix(timestamp, 0)
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month()
}

// 获取当前时间戳
func GetTimestamp() int {
	return int(time.Now().Unix())
}

// 转换时间戳
func TranTimestamp(year int, month int, day int, hour int, min int, sec int) int {
	t := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)
	return int(t.Unix())
}

// 获取今日0点时间戳
func GetTodayZeroTimestamp() int {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return int(tm1.Unix())
}

// 获取今日23点59分59秒时间戳
func GetTodayEndTimestamp() int {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return int(tm1.Unix())
}

// return %d小时 | %d时%d分%d秒 | %d分 | %d分%d秒 | %d秒
func FormatTimeLength[T Integer](sec T) string {
	h := sec / (60 * 60)
	m := sec % (60 * 60) / 60
	s := sec % 60
	if h > 0 && m == 0 && s == 0 {
		return fmt.Sprintf("%d小时", h)
	} else if h > 0 {
		return fmt.Sprintf("%d时%d分%d秒", h, m, s)
	} else if m > 0 && s == 0 {
		return fmt.Sprintf("%d分", m)
	} else if m > 0 {
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

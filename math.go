package gUtils

import (
	"fmt"
	"math"
	"strconv"
)

// 计算百分比率 FormatPercentageRate(3,10,2) -> return 30%
func FormatPercentageRate(value int, denominator int, remainNum int)string {
	return fmt.Sprintf("%.2f%%", CalcRate(value *100, denominator, remainNum))
}

//计算比率,并保留N位小数 CaclRate(3,10,2) -> return 0.30
func CalcRate(value int, denominator int, n int) float32 {
	if denominator == 0 {
		return 0
	}
	if denominator > 0 {
		s := "%." + strconv.Itoa(n) + "f"
		v := fmt.Sprintf(s, float32(value)/float32(denominator))
		value, _ := strconv.ParseFloat(v, 64)
		return float32(value)
	}
	return 0
}


// 保留N位小数点
func Decimal(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func Max(vals...int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func Min(vals...int) int {
	var min int
	for _, val := range vals {
		if val <= min {
			min = val
		}
	}
	return min
}


func MaxInt32(vals...int32) int32 {
	var max int32
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func MinInt32(vals...int32) int32 {
	var min int32
	for _, val := range vals {
		if val <= min {
			min = val
		}
	}
	return min
}

func MaxInt64(vals...int64) int64 {
	var max int64
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func MinInt64(vals...int64) int64 {
	var min int64
	for _, val := range vals {
		if val <= min {
			min = val
		}
	}
	return min
}
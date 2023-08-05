package gUtils

import (
	"fmt"
	"math"
	"strconv"
)

// 计算百分比率 FormatPercentageRate(3,10,2) -> return 30%
func FormatPercentageRate(value int, denominator int, remainNum int) string {
	if denominator == 0 {
		return "-"
	}
	return fmt.Sprintf("%.2f%%", CalcRate(value*100, denominator, remainNum))
}

// 计算比率,并保留N位小数 CaclRate(3,10,2) -> return 0.30
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

func Sum[T Number](vals ...T) T {
	var sum T
	for _, val := range vals {
		sum += val
	}
	return sum
}

func Max[T Number](vals ...T) T {
	isSet := false
	var max T
	for _, val := range vals {
		if val > max || isSet == false {
			isSet = true
			max = val
		}
	}
	return max
}

func Min[T Number](vals ...T) T {
	isSet := false
	var min T
	for _, val := range vals {
		if val <= min || isSet == false {
			isSet = true
			min = val
		}
	}
	return min
}

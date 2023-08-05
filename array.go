package gUtils

import (
	"strconv"
)

func IsInArray[T Number | string](v T, array []T) bool {
	for _, e := range array {
		if e == v {
			return true
		}
	}
	return false
}

// []string => []int
func ArrayStr2Int(data []string) []int {
	var (
		arr = make([]int, 0, len(data))
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		var num, _ = strconv.Atoi(data[i])
		arr = append(arr, num)
	}
	return arr
}

// []int => []string
func ArrayInt2Str(data []int) []string {
	var (
		arr = make([]string, 0, len(data))
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		arr = append(arr, strconv.Itoa(data[i]))
	}
	return arr
}

// 合并数组
func MergeArray(dest []interface{}, src []interface{}) (result []interface{}) {
	result = make([]interface{}, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return
}

func RemoveDuplicateArray(s []interface{}) []interface{} {
	maps := make(map[interface{}]interface{}, len(s))
	r := make([]interface{}, 0)
	for _, v := range s {
		if _, ok := maps[v]; ok {
			continue
		}
		maps[v] = true
		r = append(r, v)
	}
	return r
}

func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

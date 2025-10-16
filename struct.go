package gUtils

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

const omitempty = "omitempty"
const omit = "omit"

func StructToMapReflect(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			isOmitempty := false
			isOmit := false
			if strings.Contains(tagValue, ",") {
				for _, subTagValue := range strings.Split(tagValue, ",") {
					if subTagValue != omitempty {
						tagValue = subTagValue
					} else if subTagValue == omitempty {
						isOmitempty = true
					}
					if subTagValue == omit {
						isOmit = true
					}
				}
			}
			if isOmitempty && isOmit == false {
				switch v.Field(i).Kind() {
				case reflect.String:
					isOmit = v.Field(i).String() == ""
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					isOmit = v.Field(i).Int() == 0
				}
			}
			if !isOmit {
				//fmt.Println(tagValue, v.Field(i).String(), isOmitempty)
				out[tagValue] = v.Field(i).Interface()
			}
		}
	}
	return out, nil
}

// => age=1&name=x
func StructToForm(in interface{}, tagName string) (string, error) {
	m, err := StructToMapReflect(in, tagName)
	if err != nil {
		return "", err
	}

	var buf = bytes.Buffer{}
	for k, v := range m {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		str, err := toString(v)
		if err != nil {
			return "", err
		}
		buf.WriteString(str)
	}

	return buf.String(), nil
}

func toString(v interface{}) (string, error) {
	switch vv := v.(type) {
	case string:
		return vv, nil
	case int:
		return strconv.FormatInt(int64(vv), 10), nil
	case uint:
		return strconv.FormatInt(int64(vv), 10), nil
	case int32:
		return strconv.FormatInt(int64(vv), 10), nil
	case uint32:
		return strconv.FormatInt(int64(vv), 10), nil
	case int64:
		return strconv.FormatInt(int64(vv), 10), nil
	case uint64:
		return strconv.FormatInt(int64(vv), 10), nil
	case float32:
		return fmt.Sprintf("%f", vv), nil
	case float64:
		return fmt.Sprintf("%f", vv), nil
	default:
		return "", fmt.Errorf("unknow type to string:%v %v", v, vv)
	}
}

// 判断结构体字段是否有为空值
func CheckStructEmpty(s interface{}, ignoreFields ...string) error {
	value := reflect.ValueOf(s)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	t := value.Type()

	//if t.Kind() == reflect.Ptr {
	//	t = t.Elem()
	//}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		name := t.Field(i).Name
		if isEmpty(field) && !IsInArray(name, ignoreFields) {
			return fmt.Errorf("filed empty:%s", name)
		}
	}

	return nil
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	//case reflect.Bool:
	//	return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.String:
		return v.Len() == 0
	}
	return false
}

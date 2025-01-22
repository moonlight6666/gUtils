package gUtils

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// 生成签名源串(表单)
// return=> age=1&name=x
func MakeSignSourceForm(in interface{}, tagName string) (string, error) {
	return MakeSignSource(in, tagName, '=', '&')
}

// 生成签名源串
// isShowKey：是否显示key
// kvSeg: kv分割符
// seg: 参数分割符
// return： k{{kvSeg}}v{{seg}}k1{{kvSeg}}v1
func MakeSignSource(in interface{}, tagName string, kvSeg byte, seg byte) (string, error) {
	m, err := StructToMapReflect(in, tagName)
	if err != nil {
		return "", err
	}
	keys := make([]string, 0, len(m))

	for k, _ := range m {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var buf = bytes.Buffer{}
	for _, k := range keys {
		if buf.Len() > 0 && seg != ' ' {
			buf.WriteByte(seg)
		}
		if kvSeg != ' ' {
			buf.WriteString(k)
			buf.WriteByte(kvSeg)
		}

		str, err := toString(m[k])
		if err != nil {
			return "", err
		}
		buf.WriteString(str)
	}

	return buf.String(), nil
}

const omitempty = "omitempty"

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
			if strings.Contains(tagValue, ",") {
				for _, subTagValue := range strings.Split(tagValue, ",") {
					if subTagValue != omitempty {
						tagValue = subTagValue
					} else if subTagValue == omitempty {
						isOmitempty = true
					}
				}
			}
			isOmit := false
			if isOmitempty {
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

package gUtils

import (
	"bytes"
	"sort"
)

// 生成签名源串(表单)
// return=> age=1&name=x
func MakeSignSourceForm(in interface{}, tagName string) (string, error) {
	return MakeSignSource(in, tagName, '=', '&')
}

// 生成签名源串
// isShowKey：是否显示key
// kvSeg: kv分割符,如果等于' ',则不显示k
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

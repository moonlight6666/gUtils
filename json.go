package gUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/**
  JSON (map转json)
*/

func ToJsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

/*
泛型比较麻烦,单独做一个
*/
func Struct2Json(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

/*
*

	JSON (json转map)
*/
func StringToJson(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}

type StringOrNumber string

func (v *StringOrNumber) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)

	if len(data) > 0 && data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*v = StringOrNumber(s)
		return nil
	}

	var n json.Number
	if err := json.Unmarshal(data, &n); err != nil {
		return fmt.Errorf("uid must be string or number: %w", err)
	}

	*v = StringOrNumber(n.String())
	return nil
}

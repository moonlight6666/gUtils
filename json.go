package gUtils

import "encoding/json"

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

/**
  JSON (json转map)
*/
func StringToJson(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}

package gUtils

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// 判断结构体字段是否有为空值
func CheckStructEmpty(s interface{}, ignoreFields ...string) error {
	value := reflect.ValueOf(s)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

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

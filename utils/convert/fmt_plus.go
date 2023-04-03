package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func InferType(str string) (interface{}, error) {
	// 尝试将字符串解析为int
	i, err := strconv.Atoi(str)
	if err == nil {
		return i, nil
	}

	// 尝试将字符串解析为float
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return f, nil
	}

	// 尝试将字符串解析为带引号的string
	s, err := strconv.Unquote(str)
	if err == nil {
		return s, nil
	}

	// 如果都不匹配，则返回原始字符串
	return str, nil
}

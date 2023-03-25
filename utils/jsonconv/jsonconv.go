package jsonconv

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

/*
*
ConfigDefault(默认API行为)、
ConfigCompatibleWithStandardLibrary(支持标准库的行为，比如encoding/jjson)、
ConfigFast(通过忽略float类型数据的精度保证最高效的性能)
*/
var jjson = jsoniter.ConfigCompatibleWithStandardLibrary

// 调用 JsonToObject(jsonStr , &obj)
func JsonToObject(jsonStr string, obj any) error {
	err := jjson.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		//log.Println("error:format", "jjson", jsonStr, "obj", obj)
		return err
	}

	return nil
}

// 默认json
func ObjectToJson(data any) string {
	bytes, err := jjson.Marshal(data)
	if err != nil {
		fmt.Println("jjson err-->", err)
		return ""
	}

	return string(bytes)
}

// 转换行结构json
func ObjectToJsonIndent(data any) string {
	bytes, err := jjson.MarshalIndent(data, "", " ")
	if err != nil {
		return ""
	}
	if string(bytes) == "{}" {
		return fmt.Sprintf("%+v", data)
	}
	return string(bytes)
}

// 转下划线json
func ObjectToJsonSnake(data any) string {
	bytes, err := jjson.Marshal(JsonSnakeCase{Value: data})
	if err != nil {
		return ""
	}

	return string(bytes)
}

func SprintPrivateValue(data any) string {
	str := fmt.Sprintf("%+v", data)
	str = strings.ReplaceAll(str, " ", "\n ")
	str = strings.ReplaceAll(str, "{", "\n{\n ")
	str = strings.ReplaceAll(str, "}", "\n}")
	return str
}

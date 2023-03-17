package jsonconv

import jsoniter "github.com/json-iterator/go"

/*
*
ConfigDefault(默认API行为)、
ConfigCompatibleWithStandardLibrary(支持标准库的行为，比如encoding/json)、
ConfigFast(通过忽略float类型数据的精度保证最高效的性能)
*/
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 调用 JsonToObject(jsonStr , &obj)
func JsonToObject(jsonStr string, obj any) error {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		//log.Println("error:format", "json", jsonStr, "obj", obj)
		return err
	}

	return nil
}

// 默认json
func ObjectToJson(data any) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 转换行结构json
func ObjectToJsonIndent(data any) string {
	bytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 转下划线json
func ObjectToJsonSnake(data any) string {
	bytes, err := json.Marshal(JsonSnakeCase{Value: data})
	if err != nil {
		return ""
	}

	return string(bytes)
}

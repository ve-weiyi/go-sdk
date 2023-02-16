package convert

import jsoniter "github.com/json-iterator/go"

// 调用 JsonToObject(jsonStr , &obj)
func JsonToObject(jsonStr string, obj any) error {
	err := jsoniter.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		//log.Println("error:format", "json", jsonStr, "obj", obj)
		return err
	}

	return nil
}

// 默认json
func ObjectToJson(data any) string {
	bytes, err := jsoniter.Marshal(data)
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 转换行结构json
func ObjectToJsonIndent(data any) string {
	bytes, err := jsoniter.MarshalIndent(data, "", " ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 转下划线json
func ObjectToJsonSnake(data any) string {
	bytes, err := jsoniter.Marshal(JsonSnakeCase{Value: data})
	if err != nil {
		return ""
	}

	return string(bytes)
}

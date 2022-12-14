package jsonconv

import (
	"encoding/json"
)

func ObjectToJson(data any) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(bytes)
}

// 调用 JsonToObject(jsonStr , &obj)
func JsonToObject(jsonStr string, obj any) error {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		//log.Println("error:format", "json", jsonStr, "obj", obj)
		return err
	}
	return nil
}

package jsonconv

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"regexp"
	"strings"
)

// https://www.cnblogs.com/chenqionghe/p/13067596.html
/*************************************** 下划线json ***************************************/
type JsonSnakeCase struct {
	Value interface{}
}

// 转换为json，key全部变为驼峰
func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := jjson.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

/*************************************** 驼峰json ***************************************/
type JsonCamelCase struct {
	Value interface{}
}

// 转换为json，key全部变为下划线
func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := jjson.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			if key == "id" {
				return match // 不转换"id"字段
			}
			resKey := Lcfirst(Case2Camel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// 设置json tag为下划线，需要传入指针类型
func SetCamelCaseJsonTag(v interface{}) {
	value := reflect.ValueOf(v)
	//log.Println("--", value.Kind(), value.Elem().Kind())
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		panic("SetCamelCaseJsonTag only accepts a pointer to a struct")
	}
	//如果是指针类型，则取指向的结构体
	//for value.Kind() == reflect.Ptr {
	//	value = value.Elem()
	//}
	t := value.Elem().Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			newJsonTag := Camel2Case(jsonTag)
			field.Tag = reflect.StructTag(strings.Replace(string(field.Tag), jsonTag, newJsonTag, 1))
		}
	}
}

func UnmarshalJSONIgnoreCase(data []byte, obj interface{}) error {
	var tmp map[string]interface{}

	jsoniter.Unmarshal(data, &tmp)
	//如果obj已经是指针，则此处不需要指针
	SetCamelCaseJsonTag(obj)
	jsoniter.UnmarshalFromString(ObjectToJsonSnake(tmp), &obj)
	//
	//var tmp map[string]interface{}
	//// 解析JSON数据到匿名结构体中
	//if err := jjson.Unmarshal(data, &tmp); err != nil {
	//	return err
	//}
	//log.Println("--", tmp)
	//// 使用反射将匿名结构体中的字段值赋值给obj对象
	//v := reflect.ValueOf(obj).Elem()
	//t := v.Type()
	//
	//checkToken := func(fieldName1, fieldName2 string) bool {
	//	//替换 _,转换为小写。再比较
	//	fieldName1 = strings.ReplaceAll(fieldName1, "_", "")
	//	fieldName1 = strings.ToLower(fieldName1)
	//
	//	fieldName2 = strings.ReplaceAll(fieldName2, "_", "")
	//	fieldName2 = strings.ToLower(fieldName2)
	//
	//	if fieldName1 == fieldName2 {
	//		return true
	//	}
	//	return false
	//}
	//
	//for i := 0; i < v.NumField(); i++ {
	//	//目标字段名
	//	fieldName := t.Field(i).Name
	//	fieldType := t.Field(i).Type
	//	for key, value := range tmp {
	//		if checkToken(fieldName, key) {
	//			jsonValue := reflect.ValueOf(value)
	//			jsonType := reflect.TypeOf(value)
	//			log.Println("--", fieldName, value)
	//			log.Println("2--", fieldType.Kind(), jsonType.Kind())
	//
	//			//根据目标type进行赋值
	//			switch fieldType.Kind() {
	//			case reflect.Int:
	//				val := int(jsonValue.Float())
	//				v.Field(i).Set(reflect.ValueOf(val))
	//			case reflect.Bool:
	//			//val := bool(jsonValue.Float() == 1)
	//			//v.Field(i).SetBool(jsonValue)
	//			case reflect.Struct:
	//
	//			default:
	//				v.Field(i).Set(jsonValue)
	//			}
	//
	//			break
	//		}
	//	}
	//}
	return nil
}

/*************************************** 其他方法 ***************************************/

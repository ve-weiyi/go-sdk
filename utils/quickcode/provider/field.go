package provider

import (
	"fmt"
	"strings"
)

//type Field struct {
//	Name             string
//	Type             string
//	ColumnName       string
//	ColumnComment    string
//	MultilineComment bool
//	FieldJSONTag          string
//	GORMTag          string
//	NewTag           string
//	OverwriteTag     string
//	CustomGenType    string
//}

type Field struct {
	FieldName    string `json:"fieldName"` // Field名
	FieldType    string `json:"fieldType"` // Field数据类型
	FieldJSONTag string
	GORMTag      string
	NewTag       string
	OverwriteTag string
	//FieldDefault    string `json:"fieldDefault"`    // 默认值
	//FieldDesc       string `json:"fieldDesc"`       // 中文名
	//FieldJson       string `json:"fieldJson"`       // FieldJson
	//DataType        string `json:"dataType"`        // 数据库字段类型(长度)
	//DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment          string `json:"comment"` // 数据库字段描述
	MultilineComment bool
	//ColumnName      string `json:"columnName"`      // 数据库字段
	//FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	//DictType        string `json:"dictType"`        // 字典
	//Require         bool   `json:"require"`         // 是否必填
	//ErrorText       string `json:"errorText"`       // 校验失败文字
	//Clearable       bool   `json:"clearable"`       // 是否可清空
	//Sort            bool   `json:"sort"`            // 是否增加排序
}

// Tags ...
func (m *Field) Tags() string {
	if m.OverwriteTag != "" {
		return strings.TrimSpace(m.OverwriteTag)
	}

	var tags strings.Builder
	if gormTag := strings.TrimSpace(m.GORMTag); gormTag != "" {
		tags.WriteString(fmt.Sprintf(`gorm:"%s" `, gormTag))
	}
	if jsonTag := strings.TrimSpace(m.FieldJSONTag); jsonTag != "" {
		tags.WriteString(fmt.Sprintf(`json:"%s" `, jsonTag))
	}
	if newTag := strings.TrimSpace(m.NewTag); newTag != "" {
		tags.WriteString(newTag)
	}
	return strings.TrimSpace(tags.String())
}

// GenType ...
func (m *Field) GenType() string {
	typ := strings.TrimLeft(m.FieldType, "*")
	switch typ {
	case "string", "bytes":
		return strings.Title(typ)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return strings.Title(typ)
	case "float64", "float32":
		return strings.Title(typ)
	case "bool":
		return strings.Title(typ)
	case "time.Time":
		return "Time"
	case "json.RawMessage", "[]byte":
		return "Bytes"
	default:
		return "Field"
	}
}

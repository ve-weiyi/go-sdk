package provider

import (
	"bytes"
	"fmt"
	"go/token"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// 需要的数据
type Column struct {
	columnType      gorm.ColumnType
	Indexes         []*Index
	ColumnName      string `json:"columnName" gorm:"column:column_name"`       //列名
	ColumnType      string `json:"columnType" gorm:"column:column_type"`       //字段类型 varchar(11)
	ColumnDefault   string `json:"columnDefault" gorm:"column:column_default"` //默认值
	ColumnComment   string `json:"columnComment" gorm:"column:column_comment"` //备注
	DataType        string `json:"dataType" gorm:"column:data_type"`           //数据类型 varchar
	DataTypeLong    int    `json:"dataTypeLong" gorm:"column:data_type_long"`  //数据长度
	IsNullable      bool   `json:"isNullable" gorm:"column:is_nullable"`       //是否可空
	IsPrimaryKey    bool
	IsUnique        bool
	IsAutoIncrement bool
	dataTypeMap     map[string]func(detailType string) (dataType string) `gorm:"-"`
}

// GetDataType get data type
func (c *Column) GetDataType() (fieldtype string) {
	return dataType.Get(c.DataType, c.ColumnType)
}

func (c *Column) buildGormTag() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("column:%s;type:%s", c.columnType.Name(), c.ColumnType))

	isPriKey, ok := c.columnType.PrimaryKey()
	isValidPriKey := ok && isPriKey
	if isValidPriKey {
		buf.WriteString(";primaryKey")
		if at, ok := c.columnType.AutoIncrement(); ok {
			buf.WriteString(fmt.Sprintf(";autoIncrement:%t", at))
		}
	} else if n, ok := c.columnType.Nullable(); ok && !n {
		buf.WriteString(";not null")
	}

	for _, idx := range c.Indexes {
		if idx == nil {
			continue
		}
		if pk, _ := idx.PrimaryKey(); pk { //ignore PrimaryKey
			continue
		}
		if uniq, _ := idx.Unique(); uniq {
			buf.WriteString(fmt.Sprintf(";uniqueIndex:%s,priority:%d", idx.Name(), idx.Priority))
		} else {
			buf.WriteString(fmt.Sprintf(";index:%s,priority:%d", idx.Name(), idx.Priority))
		}
	}

	if dtValue := c.defaultTagValue(); !isValidPriKey && c.needDefaultTag(dtValue) { // cannot set default tag for primary key
		buf.WriteString(fmt.Sprintf(`;default:%s`, dtValue))
	}
	return buf.String()
}

func (c *Column) ToField(nullable, coverable, signable bool) *Field {
	fieldType := c.GetDataType()
	if signable && strings.Contains(c.ColumnType, "unsigned") && strings.HasPrefix(fieldType, "int") {
		fieldType = "u" + fieldType
	}
	switch {
	case c.ColumnName == "deleted_at" && fieldType == "time.Time":
		fieldType = "gorm.DeletedAt"
	case coverable:
		if c.IsNullable {
			fieldType = "*" + fieldType
		}
	case nullable:
		if c.IsNullable {
			fieldType = "*" + fieldType
		}
	}
	fileName := c.ColumnName
	if token.IsKeyword(fileName) {
		fileName = fileName + "_"
	}

	return &Field{
		FieldName:        fileName,
		FieldType:        fieldType,
		FieldJSONTag:     c.ColumnName,
		GORMTag:          c.buildGormTag(),
		NewTag:           "",
		OverwriteTag:     "",
		Comment:          c.ColumnComment,
		MultilineComment: strings.Contains(c.ColumnComment, "\n"),
	}
}

// needDefaultTag check if default tag needed
func (c *Column) needDefaultTag(defaultTagValue string) bool {
	if defaultTagValue == "" {
		return false
	}
	switch c.columnType.ScanType().Kind() {
	case reflect.Bool:
		return defaultTagValue != "false"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return defaultTagValue != "0"
	case reflect.String:
		return defaultTagValue != ""
	case reflect.Struct:
		return strings.Trim(defaultTagValue, "'0:- ") != ""
	}
	return c.columnType.Name() != "created_at" && c.columnType.Name() != "updated_at"
}

// defaultTagValue return gorm default tag's value
func (c *Column) defaultTagValue() string {
	value, ok := c.columnType.DefaultValue()
	if !ok {
		return ""
	}
	if value != "" && strings.TrimSpace(value) == "" {
		return "'" + value + "'"
	}
	return value
}

//func (c *Column) columnType() (v string) {
//	if cl, ok := c.columnType.ColumnType(); ok {
//		return cl
//	}
//	return c.columnType.DatabaseTypeName()
//}

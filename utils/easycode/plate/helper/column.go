package helper

import (
	"bytes"
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate/provider"
	"go/token"
	"strings"
)

func ColumnToField(c *provider.Column) *Field {
	fieldType := GetDataType(c.DataType, c.ColumnType)

	if strings.Contains(c.ColumnType, "unsigned") && strings.HasPrefix(fieldType, "int") {
		fieldType = "u" + fieldType
	}

	if c.IsNullable {
		fieldType = "*" + fieldType
	}

	fileName := c.ColumnName
	if token.IsKeyword(fileName) {
		fileName = fileName + "_"
	}

	return &Field{
		FieldName:        fileName,
		FieldType:        fieldType,
		FieldJSONTag:     c.ColumnName,
		GORMTag:          buildGormTag(c),
		NewTag:           "",
		OverwriteTag:     "",
		Comment:          c.ColumnComment,
		MultilineComment: strings.Contains(c.ColumnComment, "\n"),
	}
}

func buildGormTag(m *provider.Column) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("column:%s;type:%s", m.ColumnName, m.DataType))

	switch {
	case m.IsPrimaryKey:
		buf.WriteString(";primaryKey")
		if m.IsAutoIncrement {
			buf.WriteString(fmt.Sprintf(";autoIncrement:%t", m.IsAutoIncrement))
		}
	case m.IsNullable:
		buf.WriteString(";not null")
	}

	for _, idx := range m.Indexes {
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

	if m.ColumnDefault != "" { // cannot set default tag for primary key
		buf.WriteString(fmt.Sprintf(`;default:%s`, m.ColumnDefault))
	}
	return buf.String()
}

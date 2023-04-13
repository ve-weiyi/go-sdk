package tmpl

const Model = NotEditMark + `
package {{.Package}}

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

// TableName{{.StructName}} return the table name of <{{.TableName}}>
const TableName{{.StructName}} = "{{.TableName}}"

// {{.StructName}} mapped from table <{{.TableName}}>
type {{.StructName}} struct {
    {{range .Fields}}
	{{if .MultilineComment -}}
	/*
	{{.ColumnComment}}
    */
	{{end -}}
    {{.FieldName}} {{.FieldType}} ` + "`{{.Tags}}` " +
	"{{if not .MultilineComment}}{{if .Comment}}// {{.Comment}}{{end}}{{end}}" +
	`{{end}}
}

`
const Request = `

}
`

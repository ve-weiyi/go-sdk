package easycode

import (
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate/helper"
)

// AutoCodeStructData 初始版本自动化代码工具
type AutoCodeStructData struct {
	FileName           string
	TableName          string          `json:"tableName"`          // 表名 				auto_code
	StructName         string          `json:"structName"`         // Struct名称 		AutoCode 大写驼峰命名
	CamelName          string          `json:"camelName"`          // Struct变量名 	autoCode 小写驼峰命名
	SnakeName          string          `json:"snakeName"`          // go文件名称 		auto_code.go 下划线命名
	ChineseName        string          `json:"chineseName"`        // Struct中文名称 	创建api的描述和注释
	ApiPathPrefix      string          `json:"apiPathName"`        // api路径名			api路径前缀
	AutoCreateApiToSql bool            `json:"autoCreateApiToSql"` // 是否自动创建api
	AutoCreateResource bool            `json:"autoCreateResource"` // 是否自动创建资源标识
	AutoMoveFile       bool            `json:"autoMoveFile"`       // 是否自动移动文件
	BusinessDB         string          `json:"businessDB"`         // 业务数据库
	Fields             []*helper.Field `json:"fields,omitempty"`
	ImportPkgPaths     []string
	DictTypes          []string `json:"-"`
	Package            string   `json:"package"`
	NeedValid          bool     `json:"-"`
	NeedSort           bool     `json:"-"`
	ProjectName        string   `json:"-"`
}

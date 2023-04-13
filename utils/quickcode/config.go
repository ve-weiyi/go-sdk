package quickcode

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/quickcode/provider"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
)

// AutoCodeStruct 初始版本自动化代码工具
type AutoCodeStruct struct {
	FileName           string
	TableName          string            `json:"tableName"`          // 表名 				auto_code
	StructName         string            `json:"structName"`         // Struct名称 		AutoCode 大写驼峰命名
	CamelName          string            `json:"camelName"`          // Struct变量名 	autoCode 小写驼峰命名
	SnakeName          string            `json:"snakeName"`          // go文件名称 		auto_code.go 下划线命名
	ChineseName        string            `json:"chineseName"`        // Struct中文名称 	创建api的描述和注释
	ApiPathPrefix      string            `json:"apiPathName"`        // api路径名			api路径前缀
	AutoCreateApiToSql bool              `json:"autoCreateApiToSql"` // 是否自动创建api
	AutoCreateResource bool              `json:"autoCreateResource"` // 是否自动创建资源标识
	AutoMoveFile       bool              `json:"autoMoveFile"`       // 是否自动移动文件
	BusinessDB         string            `json:"businessDB"`         // 业务数据库
	Fields             []*provider.Field `json:"fields,omitempty"`
	ImportPkgPaths     []string
	DictTypes          []string `json:"-"`
	Package            string   `json:"package"`
	NeedValid          bool     `json:"-"`
	NeedSort           bool     `json:"-"`
	ProjectName        string   `json:"-"`
}

type Config struct {
	db *gorm.DB

	OutPath string // 输出路径
	OutFile string // 输出文件名称

}

func (cfg *Config) Revise() (err error) {
	cfg.OutPath, err = filepath.Abs(cfg.OutPath)
	if err != nil {
		return fmt.Errorf("outpath is invalid: %w", err)
	}
	if cfg.OutPath == "" {
		cfg.OutPath = "./query/"
	}
	if cfg.OutFile == "" {
		cfg.OutFile = cfg.OutPath + "/gen.go"
	} else if !strings.Contains(cfg.OutFile, "/") {
		cfg.OutFile = cfg.OutPath + "/" + cfg.OutFile
	}

	return nil
}

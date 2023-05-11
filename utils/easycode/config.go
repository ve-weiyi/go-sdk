package easycode

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate/tmpl"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
)

type Config struct {
	db *gorm.DB

	OutPath string // 输出路径
	OutFile string // 输出文件名称

	FieldNullable     bool // generate pointer when field is nullable
	FieldCoverable    bool // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
	FieldSignable     bool // detect integer field's unsigned type, adjust generated data type
	FieldWithIndexTag bool // generate with gorm index tag
	FieldWithTypeTag  bool // generate with gorm column type tag

	dataTypeMap    map[string]func(columnType gorm.ColumnType) (dataType string)
	fieldJSONTagNS func(columnName string) (tagContent string)
}

// WithDataTypeMap specify data type mapping relationship, only work when syncing table from db
func (cfg *Config) WithDataTypeMap(newMap map[string]func(columnType gorm.ColumnType) (dataType string)) {
	cfg.dataTypeMap = newMap
}

// WithJSONTagNameStrategy specify json tag naming strategy
func (cfg *Config) WithJSONTagNameStrategy(ns func(columnName string) (tagContent string)) {
	cfg.fieldJSONTagNS = ns
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

// ConvertStructs convert to base structures
func (cfg *Config) ConvertStructMetas(structs ...interface{}) (metas []*plate.PlateMeta, err error) {
	for _, st := range structs {
		if st == nil {
			continue
		}
		if base, ok := st.(*plate.PlateMeta); ok {
			metas = append(metas, base)
			continue
		}

		meta := &plate.PlateMeta{
			Key:            "",
			AutoCodePath:   cfg.OutPath,
			TemplateString: tmpl.Model,
			Data:           nil,
		}
		metas = append(metas, meta)
	}
	return
}

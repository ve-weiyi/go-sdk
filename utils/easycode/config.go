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
			Key:              "",
			OutFileName:      cfg.OutFile,
			AutoCodePath:     cfg.OutPath,
			AutoCodeMovePath: cfg.OutFile,
			TemplateString:   tmpl.Model,
			TemplatePath:     "",
			Data:             nil,
		}
		metas = append(metas, meta)
	}
	return
}

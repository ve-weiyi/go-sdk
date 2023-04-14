package easycode

import (
	"fmt"
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

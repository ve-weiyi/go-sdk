package quickcode

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/quickcode/provider"
	"gorm.io/gorm"
)

type Generator struct {
	Config
}

// UseDB set db connection
func (g *Generator) UseDB(db *gorm.DB) {
	if db != nil {
		g.db = db
	}
}

func NewGenerator(cfg Config) *Generator {
	if err := cfg.Revise(); err != nil {
		panic(fmt.Errorf("create generator fail: %w", err))
	}
	return &Generator{
		Config: cfg,
	}
}

func CreateModelFile(db *gorm.DB, dbName string, tbName string) error {
	mysqlDriver := provider.MysqlDriver{DB: db}
	columns, err := mysqlDriver.GetTableColumns(dbName, tbName)
	if err != nil {
		return nil
	}

	var fields []*provider.Field
	for _, column := range columns {
		field := column.ToField(true, true, true)
		fields = append(fields, field)
	}

	return nil
}

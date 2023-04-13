package provider

import (
	"gorm.io/gorm"
)

type Database interface {
	GetDB() (data []Db, err error)
	GetTables(dbName string) (data []Table, err error)
	GetTableColumn(dbName string, tableName string) (data []Column, err error)
}

type Db struct {
	Database string `json:"database" gorm:"column:database"`
}

type Table struct {
	TableName string `json:"tableName" gorm:"column:table_name"`
}

type Index struct {
	gorm.Index
	Priority int32 `gorm:"column:SEQ_IN_INDEX"`
}

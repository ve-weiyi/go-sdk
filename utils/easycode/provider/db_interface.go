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

// 需要的数据
type Column struct {
	gorm.ColumnType
	Indexes    []*Index
	ColumnName string `json:"columnName" gorm:"column:column_name"` //列名
	//ColumnType      string `json:"columnType" gorm:"column:column_type"`       //字段类型 varchar(11)
	ColumnDefault   string `json:"columnDefault" gorm:"column:column_default"` //默认值
	ColumnComment   string `json:"columnComment" gorm:"column:column_comment"` //备注
	DataType        string `json:"dataType" gorm:"column:data_type"`           //数据类型 varchar
	DataTypeLong    int64  `json:"dataTypeLong" gorm:"column:data_type_long"`  //数据长度
	IsNullable      bool   `json:"isNullable" gorm:"column:is_nullable"`       //是否可空
	IsPrimaryKey    bool
	IsUnique        bool
	IsAutoIncrement bool
	dataTypeMap     map[string]func(detailType string) (dataType string) `gorm:"-"`
}

type Index struct {
	gorm.Index
	Priority int32 `gorm:"column:SEQ_IN_INDEX"`
}

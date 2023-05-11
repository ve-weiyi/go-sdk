package provider

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
)

type MysqlDriver struct {
	*gorm.DB
}

// GetDB 获取数据库的所有数据库名
func (m *MysqlDriver) GetDB() (data []Db, err error) {
	var entities []Db
	sql := "SELECT SCHEMA_NAME FROM information_schema.schemata;"
	err = m.DB.Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables 获取数据库的所有表名
func (m *MysqlDriver) GetTables(dbName string) (data []Table, err error) {
	var entities []Table
	sql := `select * from information_schema.tables where table_schema = ?`
	err = m.DB.Raw(sql, dbName).Scan(&entities).Error

	return entities, err
}

// GetTableColumns  struct
func (m *MysqlDriver) GetTableColumns(dbName string, tableName string) (data []Column, err error) {
	var entities []Column
	//var metas []ColumnMetadata
	//sql := `SELECT * FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_schema = ? AND table_name = ?`
	//err = m.DB.Raw(sql,dbName, tableName).Scan(&metas).Error
	var mapType map[string]gorm.ColumnType
	var mapIndex map[string][]*Index

	types, err := m.Migrator().ColumnTypes(tableName)
	if err != nil {
		return nil, err
	}
	mapType = make(map[string]gorm.ColumnType, 0)
	for _, item := range types {
		mapType[item.Name()] = item
	}

	indexes, err := m.Migrator().GetIndexes(tableName)
	js, _ := json.MarshalIndent(&indexes, "", " ")
	log.Println("111--->", string(js))

	if err != nil {
		return nil, err
	}
	mapIndex = GroupByColumn(indexes)
	for _, entity := range types {
		col := Column{
			ColumnType: entity,
			ColumnName: entity.Name(),
		}

		col.Indexes = mapIndex[entity.Name()]
		col.ColumnFiledType, _ = entity.ColumnType()
		col.ColumnDefault, col.HasDefault = entity.DefaultValue()
		col.ColumnComment, _ = entity.Comment()
		col.DataType = entity.DatabaseTypeName()
		col.DataTypeLong, _ = entity.Length()
		col.IsNullable, _ = entity.Nullable()
		col.IsPrimaryKey, _ = entity.PrimaryKey()
		col.IsUnique, _ = entity.Unique()
		col.IsAutoIncrement, _ = entity.AutoIncrement()
		entities = append(entities, col)
	}

	return entities, nil
}

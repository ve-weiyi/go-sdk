package provider

import (
	"gorm.io/gorm"
)

type MysqlDriver struct {
	*gorm.DB
}

// GetDB 获取数据库的所有数据库名
func (m *MysqlDriver) GetDB() (data []Db, err error) {
	var entities []Db
	sql := "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA;"
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
	if err != nil {
		return nil, err
	}
	mapIndex = GroupByColumn(indexes)

	for _, entity := range types {
		col := Column{
			columnType: entity,
			Indexes:    mapIndex[entity.Name()],
			ColumnName: entity.Name(),
		}

		col.ColumnType, _ = entity.ColumnType()
		col.ColumnDefault, _ = entity.DefaultValue()
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

// GroupByColumn group columns
func GroupByColumn(indexList []gorm.Index) map[string][]*Index {
	columnIndexMap := make(map[string][]*Index, len(indexList))
	if len(indexList) == 0 {
		return columnIndexMap
	}

	for _, idx := range indexList {
		if idx == nil {
			continue
		}
		for i, col := range idx.Columns() {
			columnIndexMap[col] = append(columnIndexMap[col], &Index{
				Index:    idx,
				Priority: int32(i + 1),
			})
		}
	}
	return columnIndexMap
}

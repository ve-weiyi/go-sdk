package provider

import (
	"gorm.io/gorm"
)

type MysqlDriver struct {
	*gorm.DB
}

// GetDB 获取数据库的所有数据库名
func (s *MysqlDriver) GetDB() (data []Db, err error) {
	var entities []Db
	sql := "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;"
	err = s.DB.Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables 获取数据库的所有表名
func (s *MysqlDriver) GetTables(dbName string) (data []Table, err error) {
	var entities []Table
	sql := `select table_name as table_name from information_schema.tables where table_schema = ?`

	err = s.DB.Raw(sql, dbName).Scan(&entities).Error

	return entities, err
}

// GetTableColumns  struct
func (t *MysqlDriver) GetTableColumns(dbName string, tableName string) (data []Column, err error) {
	var entities []Column
	var metas []ColumnMetadata
	var mapType map[string]gorm.ColumnType
	var mapIndex map[string][]*Index
	types, err := t.Migrator().ColumnTypes(tableName)
	if err != nil {
		return nil, err
	}
	mapType = make(map[string]gorm.ColumnType, 0)
	for _, item := range types {
		mapType[item.Name()] = item
	}

	indexes, err := t.Migrator().GetIndexes(tableName)
	if err != nil {
		return nil, err
	}
	mapIndex = GroupByColumn(indexes)

	sql := `SELECT * FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?`
	err = t.DB.Raw(sql, tableName, dbName).Scan(&metas).Error

	for _, entity := range metas {
		var dataTypeLong int
		switch entity.DataType {
		case "longtext":
			dataTypeLong = entity.CharacterMaximumLength
		case "varchar":
			dataTypeLong = entity.CharacterMaximumLength
		case "double":
			dataTypeLong = entity.NumericPrecision
		case "decimal":
			dataTypeLong = entity.NumericPrecision
		case "tinyint":
			dataTypeLong = entity.NumericPrecision
		case "int":
			dataTypeLong = entity.NumericPrecision
		default:
			dataTypeLong = entity.NumericPrecision + entity.CharacterMaximumLength
		}
		//log.Println(jsonconv.ObjectToJsonIndent(entity))
		col := Column{
			columnType:      mapType[entity.ColumnName],
			Indexes:         mapIndex[entity.ColumnName],
			ColumnName:      entity.ColumnName,
			ColumnType:      entity.ColumnType,
			ColumnDefault:   entity.ColumnDefault,
			ColumnComment:   entity.ColumnComment,
			DataType:        entity.DataType,
			DataTypeLong:    dataTypeLong,
			IsNullable:      entity.IsNullable == "YES",
			IsPrimaryKey:    entity.ColumnKey == "PRI",
			IsUnique:        entity.ColumnKey == "UNI",
			IsAutoIncrement: entity.Extra == "auto_increment",
			dataTypeMap:     nil,
		}
		entities = append(entities, col)
	}

	return entities, nil
}

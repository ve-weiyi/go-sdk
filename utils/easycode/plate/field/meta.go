package field

import (
	"gorm.io/gorm"
)

// 数据库表数据
type ColumnMetadata struct {
	TableCatalog           string `gorm:"column:TABLE_CATALOG" json:"tableCatalog"`                      // 列所属的数据库名称。
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`                        // 列所属的模式名称。
	TableName              string `gorm:"column:TABLE_NAME" json:"tableName"`                            // 列所属的表名称。
	OrdinalPosition        int    `gorm:"column:ORDINAL_POSITION" json:"ordinalPosition"`                // 列在表中的位置。
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"columnName"`                          // 列的名称。
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"columnType"`                          // 列的类型和长度。
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`                    // 列的默认值。
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"columnComment"`                    // 列的注释信息。
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"columnKey"`                            // 列是否为主键或唯一键的一部分。
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"isNullable"`                          // 列是否允许为空。
	DataType               string `gorm:"column:DATA_TYPE" json:"dataType"`                              // 列的数据类型。
	CharacterMaximumLength int    `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"` // 字符类型列的最大长度。
	NumericPrecision       int    `gorm:"column:NUMERIC_PRECISION" json:"numericPrecision"`              // 数值类型列的精度。
	NumericScale           int    `gorm:"column:NUMERIC_SCALE" json:"numericScale"`                      // 数值类型列的小数位数。
	DatetimePrecision      int    `gorm:"column:DATETIME_PRECISION" json:"datetimePrecision"`            // 日期时间类型列的精度。
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"characterSetName"`             // 字符类型列的字符集名称。
	CollationName          string `gorm:"column:COLLATION_NAME" json:"collationName"`                    // 字符类型列的排序规则名称。
	Extra                  string `gorm:"column:EXTRA" json:"extra"`                                     // 列是否具有附加属性，如自动递增。
	Privileges             string `gorm:"column:PRIVILEGES" json:"privileges"`                           // 与列相关的权限信息。
}

// GetTableColumns  struct
func GetTableColumns(t *gorm.DB, schemaName string, tableName string) (result []gorm.ColumnType, err error) {
	types, err := t.Migrator().ColumnTypes(tableName)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// GetTableIndex  index
func GetTableIndex(t *gorm.DB, schemaName string, tableName string) (indexes []gorm.Index, err error) {
	return t.Migrator().GetIndexes(tableName)
}

func GetColumnMeta(t *gorm.DB, dbName string, tableName string) (data []ColumnMetadata, err error) {
	var metas []ColumnMetadata
	sql := `
	SELECT *
	FROM INFORMATION_SCHEMA.COLUMNS c
	WHERE table_name = ?
	  AND table_schema = ?
	`
	err = t.Raw(sql, tableName, dbName).Scan(&metas).Error

	return metas, err
}

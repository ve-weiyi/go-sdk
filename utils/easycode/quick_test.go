package easycode

import (
	"encoding/json"
	"fmt"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"path"
	"testing"
)

// GEN 自动生成 GORM 模型结构体文件及使用示例 https://blog.csdn.net/Jeffid/article/details/126898000
const dsn = "root:mysql7914@(127.0.0.1:3306)/blog-plus?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	var err error
	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "tb_",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	log.Println("mysql connection done")
}

func TestIndex(t *testing.T) {
	tableName := "user_account"
	index, _ := db.Migrator().GetIndexes(tableName)
	js, _ := json.MarshalIndent(&index, "", " ")
	log.Println("111--->", string(js))
}

func TestPlate(t *testing.T) {
	out := path.Join(global.GetRuntimeRoot(), "server/api", "blog")
	//out := path.Join("./autocode_template", "test")

	cfg := Config{
		db:      nil,
		OutPath: out,
		OutFile: path.Join("./autocode_template", "test"),
	}
	typeInt := "int"
	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return typeInt },
	}
	cfg.WithDataTypeMap(dataMap)
	cfg.WithJSONTagNameStrategy(func(columnName string) (tagContent string) {
		//toStringField := "time"
		//if strings.Contains(columnName, toStringField) {
		// return columnName + "\" example:\"2022-11-16T16:00:00.000Z"
		//}
		return columnName
	})

	gen := NewGenerator(cfg)
	gen.UseDB(db)
	gen.ApplyMetas(gen.GenerateMetasFromTable("blog-v2", "user_account", "账号"))
	//tableName := "auth"
	//data := plate.AutoCodeStructData{
	//	Package:        "blog",
	//	TableName:      tableName,
	//	StructName:     jsonconv.Case2Camel(tableName),
	//	ValueName:      jsonconv.Case2CamelNotFirst(tableName),
	//	JsonName:       jsonconv.Camel2Case(tableName),
	//	ChineseName:    "权限认证",
	//	Fields:         nil,
	//	ImportPkgPaths: nil,
	//}
	//gen.ApplyMetas(gen.GenerateMetasFromModel(data.Reverse()))
	//gen.RollBack()
	gen.Execute()
}

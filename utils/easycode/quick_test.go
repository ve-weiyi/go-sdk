package easycode

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
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

//
//func TestName(t *testing.T) {
//	gen := NewGenerator(Config{
//		db:      nil,
//		OutPath: path.Join("./autocode_template", "test"),
//		OutFile: path.Join("./autocode_template", "test"),
//	})
//	gen.UseDB(db)
//	metas := gen.GenerateMetasFromTable("blog-plus", "article")
//	gen.ApplyBasic(metas)
//	gen.Execute()
//}
//
//func getFields() []*helper.Field {
//	mysqlDriver := provider.MysqlDriver{DB: db}
//	columns, err := mysqlDriver.GetTableColumns("blog-plus", "article")
//	if err != nil {
//		return nil
//	}
//
//	var fields []*helper.Field
//	for _, column := range columns {
//		log.Println(jsonconv.ObjectToJsonIndent(column))
//		field := column.ToField(true, true, true)
//		fields = append(fields, field)
//	}
//
//	for _, item := range fields {
//		log.Println(jsonconv.ObjectToJsonIndent(item))
//	}
//
//	return fields
//}

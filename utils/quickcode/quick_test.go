package quickcode

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/go-sdk/utils/quickcode/inject"
	"github.com/ve-weiyi/go-sdk/utils/quickcode/provider"
	"github.com/ve-weiyi/go-sdk/utils/quickcode/tmpl"
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

func TestName(t *testing.T) {

}

func getFields() []*provider.Field {
	mysqlDriver := provider.MysqlDriver{DB: db}
	columns, err := mysqlDriver.GetTableColumns("blog-plus", "article")
	if err != nil {
		return nil
	}

	var fields []*provider.Field
	for _, column := range columns {
		log.Println(jsonconv.ObjectToJsonIndent(column))
		field := column.ToField(true, true, true)
		fields = append(fields, field)
	}

	for _, item := range fields {
		log.Println(jsonconv.ObjectToJsonIndent(item))
	}

	return fields
}

func getTempMetas() []*logic.TempMeta {
	temporaryRoot := path.Join("./autocode_template", "test")
	autoMoveRoot := path.Join("./api")
	pkg := "test"
	fileName := "article"

	tplList := []*logic.TempMeta{
		/** server start */
		&logic.TempMeta{
			TemplateString:   tmpl.Model,
			AutoPackage:      pkg,
			AutoCodePath:     fmt.Sprintf("%v/server/model/%s.go", temporaryRoot, fileName),
			AutoMoveFilePath: fmt.Sprintf("%v/model/entity/%s.go", autoMoveRoot, fileName),
		},
		&logic.TempMeta{
			TemplateString:   tmpl.Request,
			AutoPackage:      pkg,
			AutoCodePath:     fmt.Sprintf("%v/server/request/%s.go", temporaryRoot, fileName),
			AutoMoveFilePath: fmt.Sprintf("%v/model/request/%s.go", autoMoveRoot, fileName),
		},
		//&logic.TempMeta{
		//	TemplateString:   tmpl.Service,
		//	AutoPackage:      pkg,
		//	AutoCodePath:     fmt.Sprintf("%v/service/logic/%s.go", temporaryRoot, fileName),
		//	AutoMoveFilePath: fmt.Sprintf("%v/service/logic/%s.go", autoMoveRoot, fileName),
		//},
		//&logic.TempMeta{
		//	TemplateString:   tmpl.Controller,
		//	AutoPackage:      pkg,
		//	AutoCodePath:     fmt.Sprintf("%v/controller/logic/%s.go", temporaryRoot, fileName),
		//	AutoMoveFilePath: fmt.Sprintf("%v/controller/logic/%s.go", autoMoveRoot, fileName),
		//},
		//&logic.TempMeta{
		//	TemplateString:   tmpl.Router,
		//	AutoPackage:      pkg,
		//	AutoCodePath:     fmt.Sprintf("%v/router/logic/%s.go", temporaryRoot, fileName),
		//	AutoMoveFilePath: fmt.Sprintf("%v/router/logic/%s.go", autoMoveRoot, fileName),
		//},
	}
	return tplList
}

func TestGenerateTemp(t *testing.T) {
	fields := getFields()
	tplList := getTempMetas()

	camelName := "article"
	snakeName := "article"
	structName := "Article"
	data := AutoCodeStruct{
		FileName:           snakeName,
		TableName:          snakeName,
		StructName:         structName,
		CamelName:          camelName,
		SnakeName:          snakeName,
		ChineseName:        "文章",
		ApiPathPrefix:      snakeName,
		AutoCreateApiToSql: false,
		AutoCreateResource: false,
		AutoMoveFile:       false,
		BusinessDB:         "",
		Fields:             fields,
		ImportPkgPaths:     []string{""},
		DictTypes:          nil,
		Package:            "test",
		NeedValid:          false,
		NeedSort:           false,
		ProjectName:        "",
	}

	for _, item := range tplList {
		err := item.CreateTempFile(data)
		if err != nil {
			log.Println("err-->", err)
		}
	}
}

//func (g *Generator) fillModelPkgPath(filePath string) {
//	pkgs, err := packages.Load(&packages.Config{
//		Mode: packages.NeedName,
//		Dir:  filePath,
//	})
//	if err != nil {
//		g.db.Logger.Warn(context.Background(), "parse model pkg path fail: %s", err)
//		return
//	}
//	if len(pkgs) == 0 {
//		g.db.Logger.Warn(context.Background(), "parse model pkg path fail: got 0 packages")
//		return
//	}
//	g.Config.modelPkgPath = pkgs[0].PkgPath
//}

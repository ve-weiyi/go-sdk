package easycode

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate"
	"gorm.io/gorm"
	"log"
	"os"
)

type Generator struct {
	Config
	*log.Logger

	plateMetas  map[string]plate.PlateMeta
	InjectMetas interface{}
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
		Logger: log.New(os.Stderr, "", log.LstdFlags|log.Llongfile),
	}
}

// Execute generate code to output path
func (g *Generator) Execute() {
	g.Println("Start generating code.")

	//if err := g.generateModelFile(); err != nil {
	//	g.Printf("generate model struct fail: %s", err)
	//	panic("generate model struct fail")
	//}

	g.Println("Generate code done.")
}

// ApplyBasic specify models which will implement basic .diy_method
func (g *Generator) ApplyBasic(models ...interface{}) {
	//g.ApplyInterface(func() {}, models...)
}

// ApplyInterface specifies .diy_method interfaces on structures, implment codes will be generated after calling g.Execute()
// eg: g.ApplyInterface(func(model.Method){}, model.User{}, model.Company{})
//func (g *Generator) ApplyInterface(fc interface{}, models ...interface{}) {
//	structs, err := generate.ConvertStructMetas(models...)
//	if err != nil {
//		g.Printf("check struct fail: %v", err)
//		panic("check struct fail")
//	}
//	g.apply(fc, structs...)
//}

//func (g *Generator) apply(fc interface{}, structs ...*generate.IMeta) {
//
//}

//func (g *Generator) GenerateMetasFromModel(dbName, tableName string) []*generate.PlateMeta {
//
//}
//
//func (g *Generator) GenerateMetasFromTable(dbName, tableName string) []*generate.PlateMeta {
//	mysqlDriver := provider.MysqlDriver{DB: g.db}
//	columns, err := mysqlDriver.GetTableColumns(dbName, tableName)
//	if err != nil {
//		return nil
//	}
//
//	var fields []*helper.Field
//	for _, column := range columns {
//		field := column.ToField(true, true, true)
//		fields = append(fields, field)
//	}
//
//	camelName := "article"
//	snakeName := "article"
//	structName := "Article"
//	data := &generate.AutoCodeStructData{
//		FileName:           snakeName,
//		TableName:          snakeName,
//		StructName:         structName,
//		CamelName:          camelName,
//		SnakeName:          snakeName,
//		ChineseName:        "文章",
//		ApiPathPrefix:      snakeName,
//		AutoCreateApiToSql: false,
//		AutoCreateResource: false,
//		AutoMoveFile:       false,
//		BusinessDB:         "",
//		Fields:             fields,
//		ImportPkgPaths:     []string{""},
//		DictTypes:          nil,
//		Package:            "test",
//		NeedValid:          false,
//		NeedSort:           false,
//		ProjectName:        "",
//	}
//
//	temporaryRoot := g.OutPath
//	autoMoveRoot := g.OutPath
//	fileName := tableName
//	metas := []*generate.PlateMeta{
//		/** server start */
//		&generate.PlateMeta{
//			TemplateString:   tmpl.Model,
//			AutoCodePath:     fmt.Sprintf("%v/server/model/%s.go", temporaryRoot, fileName),
//			AutoCodeMovePath: fmt.Sprintf("%v/model/entity/%s.go", autoMoveRoot, fileName),
//			Data:             data,
//		},
//		//&logic.PlateMeta{
//		//	TemplateString:   tmpl.Request,
//		//	AutoCodePath:     fmt.Sprintf("%v/server/request/%s.go", temporaryRoot, fileName),
//		//	AutoCodeMovePath: fmt.Sprintf("%v/model/request/%s.go", autoMoveRoot, fileName),
//		//	Data:             data,
//		//},
//		//&logic.PlateMeta{
//		//	TemplateString:   tmpl.Service,
//		//	AutoCodePath:     fmt.Sprintf("%v/service/logic/%s.go", temporaryRoot, fileName),
//		//	AutoCodeMovePath: fmt.Sprintf("%v/service/logic/%s.go", autoMoveRoot, fileName),
//		//	Data:             data,
//		//},
//		//&logic.PlateMeta{
//		//	TemplateString:   tmpl.Controller,
//		//	AutoCodePath:     fmt.Sprintf("%v/controller/logic/%s.go", temporaryRoot, fileName),
//		//	AutoCodeMovePath: fmt.Sprintf("%v/controller/logic/%s.go", autoMoveRoot, fileName),
//		//	Data:             data,
//		//},
//		//&logic.PlateMeta{
//		//	TemplateString:   tmpl.Router,
//		//	AutoCodePath:     fmt.Sprintf("%v/router/logic/%s.go", temporaryRoot, fileName),
//		//	AutoCodeMovePath: fmt.Sprintf("%v/router/logic/%s.go", autoMoveRoot, fileName),
//		//	Data:             data,
//		//},
//	}
//
//	return metas
//}
//
//func (g *Generator) generateModelFile() error {
//
//	for _, item := range g.metas {
//		err := item.CreateTempFile()
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

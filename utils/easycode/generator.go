package easycode

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/easycode/inject"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate/field"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate/provider"
	"github.com/ve-weiyi/go-sdk/utils/easycode/plate/tmpl"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"gorm.io/gorm"
	"log"
	"os"
)

type Generator struct {
	cfg Config
	*log.Logger

	plateMetas  []*plate.PlateMeta
	InjectMetas []*inject.AstInjectMeta
}

// UseDB set db connection
func (g *Generator) UseDB(db *gorm.DB) {
	if db != nil {
		g.cfg.db = db
	}
}

func NewGenerator(cfg Config) *Generator {
	if err := cfg.Revise(); err != nil {
		panic(fmt.Errorf("create generator fail: %w", err))
	}
	return &Generator{
		cfg:    cfg,
		Logger: log.New(os.Stderr, "", log.LstdFlags|log.Llongfile),
	}
}

// Execute generate code to output path
func (g *Generator) Execute() {
	g.Println("Start generating code.")

	if err := g.generateModelFile(); err != nil {
		g.Printf("generate model struct fail: %s", err)
		panic("generate model struct fail")
	}

	g.Println("Generate code done.")
}
func (g *Generator) RollBack() {
	g.Println("Start rollback code.")

	if err := g.rollback(); err != nil {
		g.Printf("rollback model struct fail: %s", err)
		panic("rollback model struct fail")
	}

	g.Println("RollBack code done.")
}

func (g *Generator) GenFieldConfig() *field.FieldConfig {
	return &field.FieldConfig{
		DataTypeMap: g.cfg.dataTypeMap,

		FieldSignable:     g.cfg.FieldSignable,
		FieldNullable:     g.cfg.FieldNullable,
		FieldCoverable:    g.cfg.FieldCoverable,
		FieldWithIndexTag: g.cfg.FieldWithIndexTag,
		FieldWithTypeTag:  g.cfg.FieldWithTypeTag,

		FieldJSONTagNS: g.cfg.fieldJSONTagNS,
	}
}

func (g *Generator) ApplyMetas(plates []*plate.PlateMeta, injects []*inject.AstInjectMeta) {
	g.plateMetas = append(g.plateMetas, plates...)

	g.InjectMetas = append(g.InjectMetas, injects...)
}

// 创建数据库中所有表
func (g *Generator) GenerateMetasFromSchema(dbName string) ([]*plate.PlateMeta, []*inject.AstInjectMeta) {
	mysqlDriver := provider.MysqlDriver{DB: g.cfg.db}
	tbs, err := mysqlDriver.GetTables(dbName)
	if err != nil {
		return nil, nil
	}

	var pl []*plate.PlateMeta
	var in []*inject.AstInjectMeta
	for _, tb := range tbs {
		plates, injects := g.GenerateMetasFromTable(dbName, tb.TableName, tb.TableComment)
		pl = append(pl, plates...)
		in = append(in, injects...)
	}
	return pl, in
}

// 创建一个表
func (g *Generator) GenerateMetasFromTable(dbName, tableName, tableComment string) ([]*plate.PlateMeta, []*inject.AstInjectMeta) {
	mysqlDriver := provider.MysqlDriver{DB: g.cfg.db}

	columns, err := mysqlDriver.GetTableColumns(dbName, tableName)
	if err != nil {
		return nil, nil
	}

	cfg := g.GenFieldConfig()
	var fields []*provider.Field
	for _, column := range columns {
		field := column.ToField(cfg)
		fields = append(fields, field)
	}

	if tableComment == "" {
		tableComment = tableName
	}
	data := &plate.AutoCodeStructData{
		Package:        jsonconv.Case2CamelNotFirst(dbName),
		TableName:      tableName,
		StructName:     jsonconv.Case2Camel(tableName),
		ValueName:      jsonconv.Case2CamelNotFirst(tableName),
		JsonName:       jsonconv.Camel2Case(tableName),
		ChineseName:    tableComment,
		Fields:         fields,
		ImportPkgPaths: []string{
			//"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc",
			//"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity",
			//"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request",
			//"github.com/ve-weiyi/ve-admin-store/server/api/common/controller",
			//"github.com/ve-weiyi/ve-admin-store/server/api/common/model/response",
		},
	}

	return g.GenerateMetasFromModel(data)
}

// 使用data进行创建
func (g *Generator) GenerateMetasFromModel(data *plate.AutoCodeStructData) ([]*plate.PlateMeta, []*inject.AstInjectMeta) {

	temporaryRoot := g.cfg.OutPath
	fileName := data.TableName

	metaModel := &plate.PlateMeta{
		TemplateString: tmpl.Model,
		AutoCodePath:   fmt.Sprintf("%v/model/entity/%s.go", temporaryRoot, fileName),
		Data:           data,
	}

	metaService := &plate.PlateMeta{
		TemplateString: tmpl.Service,
		AutoCodePath:   fmt.Sprintf("%v/service/logic/%s.sv.go", temporaryRoot, fileName),
		Data:           data,
	}
	metaController := &plate.PlateMeta{
		TemplateString: tmpl.Controller,
		AutoCodePath:   fmt.Sprintf("%v/controller/logic/%s.ctl.go", temporaryRoot, fileName),
		Data:           data,
	}
	metaRouter := &plate.PlateMeta{
		TemplateString: tmpl.Router,
		AutoCodePath:   fmt.Sprintf("%v/router/logic/%s.rt.go", temporaryRoot, fileName),
		Data:           data,
	}
	metas := []*plate.PlateMeta{
		/** server start */
		metaModel,
		metaService,
		metaController,
		metaRouter,
	}

	var injectMetas []*inject.AstInjectMeta

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		FilePath: fmt.Sprintf("%v/controller/controller.go", temporaryRoot),
		StructMetas: []*inject.StructMeta{
			inject.NewStructMete("AppController", fmt.Sprintf(`%vController *logic.%vController //%v`, data.StructName, data.StructName, data.ChineseName)),
		},
		FuncMetas: []*inject.FuncMeta{
			inject.NewFuncMete("NewController", fmt.Sprintf(`return &AppController{
			%vController: logic.New%vController(svcCtx),
			}`, data.StructName, data.StructName)),
		},
	})

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		FilePath: fmt.Sprintf("%v/router/router.go", temporaryRoot),
		StructMetas: []*inject.StructMeta{
			inject.NewStructMete("AppRouter", fmt.Sprintf(`%vRouter *logic.%vRouter //%v`, data.StructName, data.StructName, data.ChineseName)),
		},
		FuncMetas: []*inject.FuncMeta{
			inject.NewFuncMete("NewRouter", fmt.Sprintf(`return &AppRouter{
			%vRouter: logic.New%vRouter(svcCtx),
			}`, data.StructName, data.StructName)),
		},
	})

	injectMetas = append(injectMetas, &inject.AstInjectMeta{
		FilePath: fmt.Sprintf("%v/service/service.go", temporaryRoot),
		StructMetas: []*inject.StructMeta{
			inject.NewStructMete("AppService", fmt.Sprintf(`%vService *logic.%vService //%v`, data.StructName, data.StructName, data.ChineseName)),
		},
		FuncMetas: []*inject.FuncMeta{
			inject.NewFuncMete("NewService", fmt.Sprintf(`return &AppService{
			%vService: logic.New%vService(svcCtx),
			}`, data.StructName, data.StructName)),
		},
	})
	return metas, injectMetas
}

func (g *Generator) generateModelFile() error {

	for _, item := range g.plateMetas {
		err := item.CreateTempFile()
		if err != nil {
			return err
		}
	}

	for _, item := range g.InjectMetas {
		err := item.Inject()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) rollback() error {
	for _, item := range g.plateMetas {
		err := item.RollBack()
		if err != nil {
			return err
		}
	}

	for _, item := range g.InjectMetas {
		err := item.RollBack()
		if err != nil {
			return err
		}
	}
	return nil
}

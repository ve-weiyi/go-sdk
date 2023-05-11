package tmpl

const AppService = `
package service

import (
	"{{.SvcPackage }}"
)

type AppService struct {
	svcCtx *svc.ServiceContext //持有的repository层引用
}

func NewService(svcCtx *svc.ServiceContext) *AppService {
	return &AppService{
		svcCtx: svcCtx,
	}
}
`

const Service = `
package logic

import (
	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

type {{.StructName}}Service struct {
	svcCtx *svc.ServiceContext
}

func New{{.StructName}}Service(svcCtx *svc.ServiceContext) *{{.StructName}}Service {
	return &{{.StructName}}Service{
		svcCtx: svcCtx,
	}
}

// 创建{{.StructName}}记录
func (m *{{.StructName}}Service) Create{{.StructName}}(ctx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (data *entity.{{.StructName}}, err error) {
	db:= m.svcCtx.MainDB
	err = db.Create(&{{.ValueName}}).Error
	if err != nil {
		return nil, err
	}
	return {{.ValueName}}, err
}

// 删除{{.StructName}}记录
func (m *{{.StructName}}Service) Delete{{.StructName}}(ctx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (rows int64, err error) {
	db:= m.svcCtx.MainDB	
	query := db.Delete(&{{.ValueName}})
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新{{.StructName}}记录
func (m *{{.StructName}}Service) Update{{.StructName}}(ctx *request.Context, {{.ValueName}} *entity.{{.StructName}}) (data *entity.{{.StructName}}, err error) {
	db:= m.svcCtx.MainDB	
	err = db.Save(&{{.ValueName}}).Error
	if err != nil {
		return nil, err
	}
	return {{.ValueName}}, err
}

// 根据id获取{{.StructName}}记录
func (m *{{.StructName}}Service) Find{{.StructName}}(ctx *request.Context, id int) (data *entity.{{.StructName}}, err error) {
	db:= m.svcCtx.MainDB	
	err = db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

// 批量删除{{.StructName}}记录
func (m *{{.StructName}}Service) Delete{{.StructName}}ByIds(ctx *request.Context, ids []int) (rows int64, err error) {
	db:= m.svcCtx.MainDB	
	query := db.Delete(&[]entity.{{.StructName}}{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取{{.StructName}}记录
func (m *{{.StructName}}Service) Get{{.StructName}}InfoList(ctx *request.Context, page *request.PageInfo) (list []*entity.{{.StructName}}, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db	
	db := m.svcCtx.MainDB
	var {{.ValueName}}s []*entity.{{.StructName}}
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.OrderKey != "" && page.Order != "" {
		db = db.Order(fmt.Sprintf("` + "`%v`" + ` %v", page.OrderKey, page.Order))
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&{{.ValueName}}s).Error
	return {{.ValueName}}s, total, err
}
`

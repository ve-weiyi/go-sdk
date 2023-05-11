package tmpl

const AppController = `
package controller

import (
	"{{.SvcPackage }}"
)

type AppController struct {
	svcCtx *svc.ControllerContext //持有的service引用
}

func NewController(svcCtx *svc.ControllerContext) *AppController {
	return &AppController{
		svcCtx: svcCtx,
	}
}
`
const Controller = `
package logic

import (
	"github.com/gin-gonic/gin"
	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

type {{.StructName}}Controller struct {
	controller.BaseController
}

func New{{.StructName}}Controller(ctx *svc.ControllerContext) *{{.StructName}}Controller {
	return &{{.StructName}}Controller{
		BaseController: controller.NewBaseController(ctx),
	}
}

// Create{{.StructName}} 创建{{.ChineseName}}
// @Tags	 {{.StructName}}
// @Summary  创建{{.ChineseName}}
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.{{.StructName}}		true  "创建{{.ChineseName}}"
// @Success  200   {object}  response.Response{data=entity.{{.StructName}}}  	"返回信息"
// @Router /{{.ValueName}}/create [post]
func (m *{{.StructName}}Controller) Create{{.StructName}}(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var {{.ValueName}} entity.{{.StructName}}
	err = m.ShouldBindJSON(c, &{{.ValueName}})
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.SvcCtx.{{.StructName}}Service.Create{{.StructName}}(busCtx, &{{.ValueName}});
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

// Delete{{.StructName}} 删除{{.ChineseName}}
// @Tags 	{{.StructName}}
// @Summary 删除{{.ChineseName}}
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body	 	entity.{{.StructName}} 		true "删除{{.ChineseName}}"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /{{.ValueName}}/delete [delete]
func (m *{{.StructName}}Controller) Delete{{.StructName}}(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var {{.ValueName}} entity.{{.StructName}}
	err = m.ShouldBindJSON(c, &{{.ValueName}})
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.SvcCtx.{{.StructName}}Service.Delete{{.StructName}}(busCtx, &{{.ValueName}});
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

// Update{{.StructName}} 更新{{.ChineseName}}
// @Tags 	{{.StructName}}
// @Summary 更新{{.ChineseName}}
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		entity.{{.StructName}} 		true "更新{{.ChineseName}}"
// @Success 200  {object}  	response.Response{data=entity.{{.StructName}}}  	"返回信息"
// @Router /{{.ValueName}}/update [put]
func (m *{{.StructName}}Controller) Update{{.StructName}}(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var {{.ValueName}} entity.{{.StructName}}
	err = m.ShouldBindJSON(c, &{{.ValueName}})
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.SvcCtx.{{.StructName}}Service.Update{{.StructName}}(busCtx, &{{.ValueName}});
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

// Find{{.StructName}} 用id查询{{.ChineseName}}
// @Tags 	{{.StructName}}
// @Summary 用id查询{{.ChineseName}}
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		entity.{{.StructName}} 		true "用id查询{{.ChineseName}}"
// @Success 200  {object}  	response.Response{data=entity.{{.StructName}}}  	"返回信息"
// @Router /{{.ValueName}}/find [get]
func (m *{{.StructName}}Controller) Find{{.StructName}}(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var {{.ValueName}} entity.{{.StructName}}
	err = c.ShouldBindQuery(&{{.ValueName}})
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.SvcCtx.{{.StructName}}Service.Find{{.StructName}}(busCtx, {{.ValueName}}.ID);
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

// Delete{{.StructName}}ByIds 批量删除{{.ChineseName}}
// @Tags 	{{.StructName}}
// @Summary 批量删除{{.ChineseName}}
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		request.IdsReq 			true "批量删除{{.ChineseName}}"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /{{.ValueName}}/deleteByIds [delete]
func (m *{{.StructName}}Controller) Delete{{.StructName}}ByIds(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var IDS []int
	err = m.ShouldBindJSON(c, &IDS)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	data, err := m.SvcCtx.{{.StructName}}Service.Delete{{.StructName}}ByIds(busCtx, IDS);
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	m.ResponseOk(c, data)
}

// Get{{.StructName}}List 分页获取{{.ChineseName}}列表
// @Tags 	{{.StructName}}
// @Summary 分页获取{{.ChineseName}}列表
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		blogReq.{{.StructName}}Search 	true "分页获取{{.ChineseName}}列表"
// @Success 200  {object}  	response.Response{data=response.PageResult{list=[]entity.{{.StructName}}}}  	"返回信息"
// @Router /{{.ValueName}}/list [get]
func (m *{{.StructName}}Controller) Get{{.StructName}}List(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var pageInfo request.PageInfo
	err = m.ShouldBindQuery(c, &pageInfo)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	list, total, err := m.SvcCtx.{{.StructName}}Service.Get{{.StructName}}InfoList(busCtx, &pageInfo); 
	if err != nil {
		m.ResponseError(c, err)
		return
	}
	
	m.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	})
}
`

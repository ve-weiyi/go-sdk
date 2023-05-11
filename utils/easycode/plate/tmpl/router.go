package tmpl

const AppRouter = `
package router

import (
	"{{.SvcPackage }}"
)

type AppRouter struct {
	svcCtx *svc.RouterContext //持有的controller引用
}

func NewRouter(svcCtx *svc.RouterContext) *AppRouter {
	return &AppRouter{
		svcCtx: svcCtx,
	}
}
`
const Router = `
package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/middleware"

	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

type {{.StructName}}Router struct {
	svcCtx *svc.RouterContext
}

func New{{.StructName}}Router(ctx *svc.RouterContext) *{{.StructName}}Router {
	return &{{.StructName}}Router{
		svcCtx: ctx,
	}
}

// Init{{.StructName}}Router 初始化 {{.StructName}} 路由信息
func (s *{{.StructName}}Router) Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.ValueName}}Router := Router.Group("{{.ValueName}}")
	{{.ValueName}}OperationRouter := Router.Group("{{.ValueName}}").Use(middleware.JwtToken()).Use(middleware.OperationRecord())
	var self = s.svcCtx.AppController.{{.StructName}}Controller
	{
		{{.ValueName}}Router.GET("find", self.Find{{.StructName}})    // 根据ID获取{{.StructName}}
		{{.ValueName}}Router.GET("list", self.Get{{.StructName}}List) // 获取{{.StructName}}列表
	}
	{
		{{.ValueName}}OperationRouter.POST("create", self.Create{{.StructName}})             // 新建{{.StructName}}
		{{.ValueName}}OperationRouter.DELETE("delete", self.Delete{{.StructName}})           // 删除{{.StructName}}
		{{.ValueName}}OperationRouter.PUT("update", self.Update{{.StructName}})              // 更新{{.StructName}}
		{{.ValueName}}OperationRouter.DELETE("deleteByIds", self.Delete{{.StructName}}ByIds) // 批量删除{{.StructName}}
	}
}
`

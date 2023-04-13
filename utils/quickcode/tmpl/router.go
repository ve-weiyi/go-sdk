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
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/router/svc"
	"github.com/ve-weiyi/ve-admin-store/server/middleware"

	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
	svcCtx *svc.RouterContext
}

func NewArticleRouter(ctx *svc.RouterContext) *ArticleRouter {
	return &ArticleRouter{
		svcCtx: ctx,
	}
}

// InitArticleRouter 初始化 文章 路由信息
func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	articleRouterWithoutRecord := Router.Group("article")
	var self = s.svcCtx.ArticleCtl
	{
		articleRouter.POST("create", self.CreateArticle)             // 新建Article
		articleRouter.DELETE("delete", self.DeleteArticle)           // 删除Article
		articleRouter.DELETE("deleteByIds", self.DeleteArticleByIds) // 批量删除Article
		articleRouter.PUT("update", self.UpdateArticle)              // 更新Article
	}
	{
		articleRouterWithoutRecord.GET("find", self.FindArticle)    // 根据ID获取Article
		articleRouterWithoutRecord.GET("list", self.GetArticleList) // 获取Article列表
	}
}
`

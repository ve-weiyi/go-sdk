package tmpl

const AppController = `
package controller

import (
	"{{.SvcPackage }}"
)

type AppController struct {
	svcCtx *svc.CtlContext //持有的service引用
}

func NewController(svcCtx *svc.CtlContext) *AppController {
	return &AppController{
		svcCtx: svcCtx,
	}
}
`
const Controller = `base.go
package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/common/controller"
	"github.com/ve-weiyi/ve-admin-store/server/api/common/model/response"
)

type ArticleController struct {
	controller.BaseController
}

func NewArticleController(ctx *svc.CtlContext) *ArticleController {
	return &ArticleController{
		BaseController: controller.NewBaseController(ctx),
	}
}

// CreateArticle 创建文章
// @Tags	 Article
// @Summary  创建文章
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.Article		true  "创建文章"
// @Success  200   {object}  response.Response{data=entity.Article}  	"返回信息"
// @Router /article/create [post]
func (m *ArticleController) CreateArticle(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = c.ShouldBindJSON(&article)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.ArticleService.CreateArticle(busCtx, &article); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// DeleteArticle 删除文章
// @Tags 	Article
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body	 	entity.Article 		true "删除文章"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /article/delete [delete]
func (m *ArticleController) DeleteArticle(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = c.ShouldBindJSON(&article)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.ArticleService.DeleteArticle(busCtx, &article); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// UpdateArticle 更新文章
// @Tags 	Article
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		entity.Article 		true "更新文章"
// @Success 200  {object}  	response.Response{data=entity.Article}  	"返回信息"
// @Router /article/update [put]
func (m *ArticleController) UpdateArticle(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = c.ShouldBindJSON(&article)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.ArticleService.UpdateArticle(busCtx, &article); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// FindArticle 用id查询文章
// @Tags 	Article
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		entity.Article 		true "用id查询文章"
// @Success 200  {object}  	response.Response{data=entity.Article}  	"返回信息"
// @Router /article/find [get]
func (m *ArticleController) FindArticle(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var article entity.Article
	err = c.ShouldBindQuery(&article)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.ArticleService.FindArticle(busCtx, article.ID); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// DeleteArticleByIds 批量删除文章
// @Tags 	Article
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		request.IdsReq 			true "批量删除文章"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /article/deleteByIds [delete]
func (m *ArticleController) DeleteArticleByIds(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var IDS []int
	err = c.ShouldBindJSON(&IDS)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.ArticleService.DeleteArticleByIds(busCtx, IDS); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// GetArticleList 分页获取文章列表
// @Tags 	Article
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		blogReq.ArticleSearch 	true "分页获取文章列表"
// @Success 200  {object}  	response.Response{data=response.PageResult{list=[]entity.Article}}  	"返回信息"
// @Router /article/list [get]
func (m *ArticleController) GetArticleList(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var pageInfo request.ArticleSearch
	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if list, total, err := m.SvcCtx.ArticleService.GetArticleInfoList(busCtx, &pageInfo); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		})
	}
}
`

package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/go-sdk/utils/easycode/autocode_template/test/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/api/common/controller"
)

type AuthController struct {
	controller.BaseController
}

func NewAuthController(ctx *svc.CtlContext) *AuthController {
	return &AuthController{
		BaseController: controller.NewBaseController(ctx),
	}
}

// CreateAuth 创建文章
// @Tags	 Auth
// @Summary  创建文章
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.Auth		true  "创建文章"
// @Success  200   {object}  response.Response{data=entity.Auth}  	"返回信息"
// @Router /auth/create [post]
func (m *AuthController) CreateAuth(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var auth entity.Auth
	err = c.ShouldBindJSON(&auth)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.AuthService.CreateAuth(busCtx, &auth); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// DeleteAuth 删除文章
// @Tags 	Auth
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body	 	entity.Auth 		true "删除文章"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /auth/delete [delete]
func (m *AuthController) DeleteAuth(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var auth entity.Auth
	err = c.ShouldBindJSON(&auth)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.AuthService.DeleteAuth(busCtx, &auth); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// UpdateAuth 更新文章
// @Tags 	Auth
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		entity.Auth 		true "更新文章"
// @Success 200  {object}  	response.Response{data=entity.Auth}  	"返回信息"
// @Router /auth/update [put]
func (m *AuthController) UpdateAuth(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var auth entity.Auth
	err = c.ShouldBindJSON(&auth)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.AuthService.UpdateAuth(busCtx, &auth); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// FindAuth 用id查询文章
// @Tags 	Auth
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		entity.Auth 		true "用id查询文章"
// @Success 200  {object}  	response.Response{data=entity.Auth}  	"返回信息"
// @Router /auth/find [get]
func (m *AuthController) FindAuth(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var auth entity.Auth
	err = c.ShouldBindQuery(&auth)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if data, err := m.SvcCtx.AuthService.FindAuth(busCtx, auth.ID); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// DeleteAuthByIds 批量删除文章
// @Tags 	Auth
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		request.IdsReq 			true "批量删除文章"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /auth/deleteByIds [delete]
func (m *AuthController) DeleteAuthByIds(c *gin.Context) {
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

	if data, err := m.SvcCtx.AuthService.DeleteAuthByIds(busCtx, IDS); err != nil {
		m.ResponseError(c, err)
	} else {
		m.ResponseOk(c, data)
	}
}

// GetAuthList 分页获取文章列表
// @Tags 	Auth
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		blogReq.AuthSearch 	true "分页获取文章列表"
// @Success 200  {object}  	response.Response{data=response.PageResult{list=[]entity.Auth}}  	"返回信息"
// @Router /auth/list [get]
func (m *AuthController) GetAuthList(c *gin.Context) {
	busCtx, err := m.GetContent(c)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	var pageInfo request.PageInfo
	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		m.ResponseError(c, err)
		return
	}

	if list, total, err := m.SvcCtx.AuthService.GetAuthInfoList(busCtx, &pageInfo); err != nil {
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

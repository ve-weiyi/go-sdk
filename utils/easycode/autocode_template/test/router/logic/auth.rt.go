package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/router/svc"
	"github.com/ve-weiyi/ve-admin-store/server/middleware"
)

type AuthRouter struct {
	svcCtx *svc.RouterContext
}

func NewAuthRouter(ctx *svc.RouterContext) *AuthRouter {
	return &AuthRouter{
		svcCtx: ctx,
	}
}

// InitAuthRouter 初始化 文章 路由信息
func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth").Use(middleware.OperationRecord())
	authRouterWithoutRecord := Router.Group("auth")
	var self = s.svcCtx.AppController.AuthController
	{
		authRouter.POST("create", self.CreateAuth)             // 新建Auth
		authRouter.DELETE("delete", self.DeleteAuth)           // 删除Auth
		authRouter.DELETE("deleteByIds", self.DeleteAuthByIds) // 批量删除Auth
		authRouter.PUT("update", self.UpdateAuth)              // 更新Auth
	}
	{
		authRouterWithoutRecord.GET("find", self.FindAuth)    // 根据ID获取Auth
		authRouterWithoutRecord.GET("list", self.GetAuthList) // 获取Auth列表
	}
}

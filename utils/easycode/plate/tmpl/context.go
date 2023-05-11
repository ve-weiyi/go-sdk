package tmpl

const ControllerContext = `
package svc

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
	"github.com/ve-weiyi/ve-admin-store/server/config"
)

// 注册需要用到的rpc
type ControllerContext struct {
	*service.AppService
}

func NewControllerContext(cfg *config.Config) *ControllerContext {
	ctx := svc.NewServiceContext(cfg)
	sv := service.NewService(ctx)
	if sv == nil {
		panic("sv cannot be null")
	}

	return &ControllerContext{
		AppService: sv,
	}
}

`

const RouterContext = `
package svc

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/config"
)

// 注册需要用到的api
type RouterContext struct {
	*controller.AppController
}

func NewRouterContext(cfg *config.Config) *RouterContext {
	ctx := svc.NewControllerContext(cfg)
	ctl := controller.NewController(ctx)
	if ctl == nil {
		panic("ctl cannot be null")
	}

	return &RouterContext{
		AppController: ctl,
	}
}

`

const ServiceContext = `
package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/go-sdk/utils/glog"
	"github.com/ve-weiyi/ve-admin-store/server/config"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"gorm.io/gorm"
)

// 注册需要用到的gorm、redis、model
type ServiceContext struct {
	Cfg    *config.Config
	MainDB *gorm.DB
	DBList map[string]*gorm.DB
	Cache  *redis.Client
	Log    *glog.Glogger
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	return &ServiceContext{
		Cfg:    cfg,
		MainDB: global.GVA_DB,
		DBList: global.GVA_DBList,
		Cache:  global.GVA_REDIS,
		Log:    global.GVA_LOG,
	}
}

`

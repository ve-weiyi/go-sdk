package logic

import (
	"fmt"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"

	"github.com/ve-weiyi/go-sdk/utils/easycode/autocode_template/test/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type AuthService struct {
	svcCtx *svc.ServiceContext
}

func NewAuthService(svcCtx *svc.ServiceContext) *AuthService {
	return &AuthService{
		svcCtx: svcCtx,
	}
}

// 创建Auth记录
func (m *AuthService) CreateAuth(ctx *request.Context, auth *entity.Auth) (data *entity.Auth, err error) {
	err = m.svcCtx.MainDB.Create(&auth).Error
	if err != nil {
		return nil, err
	}
	return auth, err
}

// 删除Auth记录
func (m *AuthService) DeleteAuth(ctx *request.Context, auth *entity.Auth) (rows int64, err error) {
	query := m.svcCtx.MainDB.Delete(&auth)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Auth记录
func (m *AuthService) UpdateAuth(ctx *request.Context, auth *entity.Auth) (data *entity.Auth, err error) {
	err = m.svcCtx.MainDB.Save(&auth).Error
	if err != nil {
		return nil, err
	}
	return auth, err
}

// 根据id获取Auth记录
func (m *AuthService) FindAuth(ctx *request.Context, id int) (data *entity.Auth, err error) {
	err = m.svcCtx.MainDB.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

// 批量删除Auth记录
func (m *AuthService) DeleteAuthByIds(ctx *request.Context, ids []int) (rows int64, err error) {
	query := m.svcCtx.MainDB.Delete(&[]entity.Auth{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取Auth记录
func (m *AuthService) GetAuthInfoList(ctx *request.Context, page *request.PageInfo) (list []*entity.Auth, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := m.svcCtx.MainDB.Model(&entity.Auth{})
	var auths []*entity.Auth
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.OrderKey != "" && page.Order != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.OrderKey, page.Order))
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&auths).Error
	return auths, total, err
}

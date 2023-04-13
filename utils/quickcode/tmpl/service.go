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
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type ArticleService struct {
	svcCtx *svc.ServiceContext
}

func NewArticleService(svcCtx *svc.ServiceContext) *ArticleService {
	return &ArticleService{
		svcCtx: svcCtx,
	}
}

// 创建Article记录
func (m *ArticleService) CreateArticle(ctx *request.Context, article *entity.Article) (data *entity.Article, err error) {
	err = m.svcCtx.MainDB.Create(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 删除Article记录
func (m *ArticleService) DeleteArticle(ctx *request.Context, article *entity.Article) (rows int64, err error) {
	query := m.svcCtx.MainDB.Delete(&article)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Article记录
func (m *ArticleService) UpdateArticle(ctx *request.Context, article *entity.Article) (data *entity.Article, err error) {
	err = m.svcCtx.MainDB.Save(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 根据id获取Article记录
func (m *ArticleService) FindArticle(ctx *request.Context, id uint) (data *entity.Article, err error) {
	err = m.svcCtx.MainDB.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

// 批量删除Article记录
func (m *ArticleService) DeleteArticleByIds(ctx *request.Context, ids []int) (rows int64, err error) {
	query := m.svcCtx.MainDB.Delete(&[]entity.Article{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取Article记录
func (m *ArticleService) GetArticleInfoList(ctx *request.Context, info *request.ArticleSearch) (list []*entity.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := m.svcCtx.MainDB.Model(&entity.Article{})
	var articles []*entity.Article
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&articles).Error
	return articles, total, err
}
`

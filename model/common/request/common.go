package request

// 查询条件
type Condition struct {
	Uid       int    `json:"uid" example:""`
	Username  string `json:"username" example:""`
	Keywords  string `json:"keywords" example:""`
	StartTime string `json:"startTime" example:""`
	EndTime   string `json:"endTime" example:""`
	//Page      int    `json:"page" example:"0"`
}

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keywords string `json:"keywords" form:"keywords"` //关键字
}

const (
	DEFAULT_PAGE_SIZE = 10
	MAX_PAGE_SIZE     = 100
)

func (p *PageInfo) Fix() *PageInfo {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = DEFAULT_PAGE_SIZE
	}
	if p.PageSize > MAX_PAGE_SIZE {
		p.PageSize = MAX_PAGE_SIZE
	}
	return p
}

func (p *PageInfo) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

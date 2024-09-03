package pagination

import (
	"gorm.io/gorm"
)

// Pagination 分页器
type Pagination struct {
	Total int64  `json:"total"`
	Size  int    `json:"size,omitempty;query:size"`
	Page  int    `json:"page,omitempty;query:page"`
	Sort  string `json:"sort,omitempty;query:sort"`
}

// GetOffset 获取偏移量
func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Size
}

// GetSize 获取每页数量
func (p *Pagination) GetSize() int {
	if p.Size == 0 {
		p.Size = 10
	}
	return p.Size
}

// GetPage 获取当前页
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

// GetSort 获取排序
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
		return "id desc"
	}
	return p.Sort
}

// Paginate 分页
func (p *Pagination) Paginate(model interface{}, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var total int64
	db.Model(model).Count(&total)
	p.Total = total
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetSize()).Order(p.GetSort())
	}
}

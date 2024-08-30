package pagination

// Paginator 分页器
type Paginator struct {
	Total int64 `json:"total"`
	Size  int   `json:"size"`
	Page  int   `json:"page"`
}

// Setup 设置分页器
func (paginator *Paginator) Setup() {
	if paginator.Size == 0 {
		paginator.Size = 10
	}
	if paginator.Page == 0 {
		paginator.Page = 1
	}
}

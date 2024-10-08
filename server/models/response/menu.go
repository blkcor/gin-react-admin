package response

// MenuListResponse 菜单列表响应
type MenuListResponse struct {
	BaseResponse[[]MenuGroup]
}

// MenuItem 菜单项: 菜单的基本信息
type MenuItem struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

// MenuGroup 菜单组: 一个父菜单和多个子菜单
type MenuGroup struct {
	ParentMenu MenuItem   `json:"parent_menu"`
	ChildMenus []MenuItem `json:"child_menus"`
}

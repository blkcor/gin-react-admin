package v1

import (
	"github.com/blkcor/gin-react-admin/core/db"
	"github.com/blkcor/gin-react-admin/models/model"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMenu godoc
//
// @Summary      获取菜单
// @Description  获取用户菜单接口
// @Tags         菜单相关接口
// @Produce      json
// @Success      200  {object}  response.MenuListResponse  "获取菜单成功，返回当前用户拥有的菜单信息"
// @Failure      400  {object}  map[string]interface{}  "参数错误，返回详细错误信息"
// @Failure      401  {object}  map[string]interface{}  "用户认证失败，包括用户名不存在、密码错误或验证码错误"
// @Failure      500  {object}  map[string]interface{}  "服务器内部错误，返回详细错误信息"
// @Router       /v1/menu [get]
// @Security     ApiKeyAuth
func GetMenu(context *gin.Context) {
	//首先拿到当前用户角色对应的所有menu id
	claim, err := jwt.GetClaimFromContext(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户认证失败，请重新登录",
		})
		return
	}
	roleId := claim.RoleId
	var roleMenus []model.RoleMenu
	db.DB.Where("role_id = ?", roleId).Find(&roleMenus)
	var menuIds []uint32
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuID)
	}
	//然后根据menu id查询所有menu
	var menus []model.Menu
	db.DB.Where("id in (?)", menuIds).Find(&menus)
	//组合最终结果
	//1、找到所有的一级菜单
	var flm []model.Menu
	for _, menu := range menus {
		if menu.Layer == 1 {
			flm = append(flm, menu)
		}
	}
	var slm []model.Menu
	for _, menu := range menus {
		if menu.Layer == 2 {
			slm = append(slm, menu)
		}
	}

	var result response.MenuListResponse

	for _, f := range flm {
		menuGroup := response.MenuGroup{}
		fMenuItem := response.MenuItem{
			ID:   f.ID,
			Name: f.Name,
			Icon: f.Icon,
		}
		menuGroup.ParentMenu = fMenuItem
		for _, s := range slm {
			if f.ID == s.ParentID {
				sMenuItem := response.MenuItem{
					ID:   s.ID,
					Name: s.Name,
					Icon: s.Icon,
				}
				menuGroup.ChildMenus = append(menuGroup.ChildMenus, sMenuItem)
			}
		}
		result.MenuGroups = append(result.MenuGroups, menuGroup)
	}

	context.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

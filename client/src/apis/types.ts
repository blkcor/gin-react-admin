import { BaseResponse } from './base'

/**登录请求参数 */
export interface LoginRequest {
  username: string
  password: string
  captcha: string
}

/**登录响应结果 */
export interface LoginResponse extends BaseResponse<UserInfo> {
  token: string
}

/**退出登录响应结果 */
export interface LogoutResponse extends BaseResponse<null> { }

/**用户信息 */
export interface UserInfo {
  userId: number
  username: string
  email: string
  avatar: string
  userRole: string
  roleCode: string
}

/**用户菜单响应结果 */
export interface MenuListResponse extends BaseResponse<MenuGroup[]> { }

export interface MenuGroup {
  parent_menu: MenuItem
  child_menus: MenuItem[]
}

export interface MenuItem {
  id: number
  name: string
  icon: string
  path: string
}

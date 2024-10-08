import { alovaIns } from '@/http/alova'
import { LoginRequest, LoginResponse, LogoutResponse, MenuListResponse } from './types'
const getRequest = <T>(url: string, config?: any) => alovaIns.Get<T>(url, { ...config })
const postRequest = <T>(url: string, params?: any) => alovaIns.Post<T, unknown>(url, params)
// const putRequest = <T>(url: string, params?: any) => alovaIns.Put<T, unknown>(url, params)
// const deleteRequest = <T>(url: string, params?: any) => alovaIns.Delete<T, unknown>(url, params)

export default {
  /**获取登录验证码 */
  getCaptcha: () =>
    getRequest<any>('/captcha'),

  /**登录接口 */
  login: (loginRequest: LoginRequest) => postRequest<LoginResponse>('/login', loginRequest),

  /**退出登录 */
  logout: () => getRequest<LogoutResponse>('/logout'),

  /**获取用户菜单列表*/
  getMenuList: () => getRequest<MenuListResponse>('/v1/menu/list'),
}

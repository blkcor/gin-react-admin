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

/**用户信息 */
export interface UserInfo {
  userId: number
  username: string
  email: string
  avatar: string
  userRole: string
  roleCode: string
}

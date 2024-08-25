/**
 * 封装一些通用的类型
 */

/**通用返回结果 */
export interface BaseResponse<T> {
  message: string
  data: T
  success: boolean
}

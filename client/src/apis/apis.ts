import { alovaIns } from '@/http/alova'
const getRequest = <T>(url: string, config?: any) => alovaIns.Get<T>(url, { ...config })
const postRequest = <T>(url: string, params?: any) => alovaIns.Post<T, unknown>(url, params)
const putRequest = <T>(url: string, params?: any) => alovaIns.Put<T, unknown>(url, params)
const deleteRequest = <T>(url: string, params?: any) => alovaIns.Delete<T, unknown>(url, params)

export default {
  /**获取登录验证码 */
  getCaptcha: () =>
    getRequest<any>('/captcha', {
      //设置不进行缓存
      cacheFor: 0,
    }),
}

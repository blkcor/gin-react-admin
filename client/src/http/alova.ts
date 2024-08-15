import { createAlova } from 'alova'
import fetchAdapter from 'alova/fetch'
import { toast } from 'react-toastify'
import ReactHook from 'alova/react'

//装饰器模式
function getToken() {
  let tempToken = ''
  return {
    get() {
      if (tempToken) return tempToken
      const token = localStorage.getItem('token')
      if (token) {
        tempToken = token
      }
      return tempToken
    },
    clear() {
      tempToken = ''
    },
  }
}

const computedToken = getToken()

export const alovaIns = createAlova({
  baseURL: 'http://localhost:3000/api',
  timeout: 10000,
  statesHook: ReactHook,
  requestAdapter: fetchAdapter(),
  beforeRequest({ config }) {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${computedToken.get()}`
    }
  },
  // 响应拦截器，也与axios类似
  responded: async (response) => {
    const contentType = response.headers.get('Content-Type') || ''
    // 如果响应类型是 JSON
    if (contentType.includes('application/json')) {
      const json = await response.json()
      if (response.status !== 200 || !json.success) {
        if (json.errMsg) {
          if (!computedToken.get() && response.status === 401) {
            // do nothing
          } else {
            toast.error(json.errMsg, {
              position: 'top-center',
            })
          }
          throw new Error(json.errMsg)
        } else {
          throw new Error(json.message)
        }
      }
      return json.data
    }
    // 如果响应类型是文件
    if (contentType.includes('application/octet-stream') || contentType.includes('image/') || contentType.includes('application/pdf')) {
      return response.blob()
    }
  },
})

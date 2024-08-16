export interface HomeLoaderRepsonse {
  code: number
  message: string
}
/**
 * 路由拦截器定义：主要是用户在进行页面切换的时候对token有效性进行校验
 * */
export const HomeLoader = () => {
  //token或者userInfo不存在的时候，跳转到登录页面
  const token = localStorage.getItem('token')
  const userInfo = localStorage.getItem('userInfo')
  let data: HomeLoaderRepsonse = {} as HomeLoaderRepsonse
  if (!token || !userInfo) {
    data = {
      message: '认证失败，请重新登录',
      code: 401,
    }
  } else {
    data = {
      code: 200,
      message: 'success',
    }
  }

  return data
}

import { HomeLoaderRepsonse } from '@/routers/loader'
import { FC, PropsWithChildren, useEffect } from 'react'
import { useNavigate, useRouteLoaderData } from 'react-router-dom'

interface Iprops {
  code?: string
}

const Permission: FC<PropsWithChildren<Iprops>> = (props) => {
  const loaderData = useRouteLoaderData('home') as HomeLoaderRepsonse
  const navigate = useNavigate()

  useEffect(() => {
    // 这个root是我们在前面路由中定义了 id: 'root'
    if (loaderData.code === 401) {
      //重定向到登录页
      navigate('/login')
    }
  }, [])
  return <>{props.children}</>
}

export default Permission

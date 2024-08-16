import { lazy, Suspense } from 'react'
import { createBrowserRouter } from 'react-router-dom'
import HomepageLayout from '@/layouts/home'
import { HomeLoader } from './loader'
import Permission from '@/components/Permission'

const Home = lazy(() => import('@/views/home/index'))
const About = lazy(() => import('@/views/about/index'))
const Login = lazy(() => import('@/views/login/index'))

/**
 * @param Component 懒加载的组件
 * @param code 用于判断权限的字段
 * @returns
 */
const LazyLoad = (Component: React.LazyExoticComponent<() => JSX.Element>, code?: string) => {
  return (
    <Permission code={code}>
      <Suspense fallback={<div>loading...</div>}>
        <Component />
      </Suspense>
    </Permission>
  )
}

export const router = createBrowserRouter([
  {
    path: '/login',
    element: <Login />,
  },
  {
    id: 'home',
    path: '/',
    element: <HomepageLayout />,
    loader: HomeLoader,
    children: [
      {
        path: '/',
        element: LazyLoad(Home, 'home'),
      },
      {
        path: '/about',
        element: LazyLoad(About, 'about'),
      },
      {
        path: '*',
        element: <div>404</div>,
      },
    ],
  },
])

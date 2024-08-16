import { lazy } from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'

const HomepageLayout = lazy(() => import('@/layouts/home'))
const Home = lazy(() => import('@/views/home/index'))
const About = lazy(() => import('@/views/about/index'))
const Login = lazy(() => import('@/views/login/index'))

function App() {
  const router = createBrowserRouter([
    {
      path: '/login',
      element: <Login />,
    },
    {
      path: '/',
      element: <HomepageLayout />,
      children: [
        {
          path: '/',
          element: <Home />,
        },
        {
          path: '/about',
          element: <About />,
        },
        {
          path: '*',
          element: <div>404</div>,
        },
      ],
    },
  ])

  return (
    <>
      <Theme className="h-full w-full">
        <RouterProvider router={router} />
        <ToastContainer />
      </Theme>
    </>
  )
}

export default App

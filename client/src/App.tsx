import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import HomepageLayout from '@/layouts/home'
import Home from '@/views/home/index'
import About from '@/views/about'
import Login from './views/login'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'

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

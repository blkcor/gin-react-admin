import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import './App.css'
import Layout from '@/layouts/default'
import Home from '@/views/home/index'
import About from '@/views/about'
import Login from './views/login'

function App() {
  const router = createBrowserRouter([
    {
      path: '/login',
      element: <Login />,
    },
    {
      path: '/',
      element: <Layout />,
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
  return <RouterProvider router={router} />
}

export default App

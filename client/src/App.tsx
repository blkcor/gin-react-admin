import { RouterProvider } from 'react-router-dom'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import { router } from './routers'

function App() {
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

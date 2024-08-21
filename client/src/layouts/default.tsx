import Aside from '@/components/Aside'
import Footer from '@/components/Footer'
import Header from '@/components/Header'
import { Outlet } from 'react-router-dom'

export default function Root() {
  return (
    <div className="flex min-h-screen transition-all duration-500">
      <Aside />
      <div className="flex flex-col flex-grow">
        <Header />
        <main className="flex-grow p-3 bg-[#EFEFF5] dark:bg-[#111827]">
          <Outlet />
        </main>
        <Footer />
      </div>
    </div>
  )
}

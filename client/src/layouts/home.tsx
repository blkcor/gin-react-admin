import Aside from '@/components/Aside'
import Header from '@/components/Header'
import { Outlet } from 'react-router-dom'

export default function Root() {
  return (
    <div className="flex min-h-screen">
      <Aside />
      <div className="flex flex-col flex-grow">
        <Header />
        <main className="flex-grow bg-[#EFEFF5]">
          <Outlet />
        </main>
        <footer className="p-4 text-center">this is footer!</footer>
      </div>
    </div>
  )
}

import Aside from '@/components/aside/Aside'
import Footer from '@/components/Footer'
import Header from '@/components/Header'
import { useDark } from '@/hooks/useDark'
import { systemAtom } from '@/stores/systemAtom'
import { userAtom } from '@/stores/userAtom'
import { useAtom } from 'jotai'
import { useEffect } from 'react'
import { Outlet } from 'react-router-dom'
import { useNavigate } from 'react-router-dom'

export default function Root() {
  const [user, setUser] = useAtom(userAtom)
  const [systemConfig, setSystemConfig] = useAtom(systemAtom)
  const { darkMode } = useDark()
  const navigate = useNavigate()

  useEffect(() => {
    // 从localStorage进行初始化userAtom
    if (!user.token || !user.userInfo.userId) {
      const token = localStorage.getItem('token')!
      const userInfo = JSON.parse(localStorage.getItem('userInfo')!)
      setUser({
        token,
        userInfo,
      })
    }
    // 从localStorage进行初始化systemAtom
    if (!localStorage.getItem('systemConfig')) {
      localStorage.setItem('systemConfig', JSON.stringify(systemConfig))
    } else {
      setSystemConfig(JSON.parse(localStorage.getItem('systemConfig')!))
    }

    //初始化系统背景
    if (darkMode) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
    // Save preference tolocalStorage
    localStorage.setItem('darkMode', darkMode.toString())

    // 事件监听
    const handleUnauthorized = () => {
      navigate('/login')
    }

    window.addEventListener('unauthorized', handleUnauthorized)

    return () => {
      window.removeEventListener('unauthorized', handleUnauthorized)
    }
  }, [])

  return (
    <div className="flex min-h-screen transition-all duration-500 w-full ">
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

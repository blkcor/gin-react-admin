import { userAtom } from '@/stores/userAtom'
import { TextAlignLeftIcon, MoonIcon, SunIcon, GitHubLogoIcon, EnterFullScreenIcon, ExitFullScreenIcon } from '@radix-ui/react-icons'
import { Avatar } from '@radix-ui/themes'
import { useAtomValue } from 'jotai'
import { useEffect, useState } from 'react'

const Header = () => {
  const user = useAtomValue(userAtom)

  const [darkMode, setDarkMode] = useState(() => {
    const isDarkMode = localStorage.getItem('darkMode') === 'true' || (!('darkMode' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
    return isDarkMode
  })
  const [isFullScreen, setIsFullScreen] = useState(false)

  useEffect(() => {
    // Check for user preference in localStorage or system preference
    const isDarkMode = localStorage.getItem('darkMode') === 'true' || (!('darkMode' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
    setDarkMode(isDarkMode)
  }, [])

  useEffect(() => {
    // Apply dark mode class to body
    if (darkMode) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
    // Save preference tolocalStorage
    localStorage.setItem('darkMode', darkMode.toString())
  }, [darkMode])

  const toggleDarkMode = () => {
    setDarkMode(!darkMode)
  }

  function toggleFullScreen() {
    if (!document.fullscreenElement) {
      document.documentElement.requestFullscreen()
      setIsFullScreen(true)
    } else if (document.exitFullscreen) {
      document.exitFullscreen()
      setIsFullScreen(false)
    }
  }

  return (
    <div className="w-full flex p-[15.5px] gap-3 justify-between items-center layout-border-b bg-white dark:bg-gray-800">
      {/* 隐藏侧边栏的图标 */}
      <TextAlignLeftIcon className="w-6 h-6 cursor-pointer" />
      <div className="flex-1"></div>
      {/* 功能栏 */}
      <div className="flex gap-3 items-center ">
        <div onClick={toggleDarkMode}>{darkMode ? <MoonIcon className="w-4 h-4 cursor-pointer" /> : <SunIcon className="w-4 h-4 cursor-pointer" />}</div>

        <a href="https://www.github.com/blkcor" target="_blank">
          <GitHubLogoIcon className="w-4 h-4 cursor-pointer" />
        </a>
        <button onClick={toggleFullScreen}>{isFullScreen ? <ExitFullScreenIcon className="w-4 h-4 cursor-pointer" /> : <EnterFullScreenIcon className="w-4 h-4 cursor-pointer" />}</button>
        <div className="flex items-center flex-1 gap-2">
          <Avatar className="w-8 h-8 rounded-full" src={user.userInfo.avatar} fallback="A" />
          <div className="flex flex-col  items-center">
            <span className="text-xs text-black dark:text-white font-semibold">{user.userInfo.userRole}</span>
            <span className="text-xs text-[#7E7E7E]">[{user.userInfo.roleCode}]</span>
          </div>
        </div>
      </div>
    </div>
  )
}
export default Header

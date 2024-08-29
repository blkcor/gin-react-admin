import { userAtom } from '@/stores/userAtom'
import { ExitIcon, TextAlignLeftIcon, MoonIcon, SunIcon, GitHubLogoIcon, EnterFullScreenIcon, ExitFullScreenIcon } from '@radix-ui/react-icons'
import { Avatar } from '@radix-ui/themes'
import { useAtomValue } from 'jotai'
import * as DropdownMenu from '@radix-ui/react-dropdown-menu'
import { useDark } from '@/hooks/useDark'
import { useFullScreen } from '@/hooks/useFullScreen'
import { useCollapse } from '@/hooks/useCollapse'

const Header = () => {
  const user = useAtomValue(userAtom)

  const { darkMode, toggleDarkMode } = useDark()
  const { isFullScreen, toggleFullScreen } = useFullScreen()
  const { toggleCollapse } = useCollapse()
  return (
    <div className="w-full flex p-[15.5px] gap-3 justify-between items-center layout-border-b bg-white dark:bg-gray-800">
      {/* 隐藏侧边栏的图标 */}
      <TextAlignLeftIcon onClick={toggleCollapse} className="w-6 h-6 cursor-pointer" />
      <div className="flex-1"></div>
      {/* 功能栏 */}
      <div className="flex gap-2 items-center ">
        <div className="p-1 box-border rounded-md hover:bg-gray-400/30" onClick={toggleDarkMode}>
          {darkMode ? <MoonIcon className="w-4 h-4 cursor-pointer" /> : <SunIcon className="w-4 h-4 cursor-pointer" />}
        </div>

        <a href="https://www.github.com/blkcor" target="_blank" className="p-1 box-border rounded-md hover:bg-gray-400/30">
          <GitHubLogoIcon className="w-4 h-4 cursor-pointer" />
        </a>
        <button className="p-1 box-border rounded-md hover:bg-gray-400/30" onClick={toggleFullScreen}>
          {isFullScreen ? <ExitFullScreenIcon className="w-4 h-4 cursor-pointer" /> : <EnterFullScreenIcon className="w-4 h-4 cursor-pointer" />}
        </button>
        <DropdownMenu.Root>
          <div className="flex items-center flex-1 gap-2">
            <DropdownMenu.Trigger asChild>
              <Avatar className="w-8 h-8 rounded-full cursor-pointer" src={user.userInfo.avatar} fallback="A" />
            </DropdownMenu.Trigger>
            <DropdownMenu.Portal>
              <DropdownMenu.Content
                className="mt-1 py-2 w-36 rounded-lg bg-white shadow ring-1 ring-slate-900/5 text-sm leading-6 font-semibold text-slate-700 dark:bg-slate-900 dark:text-slate-300 dark:highlight-white/5"
                sideOffset={3}
              >
                <DropdownMenu.Item className="cursor-pointer flex items-center justify-between px-3 py-1">
                  <span>Logout</span>
                  <ExitIcon />
                </DropdownMenu.Item>
              </DropdownMenu.Content>
            </DropdownMenu.Portal>
            <div className="flex flex-col  items-center">
              <span className="text-xs text-black dark:text-white font-semibold">{user.userInfo.userRole}</span>
              <span className="text-xs text-[#7E7E7E]">[{user.userInfo.roleCode}]</span>
            </div>
          </div>
        </DropdownMenu.Root>
      </div>
    </div>
  )
}
export default Header

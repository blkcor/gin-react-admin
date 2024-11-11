import { userAtom } from '@/stores/userAtom'
import Avatar from '@mui/material/Avatar'
import ExitToAppIcon from '@mui/icons-material/ExitToApp'
import PersonIcon from '@mui/icons-material/Person'
import AlignHorizontalLeftIcon from '@mui/icons-material/AlignHorizontalLeft'
import AlignHorizontalRightIcon from '@mui/icons-material/AlignHorizontalRight'
import DarkModeIcon from '@mui/icons-material/DarkMode'
import LightModeIcon from '@mui/icons-material/LightMode'
import GitHubIcon from '@mui/icons-material/GitHub'
import FullscreenIcon from '@mui/icons-material/Fullscreen'
import FullscreenExitIcon from '@mui/icons-material/FullscreenExit'
import { useAtomValue } from 'jotai'
import { Dropdown } from '@mui/base/Dropdown'
import { MenuButton } from '@mui/base/MenuButton'
import { Menu } from '@mui/base/Menu'
import { MenuItem } from '@mui/base/MenuItem'
import { useDark } from '@/hooks/useDark'
import { useFullScreen } from '@/hooks/useFullScreen'
import { useCollapse } from '@/hooks/useCollapse'

const Header = () => {
  const user = useAtomValue(userAtom)

  const { darkMode, toggleDarkMode } = useDark()
  const { isFullScreen, toggleFullScreen } = useFullScreen()
  const { collapsed, toggleCollapse } = useCollapse()
  return (
    <div className="w-full flex p-[15.5px] gap-3 justify-between items-center h-16 dark:bg-gray-800">
      {/* 隐藏侧边栏的图标 */}
      {collapsed ? (
        <AlignHorizontalRightIcon onClick={toggleCollapse} className="cursor-pointer" fontSize="small" />
      ) : (
        <AlignHorizontalLeftIcon onClick={toggleCollapse} className="cursor-pointer" fontSize="small" />
      )}
      <div className="flex-1"></div>
      {/* 功能栏 */}
      <div className="flex gap-2 items-center ">
        <div className="p-1 box-border rounded-md hover:bg-gray-400/30" onClick={toggleDarkMode}>
          {darkMode ? <DarkModeIcon className="cursor-pointer" fontSize="small" /> : <LightModeIcon className="cursor-pointer" fontSize="small" />}
        </div>

        <a href="https://www.github.com/blkcor" target="_blank" className="p-1 box-border rounded-md hover:bg-gray-400/30">
          <GitHubIcon className="cursor-pointer" fontSize="small" />
        </a>
        <button className="p-1 box-border rounded-md hover:bg-gray-400/30" onClick={toggleFullScreen}>
          {isFullScreen ? <FullscreenExitIcon fontSize="small" className="cursor-pointer" /> : <FullscreenIcon fontSize="small" className="cursor-pointer" />}
        </button>
        <Dropdown>
          <MenuButton>
            <Avatar className="w-8 h-8 rounded-full cursor-pointer" src={user.userInfo.avatar} alt="A" />
          </MenuButton>

          <Menu className="mt-2 p-2 bg-gray-100 dark:bg-gray-700 rounded-md shadow-lg">
            <MenuItem className="mb-1 flex items-center gap-2 p-2 text-sm hover:bg-gray-200 dark:hover:bg-gray-600 rounded-md cursor-pointer hover:outline hover:outline-1 hover:outline-gray-300">
              <PersonIcon fontSize="small" className="text-gray-600 dark:text-gray-300" />
              <span className="text-gray-800 dark:text-gray-300">Profile</span>
            </MenuItem>
            <MenuItem className="flex items-center gap-2 p-2 text-sm hover:bg-gray-200 dark:hover:bg-gray-600 rounded-md cursor-pointer hover:outline hover:outline-1 hover:outline-gray-300">
              <ExitToAppIcon fontSize="small" className="text-red-600 dark:text-red-400" />
              <span className="text-red-600 dark:text-red-400">Logout</span>
            </MenuItem>
          </Menu>
        </Dropdown>
        <div className="flex flex-col  items-center">
          <span className="text-xs text-black font-semibold">{user.userInfo.userRole}</span>
          <span className="text-xs text-[#7E7E7E]">[{user.userInfo.roleCode}]</span>
        </div>
      </div>
    </div>
  )
}
export default Header

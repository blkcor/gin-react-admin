import { userAtom } from '@/stores/userAtom'
import { TextAlignLeftIcon, MoonIcon, SunIcon, GitHubLogoIcon, EnterFullScreenIcon } from '@radix-ui/react-icons'
import { Avatar } from '@radix-ui/themes'
import { useAtomValue } from 'jotai'

const Header = () => {
  const user = useAtomValue(userAtom)
  return (
    <div className="w-full flex py-2 px-4 gap-3 justify-between items-center border-solid border-b-2 border-middle">
      {/* 隐藏侧边栏的图标 */}
      <TextAlignLeftIcon className="w-6 h-6 cursor-pointer" />
      {/*TODO: tabs */}
      <div className="min-w-[70%] bg-red-200"></div>
      {/* 功能栏 */}
      <div className="flex gap-3 items-center ">
        <MoonIcon className="w-4 h-4 cursor-pointer" />
        <SunIcon className="w-4 h-4 cursor-pointer" />
        <GitHubLogoIcon className="w-4 h-4 cursor-pointer" />
        <EnterFullScreenIcon className="w-4 h-4 cursor-pointer" />
        <div className="flex items-center flex-1 gap-2">
          <Avatar className="w-8 h-8 rounded-full" src={user.userInfo.avatar} fallback="A" />
          <div className="flex flex-col  items-center">
            <span className="text-xs text-black font-semibold">{user.userInfo.userRole}</span>
            <span className="text-xs text-[#7E7E7E]">[{user.userInfo.roleCode}]</span>
          </div>
        </div>
      </div>
    </div>
  )
}
export default Header

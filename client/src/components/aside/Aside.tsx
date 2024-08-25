import logo from '@/assets/images/logo_with_no_word_transparent.svg'
import { useRequest } from 'alova/client'
import api from '@/apis/apis'
import { toast } from 'react-toastify'
import AsideMenu from './AsideMenu'
import { useState } from 'react'
import { MenuGroup } from '@/apis/types'

const Aside = () => {
  const [menuGroup, setMenuGroup] = useState<MenuGroup[]>()
  const { loading, onError, onSuccess } = useRequest(api.getMenuList, {
    immediate: true,
  })

  onSuccess((event) => {
    console.log(event.data)
    setMenuGroup(event.data.data)
  })

  onError((event) => {
    toast.error('获取菜单失败: ' + event.error, {
      position: 'top-center',
    })
  })

  return (
    <aside className="w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700">
      <div className="flex flex-col">
        <div className="flex items-center justify-center h-16 border-b border-gray-200 dark:border-gray-700">
          {/* logo区域 */}
          <div className="flex items-center p-2 gap-2">
            <img src={logo} alt="logo" className="w-10 h-10 " />
            <span className="text-2xs font-bold text-gray-800 dark:text-white">Gin React Admin</span>
          </div>
        </div>
        {loading ? (
          <div className="flex items-center justify-center flex-1">
            <p>加载中...</p>
          </div>
        ) : menuGroup ? (
          <AsideMenu menuData={menuGroup} />
        ) : (
          <div className="flex items-center justify-center flex-1">
            <p>无法加载菜单</p>
          </div>
        )}
      </div>
    </aside>
  )
}

export default Aside

import logo from '@/assets/images/logo_with_no_word_transparent.svg'
import { useRequest } from 'alova/client'
import api from '@/apis/apis'
import { toast } from 'react-toastify'
import AsideMenu from './AsideMenu'
import { useState } from 'react'
import { MenuGroup } from '@/apis/types'

import { motion, AnimatePresence } from 'framer-motion'
import { useCollapse } from '@/hooks/useCollapse'

const Aside = () => {
  const { loading, onError, onSuccess } = useRequest(api.getMenuList, {
    force: true,
  })

  onSuccess((event) => {
    setMenuGroup(event.data.data)
  })

  onError((event) => {
    toast.error('获取菜单失败: ' + event.error, {
      position: 'top-center',
    })
  })

  const [menuGroup, setMenuGroup] = useState<MenuGroup[]>()
  const { collapsed } = useCollapse()

  return (
    <motion.aside
      className={`bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 `}
      animate={{
        width: collapsed ? '5rem' : '16rem',
      }}
    >
      <div className="flex flex-col">
        <div className="flex items-center justify-center h-16 border-b border-gray-200 dark:border-gray-700">
          {/* logo区域 */}
          <motion.div layout className="flex items-center p-2 gap-2 ">
            <img src={logo} alt="logo" className="w-10 h-10 flex-shrink-0" />
            <AnimatePresence>
              <motion.span
                style={{ minWidth: collapsed ? '0px' : '160px' }}
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                transition={{ duration: 0.1, ease: 'easeInOut' }}
                className="text-2xs font-bold text-gray-800 dark:text-white flex-auto"
              >
                {!collapsed && 'Gin React Admin'}
              </motion.span>
            </AnimatePresence>
          </motion.div>
        </div>
        {loading ? (
          <div className="flex items-center justify-center flex-1">
            <p>加载中...</p>
          </div>
        ) : menuGroup ? (
          <AsideMenu menuData={menuGroup} collapsed={collapsed} />
        ) : (
          <div className="flex items-center justify-center flex-1">
            <p>无法加载菜单</p>
          </div>
        )}
      </div>
    </motion.aside>
  )
}

export default Aside

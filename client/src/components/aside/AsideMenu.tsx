import React, { useState } from 'react'
import MenuGroup from './MenuGroup'
import { MenuGroup as MenuGroupType } from '@/apis/types'

interface AsideMenuProps {
  menuData: MenuGroupType[]
  collapsed: boolean
}

const AsideMenu: React.FC<AsideMenuProps> = ({ menuData, collapsed }) => {
  const [openMenuId, setOpenMenuId] = useState<number | null>(null)

  const handleMenuToggle = (menuId: number) => {
    setOpenMenuId(openMenuId === menuId ? null : menuId)
  }

  return (
    <div className="h-full overflow-y-auto">
      {menuData.map((group) => (
        <MenuGroup
          collapsed={collapsed}
          key={group.parent_menu.id}
          parentMenu={group.parent_menu}
          childMenus={group.child_menus}
          isOpen={openMenuId === group.parent_menu.id}
          onToggle={() => handleMenuToggle(group.parent_menu.id)}
        />
      ))}
    </div>
  )
}

export default AsideMenu

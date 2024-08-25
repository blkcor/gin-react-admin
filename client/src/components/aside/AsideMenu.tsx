import React, { useState } from 'react'
import MenuGroup from './MenuGroup'
import { MenuGroup as MenuGroupType } from '@/apis/types'

interface AsideMenuProps {
  menuData: MenuGroupType[]
}

const AsideMenu: React.FC<AsideMenuProps> = ({ menuData }) => {
  const [openMenuId, setOpenMenuId] = useState<number | null>(null)

  const handleMenuToggle = (menuId: number) => {
    setOpenMenuId(openMenuId === menuId ? null : menuId)
  }

  return (
    <div className="h-full overflow-y-auto">
      {menuData.map((group) => (
        <MenuGroup
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

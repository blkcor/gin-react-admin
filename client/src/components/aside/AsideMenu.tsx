import React from 'react'
import { ScrollArea } from '@radix-ui/react-scroll-area'
import MenuGroup from './MenuGroup'
import { MenuGroup as MG } from '@/apis/types'

interface AsideMenuProps {
  menuData: MG[]
}

const AsideMenu: React.FC<AsideMenuProps> = ({ menuData }) => {
  return (
    <ScrollArea className="w-64">
      <div className="p-4">
        {menuData.map((group) => (
          <MenuGroup key={group.parent_menu.id} parentMenu={group.parent_menu} childMenus={group.child_menus} />
        ))}
      </div>
    </ScrollArea>
  )
}

export default AsideMenu

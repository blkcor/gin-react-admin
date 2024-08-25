import React from 'react'
import * as Collapsible from '@radix-ui/react-collapsible'
import MenuItem from './MenuItem'
import { ChevronDownIcon } from '@radix-ui/react-icons'
import { Icon } from '@iconify/react'

interface MenuGroupProps {
  parentMenu: {
    id: number
    name: string
    icon: string
  }
  childMenus: Array<{
    id: number
    name: string
    icon: string
  }>
}

const MenuGroup: React.FC<MenuGroupProps> = ({ parentMenu, childMenus }) => {
  return (
    <Collapsible.Root className="mb-4">
      <Collapsible.Trigger className="flex items-center justify-between w-full p-2  rounded">
        <div className="flex items-center gap-2">
          <Icon icon={'carbon:' + parentMenu.icon} width={20} height={20} />
          <span>{parentMenu.name}</span>
        </div>
        <ChevronDownIcon />
      </Collapsible.Trigger>
      <Collapsible.Content>
        <div className="pl-4 mt-2">
          {childMenus.map((item) => (
            <MenuItem key={item.id} item={item} />
          ))}
        </div>
      </Collapsible.Content>
    </Collapsible.Root>
  )
}

export default MenuGroup

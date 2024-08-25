import React from 'react'
import * as Collapsible from '@radix-ui/react-collapsible'
import MenuItem from './MenuItem'
import { ChevronRightIcon } from '@radix-ui/react-icons'
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
  isOpen: boolean
  onToggle: () => void
}

const MenuGroup: React.FC<MenuGroupProps> = ({ parentMenu, childMenus, isOpen, onToggle }) => {
  return (
    <Collapsible.Root className="mb-4" open={isOpen} onOpenChange={onToggle}>
      <Collapsible.Trigger className="flex items-center justify-between w-full p-4 gap-3 rounded hover:bg-[#ECF5FF] dark:hover:bg-[#343435]">
        <div className="flex items-center gap-2">
          <Icon icon={'carbon:' + parentMenu.icon} className="text-xl" />
          <span className=" text-[14px]">{parentMenu.name}</span>
        </div>
        <ChevronRightIcon className={`transform transition-transform duration-200 ${isOpen ? 'rotate-90' : ''}`} />
      </Collapsible.Trigger>
      <Collapsible.Content className="CollapsibleContent">
        <div className="mt-2">
          {childMenus.map((item) => (
            <MenuItem key={item.id} item={item} />
          ))}
        </div>
      </Collapsible.Content>
    </Collapsible.Root>
  )
}

export default MenuGroup

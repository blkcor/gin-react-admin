import React from 'react'
import { Icon } from '@iconify/react'

interface MenuItemProps {
  item: {
    id: number
    name: string
    icon: string
  }
}

const MenuItem: React.FC<MenuItemProps> = ({ item }) => {
  return (
    <div className="flex items-center py-4 pl-10 gap-3 rounded cursor-pointer hover:bg-[#ECF5FF] dark:hover:bg-[#343435]">
      <Icon icon={'carbon:' + item.icon} className="text-xl" />

      <span className="text-[14px]">{item.name}</span>
    </div>
  )
}

export default MenuItem
